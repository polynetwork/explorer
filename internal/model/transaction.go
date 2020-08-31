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

package model

type ChainInfo struct {
	Id        uint32           `gorm:"column:id"`
	Name      string           `gorm:"column:xname"`
	XType     uint32           `gorm:"column:xtype"`
	Height    uint32           `gorm:"column:height"`
	In        uint32           `gorm:"column:txin"`
	Out       uint32           `gorm:"column:txout"`
	Contracts []*ChainContract `gorm:"ForeignKey:contracts"`
	Tokens    []*ChainToken    `gorm:"ForeignKey:tokens"`
}

type ChainContract struct {
	Id       uint32 `gorm:"column:id"`
	Contract string `gorm:"column:contract"`
}

type ChainToken struct {
	Id        int32  `json:"id" gorm:"column:id"`
	Token     string `json:"token" gorm:"column:xtoken"`
	Hash      string `json:"hash" gorm:"column:hash"`
	Name      string `json:"tokenname" gorm:"column:name"`
	Type      string `json:"tokentype" gorm:"column:xtype"`
	Precision uint64 `json:"precision" gorm:"column:xprecision"`
	Desc      string `json:"desc" gorm:"column:desc"`
}

type CrossChainToken struct {
	Name      string             `json:"name"`
	Tokens    []*ChainToken      `json:"tokens"`
}

type CrossChainTxStatus struct {
	Id        uint32    `json:"chainid"`
	TT        uint32    `json:"timestamp"`
	TxNumber  uint32    `json:"txnumber"`
}

type CrossChainAddressNum struct {
	Id        uint32    `json:"chainid"`
	AddNum    uint32    `json:"addressnumber"`
}

type FChainTx struct {
	Chain        uint32 `json:"chainid" gorm:"column:chain_id"`
	TxHash       string `json:"txhash" gorm:"column:txhash"`
	State        byte   `json:"state" gorm:"column:state"`
	TT           uint32 `json:"timestamp" gorm:"column:tt"`
	Fee          uint64 `json:"fee" gorm:"column:fee"`
	Height       uint32 `json:"blockheight" gorm:"column:height"`
	User         string `json:"user" gorm:"column:xuer"`
	TChain       uint32 `json:"tchainid" gorm:"column:tchain"`
	Contract     string `json:"contract" gorm:"column:contract"`
	Key          string `json:"key" gorm:"column:xkey"`
	Param        string `json:"value" gorm:"column:xparam"`
	Transfer     *FChainTransfer
}

type FChainTransfer struct {
	TxHash       string `json:"txhash" gorm:"column:txhash"`
	Asset        string `json:"asset" gorm:"column:asset"`
	From         string `json:"from" gorm:"column:xfrom"`
	To           string `json:"to" gorm:"column:xto"`
	Amount       uint64 `json:"amount" gorm:"column:amount"`
	ToChain      uint32 `json:"tochainid" gorm:"column:tochainid"`
	ToAsset      string `json:"toasset" gorm:"column:toasset"`
	ToUser       string `json:"touser" gorm:"column:touser"`
}

type MChainTx struct {
	Chain   uint32 `json:"chainid" gorm:"column:chain_id"`
	TxHash  string `json:"txhash" gorm:"column:txhash"`
	State   byte   `json:"state" gorm:"column:state"`
	TT      uint32 `json:"timestamp" gorm:"column:tt"`
	Fee     uint64 `json:"fee" gorm:"column:fee"`
	Height  uint32 `json:"blockheight" gorm:"column:height"`
	FChain  uint32 `json:"fchain" gorm:"column:fchain"`
	FTxHash string `json:"ftxhash" gorm:"column:ftxhash"`
	TChain  uint32 `json:"tchain" gorm:"column:tchain"`
	Key     string `json:"key" gorm:"column:xkey"`
}

type TChainTx struct {
	Chain        uint32 `json:"chainid" gorm:"column:chain_id"`
	TxHash       string `json:"txhash" gorm:"column:txhash"`
	State        byte   `json:"state" gorm:"column:state"`
	TT           uint32 `json:"timestamp" gorm:"column:tt"`
	Fee          uint64 `json:"fee" gorm:"column:fee"`
	Height       uint32 `json:"blockheight" gorm:"column:height"`
	FChain       uint32 `json:"fchain" gorm:"column:fchain"`
	Contract     string `json:"contract" gorm:"column:contract"`
	RTxHash      string `json:"rtxhash" gorm:"column:rtxhash"`
	Transfer     *TChainTransfer
}

type TChainTransfer struct {
	TxHash       string `json:"txhash" gorm:"column:txhash"`
	Asset        string `json:"asset" gorm:"column:asset"`
	From         string `json:"from" gorm:"column:xfrom"`
	To           string `json:"to" gorm:"column:xto"`
	Amount       uint64 `json:"amount" gorm:"column:amount"`
}

type TokenTx struct {
	TxHash       string `json:"txhash"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Amount       uint64  `json:"amount"`
	TT           uint32   `json:"timestamp"`
	Height       uint32  `json:"blockheight"`
	Direct       uint32  `json:"direct"`
}

type AddressTx struct {
	TxHash       string `json:"txhash"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Asset        string  `json:"asset"`
	Amount       uint64  `json:"amount"`
	TT           uint32   `json:"timestamp"`
	Height       uint32  `json:"blockheight"`
	Direct       uint32  `json:"direct"`
}

type AssetStatistic struct {
	Name         string    `json:"name"`
	Addressnum   uint32    `json:"addressnumber"`
	Amount       uint64    	`json:"amount"`
	Amount_btc   uint64    `json:"amount_btc"`
	Amount_usd   uint64    `json:"amount_usd"`
	TxNum        uint32    `json:"txnumber"`
	LatestUpdate uint32    `json:"latestupdate"`
}

type AssetAddressNum struct {
	Name      string    `json:"asset"`
	AddNum    uint32    `json:"addressnumber"`
}

type AssetTxInfo struct {
	Name      string    `json:"asset"`
	Amount    uint64    `json:"amount"`
	TxNum     uint32    `json:"txnumber"`
}
