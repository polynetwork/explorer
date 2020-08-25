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

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/wire"
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/log"
)

type BTCTools struct {
	restclient *RestClient
	urls     []string
	users    []string
	passwds  []string
	node     int
}

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type Response struct {
	Result json.RawMessage   `json:"result"`
	Error  *btcjson.RPCError `json:"error"` //maybe wrong
	Id     int               `json:"id"`
}

type BlockHeader struct {
	Hash   string `json:"hash"`
	Time   uint32 `json:"time"`
	Height uint32 `json:"height"`
}

func NewBtcTools(c *conf.Config) *BTCTools {
	restclient := NewRestClient()
	restclient.SetAddr(c.Bitcoin.Rawurl[0])
	restclient.SetAuth(c.Bitcoin.User[0], c.Bitcoin.Passwd[0])
	tool := &BTCTools{
		restclient: restclient,
		urls: c.Bitcoin.Rawurl,
		users: c.Bitcoin.User,
		passwds: c.Bitcoin.Passwd,
		node: 0,
	}
	return tool
}

func (self *BTCTools) GetCurrentHeight() (uint32, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblockcount",
		Params:  nil,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - failed to marshal request: %v", err)
	}

	cur := self.node
	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	for err != nil {
		log.Errorf("BTCTools.GetCurrentHeight err:%s, url: %s", err.Error(), self.urls[self.node])
		self.node ++
		self.node = self.node % len(self.urls)
		if self.node == cur {
			break
		}
		self.restclient.SetAddr(self.urls[self.node])
		self.restclient.SetAuth(self.users[self.node], self.passwds[self.node])
		repdata, err = self.restclient.SendRestRequestWithAuth(reqdata)
	}

	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - failed to unmarshal response: %v", err)
	}
	if rep.Error != nil {
		return 0, fmt.Errorf("GetCurrentHeight - response shows failure: %v", rep.Error.Message)
	}
	var blockCount uint32
	err = json.Unmarshal(rep.Result, &blockCount)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - failed to parse height: %v", err)
	}
	return blockCount, nil
}

func (self *BTCTools) GetBlockHash(height uint32) (string, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblockhash",
		Params:  []interface{}{height},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("GetBlockHash - failed to marshal request: %v", err)
	}

	cur := self.node
	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	for err != nil {
		log.Errorf("BTCTools.GetBlockHash err:%s, url: %s", err.Error(), self.urls[self.node])
		self.node ++
		self.node = self.node % len(self.urls)
		if self.node == cur {
			break
		}
		self.restclient.SetAddr(self.urls[self.node])
		self.restclient.SetAuth(self.users[self.node], self.passwds[self.node])
		repdata, err = self.restclient.SendRestRequestWithAuth(reqdata)
	}

	if err != nil {
		return "", fmt.Errorf("GetBlockHash - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return "", fmt.Errorf("GetBlockHash - failed to unmarshal response: %v", err)
	}

	if rep.Error != nil {
		return "", fmt.Errorf("GetBlockHash - response shows failure: %v", rep.Error.Message)
	}
	var hash string
	err = json.Unmarshal(rep.Result, &hash)
	if err != nil {
		return "", fmt.Errorf("GetCurrentHeight - failed to parse height: %v", err)
	}
	return hash, nil
}

func (self *BTCTools) GetBlockHeader(hash string) (*wire.BlockHeader, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblockheader",
		Params:  []interface{}{hash},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetBlockHeader - failed to marshal request: %v", err)
	}

	cur := self.node
	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	for err != nil {
		log.Errorf("BTCTools.GetBlockHeader err:%s, url: %s", err.Error(), self.urls[self.node])
		self.node ++
		self.node = self.node % len(self.urls)
		if self.node == cur {
			break
		}
		self.restclient.SetAddr(self.urls[self.node])
		self.restclient.SetAuth(self.users[self.node], self.passwds[self.node])
		repdata, err = self.restclient.SendRestRequestWithAuth(reqdata)
	}

	if err != nil {
		return nil, fmt.Errorf("GetBlockHeader - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetBlockHeader - failed to unmarshal response: %v", err)
	}

	if rep.Error != nil {
		return nil, fmt.Errorf("GetBlockHeader - response shows failure: %v", rep.Error.Message)
	}

	blockheader := wire.BlockHeader{}
	err = blockheader.BtcDecode(bytes.NewBuffer(rep.Result), wire.ProtocolVersion, wire.LatestEncoding)
	//err = json.Unmarshal(rep.Result, &blockheader)
	if err != nil {
		return nil, fmt.Errorf("GetCurrentHeight - failed to parse height: %v", err)
	}
	return &blockheader, nil
}

