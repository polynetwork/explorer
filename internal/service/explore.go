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
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	myerror "github.com/polynetwork/explorer/internal/server/restful/error"
	"math/big"
	"sort"
	"strings"
)

// GetExplorerInfo shows explorer information, such as current blockheight (the number of blockchain and so on) on the home page.
func (exp *Service) GetExplorerInfo(start uint32, end uint32) (int64, string) {
	log.Infof("GetExplorerInfo, start: %d, end: %d", start, end)
	// get all chains
	chainInfoResps := exp.outputChainInfos(exp.chain)
	crosschainTokens := exp.outputCrossChainTokens(exp.tokens)
	// get status
	allFchainStatus, err := exp.dao.SelectFChainTxByTime(start, end)
	if err != nil {
		log.Errorf("GetExplorerInfo: SelectFChainTxByTime %s", err)
		return myerror.DB_LOADDATA_FAILED, ""
	}
	allTchainStatus, err := exp.dao.SelectTChainTxByTime(start, end)
	if err != nil {
		log.Errorf("GetExplorerInfo: SelectTChainTxByTime %s", err)
		return myerror.DB_LOADDATA_FAILED, ""
	}
	allChainAddressNum, err := exp.dao.SelectChainAddressNum()
	if err != nil {
		log.Errorf("GetExplorerInfo: SelectChainAddresses %s", err)
		return myerror.DB_LOADDATA_FAILED, ""
	}
	for _, chainInfo := range chainInfoResps {
		fchainStatus := make([]*model.CrossChainTxStatus,0)
		for _, item := range allFchainStatus {
			if item.Id == chainInfo.Id {
				fchainStatus = append(fchainStatus, item)
			}
		}
		tchainStatus := make([]*model.CrossChainTxStatus, 0)
		for _, item := range allTchainStatus {
			if item.Id == chainInfo.Id {
				tchainStatus = append(tchainStatus, item)
			}
		}
		chainAddressNum := uint32(0)
		for _, item := range allChainAddressNum {
			if item.Id == chainInfo.Id {
				chainAddressNum = item.AddNum
			}
		}
		chainInfo.OutCrossChainTxStatus = exp.outputCrossChainTxStatus(fchainStatus, start, end)
		chainInfo.InCrossChainTxStatus = exp.outputCrossChainTxStatus(tchainStatus, start, end)
		chainInfo.Addresses = chainAddressNum
	}

	// get cross chain tx
	mChainInfo := exp.GetChain(common.CHAIN_POLY)
	if mChainInfo == nil {
		log.Errorf("Can't get muti chain info")
		return myerror.DB_LOADDATA_FAILED, ""
	}
	explorerInfoResp := &model.ExplorerInfoResp{
		Chains:        chainInfoResps,
		CrossTxNumber: mChainInfo.In,
		Tokens:   crosschainTokens,
	}
	explorerInfoJsonResp, _ := json.Marshal(explorerInfoResp)
	return myerror.SUCCESS, string(explorerInfoJsonResp)
}

func (exp *Service) GetTokenTxList(token string, start uint32, end uint32) (int64, string) {
	log.Infof("GetTokenTxList, token: %s", token)
	tokenTxList, err := exp.dao.SelectTokenTxList(token, start, end - start + 1)
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	tokenTxTotal, err := exp.dao.SelectTokenTxTotal(token)
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	tokenTxListResp := exp.outputTokenTxList(token, tokenTxList, *tokenTxTotal)
	tokenTxListJsonResp, _ := json.Marshal(tokenTxListResp)
	return myerror.SUCCESS, string(tokenTxListJsonResp)
}

func (exp *Service) GetAddressTxList(chainId uint32, addr string, start uint32, end uint32) (int64, string) {
	log.Infof("GetAddressTxList, chainid: %d, address: %s", chainId, addr)
	addressTxList, err := exp.dao.SelectAddressTxList(chainId, addr, start, end - start + 1)
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	addressTxTotal, err := exp.dao.SelectAddressTxTotal(chainId, addr)
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	addressTxListResp := exp.outputAddressTxList(addressTxList, *addressTxTotal)
	addressTxListJsonResp, _ := json.Marshal(addressTxListResp)
	return myerror.SUCCESS, string(addressTxListJsonResp)
}

