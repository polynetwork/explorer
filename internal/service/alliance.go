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
	"encoding/json"
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	"github.com/polynetwork/poly/consensus/vbft/config"
	"github.com/polynetwork/poly/core/types"
	"time"
)

// loadAllianceCrossTxFromChain synchronizes cross txs from Alliance network
func (srv *Service) LoadAllianceCrossTxFromChain(context *ctx.Context) {
	sdk := srv.allianceClient
	chainInfo := srv.GetChain(srv.c.Alliance.ChainId)
	t := time.NewTicker(srv.c.Alliance.BlockDuration * time.Second)
	for {
		select {
		case <-t.C:
			currentHeight, err := sdk.GetCurrentBlockHeight()
			if err != nil {
				log.Errorf("LoadAllianceCrossTxFromChain: get current block height %s", err)
				continue
			}
			log.Infof("chain %s current height: %d, parser height: %d", chainInfo.Name, currentHeight, chainInfo.Height)
			for currentHeight > chainInfo.Height {
				tx, err := srv.dao.BeginTran()
				if err != nil {
					log.Errorf("loadAllianceCrossTxFromChain: beginTran", err)
					break
				}
				in, out, err := srv.saveAllianceCrossTxsByHeight(tx, chainInfo)
				if err != nil {
					tx.Rollback()
					log.Errorf("LoadAllianceCrossTxFromChain: can not saveOntCrossTxsByHeight %s", err)
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
					log.Errorf("LoadAllianceCrossTxFromChain: update ChainInfo %s", err)
					break
				}
				if err = tx.Commit(); err != nil {
					tx.Rollback()
					chainInfo.Height --
					chainInfo.In -= in
					chainInfo.Out -= out
					log.Errorf("LoadAllianceCrossTxFromChain: commit tx", err)
					break
				}
			}
		case <-context.Context.Done():
			log.Info("stop monitoring Alliance Network")
			return
		}
	}
}

func (srv *Service) saveAllianceCrossTxsByHeight(tx *sql.Tx, chainInfo *model.ChainInfo) (in uint32, out uint32, err error) {
	sdk := srv.allianceClient
	localHeight := chainInfo.Height
	in = 0
	out = 0
	block, err := sdk.GetBlockByHeight(localHeight)
	if err != nil {
		return
	}
	srv.parserValidator(block.Header)
	tt := block.Header.Timestamp
	events, err := sdk.GetSmartContractEventByBlock(chainInfo.Height)
	if err != nil {
		return
	}
	//log.Debugf("chain id: %d, block height: %d, events num: %d", chainInfo.Id, chainInfo.Height, len(events))
	for _, event := range events {
		for _, notify := range event.Notify {
			for _, contact := range chainInfo.Contracts {
				if notify.ContractAddress != contact.Contract {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				log.Infof("chain: %s, tx hash: %s, method: %s, state: %d, gas: %d\n", chainInfo.Name, event.TxHash, contractMethod, event.State, 0)
				if contractMethod != "makeProof" && contractMethod != "btcTxToRelay" {
					continue
				}
				if len(states) < 4 {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))
				if !srv.IsMonitorChain(fchainid){
					continue
				}
				if !srv.IsMonitorChain(tchainid) {
					continue
				}
				mctx := &model.MChainTx{}
				mctx.Chain = chainInfo.Id
				mctx.TxHash = event.TxHash
				mctx.State = event.State
				mctx.Fee = 0
				mctx.TT = tt
				mctx.Height = chainInfo.Height
				mctx.FChain = fchainid
				mctx.TChain = tchainid
				if tchainid == srv.c.Bitcoin.ChainId {
					if len(states) < 5 {
						continue
					}
					if fchainid == srv.c.Ethereum.ChainId || fchainid == srv.c.Cosmos.ChainId {
						mctx.FTxHash = states[4].(string)
						mctx.Key = states[3].(string)
					} else {
						mctx.FTxHash = common.HexStringReverse(states[4].(string))
						mctx.Key = states[3].(string)
					}
				} else {
					if fchainid == srv.c.Ethereum.ChainId || fchainid == srv.c.Cosmos.ChainId || fchainid == common.CHAIN_BSC || fchainid == common.CHAIN_HECO || fchainid == common.CHAIN_O3 || fchainid == common.CHAIN_OK {
						mctx.FTxHash = states[3].(string)
					} else {
						mctx.FTxHash = common.HexStringReverse(states[3].(string))
					}
				}
				err = srv.dao.TxInsertMChainTxAndCache(tx, mctx)
				if err != nil {
					log.Errorf("saveAllianceCrossTxsByHeight: InsertMChainTx %s", err)
					return 0, 0, err
				}
				in++
				out++
			}
		}
	}
	return in, out, nil
}

func (srv *Service) parserValidator(header *types.Header) {
	blkInfo := &vconfig.VbftBlockInfo{}
	if err := json.Unmarshal(header.ConsensusPayload, blkInfo); err != nil {
		log.Errorf("parserValidator - unmarshal blockInfo error: %s", err)
		return
	}
	if blkInfo.NewChainConfig == nil {
		return
	}
	var bookkeepers  []string
	for _, peer := range blkInfo.NewChainConfig.Peers {
		bookkeepers = append(bookkeepers, peer.ID)
	}
}
