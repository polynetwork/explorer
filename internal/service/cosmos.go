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
	"strconv"
	"strings"
	"time"
)

const (
	_cosmos_crosschainlock    = "make_from_cosmos_proof"
	_cosmos_crosschainunlock  = "verify_to_cosmos_proof"
	_cosmos_lock = "lock"
	_cosmos_unlock = "unlock"
)

func (srv *Service) LoadCosmosCrossTxFromChain(context *ctx.Context) {
	chainInfo := srv.GetChain(srv.c.Cosmos.ChainId)
	chainClient := srv.cosmosClient
	t := time.NewTicker(srv.c.Cosmos.BlockDuration * time.Second)
	for {
		select {
		case <-t.C:
			status, err := chainClient.Status()
			if err != nil {
				log.Errorf("LoadCosmosCrossTxFromChain: get current block status %s", err)
				continue
			}
			log.Infof("chain %s current height: %d, parser height: %d", chainInfo.Name, status.SyncInfo.LatestBlockHeight, chainInfo.Height)
			for status.SyncInfo.LatestBlockHeight > int64(chainInfo.Height) {
				tx, err := srv.dao.BeginTran()
				if err != nil {
					log.Errorf("LoadCosmosCrossTxFromChain: BeginTran ", err)
					break
				}
				in, out, err := srv.saveCosmosCrossTxsByHeight(tx, chainInfo)
				if err != nil {
					log.Errorf("LoadCosmosCrossTxFromChain: saveEthCrossTxsByHeight %s", err)
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
					log.Errorf("LoadCosmosCrossTxFromChain: update ChainInfo %s", err)
					break
				}
				if err = tx.Commit(); err != nil {
					tx.Rollback()
					chainInfo.Height --
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("LoadCosmosCrossTxFromChain: commit tx", err)
					break
				}
			}
		case <-context.Context.Done():
			log.Info("stop monitoring cosmos Network")
			return
		}
	}
}