func (self *BTCTools) GetTx(hash string) (*btcjson.TxRawResult, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getrawtransaction",
		Params:  []interface{}{hash, true},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetTx - failed to marshal request: %v", err)
	}

	cur := self.node
	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	for err != nil {
		log.Errorf("BTCTools.GetTx err:%s, url: %s", err.Error(), self.urls[self.node])
		self.node ++
		self.node = self.node % len(self.urls)
		if self.node == cur {
			break
		}
		self.restclient.SetAddr(self.urls[self.node])
		self.restclient.SetAuth(self.users[self.node], self.passwds[self.node])
		repdata, err = self.restclient.SendRestRequestWithAuth(reqdata)
	}
	if err != nil {
		return nil, fmt.Errorf("GetTx - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetTx - failed to unmarshal response: %v", err)
	}
	if rep.Error != nil {
		return nil, fmt.Errorf("GetTx - response shows failure: %v", rep.Error.Message)
	}

	txRawResult := &btcjson.TxRawResult{}
	err = json.Unmarshal(rep.Result, txRawResult)
	if err != nil {
		return nil, fmt.Errorf("GetTx - Unmarshal Result: %v", err)
	}
	return txRawResult, nil
}

func (self *BTCTools) GetTxOut(hash string, vout uint32) (*btcjson.GetTxOutResult, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "gettxout",
		Params:  []interface{}{hash, vout, false},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - failed to marshal request: %v", err)
	}

	cur := self.node
	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	for err != nil{
		log.Errorf("BTCTools.GetTxOut err:%s, url: %s", err.Error(), self.urls[self.node])
		self.node ++
		self.node = self.node % len(self.urls)
		if self.node == cur {
			break
		}
		self.restclient.SetAddr(self.urls[self.node])
		self.restclient.SetAuth(self.users[self.node], self.passwds[self.node])
		repdata, err = self.restclient.SendRestRequestWithAuth(reqdata)
	}
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - failed to unmarshal response: %v", err)
	}

	txOutResult := &btcjson.GetTxOutResult{}
	err = json.Unmarshal(rep.Result, txOutResult)
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - Unmarshal Result: %v", err)
	}
	return txOutResult, nil
}

func (self *BTCTools) GetBlock(hash string) (*btcjson.BlockDetails, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblock",
		Params:  []interface{}{hash, true},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetBlock - failed to marshal request: %v", err)
	}

	cur := self.node
	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	for err != nil {
		log.Errorf("BTCTools.GetBlock err:%s, url: %s", err.Error(), self.urls[self.node])
		self.node ++
		self.node = self.node % len(self.urls)
		if self.node == cur {
			break
		}
		self.restclient.SetAddr(self.urls[self.node])
		self.restclient.SetAuth(self.users[self.node], self.passwds[self.node])
		repdata, err = self.restclient.SendRestRequestWithAuth(reqdata)
	}
	if err != nil {
		return nil, fmt.Errorf("GetBlock - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetBlock - failed to unmarshal response: %v", err)
	}
	if rep.Error != nil {
		return nil, fmt.Errorf("GetBlock - response shows failure: %v", rep.Error.Message)
	}

	block := &btcjson.BlockDetails{}
	json.Unmarshal(rep.Result, block)
	return block, nil
}

func (self *BTCTools) GetBlockHeaderByHeight(height uint32) (*wire.BlockHeader, error) {
	hash, err := self.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	header, err := self.GetBlockHeader(hash)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (self *BTCTools) PaserRawTx(rawTx []byte) (*wire.MsgTx, string, error) {
	tx := wire.MsgTx{}
	err := tx.BtcDecode(bytes.NewBuffer(rawTx), wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		return nil, "", fmt.Errorf("PaserRawTx - failed: %v", err)
	}
	return &tx, tx.TxHash().String(), nil
}