// TODO GetCrossTxList gets Cross transaction list from start to end (to be optimized)
func (exp *Service) GetCrossTxList(start int, end int) (int64, string) {
	log.Infof("GetCrossTxList, start: %d, end: %d", start, end)
	mChainTxs, err := exp.dao.SelectMChainTxByLimit(start, end-start+1, common.CHAIN_ETH)
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	crossTxListResp := exp.outputCrossTxList(mChainTxs)
	crossTxsJsonResp, _ := json.Marshal(crossTxListResp)
	return myerror.SUCCESS, string(crossTxsJsonResp)
}

// GetCrossTx gets cross tx by Tx
func (exp *Service) GetCrossTx(hash string) (int64, string) {
	log.Infof("GetCrossTx: hash: %s", hash)
	crosstx := &model.CrossTxResp{
		Fchaintx_valid: false,
		Mchaintx_valid: false,
		Tchaintx_valid: false,
		Transfer: &model.CrossTransferResp{
			CrossTxType: 0,
		},
	}
	fChainTx := new(model.FChainTx)
	mChainTx := new(model.MChainTx)
	tChainTx := new(model.TChainTx)
	var err error
	log.Infof("*********************GetCrossTx phase 1************************")
	if fChainTx, err = exp.dao.FChainTx(hash, common.CHAIN_ETH); err != nil {
		log.Errorf("GetCrossTx: get fChainTx %s", err)
		return myerror.DB_CONNECTTION_FAILED, ""
	} else if fChainTx != nil {
		crosstx.Fchaintx_valid = true
		if mChainTx, err = exp.dao.MChainTxByFTx(fChainTx.TxHash); err != nil {
			log.Errorf("GetCrossTx: get mChainTx %s", err)
			return myerror.DB_CONNECTTION_FAILED, ""
		} else if mChainTx != nil {
			crosstx.Mchaintx_valid = true
			if tChainTx, err = exp.dao.TChainTxByMTx(mChainTx.TxHash); err != nil {
				log.Errorf("GetCrossTx: get tChainTx %s", err)
				return myerror.DB_CONNECTTION_FAILED, ""
			} else if tChainTx != nil && tChainTx.State == 1 {
				crosstx.Tchaintx_valid = true
			}
		}
	}

	log.Infof("*********************GetCrossTx phase 2************************")
	if !crosstx.Fchaintx_valid {
		if mChainTx, err = exp.dao.MChainTx(hash); err != nil {
			return myerror.DB_CONNECTTION_FAILED, ""
		} else if mChainTx != nil {
			crosstx.Mchaintx_valid = true
			if fChainTx, err = exp.dao.FChainTx(mChainTx.FTxHash, common.CHAIN_POLY); err != nil {
				log.Errorf("GetCrossTx: get fChainTx %s", err)
				return myerror.DB_CONNECTTION_FAILED, ""
			} else if fChainTx != nil {
				crosstx.Fchaintx_valid = true
				if tChainTx, err = exp.dao.TChainTxByMTx(mChainTx.TxHash); err != nil {
					log.Errorf("GetCrossTx: get tChainTx %s", err)
					return myerror.DB_CONNECTTION_FAILED, ""
				} else if tChainTx != nil && tChainTx.State == 1 {
					crosstx.Tchaintx_valid = true
				}
			}
		}
	}

	log.Infof("*********************GetCrossTx phase 3************************")
	if !(crosstx.Fchaintx_valid || crosstx.Mchaintx_valid) {
		if tChainTx, err = exp.dao.TChainTx(hash); err != nil {
			return myerror.DB_CONNECTTION_FAILED, ""
		} else if tChainTx != nil {
			crosstx.Tchaintx_valid = true
			if mChainTx, err = exp.dao.MChainTx(tChainTx.RTxHash); err != nil {
				log.Errorf("GetCrossTx: get mChainTx %s", err)
				return myerror.DB_CONNECTTION_FAILED, ""
			} else if mChainTx != nil {
				crosstx.Mchaintx_valid = true
				if fChainTx, err = exp.dao.FChainTx(mChainTx.FTxHash, common.CHAIN_POLY); err != nil {
					log.Errorf("GetCrossTx: get fChainTx %s", err)
					return myerror.DB_CONNECTTION_FAILED, ""
				} else if fChainTx != nil {
					if tChainTx.State == 0 {
						crosstx.Tchaintx_valid = false
					}
					crosstx.Fchaintx_valid = true
				}
			}
		}
	}

	log.Infof("*********************GetCrossTx phase 4************************")
	outputType := 0
	if crosstx.Fchaintx_valid {
		xx, _ := json.Marshal(fChainTx)
		log.Infof("f chain tx: %s", string(xx))
		crosstx.Transfer = exp.outputCrossTransfer(fChainTx.Chain, fChainTx.User, fChainTx.TT, fChainTx.Transfer)
		crosstx.Fchaintx = exp.outputFChainTx(fChainTx)
		outputType = 1
	}

	if crosstx.Fchaintx_valid && crosstx.Mchaintx_valid {
		xx, _ := json.Marshal(mChainTx)
		log.Infof("m chain tx: %s", string(xx))
		crosstx.Mchaintx = exp.outputMChainTx(mChainTx)
		outputType = 2
	}

	if crosstx.Fchaintx_valid && crosstx.Mchaintx_valid && crosstx.Tchaintx_valid {
		xx, _ := json.Marshal(tChainTx)
		log.Infof("t chain tx: %s", string(xx))
		crosstx.Tchaintx = exp.outputTChainTx(tChainTx)
		outputType = 3
	}

	if outputType == 0 {
		crosstx.Fchaintx_valid = false
		crosstx.Mchaintx_valid = false
		crosstx.Tchaintx_valid = false
	} else if outputType == 1 {
		crosstx.Mchaintx_valid = false
		crosstx.Tchaintx_valid = false
	} else if outputType == 2 {
		crosstx.Tchaintx_valid = false
	}
	crossTxJson, _ := json.Marshal(crosstx)
	return myerror.SUCCESS, string(crossTxJson)
}

