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
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/ethtools/btcx"
	"github.com/polynetwork/explorer/internal/ethtools/eccm"
	"github.com/polynetwork/explorer/internal/ethtools/lockproxy"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	"math/big"
	"strings"
	"time"
)

const (
	_eth_crosschainlock   = "CrossChainLockEvent"
	_eth_crosschainunlock = "CrossChainUnlockEvent"
	_eth_lock = "LockEvent"
	_eth_unlock = "UnlockEvent"
)

// loadOntCrossTxFromChain synchronizes cross txs from Ethereum network
func (srv *Service) LoadEthCrossTxFromChain(context *ctx.Context) {
	chainInfo := srv.GetChain(srv.c.Ethereum.ChainId)
	chainClient := srv.ethClient
	t := time.NewTicker(srv.c.Ethereum.BlockDuration * time.Second)
	ethContext := ctx.New()
	for {
		select {
		case <-t.C:
			currentHeight, err := chainClient.GetCurrentBlockHeight(ethContext)
			if err != nil {
				log.Errorf("loadEthCrossTxFromChain: get current block height %s", err)
				continue
			}
			log.Infof("chain %s current height: %d, parser height: %d", chainInfo.Name, currentHeight, chainInfo.Height)
			for currentHeight > int64(chainInfo.Height) {
				tx, err := srv.dao.BeginTran()
				if err != nil {
					log.Errorf("loadEthCrossTxFromChain: BeginTran ", err)
					break
				}
				in, out, err := srv.saveEthCrossTxsByHeight(tx, ethContext, chainInfo)
				if err != nil {
					log.Errorf("loadEthCrossTxFromChain: saveEthCrossTxsByHeight %s", err)
					tx.Rollback()
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
					log.Errorf("loadEthCrossTxFromChain: update ChainInfo %s", err)
					break
				}
				if err = tx.Commit(); err != nil {
					tx.Rollback()
					chainInfo.Height --
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("LoadEthCrossTxFromChain: commit tx", err)
					break
				}
			}
		case <-context.Context.Done():
			log.Info("stop monitoring Ethereum Network")
			return
		}
	}

}

