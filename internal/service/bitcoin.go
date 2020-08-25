/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package service

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	mccommon "github.com/polynetwork/poly/common"
	"golang.org/x/crypto/ripemd160"
	"time"
)

func (self *Service) MonitorBtcChainFromAlliance(context *ctx.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("recover info:", r)
		}
	}()

	chainInfo := self.GetChain(self.c.Bitcoin.ChainId)
	updateTicker := time.NewTicker(self.c.Bitcoin.BlockDuration * time.Second)
	for {
		select {
		case <-updateTicker.C:
			txs, err := self.dao.SelectAllianceTx(chainInfo.Height, chainInfo.Id)
			log.Infof("chain %s current height: %d", chainInfo.Name, chainInfo.Height)
			if err != nil {
				log.Error(err)
				continue
			}
			for _, mtx := range txs {
				tx, err := self.dao.BeginTran()
				if err != nil {
					log.Errorf("MonitorBtcChainFromAlliance: BeginTran ", err)
					break
				}
				in, out, err := self.parseTxFromAlliance(tx, mtx)
				if err != nil {
					tx.Rollback()
					log.Errorf("MonitorBtcChainFromAlliance: parseTxFromAlliance %s", err)
					break
				}
				oldHeight := chainInfo.Height
				chainInfo.Height = mtx.Height
				chainInfo.In += in
				chainInfo.Out += out
				err = self.dao.TxUpdateChainInfoById(tx, chainInfo)
				if err != nil {
					tx.Rollback()
					chainInfo.Height = oldHeight
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("MonitorBtcChainFromAlliance: update chain info failed: %s", err)
					break
				}
				if err = tx.Commit(); err != nil {
					tx.Rollback()
					chainInfo.Height = oldHeight
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("MonitorBtcChainFromAlliance: commit tx", err)
					break
				}
			}
			self.checkTx()
		case <-context.Context.Done():
			log.Info("stop monitoring Bitcoin Network")
			return
		}
	}
}

