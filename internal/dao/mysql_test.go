package dao

import (
	"fmt"
	"github.com/polynetwork/explorer/internal/conf"
	"testing"
	"time"
)

func NewDao() *Dao {
	if err := conf.DefConfig.Init("../../config/config.json"); err != nil {
		panic(err)
	}
	d := NEW(conf.DefConfig)
	return d
}

func TestSelectMChainTxByLimit(t *testing.T) {
	dao := NewDao()
	dao.SelectMChainTxByLimit(0, 10)
}

func TestSelectMChainTxByLimitPerf(t *testing.T) {
	dao := NewDao()
	querys := 10000
	start := time.Now().Unix()
	for i := 0;i < querys;i ++ {
		_, err := dao.SelectMChainTxByLimit(0, 10)
		if err != nil {
			fmt.Printf("SelectMChainTxByLimit err: %s\n", err.Error())
		}
	}
	end := time.Now().Unix()
	fmt.Printf("querys: %d, times: %d\n", querys, end - start)
}
