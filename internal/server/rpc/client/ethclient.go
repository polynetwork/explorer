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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/log"
	"math/big"
)

type EthereumClient struct {
	rpcClient *rpc.Client
	Client    *ethclient.Client
	urls     []string
	node     int
}

func NewEthereumClient(c *conf.Config) (client *EthereumClient) {
	rpcClient, err := rpc.Dial(c.Ethereum.Rawurl[0])
	if err != nil {
		log.Error("can't connect to ethereum", err)
		panic(err)
	}
	rawClient, err := ethclient.Dial(c.Ethereum.Rawurl[0])
	if err != nil {
		log.Error("can't connect to ethereum", err)
		panic(err)
	}
	return &EthereumClient{
		rpcClient: rpcClient,
		Client: rawClient,
		urls: c.Ethereum.Rawurl,
		node: 0,
	}
}

// GetHeaderByNumber returns the given header
func (ec *EthereumClient) GetHeaderByNumber(ctx *ctx.Context, number int64) (header *types.Header, err error) {
	cur := ec.node
	if number < 0 {
		header, err = ec.Client.HeaderByNumber(ctx.Context, nil)
	} else {
		header, err = ec.Client.HeaderByNumber(ctx.Context, big.NewInt(number))
	}
	for err != nil {
		log.Errorf("EthereumClient.GetHeaderByNumber err:%s, url: %s", err.Error(), ec.urls[ec.node])
		ec.node ++
		ec.node = ec.node % len(ec.urls)
		if ec.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		ec.rpcClient, _ = rpc.Dial(ec.urls[ec.node])
		ec.Client, _ = ethclient.Dial(ec.urls[ec.node])

		if number < 0 {
			header, err = ec.Client.HeaderByNumber(ctx.Context, nil)
		} else {
			header, err = ec.Client.HeaderByNumber(ctx.Context, big.NewInt(number))
		}
	}
	return header, err
}

// GetCurrentBlockHeight returns current block height
func (ec *EthereumClient) GetCurrentBlockHeight(ctx *ctx.Context) (height int64, _ error) {
	var result hexutil.Big
	cur := ec.node
	err := ec.rpcClient.CallContext(ctx.Context, &result, "eth_blockNumber")
	for err != nil {
		log.Errorf("EthereumClient.GetCurrentBlockHeight err:%s, url: %s", err.Error(), ec.urls[ec.node])
		ec.node ++
		ec.node = ec.node % len(ec.urls)
		if ec.node == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		ec.rpcClient, _ = rpc.Dial(ec.urls[ec.node])
		ec.Client, _ = ethclient.Dial(ec.urls[ec.node])
		err = ec.rpcClient.CallContext(ctx.Context, &result, "eth_blockNumber")
	}
	return (*big.Int)(&result).Int64(), err
}

func (ec *EthereumClient) GetTransactionByHash(ctx *ctx.Context, hash common.Hash) (*types.Transaction, error) {
	cur := ec.node
	tx, _, err := ec.Client.TransactionByHash(ctx.Context, hash)
	for err != nil {
		log.Errorf("EthereumClient.GetTransactionByHash err:%s, url: %s", err.Error(), ec.urls[ec.node])
		ec.node ++
		ec.node = ec.node % len(ec.urls)
		if ec.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		ec.rpcClient, _ = rpc.Dial(ec.urls[ec.node])
		ec.Client, _ = ethclient.Dial(ec.urls[ec.node])
		tx, _, err = ec.Client.TransactionByHash(ctx.Context, hash)
	}
	return tx, err
}

func (ec *EthereumClient) GetTransactionReceipt(ctx *ctx.Context, hash common.Hash) (*types.Receipt, error) {
	cur := ec.node
	receipt, err := ec.Client.TransactionReceipt(ctx.Context, hash)
	for err != nil {
		log.Errorf("EthereumClient.GetTransactionReceipt err:%s, url: %s", err.Error(), ec.urls[ec.node])
		ec.node ++
		ec.node = ec.node % len(ec.urls)
		if ec.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		ec.rpcClient, _ = rpc.Dial(ec.urls[ec.node])
		ec.Client, _ = ethclient.Dial(ec.urls[ec.node])
		receipt, err = ec.Client.TransactionReceipt(ctx.Context, hash)
	}
	return receipt, nil
}

// Close client
func (ec *EthereumClient) Close() {
	ec.rpcClient.Close()
	ec.Client.Close()
}
