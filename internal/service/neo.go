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
	"fmt"
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	"math/big"
	"strconv"
	"time"
)

const (
	_neo_crosschainlock = "CrossChainLockEvent"
	_neo_crosschainunlock = "CrossChainUnlockEvent"
	_neo_lock               = "Lock"
	_neo_lock2              = "LockEvent"
	_neo_unlock             = "UnlockEvent"
	_neo_unlock2            = "Unlock"
)

// loadOntCrossTxFromChain synchronizes cross txs from Neo network
func (srv *Service) LoadNeoCrossTxFromChain(context *ctx.Context) {
	chainInfo := srv.GetChain(srv.c.Neo.ChainId)
	neoClient := srv.neoClient
	t := time.NewTicker(srv.c.Neo.BlockDuration * time.Second)
	for {
		select {
		case <-t.C:
			countrep := neoClient.GetBlockCount()
			if countrep.ErrorResponse.Error.Message != "" {
				log.Errorf("loadNeoCrossTxFromChain, GetBlockCount err: %v", countrep.ErrorResponse)
				continue
			}
			log.Infof("chain %s current height: %d, parser height: %d", chainInfo.Name, countrep.Result, chainInfo.Height)
			for uint32(countrep.Result) > chainInfo.Height + 1 {
				tx, err := srv.dao.BeginTran()
				if err != nil {
					log.Errorf("loadNeoCrossTxFromChain: BeginTran ", err)
					break
				}
				in, out, err := srv.saveNeoCrossTxsByHeight(tx, chainInfo)
				if err != nil {
					tx.Rollback()
					log.Errorf("loadNeoCrossTxFromChain: %s", err)
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
					log.Errorf("loadNeoCrossTxFromChain: update ChainInfo %s", err)
					break
				}
				if err = tx.Commit(); err != nil {
					tx.Rollback()
					chainInfo.Height --
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("LoadNeoCrossTxFromChain: commit tx", err)
					break
				}
			}
		case <-context.Context.Done():
			log.Info("stop monitoring NEO Network")
			return
		}
	}
}

