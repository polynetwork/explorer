package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/explorer/internal/ctx"
	"testing"
)

func TestService_GetCurrentBlockHeight(t *testing.T) {
	ethContext := ctx.New()
	height, err := srv.ethClient.GetCurrentBlockHeight(ethContext)
	if err != nil {
		t.Errorf("TestService_GetCurrentBlockHeight %v", err)
	}
	fmt.Println(height)
}

func TestService_GetHeaderByNumber(t *testing.T) {
	ethContext := ctx.New()
	header, err := srv.ethClient.GetHeaderByNumber(ethContext, 1)
	if err != nil {
		t.Errorf("TestService_GetHeaderByNumber %v", err)
	}
	fmt.Println(header.Hash().Hex())
}

func TestService_GetHeaderByHash(t *testing.T) {
	ethContext := ctx.New()
	hash := common.HexToHash("0x41800b5c3f1717687d85fc9018faac0a6e90b39deaa0b99e7fe4fe796ddeb26a")
	header, err := srv.ethClient.GetHeaderByHash(ethContext, hash)
	if err != nil {
		t.Errorf("TestService_GetHeaderByHash %v", err)
	}
	fmt.Println(header.Number)
}

func TestService_GetBlockByNumber(t *testing.T) {
	ethContext := ctx.New()
	block, err := srv.ethClient.GetBlockByNumber(ethContext, 1)
	if err != nil {
		t.Errorf("TestService_GetHeaderByNumber %v", err)
	}
	fmt.Println(block.Hash().Hex())
}

func TestService_GetBlockByHash(t *testing.T) {
	ethContext := ctx.New()
	hash := common.HexToHash("0x41800b5c3f1717687d85fc9018faac0a6e90b39deaa0b99e7fe4fe796ddeb26a")
	block, err := srv.ethClient.GetBlockByHash(ethContext, hash)
	if err != nil {
		t.Errorf("TestService_GetHeaderByHash %v", err)
	}
	fmt.Println(block.Number())
}
