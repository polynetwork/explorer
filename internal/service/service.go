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
	"encoding/hex"
	cosmos_types "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/dao"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	restclient "github.com/polynetwork/explorer/internal/server/restful/client"
	"github.com/polynetwork/explorer/internal/server/rpc/client"
	ontcommon "github.com/ontio/ontology/common"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type Service struct {
	c              *conf.Config
	dao            *dao.Dao
	neoClient      *client.NeoClient
	ethClient      *client.EthereumClient
	ontClient      *client.OntologySDK
	allianceClient *client.AllianceSDK
	bitcoinClient  *restclient.BTCTools
	cosmosClient   *client.CosmosClient
	chain          []*model.ChainInfo
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:              c,
		dao:            dao.NEW(c),
		neoClient:      client.NewNeoClient(c),
		ethClient:      client.NewEthereumClient(c),
		ontClient:      client.NewOntologySDK(c),
		allianceClient: client.NewAllianceSDK(c),
		bitcoinClient:  restclient.NewBtcTools(c),
		cosmosClient:   client.NewCosmosClient(c),
		chain: make([]*model.ChainInfo, 0),
	}
	return s
}

// Ping Service
func (s *Service) Ping() (err error) {
	return s.dao.Ping()
}

// Close Service
func (s *Service) Close() {
	s.dao.Close()
	s.ethClient.Close()
}

func (exp *Service) GetChainInfos() []*model.ChainInfo {
	// get all chains
	chainInfos, err := exp.dao.SelectAllChainInfos()
	if err != nil {
		panic( err)
	}
	if chainInfos == nil {
		panic("GetExplorerInfo: can't get AllChainInfos")
	}

	// get all tokens and contracts
	for _, chainInfo := range chainInfos {
		chainContracts, err := exp.dao.SelectContractById(chainInfo.Id)
		if err != nil {
			panic(err)
		}
		chainInfo.Contracts = chainContracts

		chainTokens, err := exp.dao.SelectTokenById(chainInfo.Id)
		if err != nil {
			panic(err)
		}
		chainInfo.Tokens = chainTokens
	}
	return chainInfos
}

func (exp *Service) Start(context *ctx.Context) {
	exp.CheckChains(context)
	t := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-t.C:
			exp.CheckChains(context)
		case <-context.Context.Done():
			log.Infof("stop service start routine!")
			return
		}
	}
}

func (exp *Service) CheckChains(context *ctx.Context) {
	chainInfoNew := exp.GetChainInfos()
	chainInfoOld := exp.chain
	exp.chain = chainInfoNew
	if exp.c.Server.Master == 0 {
		return
	}
	for _, chainNew := range chainInfoNew {
		exist := false
		for _, chainOld := range chainInfoOld {
			if chainOld.Id == chainNew.Id {
				exist = true
				break
			}
		}
		if exist == true {
			continue
		}
		if chainNew.Id == common.CHAIN_POLY {
			go exp.LoadAllianceCrossTxFromChain(context)
		} else if chainNew.Id == common.CHAIN_BTC {
			go exp.MonitorBtcChainFromAlliance(context)
		} else if chainNew.Id == common.CHAIN_ETH {
			go exp.LoadEthCrossTxFromChain(context)
		} else if chainNew.Id == common.CHAIN_ONT {
			go exp.LoadOntCrossTxFromChain(context)
		} else if chainNew.Id == common.CHAIN_NEO {
			go exp.LoadNeoCrossTxFromChain(context)
		} else if chainNew.Id == common.CHAIN_COSMOS {
			go exp.LoadCosmosCrossTxFromChain(context)
		}
	}
}

func (exp *Service) TxType2Name(ttype uint32) string {
	return "cross chain transfer"
}

func (exp *Service) GetChain(chainId uint32) *model.ChainInfo {
	for _, chainInfo := range exp.chain {
		if chainInfo.Id == chainId {
			return chainInfo
		}
	}
	return nil
}

func (exp *Service) ChainId2Name(chainId uint32) string {
	for _, chainInfo := range exp.chain {
		if chainInfo.Id == chainId {
			return chainInfo.Name
		}
	}
	return "unknow chain"
}

func (exp *Service) AssetInfo(tokenHash string) (string, string) {
	for _, chainInfo := range exp.chain {
		for _, token := range chainInfo.Tokens {
			if token.Hash == tokenHash {
				return token.Name, token.Type
			}
		}
	}
	return "unknow token", "unknow token"
}

func (exp *Service) GetToken(tokenHash string) (*model.ChainToken) {
	for _, chainInfo := range exp.chain {
		for _, token := range chainInfo.Tokens {
			if token.Hash == tokenHash {
				return token
			}
		}
	}
	return nil
}

func (exp *Service) SearchToken(name string, chainId uint32) (*model.ChainToken) {
	for _, chainInfo := range exp.chain {
		if chainInfo.Id != chainId {
			continue
		}
		for _, token := range chainInfo.Tokens {
			if token.Name == name {
				return token
			}
		}
		break
	}
	return nil
}

func (exp *Service) Hash2Address(chainId uint32, value string) string {
	if chainId == common.CHAIN_ETH {
		addr := ethcommon.HexToAddress(value)
		return strings.ToLower(addr.String()[2:])
	} else if chainId == common.CHAIN_COSMOS {
		addr, _ := cosmos_types.AccAddressFromHex(value)
		return addr.String()
	} else if chainId == common.CHAIN_BTC {
		addrHex, _ := hex.DecodeString(value)
		return string(addrHex)
	} else if chainId == common.CHAIN_NEO {
		addrHex, _ := hex.DecodeString(value)
		addr, _ := helper.UInt160FromBytes(addrHex)
		return helper.ScriptHashToAddress(addr)
	} else if chainId == common.CHAIN_ONT {
		value = common.HexStringReverse(value)
		addr, _ := ontcommon.AddressFromHexString(value)
		return addr.ToBase58()
	}
	return value
}

func (exp *Service) FormatAmount(precision uint64, amount uint64) string {
	precision_new := decimal.New(int64(precision), 0)
	amount_new := decimal.New(int64(amount), 0)
	return amount_new.Div(precision_new).String()
}

func (exp *Service) FormatFee(chain uint32, fee uint64) string {
	if chain == common.CHAIN_BTC {
		precision_new := decimal.New(int64(100000000), 0)
		fee_new := decimal.New(int64(fee), 0)
		return fee_new.Div(precision_new).String()
	} else if chain == common.CHAIN_ONT {
		precision_new := decimal.New(int64(1000000000), 0)
		fee_new := decimal.New(int64(fee), 0)
		return fee_new.Div(precision_new).String()
	} else if chain == common.CHAIN_ETH {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		fee_new := decimal.New(int64(fee), 0)
		return fee_new.Div(precision_new).String()
	} else {
		precision_new := decimal.New(int64(1), 0)
		fee_new := decimal.New(int64(fee), 0)
		return fee_new.Div(precision_new).String()
	}
}

func (exp *Service) DayOfTime(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	return uint32(end_t_new.Unix())
}

func (exp *Service) DayOfTimeUp(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	time_t_unix := uint32(end_t_new.Unix())
	if t > time_t_unix {
		time_t_unix = uint32(end_t_new.AddDate(0, 0, 1).Unix())
	}
	return time_t_unix
}

func (exp *Service) DayOfTimeAddOne(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	time_t_unix := uint32(end_t_new.AddDate(0, 0, 1).Unix())
	return time_t_unix
}

func (exp *Service) DayOfTimeSubOne(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	time_t_unix := uint32(end_t_new.AddDate(0, 0, -1).Unix())
	return time_t_unix
}

