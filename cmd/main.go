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

package main

import (
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/ctx"
	"github.com/polynetwork/explorer/internal/log"
	restful "github.com/polynetwork/explorer/internal/server/restful/server"
	"os"
	"os/signal"
	"syscall"
)

var (
	rt *restful.RestServer
)

func init() {
	//flag.StringVar(&confPath, "conf", "", "set the config file path")
	log.InitLog(1, "../log/")
	err := conf.DefConfig.Init("../config/config.json")
	if err != nil {
		log.Error(err)
		return
	}
}
func main() {
	context := ctx.New().WithCancel()
	startRest(context)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("get a signal %s", s.String())
		switch s {
		// shutdown service
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			context.Cancel()
			rt.Stop(context)
			return
		default:
			return
		}
	}
}

func startRest(context *ctx.Context) {
	rt = restful.InitRestServer(context)
	go rt.Start()
	log.Infof("restful start success")
}