func (srv *Service) saveEthCrossTxsByHeight(tx *sql.Tx, ctx *ctx.Context, chainInfo *model.ChainInfo) (uint32, uint32, error) {
	sdk := srv.ethClient
	var in uint32 = 0
	var out uint32 = 0

	blockHeader, err := sdk.GetHeaderByNumber(ctx, int64(chainInfo.Height))
	if err == ethereum.NotFound {
		log.Errorf("saveEthCrossTxsByHeight: GetHeaderByNumber %v", err)
		chainInfo.Height--
		return in, out, nil
	}
	if err != nil {
		log.Errorf("saveEthCrossTxsByHeight: GetHeaderByNumber %v", err)
		return in, out, err
	}
	tt := blockHeader.Time
	for _, contract := range chainInfo.Contracts {
		eccmLockEvents, eccmUnLockEvents, err := srv.getECCMEventByBlockNumber(ctx, contract.Contract, uint64(chainInfo.Height))
		if err != nil {
			return 0, 0, err
		}
		ethLockEvents := make([]*model.LockEvent, 0)
		ethUnlockEvents := make([]*model.UnlockEvent, 0)
		{
			lockEventsT, unLockEventsT, err := srv.getProxyEventByBlockNumber(srv.c.Ethereum.Proxy, uint64(chainInfo.Height))
			if err != nil {
				return 0, 0, err
			}
			ethLockEvents = append(ethLockEvents, lockEventsT...)
			ethUnlockEvents = append(ethUnlockEvents, unLockEventsT...)
		}
		{
			lockEventsT, unLockEventsT, err := srv.getBTCXEventByBlockNumber(srv.c.Ethereum.BTCX, uint64(chainInfo.Height))
			if err != nil {
				return 0, 0, err
			}
			ethLockEvents = append(ethLockEvents, lockEventsT...)
			ethUnlockEvents = append(ethUnlockEvents, unLockEventsT...)
		}
		// save lockEvent to db
		for _, lockEvent := range eccmLockEvents {
			if lockEvent.Method == _eth_crosschainlock {
				log.Infof("from chain: %s, txhash: %s, txid: %s\n", chainInfo.Name, lockEvent.TxHash, lockEvent.Txid)
				fctx := &model.FChainTx{}
				fctx.Chain = srv.c.Ethereum.ChainId
				fctx.TxHash = lockEvent.Txid
				fctx.State = 1
				fctx.Fee = lockEvent.Fee
				fctx.TT = uint32(tt)
				fctx.Height = chainInfo.Height
				fctx.User = srv.Hash2Address(common.CHAIN_ETH, lockEvent.User)
				fctx.TChain = uint32(lockEvent.Tchain)
				fctx.Contract = lockEvent.Contract
				fctx.Key = lockEvent.TxHash
				fctx.Param = hex.EncodeToString(lockEvent.Value)
				for _, v := range ethLockEvents {
					if v.TxHash == lockEvent.TxHash {
						toAssetHash := v.ToAssetHash
						if v.ToChainId == common.CHAIN_ONT {
							toAssetHash = common.HexStringReverse(toAssetHash)
						}
						fctransfer := &model.FChainTransfer{}
						fctransfer.TxHash = lockEvent.Txid
						fctransfer.From = srv.Hash2Address(common.CHAIN_ETH, v.FromAddress)
						fctransfer.To = srv.Hash2Address(common.CHAIN_ETH, lockEvent.Contract)
						fctransfer.Asset = strings.ToLower(v.FromAssetHash)
						fctransfer.Amount = v.Amount
						fctransfer.ToChain = v.ToChainId
						fctransfer.ToAsset = toAssetHash
						fctransfer.ToUser = srv.Hash2Address(v.ToChainId, v.ToAddress)
						fctx.Transfer = fctransfer
						break
					}
				}
				if !srv.IsMonitorChain(fctx.TChain) {
					continue
				}
				if fctx.Transfer != nil && !srv.IsMonitorChain(fctx.Transfer.ToChain) {
					continue
				}
				err = srv.dao.TxInsertFChainTxAndCache(tx, fctx)
				if err != nil {
					log.Errorf("saveEthCrossTxsByHeight: InsertFChainTx %s", err)
					return 0, 0, err
				}
				out++
			}
		}
		// save unLockEvent to db
		for _, unLockEvent := range eccmUnLockEvents {
			if unLockEvent.Method == _eth_crosschainunlock {
				log.Infof("to chain: %s, txhash: %s\n", chainInfo.Name, unLockEvent.TxHash)
				tctx := &model.TChainTx{}
				tctx.Chain = srv.c.Ethereum.ChainId
				tctx.TxHash = unLockEvent.TxHash
				tctx.State = 1
				tctx.Fee = unLockEvent.Fee
				tctx.TT = uint32(tt)
				tctx.Height = chainInfo.Height
				tctx.FChain = unLockEvent.FChainId
				tctx.Contract = unLockEvent.Contract
				tctx.RTxHash = unLockEvent.RTxHash
				for _, v := range ethUnlockEvents {
					if v.TxHash == unLockEvent.TxHash {
						tctransfer := &model.TChainTransfer{}
						tctransfer.TxHash = unLockEvent.TxHash
						tctransfer.From = srv.Hash2Address(common.CHAIN_ETH, unLockEvent.Contract)
						tctransfer.To = srv.Hash2Address(common.CHAIN_ETH, v.ToAddress)
						tctransfer.Asset = strings.ToLower(v.ToAssetHash)
						tctransfer.Amount = v.Amount
						tctx.Transfer = tctransfer
						break
					}
				}
				if !srv.IsMonitorChain(tctx.FChain) {
					continue
				}
				err = srv.dao.TxInsertTChainTxAndCache(tx, tctx)
				if err != nil {
					log.Errorf("saveOntCrossTxsByHeight: InsertTChainTx %s", err)
					return 0, 0, err
				}
				in++
			}
		}
	}
	return in, out, nil
}

func (srv *Service) getECCMEventByBlockNumber(ctx *ctx.Context, contractAddr string, height uint64) ([]*model.ECCMLockEvent, []*model.ECCMUnlockEvent, error) {
	lockAddress := ethcommon.HexToAddress(contractAddr)
	lockContract, err := eccm_abi.NewEthCrossChainManager(lockAddress, srv.ethClient.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	ethLockEvents := make([]*model.ECCMLockEvent, 0)
	lockEvents, err := lockContract.FilterCrossChainEvent(opt, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}

	for lockEvents.Next() {
		evt := lockEvents.Event
		Fee := srv.GetConsumeGas(ctx, evt.Raw.TxHash)
		ethLockEvents = append(ethLockEvents, &model.ECCMLockEvent{
			Method:   _eth_crosschainlock,
			Txid:     hex.EncodeToString(evt.TxId),
			TxHash:   evt.Raw.TxHash.String()[2:],
			User:     strings.ToLower(evt.Sender.String()[2:]),
			Tchain:   uint32(evt.ToChainId),
			Contract: strings.ToLower(evt.ProxyOrAssetContract.String()[2:]),
			Value:    evt.Rawdata,
			Height:   height,
			Fee: Fee,
		})
	}

	// ethereum unlock events from given block
	ethUnlockEvents := make([]*model.ECCMUnlockEvent, 0)
	unlockEvents, err := lockContract.FilterVerifyHeaderAndExecuteTxEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
	}

	for unlockEvents.Next() {
		evt := unlockEvents.Event
		Fee := srv.GetConsumeGas(ctx, evt.Raw.TxHash)
		ethUnlockEvents = append(ethUnlockEvents, &model.ECCMUnlockEvent{
			Method:    _eth_crosschainunlock,
			TxHash:    evt.Raw.TxHash.String()[2:],
			RTxHash:   common.HexStringReverse(hex.EncodeToString(evt.CrossChainTxHash)),
			Contract:  hex.EncodeToString(evt.ToContract),
			FChainId:  uint32(evt.FromChainID),
			Height:    height,
			Fee: Fee,
		})
	}
	return ethLockEvents, ethUnlockEvents, nil
}

