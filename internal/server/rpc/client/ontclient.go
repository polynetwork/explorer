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
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/core/types"
	sdkcom "github.com/ontio/ontology-go-sdk/common"
)

type OntologySDK struct {
	sdk      *sdk.OntologySdk
	urls     []string
	node     int
}

func NewOntologySDK(c *conf.Config) *OntologySDK {
	rawsdk := sdk.NewOntologySdk()
	rawsdk.NewRpcClient().SetAddress(c.Ontology.Rawurl[0])
	return &OntologySDK{
		sdk: rawsdk,
		urls: c.Ontology.Rawurl,
		node: 0,
	}
}

func (client *OntologySDK) NextClient() int {
	client.node ++
	client.node = client.node % len(client.urls)
	rawsdk := sdk.NewOntologySdk()
	rawsdk.NewRpcClient().SetAddress(client.urls[client.node])
	client.sdk = rawsdk
	return client.node
}

func (client *OntologySDK) GetCurrentBlockHeight() (uint32, error) {
	cur := client.node
	height, err := client.sdk.GetCurrentBlockHeight()
	for err != nil {
		log.Errorf("OntologySDK.GetCurrentBlockHeight err:%s, url: %s", err.Error(), client.urls[client.node])
		next := client.NextClient()
		if next == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		height, err = client.sdk.GetCurrentBlockHeight()
	}
	return height, err
}

func (client *OntologySDK) GetBlockByHeight(height uint32) (*types.Block, error) {
	cur := client.node
	block, err := client.sdk.GetBlockByHeight(height)
	for err != nil {
		log.Errorf("OntologySDK.GetBlockByHeight err:%s, url: %s", err.Error(), client.urls[client.node])
		next := client.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		block, err = client.sdk.GetBlockByHeight(height)
	}
	return block, err
}

func (client *OntologySDK) GetSmartContractEventByBlock(height uint32) ([]*sdkcom.SmartContactEvent, error) {
	cur := client.node
	event, err := client.sdk.GetSmartContractEventByBlock(height)
	for err != nil {
		log.Errorf("OntologySDK.GetBlockByHeight err:%s, url: %s", err.Error(), client.urls[client.node])
		next := client.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		event, err = client.sdk.GetSmartContractEventByBlock(height)
	}
	return event, err
}