func (self *Service) parseTxFromAlliance(tx *sql.Tx, mtx *model.MChainTx) (uint32, uint32, error) {
	var in uint32 = 0
	var out uint32 = 0
	if mtx.FChain == self.c.Bitcoin.ChainId {
		rtxhash, _ := hex.DecodeString(mtx.FTxHash)
		btctx, err := self.bitcoinClient.GetTx(hex.EncodeToString(rtxhash))
		if err != nil {
			log.Errorf("parseTxFromAlliance: GetTx mtx.FTxHash: %s, err: %v", rtxhash, err)
			return 0, 0, err
		}
		log.Infof("from chain: %s, tx hash: %s\n", self.c.Bitcoin.Name, hex.EncodeToString(rtxhash))
		var fromAddress string
		var inputAmount float64
		var outputAmount float64
		//inputAddr := make(map[string]uint32)
		inputAddr := make([]string, 0)
		inputAddrExist := make(map[string]bool)
		for _, vin := range btctx.Vin {
			txOutRes, err := self.bitcoinClient.GetTx(vin.Txid)
			if err != nil {
				log.Errorf("parseTxFromAlliance: GetTxOut txid %s %v", vin.Txid, err)
				return 0, 0, err
			}
			inputAmount += txOutRes.Vout[vin.Vout].Value
			address := txOutRes.Vout[vin.Vout].ScriptPubKey.Addresses[0]
			_, ok := inputAddrExist[address]
			if ok {
				continue
			}
			inputAddr = append(inputAddr, txOutRes.Vout[vin.Vout].ScriptPubKey.Addresses[0])
			inputAddrExist[address] = true
		}

		for _, vout := range btctx.Vout {
			outputAmount += vout.Value
		}
		gas, err := btcutil.NewAmount(inputAmount - outputAmount)
		if err != nil {
			log.Errorf("parseTxFromAlliance: btcutil.NewAmount %v", err)
		}

		amount, err := btcutil.NewAmount(outputAmount)
		if err != nil {
			log.Errorf("parseTxFromAlliance: btcutil.NewAmount %v", err)
		}

		block, err := self.bitcoinClient.GetBlock(btctx.BlockHash)
		if err != nil {
			log.Errorf("parseTxFromAlliance: GetBlock %s, %v", btctx.BlockHash, err)
			return 0, 0, err
		}
		_, toaddress, tchain, _, extData, err := self.ParseBtcCrossTransfer(btctx)
		if err != nil {
			log.Errorf("parseTxFromAlliance: ParseCrossTransfer %s", err)
			return 0, 0, err
		}

		rk := GetUtxoKey1(&btctx.Vout[0].ScriptPubKey)
		/*
		toAddr := GetUtxoKey(&btctx.Vout[0].ScriptPubKey)
		outputAddr := make([]string, 0)
		outputAddr = append(outputAddr, toAddr)
		*/
		outputAddr := btctx.Vout[0].ScriptPubKey.Addresses

		fromAddressByte, err := json.Marshal(inputAddr)
		if err != nil {
			log.Errorf("parseTxFromAlliance: Marshal %s", err)
			return 0, 0, err
		}
		fromAddress = string(fromAddressByte)

		toAddressByte, err := json.Marshal(outputAddr)
		if err != nil {
			log.Errorf("parseTxFromAlliance: Marshal %s", err)
			return 0, 0, err
		}
		toAddress := string(toAddressByte)

		fctx := &model.FChainTx{}
		fctx.Chain = self.c.Bitcoin.ChainId
		fctx.TxHash = btctx.Hash
		fctx.State = 1
		fctx.Fee = uint64(gas)
		fctx.TT = uint32(block.Time)
		fctx.Height = uint32(block.Height)
		fctx.User = fromAddress
		fctx.TChain = uint32(tchain)
		fctx.Contract = rk
		fctx.Key = ""
		fctx.Param = hex.EncodeToString(extData)
		fxtransfer := &model.FChainTransfer{}
		fxtransfer.TxHash = btctx.Hash
		fxtransfer.From = fromAddress
		fxtransfer.To = toAddress
		fxtransfer.Asset = common.BTC_TOKEN_HASH
		fxtransfer.Amount = uint64(amount)
		fxtransfer.ToChain = uint32(tchain)
		token := self.SearchToken(common.BTC_TOKEN_NAME, uint32(tchain))
		tokenHash := ""
		if token != nil {
			tokenHash = token.Hash
		}
		fxtransfer.ToAsset = tokenHash
		fxtransfer.ToUser = self.Hash2Address(uint32(tchain), toaddress)
		fctx.Transfer = fxtransfer
		err = self.dao.TxInsertFChainTx(tx, fctx)
		if err != nil {
			log.Errorf("parseTxFromAlliance: InsertFChainTx %s", err)
			return 0, 0, err
		}
		out++
	} else if mtx.TChain == self.c.Bitcoin.ChainId {
		rawTx, _ := hex.DecodeString(mtx.Key)
		btcTx, txHash, err := self.bitcoinClient.PaserRawTx(rawTx)
		if err != nil {
			log.Errorf("Paser Btc Raw tx err: %v", err)
			return 0, 0, err
		}
		log.Infof("to chain: %s, tx hash: %s\n", self.c.Bitcoin.Name, txHash)
		btctx, err := self.bitcoinClient.GetTx(btcTx.TxIn[0].PreviousOutPoint.Hash.String())
		rk := GetUtxoKey1(&btctx.Vout[0].ScriptPubKey)
		/*
		fromAddr := GetUtxoKey(&btctx.Vout[0].ScriptPubKey)
		inputAddr := make([]string, 0)
		inputAddr = append(inputAddr, fromAddr)
		*/
		inputAddr := btctx.Vout[0].ScriptPubKey.Addresses
		fromAddressByte, err := json.Marshal(inputAddr)
		if err != nil {
			log.Errorf("parseTxFromAlliance: Marshal %s", err)
			return 0, 0, err
		}
		fromAddress := string(fromAddressByte)
		tctx := &model.TChainTx{}
		tctx.Chain = self.c.Bitcoin.ChainId
		tctx.TxHash = txHash
		tctx.State = 0
		tctx.Fee = 0
		tctx.TT = 0
		tctx.Height = 0
		tctx.FChain = mtx.FChain
		tctx.Contract = rk
		tctx.RTxHash = mtx.TxHash
		tctransfer := &model.TChainTransfer{}
		tctransfer.TxHash = txHash
		tctransfer.Asset = common.BTC_TOKEN_HASH
		tctransfer.From = fromAddress
		tctransfer.To = ""
		tctransfer.Amount = 0
		tctx.Transfer = tctransfer
		err = self.dao.TxInsertTChainTx(tx, tctx)
		if err != nil {
			log.Errorf("parseTxFromAlliance: InsertTChainTx %s", err)
			return 0, 0, err
		}
		in++
	}
	return in, out, nil
}