func (exp *Service) outputChainInfos(chainInfos []*model.ChainInfo) []*model.ChainInfoResp {
	chainInfoResps := make([]*model.ChainInfoResp, 0)
	for _, chainInfo := range chainInfos {
		chainInfoResp := &model.ChainInfoResp{
			Id:        chainInfo.Id,
			Name:      chainInfo.Name,
			Height:    chainInfo.Height,
			In:        chainInfo.In,
			Out:       chainInfo.Out,
		}
		chainInfoResp.Tokens = exp.outputChainTokens(chainInfo.Tokens)
		chainInfoResp.Contracts = exp.outputChainContracts(chainInfo.Contracts)
		chainInfoResps = append(chainInfoResps, chainInfoResp)
	}
	return chainInfoResps
}

func (exp *Service) outputChainContracts(chainContracts []*model.ChainContract) []*model.ChainContractResp {
	chainContractResps := make([]*model.ChainContractResp, 0)
	for _, chainContract := range chainContracts {
		chainContractResp := &model.ChainContractResp{
			Id:        chainContract.Id,
			Contract:      chainContract.Contract,
		}
		chainContractResps = append(chainContractResps, chainContractResp)
	}
	return chainContractResps
}

func (exp *Service) outputChainTokens(chainTokens []*model.ChainToken) []*model.ChainTokenResp {
	chainTokenResps := make([]*model.ChainTokenResp, 0)
	for _, chainToken := range chainTokens {
		chainTokenResp := &model.ChainTokenResp{
			Chain:     chainToken.Id,
			ChainName: exp.ChainId2Name(uint32(chainToken.Id)),
			Hash:      chainToken.Hash,
			Token:     chainToken.Token,
			Name:      chainToken.Name,
			Type:      chainToken.Type,
			Precision: chainToken.Precision,
			Desc:      chainToken.Desc,
		}
		chainTokenResps = append(chainTokenResps, chainTokenResp)
	}
	return chainTokenResps
}

func (exp *Service) outputCrossChainTokens(crossChainTokens []*model.CrossChainToken) []*model.CrossChainTokenResp {
	crossChainTokenResps := make([]*model.CrossChainTokenResp, 0)
	for _, item := range crossChainTokens {
		crossChainTokenResp := &model.CrossChainTokenResp{
			Name: item.Name,
			Tokens: exp.outputChainTokens(item.Tokens),
		}
		crossChainTokenResps = append(crossChainTokenResps, crossChainTokenResp)
	}
	return crossChainTokenResps
}

