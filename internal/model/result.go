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

// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta

package model

type ExplorerInfoReq struct {
	Start        string    `json:"start"`
	End          string    `json:"end"`
}

// swagger:parameters ExplorerInfoRequest
type ExplorerInfoRequest struct {
	// in: body
	Body ExplorerInfoReq
}

type ExplorerInfoResp struct {
	Chains        []*ChainInfoResp `json:"chains"`
	CrossTxNumber uint32           `json:"crosstxnumber"`
	Tokens        []*CrossChainTokenResp `json:"tokens"`
}

// getexplorerinfo response
// swagger:response ExplorerInfoResponse
type ExplorerInfoResponse struct {
	// response body
	// in: body
	Body struct {
		Code          int                    `json:"code"`
		Action        string                 `json:"action"`
		Desc          string                 `json:"desc"`
		Version       string                 `json:"version"`
		Result        ExplorerInfoResp       `json:"result"`
	}
}

type ChainInfoResp struct {
	Id        uint32               `json:"chainid"`
	Name      string               `json:"chainname"`
	Height    uint32               `json:"blockheight"`
	In        uint32               `json:"in"`
	InCrossChainTxStatus []*CrossChainTxStatusResp    `json:"incrosschaintxstatus"`
	Out       uint32               `json:"out"`
	OutCrossChainTxStatus []*CrossChainTxStatusResp    `json:"outcrosschaintxstatus"`
	Addresses uint32               `json:"addresses"`
	Contracts []*ChainContractResp `json:"contracts"`
	Tokens    []*ChainTokenResp    `json:"tokens"`
}

type CrossChainTxStatusResp struct {
	TT        uint32    `json:"timestamp"`
	TxNumber  uint32    `json:"txnumber"`
}

type ChainContractResp struct {
	Id       uint32 `json:"chainid"`
	Contract string `json:"contract"`
}

type ChainTokenResp struct {
	Chain       int32  `json:"chainid"`
	ChainName   string    `json:"chainname"`
	Hash        string `json:"hash"`
	Token       string  `json:"token"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Precision   uint64 `json:"precision"`
	Desc        string `json:"desc"`
}

type CrossChainTokenResp struct {
	Name      string             `json:"name"`
	Tokens    []*ChainTokenResp  `json:"tokens"`
}

type FChainTxResp struct {
	ChainId    uint32    `json:"chainid"`
	ChainName  string    `json:"chainname"`
	TxHash     string    `json:"txhash"`
	State      byte      `json:"state"`
	TT         uint32    `json:"timestamp"`
	Fee        string    `json:"fee"`
	Height     uint32    `json:"blockheight"`
	User       string    `json:"user"`
	TChainId   uint32    `json:"tchainid"`
	TChainName string    `json:"tchainname"`
	Contract   string    `json:"contract"`
	Key        string    `json:"key"`
	Param      string    `json:"param"`
	Transfer   *FChainTransferResp `json:"transfer"`
}

type FChainTransferResp struct {
	TokenHash    string `json:"tokenhash"`
	TokenName    string `json:"tokenname"`
	TokenType    string `json:"tokentype"`
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	ToChain      uint32 `json:"tchainid"`
	ToChainName  string `json:"tchainname"`
	ToTokenHash  string `json:"totokenhash"`
	ToTokenName  string `json:"totokenname"`
	ToTokenType  string `json:"totokentype"`
	ToUser       string `json:"tuser"`
}

type MChainTxResp struct {
	ChainId    uint32 `json:"chainid"`
	ChainName  string `json:"chainname"`
	TxHash     string `json:"txhash"`
	State      byte   `json:"state"`
	TT         uint32 `json:"timestamp"`
	Fee        string `json:"fee"`
	Height     uint32 `json:"blockheight"`
	FChainId   uint32 `json:"fchainid"`
	FChainName string `json:"fchainname"`
	FTxHash    string `json:"ftxhash"`
	TChainId   uint32 `json:"tchainid"`
	TChainName string `json:"tchainname"`
	Key        string `json:"key"`
}

type TChainTxResp struct {
	ChainId    uint32    `json:"chainid"`
	ChainName  string    `json:"chainname"`
	TxHash     string    `json:"txhash"`
	State      byte      `json:"state"`
	TT         uint32    `json:"timestamp"`
	Fee        string    `json:"fee"`
	Height     uint32    `json:"blockheight"`
	FChainId   uint32    `json:"fchainid"`
	FChainName string    `json:"fchainname"`
	Contract   string    `json:"contract"`
	RTxHash    string    `json:"mtxhash"`
	Transfer   *TChainTransferResp `json:"transfer"`
}

type TChainTransferResp struct {
	TokenHash    string `json:"tokenhash"`
	TokenName    string `json:"tokenname"`
	TokenType    string `json:"tokentype"`
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
}

type CrossTransferResp struct {
	CrossTxType  uint32 `json:"crosstxtype"`
	CrossTxName  string `json:"crosstxname"`
	TT           uint32    `json:"timestamp"`
	FromChainId  uint32 `json:"fromchainid"`
	FromChain    string `json:"fromchainname"`
	FromAddress  string `json:"fromaddress"`
	ToChainId    uint32 `json:"tochainid"`
	ToChain      string `json:"tochainname"`
	ToAddress    string `json:"toaddress"`
	TokenHash    string `json:"tokenhash"`
	TokenName    string `json:"tokenname"`
	TokenType    string `json:"tokentype"`
	Amount       string `json:"amount"`
}

// swagger:parameters CrossTxReq
type CrossTxReq struct {
	// in: query
	TxHash    string       `json:"txhash"`
}

type CrossTxResp struct {
	Transfer       *CrossTransferResp `json:"crosstransfer"`
	Fchaintx       *FChainTxResp      `json:"fchaintx"`
	Fchaintx_valid bool               `json:"fchaintx_valid"`
	Mchaintx       *MChainTxResp      `json:"mchaintx"`
	Mchaintx_valid bool               `json:"mchaintx_valid"`
	Tchaintx       *TChainTxResp      `json:"tchaintx"`
	Tchaintx_valid bool               `json:"tchaintx_valid"`
}

// getcrosstx response
// swagger:response CrossTxResponse
type CrossTxResponse struct {
	// response body
	// in: body
	Body struct {
		Code          int                    `json:"code"`
		Action        string                 `json:"action"`
		Desc          string                 `json:"desc"`
		Version       string                 `json:"version"`
		Result        CrossTxResp            `json:"result"`
	}
}

type CrossTxListReq struct {
	Start        string    `json:"start"`
	End          string    `json:"end"`
}

// swagger:parameters CrossTxListRequest
type CrossTxListRequest struct {
	// in: body
	Body CrossTxListReq
}

type CrossTxOutlineResp struct {
	TxHash     string        `json:"txhash"`
	State      byte          `json:"state"`
	TT         uint32        `json:"timestamp"`
	Fee        uint64        `json:"fee"`
	Height     uint32        `json:"blockheight"`
	FChainId   uint32        `json:"fchainid"`
	FChainName string        `json:"fchainname"`
	TChainId   uint32        `json:"tchainid"`
	TChainName string        `json:"tchainname"`
}

type CrossTxListResp struct {
	CrossTxList       []*CrossTxOutlineResp     `json:"crosstxs"`
}

// getcrosstxlist response
// swagger:response CrossTxListResponse
type CrossTxListResponse struct {
	// response body
	// in: body
	Body struct {
		Code          int                    `json:"code"`
		Action        string                 `json:"action"`
		Desc          string                 `json:"desc"`
		Version       string                 `json:"version"`
		Result        CrossTxListResp         `json:"result"`
	}
}

type TokenTxListReq struct {
	Token       string     `json:"token"`
}

// swagger:parameters TokenTxListRequest
type TokenTxListRequest struct {
	// in: body
	Body TokenTxListReq
}


type TokenTxResp struct {
	TxHash       string `json:"txhash"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Amount       string  `json:"amount"`
	TT           uint32   `json:"timestamp"`
	Height       uint32  `json:"blockheight"`
	Direct       uint32  `json:"direct"`
}

