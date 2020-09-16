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
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	"math/big"
	"strconv"
	"time"
)

const (
	_ont_crosschainlock = "makeFromOntProof"
	_ont_crosschainunlock = "verifyToOntProof"
	_ont_lock               = "lock"
	_ont_unlock               = "unlock"
)

// loadOntCrossTxFromChain synchronizes cross txs from Ontology network
func (srv *Service) LoadOntCrossTxFromChain(context *ctx.Context) {
	sdk := srv.ontClient
	chainInfo := srv.GetChain(srv.c.Ontology.ChainId)
	t := time.NewTicker(srv.c.Ontology.BlockDuration * time.Second)
	for {
		select {
		case <-t.C:
			currentHeight, err := sdk.GetCurrentBlockHeight()
			if err != nil {
				log.Errorf("loadOntCrossTxFromChain: get current block height %s", err)
				continue
			}
			log.Infof("chain %s current height: %d, parser height: %d", chainInfo.Name, currentHeight, chainInfo.Height)
			for currentHeight > chainInfo.Height {
				tx, err := srv.dao.BeginTran()
				if err != nil {
					log.Errorf("loadOntCrossTxFromChain: BeginTran ", err)
					break
				}
				in, out, err := srv.saveOntCrossTxsByHeight(tx, chainInfo)
				if err != nil {
					tx.Rollback()
					log.Errorf("loadOntCrossTxFromChain: saveOntCrossTxsByHeight %s", err)
					break
				}
				chainInfo.Height++
				chainInfo.In += in
				chainInfo.Out += out
				err = srv.dao.TxUpdateChainInfoById(tx, chainInfo)
				if err != nil {
					tx.Rollback()
					chainInfo.Height --
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("loadOntCrossTxFromChain: update ChainInfo %s", err)
					break
				}
				if err = tx.Commit(); err != nil {
					tx.Rollback()
					chainInfo.Height --
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("loadOntCrossTxFromChain: commit tx", err)
					break
				}
			}
		case <-context.Context.Done():
			log.Info("stop monitoring Alliance Network")
			return
		}
	}
}

