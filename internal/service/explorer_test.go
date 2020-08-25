package service

import (
	"bytes"
	"fmt"
	"github.com/polynetwork/explorer/internal/model"
	"io/ioutil"
	"net/http"
	"testing"
)

var clientMocker = &http.Client{}

func BenchmarkGetCrossTxList(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Get("http://localhost:30334/api/v1/getexplorerinfo")
		if err != nil {
			b.Error(err)
		}
		//fmt.Println(res)
	}
	b.StopTimer()
}

func BenchmarkGetCrossTx(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Get("http://localhost:30334/api/v1/getcrosstx/9765f91b8b1036a5ae29e9e872b4b937a201098e1317ba7e36261ff7cdc49285")
		if err != nil {
			b.Error(err)
		}
		//fmt.Println(res)
	}
	b.StopTimer()
}

func BenchmarkGetInfo(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		params := "{\"start\":\"0\",\"end\":\"20\"}"
		_, err := Post("http://localhost:30334/api/v1/getcrosstxlist", params)
		if err != nil {
			b.Error(err)
		}
		//fmt.Println(res)
	}
	b.StopTimer()
}

func Get(params string) (string, error) {
	req, err := http.NewRequest("GET", params, nil)
	if err != nil {
		return "", fmt.Errorf("GET %v", err)
	}
	resp, err := clientMocker.Do(req)
	if err != nil {
		return "", fmt.Errorf("GET %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), nil
	}
	return "", err
}

func Post(addr string, params string) (string, error) {
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer([]byte(params)))
	if err != nil {
		return "", fmt.Errorf("POST %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := clientMocker.Do(req)
	if err != nil {
		return "", fmt.Errorf("GET %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), nil
	}
	return "", err
}

func TestOutputCrossChainTxStatus(t *testing.T) {
	exp := New(nil)
	status := make([]*model.CrossChainTxStatus, 0)
	status = append(status, &model.CrossChainTxStatus{
		TT : 1593014400,
		TxNumber: 7,
	})
	exp.outputCrossChainTxStatus(status, 1592274867, 1593534067, 7)
}
