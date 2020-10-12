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

package restful

import (
	"crypto/tls"
	"encoding/json"
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/log"
	ierror "github.com/polynetwork/explorer/internal/server/restful/error"
	"golang.org/x/net/netutil"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const TLS_PORT int = 443
const MAX_REQUEST_BODY_SIZE = 1 << 20

type handler func(map[string]interface{}) map[string]interface{}
type Action struct {
	sync.RWMutex
	name    string
	handler handler
}
type RestServer struct {
	router   *Router
	listener net.Listener
	server   *http.Server
	postMap  map[string]Action //post method map
	getMap   map[string]Action //get method map
}

const (
	GET_EXPLORER_INFO = "/api/v1/getexplorerinfo"
	GET_CROSSTX       = "/api/v1/getcrosstx"
	GET_CROSSTX_LIST  = "/api/v1/getcrosstxlist"
	GET_TOKENTX_LIST  = "/api/v1/gettokentxlist"
	GET_ADDRESSTX_LIST = "/api/v1/getaddresstxlist"
	GET_LATEST_VALIDATOR = "/api/v1/getlatestvalidator"
	GET_ASSET_STATISTIC  = "/api/v1/getassetstatistic"
	GET_TRANSFER_STATISTIC  = "/api/v1/gettransferstatistic"
)

//init restful server
func InitRestServer(context *ctx.Context) *RestServer {
	rt := &RestServer{}
	InitInterface()
	StartMonitorService(context)
	rt.router = NewRouter()
	rt.registryMethod()
	rt.initGetHandler()
	rt.initPostHandler()
	return rt
}

//start server
func (this *RestServer) Start() error {
	retPort := int(conf.DefConfig.Server.RestPort)
	if retPort == 0 {
		log.Fatal("Not configure HttpRestPort port ")
		return nil
	}
	log.Infof("restful port: %d", retPort)
	tlsFlag := false
	if tlsFlag || retPort%1000 == TLS_PORT {
		var err error
		this.listener, err = this.initTlsListen()
		if err != nil {
			log.Error("Https Cert: ", err.Error())
			return err
		}
	} else {
		var err error
		this.listener, err = net.Listen("tcp", ":"+strconv.Itoa(retPort))
		if err != nil {
			log.Fatal("net.Listen: ", err.Error())
			return err
		}
	}
	this.server = &http.Server{Handler: this.router}
	//set LimitListener number
	if conf.DefConfig.Server.HttpMaxConnections > 0 {
		this.listener = netutil.LimitListener(this.listener, conf.DefConfig.Server.HttpMaxConnections)
	}
	err := this.server.Serve(this.listener)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
		return err
	}

	return nil
}

// resigtry handler method
func (this *RestServer) registryMethod() {
	getMethodMap := map[string]Action{
		GET_CROSSTX:       {name: "getcrosstx", handler: GetCrossTx},
		GET_LATEST_VALIDATOR:       {name: "getlatestvalidator", handler: GetLatestValidator},
		GET_ASSET_STATISTIC:       {name: "getassetstatistic", handler: GetAssetStatistic},
		GET_TRANSFER_STATISTIC:       {name: "gettransferstatistic", handler: GetTransferStatistic},
	}
	postMethodMap := map[string]Action{
		GET_CROSSTX_LIST: {name: "getcrosstxlist", handler: GetCrossTxList},
		GET_EXPLORER_INFO: {name: "getexplorerinfo", handler: GetExplorerInfo},
		GET_TOKENTX_LIST: {name: "gettokentxlist", handler: GetTokenTxList},
		GET_ADDRESSTX_LIST: {name: "getaddresstxlist", handler: GetAddressTxList},
	}
	this.postMap = postMethodMap
	this.getMap = getMethodMap
}

func (this *RestServer) getPath(url string) string {
	/*
	if strings.Contains(url, strings.TrimRight(GET_CROSSTX, ":txhash")) {
		return GET_CROSSTX
	}
	*/
	return url
}

//get request params
func (this *RestServer) getParams(r *http.Request, url string, req map[string]interface{}) map[string]interface{} {
	switch url {
	case GET_CROSSTX:
		req["txhash"] = getParam(r, "txhash")
	default:
	}
	return req
}

//init get handler
func (this *RestServer) initGetHandler() {

	for k, _ := range this.getMap {
		this.router.Get(k, func(w http.ResponseWriter, r *http.Request) {

			var req = make(map[string]interface{})
			var resp map[string]interface{}

			url := this.getPath(r.URL.Path)
			if h, ok := this.getMap[url]; ok {
				req = this.getParams(r, url, req)
				resp = h.handler(req)
				resp["action"] = h.name
			} else {
				resp = ResponsePack(ierror.REST_METHOD_INVALID)
			}
			this.response(w, resp)
		})
	}
}

//init post handler
func (this *RestServer) initPostHandler() {
	for k, _ := range this.postMap {
		this.router.Post(k, func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(io.LimitReader(r.Body, MAX_REQUEST_BODY_SIZE))
			defer r.Body.Close()
			var req = make(map[string]interface{})
			var resp map[string]interface{}
			url := this.getPath(r.URL.Path)
			if h, ok := this.postMap[url]; ok {
				if err := decoder.Decode(&req); err == nil {
					req = this.getParams(r, url, req)
					req["ip"] = r.RemoteAddr
					resp = h.handler(req)
					resp["action"] = h.name
				} else {
					resp = ResponsePack(ierror.REST_ILLEGAL_DATAFORMAT)
					resp["action"] = h.name
				}
			} else {
				resp = ResponsePack(ierror.REST_METHOD_INVALID)
			}
			this.response(w, resp)
		})
	}
	//Options
	for k, _ := range this.postMap {
		this.router.Options(k, func(w http.ResponseWriter, r *http.Request) {
			this.write(w, []byte{})
		})
	}

}
func (this *RestServer) write(w http.ResponseWriter, data []byte) {
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}

//response
func (this *RestServer) response(w http.ResponseWriter, resp map[string]interface{}) {
	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("HTTP Handle - json.Marshal: %v", err)
		return
	}
	this.write(w, data)
}

//stop restful server
func (this *RestServer) Stop(context *ctx.Context) {
	if this.server != nil {
		StopService()
		this.server.Shutdown(context.Context)
		log.Error("Close restful ")
	}
}

//restart server
func (this *RestServer) Restart(context *ctx.Context, cmd map[string]interface{}) map[string]interface{} {
	go func() {
		time.Sleep(time.Second)
		this.Stop(context)
		time.Sleep(time.Second)
		go this.Start()
	}()

	var resp = ResponsePack(ierror.SUCCESS)
	return resp
}

//init tls
func (this *RestServer) initTlsListen() (net.Listener, error) {

	certPath := conf.DefConfig.Server.HttpCertPath
	keyPath := conf.DefConfig.Server.HttpKeyPath

	// load cert
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Error("load keys fail", err)
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	restPort := strconv.Itoa(int(conf.DefConfig.Server.RestPort))
	log.Info("TLS listen port is ", restPort)
	listener, err := tls.Listen("tcp", ":"+restPort, tlsConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return listener, nil
}
