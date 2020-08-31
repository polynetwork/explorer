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
	"fmt"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/log"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/polynetwork/cosmos-poly-module/btcx"
	"github.com/polynetwork/cosmos-poly-module/ccm"
	"github.com/polynetwork/cosmos-poly-module/ft"
	"github.com/polynetwork/cosmos-poly-module/headersync"
	"github.com/polynetwork/cosmos-poly-module/lockproxy"
)

type CosmosClient struct {
	client   *tmclient.HTTP
	urls     []string
	node     int
	cdc      *codec.Codec
}

func NewCosmosClient(c *conf.Config) (client *CosmosClient) {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("swth", "swthpub")
	config.SetBech32PrefixForValidator("swthvaloper", "swthvaloperpub")
	config.SetBech32PrefixForConsensusNode("swthvalcons", "swthvalconspub")
	rawClient, err := tmclient.New(c.Cosmos.Rawurl[0], "/websocket")
	if err != nil {
		panic(err)
	}
	cdc := NewCDC()
	return &CosmosClient{
		client:rawClient,
		urls: c.Cosmos.Rawurl,
		node: 0,
		cdc: cdc,
	}
}

func NewCDC() *codec.Codec {
	cdc := codec.New()
	bank.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	btcx.RegisterCodec(cdc)
	ccm.RegisterCodec(cdc)
	ft.RegisterCodec(cdc)
	headersync.RegisterCodec(cdc)
	lockproxy.RegisterCodec(cdc)
	return cdc
}

func (client *CosmosClient) Status() (*ctypes.ResultStatus, error) {
	cur := client.node
	result, err := client.client.Status()
	for err != nil || result == nil {
		log.Errorf("CosmosClient.Status err:%s, url: %s", err.Error(), client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		client.client, err = tmclient.New(client.urls[client.node], "/websocket")
		if err != nil {
			panic(err)
		}
		result, err = client.client.Status()
	}
	return result, err
}

func (client *CosmosClient) Block(height *int64) (*ctypes.ResultBlock, error) {
	cur := client.node
	result, err := client.client.Block(height)
	for err != nil || result == nil {
		log.Errorf("CosmosClient.Block err:%s, url: %s", err.Error(), client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		client.client, err = tmclient.New(client.urls[client.node], "/websocket")
		if err != nil {
			panic(err)
		}
		result, err = client.client.Block(height)
	}
	return result, err
}

func (client *CosmosClient) TxSearch(query string, prove bool, page, perPage int, orderBy string) (*ctypes.ResultTxSearch, error) {
	cur := client.node
	result, err := client.client.TxSearch(query, prove, page, perPage, orderBy)
	for err != nil || result == nil {
		log.Errorf("CosmosClient.TxSearch err:%s, url: %s", err.Error(), client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		client.client, err = tmclient.New(client.urls[client.node], "/websocket")
		if err != nil {
			panic(err)
		}
		result, err = client.client.TxSearch(query, prove, page, perPage, orderBy)
	}
	return result, err
}

func (client *CosmosClient) GetGas(tx []byte) uint64 {
	decoder := auth.DefaultTxDecoder(client.cdc)
	stdTx, err := decoder(tx)
	if err != nil {
		log.Errorf("cosmos client GetGas err: %s", err.Error())
		return 0
	}
	aa, ok := stdTx.(authtypes.StdTx)
	if !ok {
		log.Errorf("This is not cosmos std tx!")
		return 0
	}
	amount := aa.Fee.Amount.AmountOf("swth").BigInt()
	gas := big.NewInt(int64(aa.Fee.Gas))
	return amount.Div(amount, gas).Uint64()
}
