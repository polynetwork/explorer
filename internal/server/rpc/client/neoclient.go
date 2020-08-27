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
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/log"
)

type NeoClient struct {
	client  *rpc.RpcClient
	urls     []string
	node     int
}

func NewNeoClient(c *conf.Config) (client *NeoClient) {
	rawClient := rpc.NewClient(c.Neo.Rawurl[0])
	//return &NeoClient{rawClient}
	return &NeoClient{
		client: rawClient,
		urls: c.Neo.Rawurl,
		node: 0,
	}
}

func (client *NeoClient) GetBlockCount() rpc.GetBlockCountResponse {
	cur := client.node
	res := client.client.GetBlockCount()
	for res.ErrorResponse.Error.Message != "" {
		log.Errorf("NeoClient.GetBlockCount err:%s, url: %s", res.ErrorResponse.Error.Message, client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return rpc.GetBlockCountResponse{
				ErrorResponse: rpc.ErrorResponse{
					Error : rpc.RpcError {
						Code: -1,
						Message: "all node is not working!",
					},

				},
			}
		}
		client.client = rpc.NewClient(client.urls[client.node])
		res = client.client.GetBlockCount()
	}
	return res
}

func (client *NeoClient) GetBlockByIndex(index uint32) rpc.GetBlockResponse {
	cur := client.node
	res := client.client.GetBlockByIndex(index)
	for res.ErrorResponse.Error.Message != "" {
		log.Errorf("NeoClient.GetBlockByIndex err:%s, url: %s", res.ErrorResponse.Error.Message, client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return rpc.GetBlockResponse{
				ErrorResponse: rpc.ErrorResponse{
					Error : rpc.RpcError {
						Code: -1,
						Message: "all node is not working!",
					},

				},
			}
		}
		client.client = rpc.NewClient(client.urls[client.node])
		res = client.client.GetBlockByIndex(index)
	}
	return res
}

func (client *NeoClient) GetApplicationLog(txId string) rpc.GetApplicationLogResponse {
	cur := client.node
	res := client.client.GetApplicationLog(txId)
	for res.ErrorResponse.Error.Message != "" {
		log.Errorf("NeoClient.GetApplicationLog err:%s, url: %s", res.ErrorResponse.Error.Message, client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return rpc.GetApplicationLogResponse{
				ErrorResponse: rpc.ErrorResponse{
					Error : rpc.RpcError {
						Code: -1,
						Message: "all node is not working!",
					},

				},
			}
		}
		client.client = rpc.NewClient(client.urls[client.node])
		res = client.client.GetApplicationLog(txId)
	}
	return res
}