func (srv *Service) saveOntCrossTxsByHeight(tx *sql.Tx, chainInfo *model.ChainInfo) (in uint32, out uint32, err error) {
	sdk := srv.ontClient
	localHeight := chainInfo.Height
	in = 0
	out = 0
	block, err := sdk.GetBlockByHeight(localHeight)
	if err != nil {
		return
	}
	tt := block.Header.Timestamp
	events, err := sdk.GetSmartContractEventByBlock(localHeight)
	for _, event := range events {
		for _, notify := range event.Notify {
			for _, contract := range chainInfo.Contracts {
				if notify.ContractAddress != contract.Contract {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				switch contractMethod {
				case _ont_crosschainlock:
					log.Infof("from chain: %s, txhash: %s\n", chainInfo.Name, event.TxHash)
					fctransfer := &model.FChainTransfer{}
					if len(states) < 7 {
						continue
					}
					for _, notifynew := range event.Notify {
						statesnew := notifynew.States.([]interface{})
						//log.Errorf("%v", statesnew)
						method, ok := statesnew[0].(string)
						if !ok {
							continue
						}
						contractMethod := srv.parseOntolofyMethod(method)
						if contractMethod == _ont_lock {
							//
							if len(statesnew) < 7 {
								continue
							}
							fctransfer.TxHash = event.TxHash
							fctransfer.From = srv.Hash2Address(common.CHAIN_ONT, statesnew[2].(string))
							fctransfer.To = srv.Hash2Address(common.CHAIN_ONT, states[5].(string))
							fctransfer.Asset = common.HexStringReverse(statesnew[1].(string))
							if len(fctransfer.Asset) < 20 {
								continue
							}
							//amount, _ := strconv.ParseUint(common.HexStringReverse(statesnew[6].(string)), 16, 64)
							amount, _ := new(big.Int).SetString(common.HexStringReverse(statesnew[6].(string)), 16)
							fctransfer.Amount = amount
							toChain, _ := strconv.ParseUint(statesnew[3].(string), 16, 32)
							fctransfer.ToChain = uint32(toChain)
							if !srv.IsMonitorChain(fctransfer.ToChain) {
								continue
							}
							fctransfer.ToAsset = statesnew[4].(string)
							fctransfer.ToUser = srv.Hash2Address(uint32(toChain), statesnew[5].(string))
							break
						}
					}
					fctx := &model.FChainTx{}
					fctx.Chain = srv.c.Ontology.ChainId
					fctx.TxHash = event.TxHash
					fctx.State = event.State
					fctx.Fee = event.GasConsumed
					fctx.TT = tt
					fctx.Height = chainInfo.Height
					fctx.User = fctransfer.From
					fctx.TChain = uint32(states[2].(float64))
					fctx.Contract = common.HexStringReverse(states[5].(string))
					fctx.Key = states[4].(string)
					fctx.Param = states[6].(string)
					fctx.Transfer = fctransfer
					if !srv.IsMonitorChain(fctx.TChain) {
						continue
					}
					if fctx.Transfer != nil && !srv.IsMonitorChain(fctx.Transfer.ToChain) {
						continue
					}
					err := srv.dao.TxInsertFChainTxAndCache(tx, fctx)
					if err != nil {
						log.Errorf("saveOntCrossTxsByHeight: InsertFChainTx %s", err)
						return 0, 0, err
					}
					out++
				case _ont_crosschainunlock:
					log.Infof("to chain: %s, txhash: %s\n", chainInfo.Name, event.TxHash)
					tctransfer := &model.TChainTransfer{}
					if len(states) < 6 {
						continue
					}
					for _, notifynew := range event.Notify {
						statesnew := notifynew.States.([]interface{})
						//log.Errorf("%v", statesnew)
						method, ok := statesnew[0].(string)
						if !ok {
							continue
						}
						contractMethod := srv.parseOntolofyMethod(method)
						if contractMethod == _ont_unlock {
							//
							if len(statesnew) < 4 {
								continue
							}
							tctransfer.TxHash = event.TxHash
							tctransfer.From = srv.Hash2Address(common.CHAIN_ONT, states[5].(string))
							tctransfer.To = srv.Hash2Address(common.CHAIN_ONT, statesnew[2].(string))
							tctransfer.Asset = common.HexStringReverse(statesnew[1].(string))
							if len(tctransfer.Asset) < 20 {
								continue
							}
							//amount, _ := strconv.ParseUint(common.HexStringReverse(statesnew[3].(string)), 16, 64)
							amount, _ := new(big.Int).SetString(common.HexStringReverse(statesnew[3].(string)), 16)
							tctransfer.Amount = amount
							break
						}
					}
					tctx := &model.TChainTx{}
					tctx.Chain = srv.c.Ontology.ChainId
					tctx.TxHash = event.TxHash
					tctx.State = event.State
					tctx.Fee = event.GasConsumed
					tctx.TT = tt
					tctx.Height = chainInfo.Height
					tctx.FChain = uint32(states[3].(float64))
					tctx.Contract = common.HexStringReverse(states[5].(string))
					tctx.RTxHash = common.HexStringReverse(states[1].(string))
					tctx.Transfer = tctransfer
					if !srv.IsMonitorChain(tctx.FChain) {
						continue
					}
					err = srv.dao.TxInsertTChainTxAndCache(tx, tctx)
					if err != nil {
						log.Errorf("saveOntCrossTxsByHeight: InsertTChainTx %s", err)
						return 0, 0, err
					}
					in++
				default:
					log.Warnf("ignore method: %s", contractMethod)
					continue
				}
			}
		}
	}
	return
}

func (srv *Service) parseOntolofyMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}