func (self *Service) checkTx() {
	txs, err := self.dao.SelectBitcoinTxUnConfirm(self.c.Bitcoin.ChainId)
	if err != nil {
		return
	}

	for i := 0; i < len(txs); i++ {
	LOOP:
		txhash := txs[i]
		btctx, err := self.bitcoinClient.GetTx(txhash)
		if err != nil {
			continue
		}
		block, err := self.bitcoinClient.GetBlock(btctx.BlockHash)
		if err != nil {
			continue
		}
		var inputAmount float64
		for _, vin := range btctx.Vin {
			txOutRes, err := self.bitcoinClient.GetTx(vin.Txid)
			if err != nil {
				log.Errorf("parseTxFromAlliance: mtx.TChain %v", err)
				if i < len(txs)-1 {
					i++
					goto LOOP
				}
				return
			}
			inputAmount += txOutRes.Vout[vin.Vout].Value
		}
		var outputAmount float64
		for _, vout := range btctx.Vout {
			outputAmount += vout.Value
		}
		gas, err := btcutil.NewAmount(inputAmount - outputAmount)
		if err != nil {
			log.Errorf("parseTxFromAlliance: btcutil.NewAmount %v", err)
		}
		/*
		toAddr := GetUtxoKey1(&btctx.Vout[0].ScriptPubKey)
		outputAddr := make([]string, 0)
		outputAddr = append(outputAddr, toAddr)
		*/
		outputAddr := btctx.Vout[0].ScriptPubKey.Addresses
		toAddressByte, err := json.Marshal(outputAddr)
		if err != nil {
			log.Errorf("parseTxFromAlliance: Marshal %s", err)
			return
		}
		toAddress := string(toAddressByte)

		log.Infof("to chain update: %s, tx hash: %s, to address: %s, output amount:%d\n", self.c.Bitcoin.Name, txhash, toAddress, uint64(outputAmount))
		err = self.dao.UpdateBitcoinTxConfirmed(txhash, uint32(block.Height), uint32(block.Time), uint64(gas), toAddress, uint64(outputAmount))
		if err != nil {
			log.Errorf("checkTx: UpdateBitcoinTxConfirmed transaction %s %v", txhash, err)
		}
	}
}

func (self *Service) ParseBtcCrossTransfer(rawTx *btcjson.TxRawResult) (xtype uint32, toAddress string, tChainid uint64, amount int64, extData []byte, err error) {
	extData, err = hex.DecodeString(rawTx.Vout[1].ScriptPubKey.Hex)
	if err != nil {
		return 0, "", 0, 0, nil, err
	}
	source := mccommon.NewZeroCopySource(extData[3:])
	tchainid, eof := source.NextUint64()
	if eof {
		return
	}
	amount, eof = source.NextInt64()
	if eof {
		return
	}
	addressBytes, eof := source.NextVarBytes()
	if eof {
		return
	}
	return 1, hex.EncodeToString(addressBytes), tchainid, amount, extData, nil
}

func GetUtxoKey1(scriptPk *btcjson.ScriptPubKeyResult) string {
	scriptPkBytes, _ := hex.DecodeString(scriptPk.Hex)
	switch scriptPk.Type {
	case "multisig":
		return hex.EncodeToString(btcutil.Hash160(scriptPkBytes))
	case "scripthash":
		return hex.EncodeToString(scriptPkBytes[2:22])
	case "witness_v0_scripthash":
		hasher := ripemd160.New()
		hasher.Write(scriptPkBytes[2:34])
		return hex.EncodeToString(hasher.Sum(nil))
	default:
		return ""
	}
}

func GetUtxoKey(scriptPk *btcjson.ScriptPubKeyResult) string {
	scriptPkBytes, _ := hex.DecodeString(scriptPk.Hex)
	switch scriptPk.Type {
	case "scripthash":
		add, _ := btcutil.NewAddressScriptHash(scriptPkBytes, &chaincfg.TestNet3Params)
		return add.EncodeAddress()
	case "witness_v0_scripthash":
		add, _ := btcutil.NewAddressWitnessScriptHash(scriptPkBytes[2:34], &chaincfg.TestNet3Params)
		return add.EncodeAddress()
	default:
		return "unsupport"
	}
}