func (exp *Service) outputCrossChainTxStatus(status []*model.CrossChainTxStatus, start uint32, end uint32) []*model.CrossChainTxStatusResp {
	status_new := make([]*model.CrossChainTxStatusResp, 0)
	current_day := exp.DayOfTime(start)
	end_day := exp.DayOfTime(end)
	for current_day <= end_day {
		status_new = append(status_new, &model.CrossChainTxStatusResp{
			TT:       current_day,
			TxNumber: 0,
		})
		current_day = exp.DayOfTimeAddOne(current_day)
	}
	i := 0
	j := 0
	for i < len(status) && j < len(status_new) {
		if status[i].TT == status_new[j].TT {
			status_new[j].TxNumber = status[i].TxNumber
			i ++
			j ++
		} else {
			j ++
		}
	}
	return status_new
}

func (exp *Service) outputCrossChainTxStatus1(status []*model.CrossChainTxStatus, start uint32, end uint32, total uint32) []*model.CrossChainTxStatus {
	if len(status) == 0 {
		return nil
	}
	current_txnumber := total
	current_tt := uint32(0)
	status_new := make([]*model.CrossChainTxStatus, 0)
	for _, s := range status {
		status_new = append(status_new, &model.CrossChainTxStatus{
			TT: s.TT,
			TxNumber: current_txnumber,
		})
		current_txnumber = current_txnumber - s.TxNumber
		current_tt = s.TT
	}
	status_new = append(status_new, &model.CrossChainTxStatus{
		TT: exp.DayOfTimeSubOne(current_tt),
		TxNumber: current_txnumber,
	})

	status_new1 := make([]*model.CrossChainTxStatus, 0)
	current_txnumber = status_new[0].TxNumber
	current_tt = status_new[0].TT
	for _, s := range status_new {
		for s.TT < current_tt {
			status_new1 = append(status_new1, &model.CrossChainTxStatus{
				TT:       current_tt,
				TxNumber: current_txnumber,
			})
			current_tt = exp.DayOfTimeSubOne(current_tt)
		}

		current_txnumber = s.TxNumber
		status_new1 = append(status_new1, &model.CrossChainTxStatus{
			TT:       current_tt,
			TxNumber: current_txnumber,
		})
		current_tt = exp.DayOfTimeSubOne(current_tt)
	}

	status_new = make([]*model.CrossChainTxStatus, 0)
	ss := status_new1[len(status_new1) - 1]
	current_txnumber = ss.TxNumber
	current_tt = exp.DayOfTime(start)
	for current_tt < ss.TT {
		status_new = append(status_new, &model.CrossChainTxStatus{
			TT:       current_tt,
			TxNumber: current_txnumber,
		})
		current_tt = exp.DayOfTimeAddOne(current_tt)
	}
	for i := 0;i < len(status_new1);i ++ {
		bb := status_new1[len(status_new1) - 1 - i]
		current_tt = bb.TT
		current_txnumber = bb.TxNumber
		if current_tt > exp.DayOfTimeUp(end) {
			break
		}
		status_new = append(status_new, bb)
	}
	for current_tt < exp.DayOfTimeUp(end) {
		status_new = append(status_new, &model.CrossChainTxStatus{
			TT:       current_tt,
			TxNumber: current_txnumber,
		})
		current_tt = exp.DayOfTimeAddOne(current_tt)
	}
	return status_new
}

func (exp *Service) outputCrossTransfer(chainid uint32, user string, tt uint32, transfer *model.FChainTransfer) *model.CrossTransferResp {
	if transfer == nil {
		return nil
	}
	crossTransfer := new(model.CrossTransferResp)
	crossTransfer.CrossTxType = 1
	crossTransfer.CrossTxName = exp.TxType2Name(crossTransfer.CrossTxType)
	crossTransfer.TT = tt
	crossTransfer.FromChainId = chainid
	crossTransfer.FromChain = exp.ChainId2Name(crossTransfer.FromChainId)
	crossTransfer.FromAddress = user
	crossTransfer.ToChainId = transfer.ToChain
	crossTransfer.ToChain = exp.ChainId2Name(crossTransfer.ToChainId)
	crossTransfer.ToAddress = transfer.ToUser
	token := exp.GetToken(transfer.Asset)
	if token != nil {
		crossTransfer.TokenHash = token.Hash
		crossTransfer.TokenName = token.Name
		crossTransfer.TokenType = token.Type
		crossTransfer.Amount = exp.FormatAmount(token.Precision, transfer.Amount)
	} else {
		crossTransfer.TokenName = transfer.Asset
		crossTransfer.Amount = transfer.Amount.String()
	}
	return crossTransfer
}

