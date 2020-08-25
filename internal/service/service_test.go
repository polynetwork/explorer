package service

import (
	"fmt"
	"github.com/polynetwork/explorer/internal/conf"
	"testing"
)

var srv *Service

func TestMain(m *testing.M) {
	err := conf.DefConfig.Init("../../config/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	srv = New(conf.DefConfig)
	m.Run()
}
