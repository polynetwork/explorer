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

package common

import (
	"encoding/hex"
	ontcommon "github.com/ontio/ontology/common"
	"strconv"
)

const (
	CHAIN_POLY   = 0
	CHAIN_BTC    = 1
	CHAIN_ETH    = 2
	CHAIN_ONT    =  3
	//CHAIN_NEO    = 5
	CHAIN_NEO    = 5
	CHAIN_COSMOS = 15
	CHAIN_BSC = 79
	CHAIN_HECO = 7
)

const (
	BTC_TOKEN_NAME = "btc"
	BTC_TOKEN_HASH = "0000000000000000000000000000000000000011"
	UNISWAP_NAME = "UNI_V2_ETH_WBTC"
)

func HexString2Base58Address(address string) string {
	addr, err := ontcommon.AddressFromHexString(address)
	if err != nil {
		return ""
	}
	return addr.ToBase58()
}

func HexBytes2Base58Address(address []byte) string {
	addr, err := ontcommon.AddressParseFromBytes(address)
	if err != nil {
		return ""
	}
	return addr.ToBase58()
}

func String2Float64(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return v
}

func HexReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := HexReverse(aa)
	return hex.EncodeToString(bb)
}