func (exp *Service) outputFChainTx(fChainTx *model.FChainTx) *model.FChainTxResp {
	fChainTxResp := &model.FChainTxResp{
		ChainId:    fChainTx.Chain,
		ChainName:  exp.ChainId2Name(fChainTx.Chain),
		TxHash:     fChainTx.TxHash,
		State:      fChainTx.State,
		TT:         fChainTx.TT,
		Fee:        exp.FormatFee(fChainTx.Chain,fChainTx.Fee),
		Height:     fChainTx.Height,
		User:       fChainTx.User,
		TChainId:   fChainTx.TChain,
		TChainName: exp.ChainId2Name(fChainTx.TChain),
		Contract:   fChainTx.Contract,
		Key:        fChainTx.Key,
		Param:      fChainTx.Param,
	}
	if fChainTx.Transfer != nil {
		fChainTxResp.Transfer = &model.FChainTransferResp{
			From:        fChainTx.Transfer.From,
			To:          fChainTx.Transfer.To,
			Amount:      fChainTx.Transfer.Amount.String(),
			ToChain:     fChainTx.Transfer.ToChain,
			ToChainName: exp.ChainId2Name(fChainTx.Transfer.ToChain),
			ToUser:      fChainTx.Transfer.ToUser,
		}
		token := exp.GetToken(fChainTx.Transfer.Asset)
		fChainTxResp.Transfer.TokenHash = fChainTx.Transfer.Asset
		if token != nil {
			fChainTxResp.Transfer.TokenHash = token.Hash
			fChainTxResp.Transfer.TokenName = token.Name
			fChainTxResp.Transfer.TokenType = token.Type
			fChainTxResp.Transfer.Amount = exp.FormatAmount(token.Precision, fChainTx.Transfer.Amount)
		} else {
			fChainTxResp.Transfer.TokenName = fChainTx.Transfer.Asset
			fChainTxResp.Transfer.Amount = fChainTx.Transfer.Amount.String()
		}
		totoken := exp.GetToken(fChainTx.Transfer.ToAsset)
		fChainTxResp.Transfer.ToTokenHash = fChainTx.Transfer.ToAsset
		if totoken != nil {
			fChainTxResp.Transfer.ToTokenHash = totoken.Hash
			fChainTxResp.Transfer.ToTokenName = totoken.Name
			fChainTxResp.Transfer.ToTokenType = totoken.Type
		} else {
			fChainTxResp.Transfer.ToTokenName = fChainTx.Transfer.ToAsset
		}
	}
	if fChainTx.Chain == common.CHAIN_ETH {
		fChainTxResp.TxHash = "0x" + fChainTx.Key
	} else if fChainTx.Chain == common.CHAIN_COSMOS {
		fChainTxResp.TxHash = strings.ToUpper(fChainTxResp.TxHash)
	}
	return fChainTxResp
}

func (exp *Service) outputMChainTx(mChainTx *model.MChainTx) *model.MChainTxResp{
	mChainTxResp := &model.MChainTxResp{
		ChainId:    mChainTx.Chain,
		ChainName:  exp.ChainId2Name(mChainTx.Chain),
		TxHash:     mChainTx.TxHash,
		State:      mChainTx.State,
		TT:         mChainTx.TT,
		Fee:        exp.FormatFee(mChainTx.Chain, mChainTx.Fee),
		Height:     mChainTx.Height,
		FChainId:   mChainTx.FChain,
		FChainName: exp.ChainId2Name(mChainTx.FChain),
		FTxHash:    mChainTx.FTxHash,
		TChainId:   mChainTx.TChain,
		TChainName: exp.ChainId2Name(mChainTx.TChain),
		Key:        mChainTx.Key,
	}
	return mChainTxResp
}