func (srv *Service) saveNeoCrossTxsByHeight(sqlTx *sql.Tx, chainInfo *model.ChainInfo) (in uint32, out uint32, err error) {
	neoClient := srv.neoClient
	blockResp := neoClient.GetBlockByIndex(chainInfo.Height)
	if blockResp.ErrorResponse.Error.Message != "" {
		return 0, 0, fmt.Errorf("get Block: %v", blockResp.ErrorResponse.Error.Message)
	}

	block := blockResp.Result
	tt := block.Time

	in = 0
	out = 0
	for _, tx := range block.Tx {
		if tx.Type != "InvocationTransaction" {
			continue
		}
		appLogResp := neoClient.GetApplicationLog(tx.Txid)
		if appLogResp.ErrorResponse.Error.Message != "" {
			continue
		}
		appLog := appLogResp.Result
		for _, exeitem := range appLog.Executions {
			for _, notify := range exeitem.Notifications {
				for _, contract := range chainInfo.Contracts {
					if notify.Contract[2:] != contract.Contract {
						continue
					}
					if len(notify.State.Value) <= 0 {
						continue
					}
					contractMethod := srv.parseNeoMethod(notify.State.Value[0].Value)
					switch contractMethod {
					case _neo_crosschainlock:
						log.Infof("from chain: %s, txhash: %s\n", chainInfo.Name, tx.Txid[2:])
						fctransfer := &model.FChainTransfer{}
						if len(notify.State.Value) < 6 {
							continue
						}
						//
						for _, notifynew := range exeitem.Notifications {
							contractMethodNew := srv.parseNeoMethod(notifynew.State.Value[0].Value)
							if contractMethodNew == _neo_lock || contractMethodNew == _neo_lock2 {
								if len(notifynew.State.Value) < 7 {
									continue
								}
								fctransfer.TxHash = tx.Txid[2:]
								fctransfer.From = srv.Hash2Address(common.CHAIN_NEO, notifynew.State.Value[2].Value)
								fctransfer.To = srv.Hash2Address(common.CHAIN_NEO, notify.State.Value[2].Value)
								fctransfer.Asset = common.HexStringReverse(notifynew.State.Value[1].Value)
								//amount, _ := strconv.ParseUint(common.HexStringReverse(notifynew.State.Value[6].Value), 16, 64)
								amount := big.NewInt(0)
								if notifynew.State.Value[6].Type == "Integer" {
									amount, _ = new(big.Int).SetString(notifynew.State.Value[6].Value, 10)
								} else {
									amount, _ = new(big.Int).SetString(common.HexStringReverse(notifynew.State.Value[6].Value), 16)
								}
								fctransfer.Amount = amount
								tchainId, _ := strconv.ParseUint(notifynew.State.Value[3].Value, 10, 32)
								fctransfer.ToChain = uint32(tchainId)
								if len(notifynew.State.Value[5].Value) != 40 {
									continue
								}
								fctransfer.ToUser = srv.Hash2Address(uint32(tchainId), notifynew.State.Value[5].Value)
								if uint32(tchainId) == srv.c.Bitcoin.ChainId {
									fctransfer.ToAsset = common.BTC_TOKEN_HASH
								} else {
									fctransfer.ToAsset = notifynew.State.Value[4].Value
								}
								break
							}
						}
						fctx := &model.FChainTx{}
						fctx.Chain = srv.c.Neo.ChainId
						fctx.TxHash = tx.Txid[2:]
						fctx.State = 1
						fctx.Fee = uint64(common.String2Float64(exeitem.GasConsumed))
						fctx.TT = uint32(tt)
						fctx.Height = chainInfo.Height
						fctx.User = fctransfer.From
						toChainId, _ := strconv.ParseInt(notify.State.Value[3].Value, 10, 64)
						fctx.TChain = uint32(toChainId)
						fctx.Contract = notify.State.Value[2].Value
						fctx.Key = notify.State.Value[4].Value
						fctx.Param = notify.State.Value[5].Value
						fctx.Transfer = fctransfer
						if !srv.IsMonitorChain(fctx.TChain) {
							continue
						}
						if fctx.Transfer != nil && !srv.IsMonitorChain(fctx.Transfer.ToChain) {
							continue
						}
						err := srv.dao.TxInsertFChainTxAndCache(sqlTx, fctx)
						if err != nil {
							return 0, 0, err
						}
						out++
					case _neo_crosschainunlock:
						log.Infof("to chain: %s, txhash: %s\n", chainInfo.Name, tx.Txid[2:])
						tctransfer := &model.TChainTransfer{}
						if len(notify.State.Value) < 4 {
							continue
						}
						for _, notifynew := range exeitem.Notifications {
							contractMethodNew := srv.parseNeoMethod(notifynew.State.Value[0].Value)
							if contractMethodNew == _neo_unlock || contractMethodNew == _neo_unlock2 {
								if len(notifynew.State.Value) < 4 {
									continue
								}
								tctransfer.TxHash = tx.Txid[2:]
								tctransfer.From = srv.Hash2Address(common.CHAIN_NEO, notify.State.Value[2].Value)
								tctransfer.To = srv.Hash2Address(common.CHAIN_NEO, notifynew.State.Value[2].Value)
								tctransfer.Asset = common.HexStringReverse(notifynew.State.Value[1].Value)
								//amount, _ := strconv.ParseUint(common.HexStringReverse(notifynew.State.Value[3].Value), 16, 64)
								amount := big.NewInt(0)
								if notifynew.State.Value[3].Type == "Integer" {
									amount, _ = new(big.Int).SetString(notifynew.State.Value[3].Value, 10)
								} else {
									amount, _ = new(big.Int).SetString(common.HexStringReverse(notifynew.State.Value[3].Value), 16)
								}
								tctransfer.Amount = amount
								break
							}
						}
						tctx := &model.TChainTx{}
						tctx.Chain = srv.c.Neo.ChainId
						tctx.TxHash = tx.Txid[2:]
						tctx.State = 1
						tctx.Fee = uint64(common.String2Float64(exeitem.GasConsumed))
						tctx.TT = uint32(tt)
						tctx.Height = chainInfo.Height
						fchainId, _ := strconv.ParseUint(notify.State.Value[1].Value, 10, 32)
						tctx.FChain = uint32(fchainId)
						tctx.Contract = common.HexStringReverse(notify.State.Value[2].Value)
						tctx.RTxHash = common.HexStringReverse(notify.State.Value[3].Value)
						tctx.Transfer = tctransfer
						if !srv.IsMonitorChain(tctx.FChain) {
							continue
						}
						err = srv.dao.TxInsertTChainTxAndCache(sqlTx, tctx)
						if err != nil {
							return 0, 0, err
						}
						in++
					default:
						log.Warnf("ignore method: %s", contractMethod)
					}
				}
			}
		}
	}
	return in, out, nil
}

func (srv *Service) parseNeoMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}
