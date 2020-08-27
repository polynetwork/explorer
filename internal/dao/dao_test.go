package dao

import (
	"github.com/polynetwork/explorer/internal/conf"
	"testing"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if err := conf.DefConfig.Init("../../config/config.json"); err != nil {
		panic(err)
	}
	d = NEW(conf.DefConfig)
	m.Run()
}