func (exp *Service) outputTChainTx(tChainTx *model.TChainTx) *model.TChainTxResp {
	tChainTxResp := &model.TChainTxResp{
		ChainId:    tChainTx.Chain,
		ChainName:  exp.ChainId2Name(tChainTx.Chain),
		TxHash:     tChainTx.TxHash,
		State:      tChainTx.State,
		TT:         tChainTx.TT,
		Fee:        exp.FormatFee(tChainTx.Chain, tChainTx.Fee),
		Height:     tChainTx.Height,
		FChainId:   tChainTx.FChain,
		FChainName: exp.ChainId2Name(tChainTx.FChain),
		Contract:   tChainTx.Contract,
		RTxHash:    tChainTx.RTxHash,
	}
	if tChainTx.Transfer != nil {
		tChainTxResp.Transfer = &model.TChainTransferResp{
			From:         tChainTx.Transfer.From,
			To:           tChainTx.Transfer.To,
			Amount:       tChainTx.Transfer.Amount.String(),
		}
		token := exp.GetToken(tChainTx.Transfer.Asset)
		tChainTxResp.Transfer.TokenHash = tChainTx.Transfer.Asset
		if token != nil {
			tChainTxResp.Transfer.TokenHash = token.Hash
			tChainTxResp.Transfer.TokenName = token.Name
			tChainTxResp.Transfer.TokenType = token.Type
			tChainTxResp.Transfer.Amount = exp.FormatAmount(token.Precision, tChainTx.Transfer.Amount)
		} else {
			tChainTxResp.Transfer.TokenName = tChainTx.Transfer.Asset
			tChainTxResp.Transfer.Amount = tChainTx.Transfer.Amount.String()
		}
	}
	if tChainTx.Chain == common.CHAIN_ETH {
		tChainTxResp.TxHash = "0x" + tChainTxResp.TxHash
	} else if tChainTx.Chain == common.CHAIN_COSMOS {
		tChainTxResp.TxHash = strings.ToUpper(tChainTxResp.TxHash)
	}
	return tChainTxResp
}

func (exp *Service) outputCrossTxList(crossTxs []*model.MChainTx) *model.CrossTxListResp {
	var crossTxListResp model.CrossTxListResp
	crossTxListResp.CrossTxList = make([]*model.CrossTxOutlineResp, 0)
	for _, mChainTx := range crossTxs {
		crossTxListResp.CrossTxList = append(crossTxListResp.CrossTxList, &model.CrossTxOutlineResp{
			TxHash: mChainTx.TxHash,
			State:  mChainTx.State,
			TT:     mChainTx.TT,
			Fee:    mChainTx.Fee,
			Height: mChainTx.Height,
			FChainId: mChainTx.FChain,
			FChainName: exp.ChainId2Name(mChainTx.FChain),
			TChainId: mChainTx.TChain,
			TChainName: exp.ChainId2Name(mChainTx.TChain),
		})
	}
	return &crossTxListResp
}

func (exp *Service) outputTokenTxList(tokenHash string, tokenTxs []*model.TokenTx, tokenTxTotal uint32) *model.TokenTxListResp {
	var tokenTxListResp model.TokenTxListResp
	tokenTxListResp.Total = tokenTxTotal
	tokenTxListResp.TokenTxList = make([]*model.TokenTxResp, 0)
	token := exp.GetToken(tokenHash)
	for _, tokenTx := range tokenTxs {
		amount := tokenTx.Amount.String()
		if token != nil {
			amount = exp.FormatAmount(token.Precision, tokenTx.Amount)
		}
		tokenTxListResp.TokenTxList = append(tokenTxListResp.TokenTxList, &model.TokenTxResp{
			TxHash: tokenTx.TxHash,
			From: tokenTx.From,
			To: tokenTx.To,
			Amount: amount,
			Height: tokenTx.Height,
			TT: tokenTx.TT,
			Direct: tokenTx.Direct,
		})
	}
	return &tokenTxListResp
}

func (exp *Service) outputAddressTxList(addressTxs []*model.AddressTx, addressTxTotal uint32) *model.AddressTxListResp {
	var addressTxListResp model.AddressTxListResp
	addressTxListResp.Total = addressTxTotal
	addressTxListResp.AddressTxList = make([]*model.AddressTxResp, 0)
	for _, addressTx := range addressTxs {
		txresp := &model.AddressTxResp{
			TxHash: addressTx.TxHash,
			From: addressTx.From,
			To: addressTx.To,
			Amount: addressTx.Amount.String(),
			Height: addressTx.Height,
			TT: addressTx.TT,
			Direct: addressTx.Direct,
			TokenHash: addressTx.Asset,
		}
		token := exp.GetToken(addressTx.Asset)
		if token != nil {
			txresp.Amount = exp.FormatAmount(token.Precision, addressTx.Amount)
			txresp.TokenName = token.Name
			txresp.TokenType = token.Type
		}
		addressTxListResp.AddressTxList = append(addressTxListResp.AddressTxList, txresp)
	}
	return &addressTxListResp
}