type TokenTxListResp struct {
	TokenTxList       []*TokenTxResp     `json:"tokentxs"`
	Total             uint32             `json:"total"`
}


// gettokentxlist response
// swagger:response TokenTxListResponse
type TokenTxListResponse struct {
	// response body
	// in: body
	Body struct {
		Code          int                    `json:"code"`
		Action        string                 `json:"action"`
		Desc          string                 `json:"desc"`
		Version       string                 `json:"version"`
		Result        TokenTxListResp        `json:"result"`
	}
}


type AddressTxListReq struct {
	Address       string     `json:"address"`
	Chain         string     `json:"chain"`
}

// swagger:parameters AddressTxListRequest
type AddressTxListRequest struct {
	// in: body
	Body AddressTxListReq
}


type AddressTxResp struct {
	TxHash       string `json:"txhash"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Amount       string  `json:"amount"`
	TT           uint32   `json:"timestamp"`
	Height       uint32  `json:"blockheight"`
	TokenHash    string  `json:"tokenhash"`
	TokenName    string  `json:"tokenname"`
	TokenType    string  `json:"tokentype"`
	Direct       uint32  `json:"direct"`
}

type AddressTxListResp struct {
	AddressTxList       []*AddressTxResp     `json:"addresstxs"`
	Total               uint32               `json:"total"`
}


// getaddresstxlist response
// swagger:response AddressTxListResponse
type AddressTxListResponse struct {
	// response body
	// in: body
	Body struct {
		Code          int                    `json:"code"`
		Action        string                 `json:"action"`
		Desc          string                 `json:"desc"`
		Version       string                 `json:"version"`
		Result        AddressTxListResp      `json:"result"`
	}
}

type AssetStatisticResp struct {
	Name         string    `json:"name"`
	Addressnum   uint32    `json:"addressnumber"`
	AddressnumPrecent string   `json:"addressnumber_precent"`
	Amount       string    	`json:"amount"`
	Amount_btc   string    `json:"amount_btc"`
	AmountBtcPrecent string   `json:"amount_btc_precent"`
	Amount_usd   string    `json:"amount_usd"`
	AmountUsdPrecent string   `json:"Amount_usd_precent"`
	TxNum        uint32    `json:"txnumber"`
	TxNumPrecent string    `json:"txnumber_precent"`
	LatestUpdate uint32    `json:"latestupdate"`
}

type AssetInfoResp struct {
	AmountBtcTotal  string   `json:"amount_btc_total"`
	AmountUsdTotal  string   `json:"amount_usd_total"`
	AssetStatistics  []*AssetStatisticResp  `json:"asset_statistics"`
}