func (srv *Service) saveCosmosCrossTxsByHeight(tx *sql.Tx, chainInfo *model.ChainInfo) (uint32, uint32, error) {
	var in uint32 = 0
	var out uint32 = 0
	client := srv.cosmosClient
	height := int64(chainInfo.Height)
	block, err := client.Block(&height)
	if err != nil {
		log.Errorf("saveCosmosCrossTxsByHeight: Block %v", err)
		return in, out, err
	}
	tt := block.Block.Time.Unix()

	ccmLockEvent, lockEvents, err := srv.getCosmosCCMLockEventByBlockNumber(uint64(chainInfo.Height))
	if err != nil {
		return 0, 0, err
	}

	ccmUnlockEvent, unlockEvents, err := srv.getCosmosCCMUnlockEventByBlockNumber(uint64(chainInfo.Height))
	if err != nil {
		return 0, 0, err
	}

	// save lockEvent to db
	for _, lockEvent := range ccmLockEvent {
		if lockEvent.Method == _cosmos_crosschainlock {
			log.Infof("from chain: %s, txhash: %s\n", chainInfo.Name, lockEvent.TxHash)
			fctx := &model.FChainTx{}
			fctx.Chain = srv.c.Cosmos.ChainId
			fctx.TxHash = lockEvent.TxHash
			fctx.State = 1
			fctx.Fee = lockEvent.Fee
			fctx.TT = uint32(tt)
			fctx.Height = chainInfo.Height
			fctx.User = lockEvent.User
			fctx.TChain = lockEvent.Tchain
			fctx.Contract = lockEvent.Contract
			fctx.Key = lockEvent.Txid
			fctx.Param = hex.EncodeToString(lockEvent.Value)
			for _, v := range lockEvents {
				if v.TxHash == lockEvent.TxHash {
					fctransfer := &model.FChainTransfer{}
					fctransfer.TxHash = lockEvent.TxHash
					fctransfer.From = v.FromAddress
					fctransfer.To = srv.Hash2Address(common.CHAIN_COSMOS, lockEvent.Contract)
					fctransfer.Asset = v.FromAssetHash
					fctransfer.Amount = v.Amount
					fctransfer.ToChain = v.ToChainId
					fctransfer.ToAsset = v.ToAssetHash
					fctransfer.ToUser = srv.Hash2Address(v.ToChainId, v.ToAddress)
					fctx.Transfer = fctransfer
					break
				}
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
	for _, unLockEvent := range ccmUnlockEvent {
		if unLockEvent.Method == _cosmos_crosschainunlock {
			log.Infof("to chain: %s, txhash: %s\n", chainInfo.Name, unLockEvent.TxHash)
			tctx := &model.TChainTx{}
			tctx.Chain = srv.c.Cosmos.ChainId
			tctx.TxHash = unLockEvent.TxHash
			tctx.State = 1
			tctx.Fee = unLockEvent.Fee
			tctx.TT = uint32(tt)
			tctx.Height = chainInfo.Height
			tctx.FChain = unLockEvent.FChainId
			tctx.Contract = unLockEvent.Contract
			tctx.RTxHash = unLockEvent.RTxHash
			for _, v := range unlockEvents {
				if v.TxHash == unLockEvent.TxHash {
					tctransfer := &model.TChainTransfer{}
					tctransfer.TxHash = unLockEvent.TxHash
					tctransfer.From = srv.Hash2Address(common.CHAIN_COSMOS, unLockEvent.Contract)
					tctransfer.To = v.ToAddress
					tctransfer.Asset = v.ToAssetHash
					tctransfer.Amount = v.Amount
					tctx.Transfer = tctransfer
					break
				}
			}
			err = srv.dao.TxInsertTChainTxAndCache(tx, tctx)
			if err != nil {
				log.Errorf("saveOntCrossTxsByHeight: InsertTChainTx %s", err)
				return 0, 0, err
			}
			in++
		}
	}
	return in, out, nil
}

func (srv *Service) getCosmosCCMLockEventByBlockNumber(height uint64) ([]*model.ECCMLockEvent, []*model.LockEvent, error) {
	client := srv.cosmosClient
	ccmLockEvents := make([]*model.ECCMLockEvent, 0)
	lockEvents := make([]*model.LockEvent, 0)
	query := fmt.Sprintf("tx.height=%d AND make_from_cosmos_proof.status='1'", height)
	res, err := client.TxSearch(query, false, 1, 100, "asc")
	if err != nil {
		return ccmLockEvents, lockEvents, err
	}
	if res.TotalCount != 0 {
		pages := ((res.TotalCount - 1) / 100) + 1
		for p := 1; p <= pages; p ++ {
			if p > 1 {
				res, err = client.TxSearch(query, false, p, 100, "asc")
				if err != nil {
					return ccmLockEvents, lockEvents, err
				}
			}
			for _, tx := range res.Txs {
				for _, e := range tx.TxResult.Events {
					if e.Type == _cosmos_crosschainlock {
						tchainId, _ := strconv.ParseUint(string(e.Attributes[5].Value), 10, 32)
						value, _ := hex.DecodeString(string(e.Attributes[6].Value))
						ccmLockEvents = append(ccmLockEvents, &model.ECCMLockEvent{
							Method: _cosmos_crosschainlock,
							Txid: string(e.Attributes[1].Value),
							TxHash: strings.ToLower(tx.Hash.String()),
							User: string(e.Attributes[3].Value),
							Tchain: uint32(tchainId),
							Contract: string(e.Attributes[4].Value),
							Height: height,
							Value: value,
							Fee: client.GetGas(tx.Tx),
						})
					} else if e.Type == _cosmos_lock {
						tchainId, _ := strconv.ParseUint(string(e.Attributes[1].Value), 10, 32)
						amount, _ := strconv.ParseUint(string(e.Attributes[5].Value), 10, 64)
						lockEvents = append(lockEvents, &model.LockEvent{
							Method: _cosmos_lock,
							TxHash: strings.ToLower(tx.Hash.String()),
							FromAddress: string(e.Attributes[3].Value),
							FromAssetHash: string(e.Attributes[0].Value),
							ToChainId: uint32(tchainId),
							ToAssetHash: string(e.Attributes[2].Value),
							ToAddress: string(e.Attributes[4].Value),
							Amount: amount,
						})
					}
				}
			}
		}
	}

	return ccmLockEvents, lockEvents, nil
}

func (srv *Service) getCosmosCCMUnlockEventByBlockNumber(height uint64) ([]*model.ECCMUnlockEvent, []*model.UnlockEvent, error) {
	client := srv.cosmosClient
	ccmUnlockEvents := make([]*model.ECCMUnlockEvent, 0)
	unlockEvents := make([]*model.UnlockEvent, 0)
	query := fmt.Sprintf("tx.height=%d", height)
	res, err := client.TxSearch(query, false, 1, 100, "asc")
	if err != nil {
		return ccmUnlockEvents, unlockEvents, err
	}
	if res.TotalCount != 0 {
		pages := ((res.TotalCount - 1) / 100) + 1
		for p := 1; p <= pages; p ++ {
			if p > 1 {
				res, err = client.TxSearch(query, false, p, 100, "asc")
				if err != nil {
					return ccmUnlockEvents, unlockEvents, err
				}
			}
			for _, tx := range res.Txs {
				for _, e := range tx.TxResult.Events {
					if e.Type == _cosmos_crosschainunlock {
						fchainId, _ := strconv.ParseUint(string(e.Attributes[2].Value), 10, 32)
						ccmUnlockEvents = append(ccmUnlockEvents, &model.ECCMUnlockEvent{
							Method: _cosmos_crosschainunlock,
							TxHash: strings.ToLower(tx.Hash.String()),
							RTxHash: common.HexStringReverse(string(e.Attributes[0].Value)),
							FChainId: uint32(fchainId),
							Contract: string(e.Attributes[3].Value),
							Height: height,
							Fee: client.GetGas(tx.Tx),
						})
					} else if e.Type == _cosmos_unlock {
						amount, _ := strconv.ParseUint(string(e.Attributes[2].Value), 10, 64)
						unlockEvents = append(unlockEvents, &model.UnlockEvent{
							Method: _cosmos_unlock,
							TxHash: strings.ToLower(tx.Hash.String()),
							ToAssetHash: string(e.Attributes[0].Value),
							ToAddress: string(e.Attributes[1].Value),
							Amount: amount,
						})
					}
				}
			}
		}
	}

	return ccmUnlockEvents, unlockEvents, nil
}

/*
func (srv *Service) getCosmosCCMEventByBlockNumber(height uint64) ([]*model.ECCMLockEvent, []*model.ECCMUnlockEvent, error) {
	client := srv.cosmosClient.Client
	ccmLockEvents := make([]*model.ECCMLockEvent, 0)
	ccmUnlockEvents := make([]*model.ECCMUnlockEvent, 0)
	{
		query := fmt.Sprintf("tx.height=%d AND make_from_cosmos_proof.status='1'", height)
		res, err := client.TxSearch(query, false, 1, 100, "asc")
		if err != nil {
			return ccmLockEvents, ccmUnlockEvents, err
		}
		if res.TotalCount != 0 {
			pages := ((res.TotalCount - 1) / 100) + 1
			for p := 1; p <= pages; p ++ {
				if p > 1 {
					res, err = client.TxSearch(query, false, p, 100, "asc")
					if err != nil {
						return ccmLockEvents, ccmUnlockEvents, err
					}
				}
				for _, tx := range res.Txs {
					for _, e := range tx.TxResult.Events {
						if e.Type == _cosmos_crosschainlock {
							tchainId, _ := strconv.ParseUint(string(e.Attributes[5].Value), 10, 32)
							value, _ := hex.DecodeString(string(e.Attributes[6].Value))
							ccmLockEvents = append(ccmLockEvents, &model.ECCMLockEvent{
								Method: _cosmos_crosschainlock,
								Txid: string(e.Attributes[1].Value),
								TxHash: tx.Hash.String(),
								User: string(e.Attributes[3].Value),
								Tchain: uint32(tchainId),
								Contract: string(e.Attributes[4].Value),
								Height: height,
								Value: value,
							})
						}
					}
				}
			}
		}
	}
	{
		query := fmt.Sprintf("tx.height=%d AND verify_to_cosmos_proof.status='1'", height)
		res, err := client.TxSearch(query, false, 1, 100, "asc")
		if err != nil {
			return ccmLockEvents, ccmUnlockEvents, err
		}
		if res.TotalCount != 0 {
			pages := ((res.TotalCount - 1) / 100) + 1
			for p := 1; p <= pages; p ++ {
				if p > 1 {
					res, err = client.TxSearch(query, false, p, 100, "asc")
					if err != nil {
						return ccmLockEvents, ccmUnlockEvents, err
					}
				}
				for _, tx := range res.Txs {
					for _, e := range tx.TxResult.Events {
						if e.Type == _cosmos_crosschainunlock {
							fchainId, _ := strconv.ParseUint(string(e.Attributes[2].Value), 10, 32)
							ccmUnlockEvents = append(ccmUnlockEvents, &model.ECCMUnlockEvent{
								Method: _cosmos_crosschainunlock,
								TxHash: tx.Hash.String(),
								RTxHash: string(e.Attributes[0].Value),
								FChainId: uint32(fchainId),
								Contract: string(e.Attributes[3].Value),
								Height: height,
							})
						}
					}
				}
			}
		}
	}
	return ccmLockEvents, ccmUnlockEvents, nil
}

func (srv *Service) getCosmosProxyEventByBlockNumber(height uint64) ([]*model.LockEvent, []*model.UnlockEvent, error) {
	client := srv.cosmosClient.Client
	lockEvents := make([]*model.LockEvent, 0)
	unlockEvents := make([]*model.UnlockEvent, 0)
	{
		query := fmt.Sprintf("tx.height=%d AND lock.status='1'", height)
		res, err := client.TxSearch(query, false, 1, 100, "asc")
		if err != nil {
			return lockEvents, unlockEvents, err
		}
		if res.TotalCount != 0 {
			pages := ((res.TotalCount - 1) / 100) + 1
			for p := 1; p <= pages; p ++ {
				if p > 1 {
					res, err = client.TxSearch(query, false, p, 100, "asc")
					if err != nil {
						return lockEvents, unlockEvents, err
					}
				}
				for _, tx := range res.Txs {
					for _, e := range tx.TxResult.Events {
						if e.Type == _cosmos_lock {
							fromassethash, _ := hex.DecodeString(string(e.Attributes[0].Value))
							tchainId, _ := strconv.ParseUint(string(e.Attributes[1].Value), 10, 32)
							amount, _ := strconv.ParseUint(string(e.Attributes[5].Value), 10, 32)
							lockEvents = append(lockEvents, &model.LockEvent{
								Method: _cosmos_lock,
								TxHash: tx.Hash.String(),
								FromAddress: string(e.Attributes[3].Value),
								FromAssetHash: string(fromassethash),
								ToChainId: uint32(tchainId),
								ToAssetHash: string(e.Attributes[2].Value),
								ToAddress: string(e.Attributes[4].Value),
								Amount: amount,
							})
						}
					}
				}
			}
		}
	}
	{
		query := fmt.Sprintf("tx.height=%d AND unlock.status='1'", height)
		res, err := client.TxSearch(query, false, 1, 100, "asc")
		if err != nil {
			return lockEvents, unlockEvents, err
		}
		if res.TotalCount != 0 {
			pages := ((res.TotalCount - 1) / 100) + 1
			for p := 1; p <= pages; p ++ {
				if p > 1 {
					res, err = client.TxSearch(query, false, p, 100, "asc")
					if err != nil {
						return lockEvents, unlockEvents, err
					}
				}
				for _, tx := range res.Txs {
					for _, e := range tx.TxResult.Events {
						if e.Type == _cosmos_unlock {
							amount, _ := strconv.ParseUint(string(e.Attributes[2].Value), 10, 32)
							unlockEvents = append(unlockEvents, &model.UnlockEvent{
								Method: _cosmos_unlock,
								TxHash: tx.Hash.String(),
								ToAssetHash: string(e.Attributes[0].Value),
								ToAddress: string(e.Attributes[1].Value),
								Amount: amount,
							})
						}
					}
				}
			}
		}
	}
	return lockEvents, unlockEvents, nil
}
*/