func (exp *Service) outputAssetInfo(assetStatistics []*model.AssetStatistic) *model.AssetInfoResp {
	assetInfo := new(model.AssetInfoResp)
	amountBtcTotal := new(big.Int)
	amountUsdTotal := new(big.Int)
	addressNumberTotal := uint32(0)
	txNumTotal := uint32(0)
	for _, assetStatistic := range assetStatistics {
		amountBtcTotal.Add(amountBtcTotal, assetStatistic.Amount_btc)
		amountUsdTotal.Add(amountUsdTotal, assetStatistic.Amount_usd)
		addressNumberTotal += assetStatistic.Addressnum
		txNumTotal += assetStatistic.TxNum
	}

	assetInfo.AmountBtcTotal = exp.FormatAmount(uint64(10000), amountBtcTotal)
	assetInfo.AmountUsdTotal = exp.FormatAmount(uint64(10000), amountUsdTotal)

	for _, assetStatistic := range assetStatistics {
		assetStatisticResp := &model.AssetStatisticResp{
			Name: assetStatistic.Name,
			AddressNum: assetStatistic.Addressnum,
			AddressNumPrecent: exp.Precent(uint64(assetStatistic.Addressnum), uint64(addressNumberTotal)),
			Amount: exp.FormatAmount(uint64(100), assetStatistic.Amount),
			AmountBtc: exp.FormatAmount(uint64(10000), assetStatistic.Amount_btc),
			AmountBtcPrecent: exp.Precent(assetStatistic.Amount_btc.Uint64(), amountBtcTotal.Uint64()),
			AmountUsd: exp.FormatAmount(uint64(10000), assetStatistic.Amount_usd),
			AmountUsdPrecent: exp.Precent(assetStatistic.Amount_usd.Uint64(), amountUsdTotal.Uint64()),
			TxNum: assetStatistic.TxNum,
			TxNumPrecent: exp.Precent(uint64(assetStatistic.TxNum), uint64(txNumTotal)),
			LatestUpdate: assetStatistic.LatestUpdate,
		}
		assetInfo.AssetStatistics = append(assetInfo.AssetStatistics, assetStatisticResp)
	}
	return assetInfo
}

