package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/polynetwork/explorer/internal/common"
	"math/big"
)

type NeoCrossChain struct {
	Txhash    string
	Amount    *big.Int
}

func NewNeoClient() (*rpc.RpcClient) {
	rawClient := rpc.NewClient("http://seed8.ngd.network:10332")
	return rawClient
}

func GetApplicationLog(client *rpc.RpcClient, txId string) rpc.GetApplicationLogResponse {
	res := client.GetApplicationLog(txId)
	if res.ErrorResponse.Error.Message != "" {
		panic("GetApplicationLog error")
	}
	return res
}

func main() {
	db, err := sql.Open("mysql",
		"root"+
			":"+"crossexplorer"+
			"@tcp("+"localhost:3306"+
			")/"+"cross_chain_explorer4"+
			"?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic("connect mysql failed")
	}

	var rows *sql.Rows
	if rows, err = db.Query("select txhash from tchain_transfer where chain_id = 4"); err != nil {
		panic(err)
	}
	allNoeCrossChains := make([]*NeoCrossChain, 0)
	defer rows.Close()
	for rows.Next() {
		r := new(NeoCrossChain)
		if err = rows.Scan(&r.Txhash); err != nil {
			panic(err)
		}
		allNoeCrossChains = append(allNoeCrossChains, r)
	}

	client := NewNeoClient()
	for _, neoCrossChain := range allNoeCrossChains {
		logResp := client.GetApplicationLog(neoCrossChain.Txhash)
		appLog := logResp.Result
		for _, item := range appLog.Executions {
			for _, notify := range item.Notifications {
				value := notify.State.Value
				method := value[0].Value
				xx, _ := hex.DecodeString(method)
				method = string(xx)
				if method == "UnlockEvent" {
					amount := big.NewInt(0)
					if value[3].Type == "Integer" {
						//data, _ := strconv.ParseUint(value[3].Value, 10, 64)
						amount, _ = new(big.Int).SetString((value[3].Value), 10)
						//amount = amount.SetInt64(int64(data))
					} else {
						amount, _ = new(big.Int).SetString(common.HexStringReverse(value[3].Value), 16)
					}
					neoCrossChain.Amount = amount
					fmt.Printf("new amount: %s\n", amount.String())
				}
			}
		}
	}
	for _, neoCrossChain := range allNoeCrossChains {
		if _, err = db.Exec("update tchain_transfer set amount = ? where txhash = ?", neoCrossChain.Amount.String(), neoCrossChain.Txhash); err != nil {
			panic(err)
		}
	}
}
