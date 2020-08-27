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
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/log"
	sdk "github.com/polynetwork/poly-go-sdk"
	sdkcom "github.com/polynetwork/poly-go-sdk/common"
	"github.com/polynetwork/poly/core/types"
)

type AllianceSDK struct {
	sdk *sdk.PolySdk
	urls     []string
	node     int
}

func NewAllianceSDK(c *conf.Config) *AllianceSDK {
	rawsdk := sdk.NewPolySdk()
	rawsdk.NewRpcClient().SetAddress(c.Alliance.Rawurl[0])
	return &AllianceSDK{
		sdk: rawsdk,
		urls: c.Alliance.Rawurl,
		node: 0,
	}
}

func (client *AllianceSDK) GetCurrentBlockHeight () (uint32, error) {
	cur := client.node
	height, err := client.sdk.GetCurrentBlockHeight()
	for err != nil {
		log.Errorf("AllianceSDK.GetCurrentBlockHeight err:%s, url: %s", err.Error(), client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		rawsdk := sdk.NewPolySdk()
		rawsdk.NewRpcClient().SetAddress(client.urls[client.node])
		client.sdk = rawsdk
		height, err = client.sdk.GetCurrentBlockHeight()
	}
	return height, err
}

func (client *AllianceSDK) GetBlockByHeight(height uint32) (*types.Block, error) {
	cur := client.node
	block, err := client.sdk.GetBlockByHeight(height)
	for err != nil {
		log.Errorf("AllianceSDK.GetBlockByHeight err:%s, url: %s", err.Error(), client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		rawsdk := sdk.NewPolySdk()
		rawsdk.NewRpcClient().SetAddress(client.urls[client.node])
		client.sdk = rawsdk
		block, err = client.sdk.GetBlockByHeight(height)
	}
	return block, err
}

func (client *AllianceSDK) GetSmartContractEventByBlock(height uint32) ([]*sdkcom.SmartContactEvent, error) {
	cur := client.node
	event, err := client.sdk.GetSmartContractEventByBlock(height)
	for err != nil {
		log.Errorf("AllianceSDK.GetSmartContractEventByBlock err:%s, url: %s", err.Error(), client.urls[client.node])
		client.node ++
		client.node = client.node % len(client.urls)
		if client.node == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		rawsdk := sdk.NewPolySdk()
		rawsdk.NewRpcClient().SetAddress(client.urls[client.node])
		client.sdk = rawsdk
		event, err = client.sdk.GetSmartContractEventByBlock(height)
	}
	return event, err
}