func (exp *Service) outputTransferStatistic(transferStatistics []*model.AllTransferStatistic) *model.AllTransferStatisticResp {
	allTransferStatistic := new(model.AllTransferStatisticResp)
	allTransferStatistic.ChainTransferStatistics = make([]*model.ChainTransferStatisticResp, 0)
	for _, transferStatistic := range transferStatistics {
		var chainStatistic *model.ChainTransferStatisticResp
		chainStatistic = nil
		for _, item1 := range allTransferStatistic.ChainTransferStatistics {
			if item1.Chain == transferStatistic.Chain {
				chainStatistic = item1
				break
			}
		}
		if chainStatistic == nil {
			chainStatistic = new(model.ChainTransferStatisticResp)
			chainStatistic.Chain = transferStatistic.Chain
			chainStatistic.ChainName = exp.ChainId2Name(transferStatistic.Chain)
			chainStatistic.AssetTransferStatistics = make([]*model.AssetTransferStatisticResp, 0)
			allTransferStatistic.ChainTransferStatistics = append(allTransferStatistic.ChainTransferStatistics, chainStatistic)
		}
		var assetStatistic *model.AssetTransferStatisticResp
		assetStatistic = nil
		for _, item2 := range chainStatistic.AssetTransferStatistics {
			if item2.Name == transferStatistic.Name {
				assetStatistic = item2
				break
			}
		}
		if assetStatistic == nil {
			assetStatistic = new(model.AssetTransferStatisticResp)
			assetStatistic.Amount1 = transferStatistic.Amount
			assetStatistic.Amount = exp.FormatAmount(uint64(100), transferStatistic.Amount)
			assetStatistic.Name = transferStatistic.Name
			assetStatistic.Token = transferStatistic.Token
			assetStatistic.Hash = transferStatistic.Hash
			assetStatistic.SourceName = transferStatistic.SourceName
			assetStatistic.SourceChain = transferStatistic.SourceChain
			assetStatistic.SourceChainName = exp.ChainId2Name(transferStatistic.SourceChain)
			chainStatistic.AssetTransferStatistics = append(chainStatistic.AssetTransferStatistics, assetStatistic)
		}
	}
	coinPrices := exp.coinPrice
	bitcoinPrice, ok := coinPrices["Bitcoin"]
	if !ok {
		log.Errorf("There is no coin Bitcoin!")
		return allTransferStatistic
	}
	for _, item1 := range allTransferStatistic.ChainTransferStatistics {
		other_amount_btc_total := big.NewInt(0)
		other_amount_usd_total := big.NewInt(0)
		for _, item2 := range item1.AssetTransferStatistics {
			amount_btc := big.NewInt(0)
			amount_usd := big.NewInt(0)
			if item2.Name == common.UNISWAP_NAME {
				amount_btc = exp.updateUniswap(item2.Amount1)
				item2.AmountBtc = exp.FormatAmount(uint64(10000), amount_btc)
				amount_usd = new(big.Int).Mul(amount_btc, big.NewInt(int64(bitcoinPrice)))
				item2.AmountUsd  = exp.FormatAmount(uint64(10000), amount_usd)
				item2.AmountUsd1 = amount_usd
			} else {
				coinPrice, ok := coinPrices[item2.Token]
				if !ok {
					log.Warnf("There is no coin %s!", item2.Name)
					item2.AmountUsd = exp.FormatAmount(uint64(10000), big.NewInt(0))
					item2.AmountBtc = exp.FormatAmount(uint64(10000), big.NewInt(0))
					item2.AmountUsd1 = big.NewInt(0)
				} else {
					amount_usd = new(big.Int).Mul(item2.Amount1, big.NewInt(int64(coinPrice*100)))
					item2.AmountUsd = exp.FormatAmount(uint64(10000), amount_usd)
					amount_btc = new(big.Int).Div(amount_usd, big.NewInt(int64(bitcoinPrice)))
					item2.AmountBtc = exp.FormatAmount(uint64(10000), amount_btc)
					item2.AmountUsd1 = amount_usd
				}
			}
			if amount_usd.Cmp(big.NewInt(0)) == 1 {
				other_amount_btc_total = new(big.Int).Add(other_amount_btc_total, amount_btc)
				other_amount_usd_total = new(big.Int).Add(other_amount_usd_total, amount_usd)
			}
		}
		for _, item2 := range item1.AssetTransferStatistics {
			if item2.AmountUsd1.Cmp(big.NewInt(0)) == 1 {
				item2.AmountUsdPrecent = exp.Precent(item2.AmountUsd1.Uint64(), other_amount_usd_total.Uint64())
			}
		}
		item1.AmountUsd1 = other_amount_usd_total
		item1.AmountBtc = exp.FormatAmount(uint64(10000), other_amount_btc_total)
		item1.AmountUsd = exp.FormatAmount(uint64(10000), other_amount_usd_total)
	}
	sort.Slice(allTransferStatistic.ChainTransferStatistics, func(i, j int) bool {
		return allTransferStatistic.ChainTransferStatistics[i].AmountUsd1.Cmp(allTransferStatistic.ChainTransferStatistics[j].AmountUsd1) == 1
	})
	for _, item1 := range allTransferStatistic.ChainTransferStatistics {
		sort.Slice(item1.AssetTransferStatistics, func(i, j int) bool {
			return item1.AssetTransferStatistics[i].AmountUsd1.Cmp(item1.AssetTransferStatistics[j].AmountUsd1) == 1
		})
	}
	return allTransferStatistic
}

// GetCrossTx gets cross tx by Tx
func (exp *Service) GetLatestValidator() (int64, string) {
	log.Infof("GetLatestValidator")
	validators, err := exp.dao.SelectPolyValidator()
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	validatorsJson, _ := json.Marshal(validators)
	return myerror.SUCCESS, string(validatorsJson)
}

// GetCrossTx gets cross tx by Tx
func (exp *Service) GetAssetStatistic() (int64, string) {
	log.Infof("GetAssetStatistic")
	assetStatistics, err := exp.dao.SelectAssetStatistic(math.MaxUint32)
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	assetInfo := exp.outputAssetInfo(assetStatistics)
	assetInfoJson, _ := json.Marshal(assetInfo)
	return myerror.SUCCESS, string(assetInfoJson)
}

func (exp *Service) GetTransferStatistic(chainid int) (int64, string) {
	log.Infof("GetTransferStatistic")
	transferStatistics, err := exp.dao.SelectAllTransferStatistic(chainid)
	if err != nil {
		return myerror.DB_CONNECTTION_FAILED, ""
	}
	transferInfo := exp.outputTransferStatistic(transferStatistics)
	transferInfoJson, _ := json.Marshal(transferInfo)
	return myerror.SUCCESS, string(transferInfoJson)
}