func (srv *Service) getProxyEventByBlockNumber(contractAddr string, height uint64) ([]*model.LockEvent, []*model.UnlockEvent, error) {
	proxyAddress := ethcommon.HexToAddress(contractAddr)
	lockContract, err := lock_proxy_abi.NewLockProxy(proxyAddress, srv.ethClient.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	ethLockEvents := make([]*model.LockEvent, 0)
	lockEvents, err := lockContract.FilterLockEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}

	for lockEvents.Next() {
		evt := lockEvents.Event
		ethLockEvents = append(ethLockEvents, &model.LockEvent{
			Method:   _eth_lock,
			TxHash:         evt.Raw.TxHash.String()[2:],
			FromAddress:     evt.FromAddress.String()[2:],
			FromAssetHash:   evt.FromAssetHash.String()[2:],
			ToChainId:     uint32(evt.ToChainId),
			ToAssetHash:   hex.EncodeToString(evt.ToAssetHash),
			ToAddress:  hex.EncodeToString(evt.ToAddress),
			Amount:    evt.Amount,
		})
	}

	// ethereum unlock events from given block
	ethUnlockEvents := make([]*model.UnlockEvent, 0)
	unlockEvents, err := lockContract.FilterUnlockEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
	}

	for unlockEvents.Next() {
		evt := unlockEvents.Event
		ethUnlockEvents = append(ethUnlockEvents, &model.UnlockEvent{
			Method:    _eth_unlock,
			TxHash:    evt.Raw.TxHash.String()[2:],
			ToAssetHash:    evt.ToAssetHash.String()[2:],
			ToAddress:   evt.ToAddress.String()[2:],
			Amount:  evt.Amount,
		})
	}
	return ethLockEvents, ethUnlockEvents, nil
}

func (srv *Service) getBTCXEventByBlockNumber(contractAddr string, height uint64) ([]*model.LockEvent, []*model.UnlockEvent, error) {
	proxyAddress := ethcommon.HexToAddress(contractAddr)
	lockContract, err := btcx_abi.NewBTCX(proxyAddress, srv.ethClient.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	ethLockEvents := make([]*model.LockEvent, 0)
	lockEvents, err := lockContract.FilterLockEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}

	for lockEvents.Next() {
		evt := lockEvents.Event
		ethLockEvents = append(ethLockEvents, &model.LockEvent{
			Method:   _eth_lock,
			TxHash:    evt.Raw.TxHash.String()[2:],
			FromAddress:     evt.FromAddress.String()[2:],
			FromAssetHash:   evt.FromAssetHash.String()[2:],
			ToChainId:     uint32(evt.ToChainId),
			ToAssetHash:   hex.EncodeToString(evt.ToAssetHash),
			ToAddress:  hex.EncodeToString(evt.ToAddress),
			Amount:    big.NewInt(int64(evt.Amount)),
		})
	}

	// ethereum unlock events from given block
	ethUnlockEvents := make([]*model.UnlockEvent, 0)
	unlockEvents, err := lockContract.FilterUnlockEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
	}

	for unlockEvents.Next() {
		evt := unlockEvents.Event
		ethUnlockEvents = append(ethUnlockEvents, &model.UnlockEvent{
			Method:    _eth_unlock,
			TxHash:    evt.Raw.TxHash.String()[2:],
			ToAssetHash:    evt.ToAssetHash.String()[2:],
			ToAddress:   evt.ToAddress.String()[2:],
			Amount:  big.NewInt(int64(evt.Amount)),
		})
	}
	return ethLockEvents, ethUnlockEvents, nil
}

func (srv *Service) GetConsumeGas(ctx *ctx.Context, hash ethcommon.Hash) uint64 {
	tx, err := srv.ethClient.GetTransactionByHash(ctx, hash)
	if err != nil {
		return 0
	}
	receipt, err := srv.ethClient.GetTransactionReceipt(ctx, hash)
	if err != nil {
		return 0
	}
	return tx.GasPrice().Uint64() * receipt.GasUsed
}
