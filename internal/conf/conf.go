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

package conf

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/polynetwork/explorer/internal/log"
	"io/ioutil"
	"os"
	"time"
)

// Config .
type Config struct {
	Redis    *Redis    `json:"redis"`
	Mysql    *MySQL    `json:"mysql"`
	Server   *Server   `json:"server"`
	CoinMarketCap  *CoinMarketCap  `json:"coinmarketcap"`
	Neo      *Neo      `json:"neo"`
	Ontology *Ontology `json:"ontology"`
	Ethereum *Ethereum `json:"ethereum"`
	Alliance *Alliance `json:"alliance"`
	Bitcoin  *Bitcoin  `json:"btc"`
	Cosmos   *Cosmos   `json:"cosmos"`
}

type Redis struct {
	Proto        string        `json:"proto"`
	Addr         string        `json:"addr"`
	PoolSize     int           `json:"pool_size"`
	MinIdleConns int           `json:"min_idle_conns"`
	DialTimeout  time.Duration `json:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
	Expiration   time.Duration `json:"expiration"` // day
}

type MySQL struct {
	Url      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"pwd"`
	DbName   string `json:"dbName"`
}

type Server struct {
	RestPort           int    `json:"rest_port"`
	Version            string `json:"version"`
	Master             int    `json:"master"`
	HttpMaxConnections int    `json:"http_max_connections"`
	HttpCertPath       string `json:"http_cert_path"`
	HttpKeyPath        string `json:"http_key_path"`
	AssetStatisticTimeslot  uint32 `json:"asset_statistic_time_slot"`
	TransferStatisticTimeslot  uint32 `json:"transfer_statistic_time_slot"`
	LogLevel           uint32 `json:"loglevel"`
}

type CoinMarketCap struct {
	Url      string `json:"url"`
	AppKey   string `json:"appkey"`
}

type Bitcoin struct {
	Name          string        `json:"name"`
	ChainId       uint32        `json:"chainId"`
	User          []string        `json:"user"`
	Passwd        []string        `json:"passwd"`
	Rawurl        []string        `json:"rawurl"`
	BlockDuration time.Duration `json:"block_duration"`
}
type Neo struct {
	Name          string        `json:"name"`
	ChainId       uint32        `json:"chainId"`
	Rawurl        []string        `json:"rawurl"`
	BlockDuration time.Duration `json:"block_duration"`
}
type Ethereum struct {
	Name          string        `json:"name"`
	ChainId       uint32        `json:"chainId"`
	Rawurl        []string        `json:"rawurl"`
	BlockDuration time.Duration `json:"block_duration"`
	Proxy         string        `json:"proxy"`
	BTCX          string        `json:"btcx"`
}
type Cosmos struct {
	Name          string        `json:"name"`
	ChainId       uint32        `json:"chainId"`
	Rawurl        []string      `json:"rawurl"`
	BlockDuration time.Duration `json:"block_duration"`
}

type Ontology struct {
	Name          string        `json:"name"`
	ChainId       uint32        `json:"chainId"`
	Rawurl        []string        `json:"rawurl"`
	BlockDuration time.Duration `json:"block_duration"`
}

type Alliance struct {
	Name          string        `json:"name"`
	ChainId       uint32        `json:"chainId"`
	Rawurl        []string        `json:"rawurl"`
	BlockDuration time.Duration `json:"block_duration"`
}

type Contract struct {
	Address         string  `json:"address"`
	Chain           uint32  `json:"chain"`
}
type Token struct {
	Name            string  `json:"name"`
	Contracts       []Contract  `json:"contracts"`
}

var DefConfig = NewConfig()

func NewConfig() *Config {
	return &Config{}
}

func (this *Config) Init(fileName string) error {
	err := this.loadConfig(fileName)
	if err != nil {
		return fmt.Errorf("loadConfig error:%s", err)
	}
	return nil
}

func (this *Config) loadConfig(fileName string) error {
	data, err := this.readFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		return fmt.Errorf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return nil
}

func (this *Config) readFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Error("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}
