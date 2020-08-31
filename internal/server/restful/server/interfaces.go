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

// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta


package restful

import (
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/server/restful/error"
	"github.com/polynetwork/explorer/internal/service"
	"strconv"
)

var srv *service.Service

func InitInterface() {
	srv = service.New(conf.DefConfig)
}

// swagger:route POST /api/v1/getexplorerinfo crosschain-explorer ExplorerInfoRequest
//
// get explore information
//
// get explore information
//
// Responses:
// 200: ExplorerInfoResponse

func GetExplorerInfo(cmd map[string]interface{}) map[string]interface{} {
	if cmd["start"] == nil || cmd["end"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start_str, ok := cmd["start"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end_str, ok := cmd["end"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start, err := strconv.Atoi(start_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end, err := strconv.Atoi(end_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	if start < 0 || end < 0 || start >= end {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	code, result := srv.GetExplorerInfo(uint32(start), uint32(end))
	if code != error.SUCCESS {
		return ResponsePack(code)
	}
	resp := ResponsePack(error.SUCCESS)
	resp["result"] = result
	return resp
}

// swagger:route GET /api/v1/getcrosstx crosschain-explorer CrossTxReq
//
// get cross chain transaction information
//
// get cross chain transaction information
//
// Responses:
// 200: CrossTxResponse

func GetCrossTx(cmd map[string]interface{}) map[string]interface{} {
	if cmd["txhash"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	txhash, ok := cmd["txhash"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	code, result := srv.GetCrossTx(txhash)
	if code != error.SUCCESS {
		return ResponsePack(code)
	}
	resp := ResponsePack(error.SUCCESS)
	resp["result"] = result
	return resp
}

// swagger:route POST /api/v1/getcrosstxlist crosschain-explorer  CrossTxListRequest
//
// get cross chain transaction list
//
// get cross chain transaction list
//
// Responses:
// 200: CrossTxListResponse

func GetCrossTxList(cmd map[string]interface{}) map[string]interface{} {
	if cmd["start"] == nil || cmd["end"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start_str, ok := cmd["start"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end_str, ok := cmd["end"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start, err := strconv.Atoi(start_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end, err := strconv.Atoi(end_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	if start < 0 || end < 0 || start >= end {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	code, result := srv.GetCrossTxList(start, end)
	if code != error.SUCCESS {
		return ResponsePack(code)
	}
	resp := ResponsePack(error.SUCCESS)
	resp["result"] = result
	return resp
}

// swagger:route POST /api/v1/gettokentxlist crosschain-explorer TokenTxListRequest
//
// get cross chain transaction list of a token
//
// get cross chain transaction list of a token
//
// Responses:
// 200: TokenTxListResponse

func GetTokenTxList(cmd map[string]interface{}) map[string]interface{} {
	if cmd["token"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	token, ok := cmd["token"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	if cmd["start"] == nil || cmd["end"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start_str, ok := cmd["start"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end_str, ok := cmd["end"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start, err := strconv.Atoi(start_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end, err := strconv.Atoi(end_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	if start < 0 || end < 0 || start >= end {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	code, result := srv.GetTokenTxList(token, uint32(start), uint32(end))
	if code != error.SUCCESS {
		return ResponsePack(code)
	}
	resp := ResponsePack(error.SUCCESS)
	resp["result"] = result
	return resp
}

// swagger:route POST /api/v1/getaddresstxlist crosschain-explorer AddressTxListRequest
//
// get cross chain transaction list of a address
//
// get cross chain transaction list of a address
//
// Responses:
// 200: AddressTxListResponse

func GetAddressTxList(cmd map[string]interface{}) map[string]interface{} {
	if cmd["address"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	address, ok := cmd["address"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	if cmd["chain"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	chainId_str, ok := cmd["chain"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	chainId, err := strconv.Atoi(chainId_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	if cmd["start"] == nil || cmd["end"] == nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start_str, ok := cmd["start"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end_str, ok := cmd["end"].(string)
	if !ok {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	start, err := strconv.Atoi(start_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	end, err := strconv.Atoi(end_str)
	if err != nil {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	if start < 0 || end < 0 || start >= end {
		return ResponsePack(error.REST_PARAM_INVALID)
	}
	code, result := srv.GetAddressTxList(uint32(chainId), address, uint32(start), uint32(end))
	if code != error.SUCCESS {
		return ResponsePack(code)
	}
	resp := ResponsePack(error.SUCCESS)
	resp["result"] = result
	return resp
}

func GetLatestValidator(cmd map[string]interface{}) map[string]interface{} {
	code, result := srv.GetLatestValidator()
	if code != error.SUCCESS {
		return ResponsePack(code)
	}
	resp := ResponsePack(error.SUCCESS)
	resp["result"] = result
	return resp
}

func GetAssetStatistic(cmd map[string]interface{}) map[string]interface{} {
	code, result := srv.GetAssetStatistic()
	if code != error.SUCCESS {
		return ResponsePack(code)
	}
	resp := ResponsePack(error.SUCCESS)
	resp["result"] = result
	return resp
}

func StartMonitorService(context *ctx.Context) {
	go srv.Start(context)
}

func StopService() {
	srv.Close()
}
