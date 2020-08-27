package dao

import (
	"encoding/hex"
	"github.com/polynetwork/explorer/internal/model"
	"testing"
)

func TestDao_InsertFChainTxAndCache(t *testing.T) {
	chainContract := model.ChainContract{
		Id:       0,
		Contract: "1edC3Bce037B4F367c24e15B6772FCB191fE0f04",
	}
	fctx := &model.FChainTx{}
	fctx.Chain = 6
	fctx.TxHash = "9765f91b8b1036a5ae29e9e872b4b937a201098e1317ba7e36261ff7cdc49285"
	fctx.State = 1
	fctx.Fee = 0
	fctx.TT = uint32(1574307075)
	fctx.Height = 1
	fctx.Key = ""
	fctx.TokenAddress = "0de7a8fef8c2740ffcf51d2a0ebafbd5cfefd882"
	fctx.Contract = chainContract.Contract
	fctx.Value = hex.EncodeToString([]byte("123"))
	fctx.TChain = 7
	fctx.Info = "{\"crosstxtype\":1,\"crosstxname\":\"\",\"fromchainid\":2,\"fromchainname\":\"\",\"fromaddress\":\"AGjD4Mo25kzcStyh1stp7tXkUuMopD43NT\",\"tochainid\":3,\"tochainname\":\"\",\"toaddress\":\"AGjD4Mo25kzcStyh1stp7tXkUuMopD43NT\",\"tokenaddress\":\"8f40cdf73ca51abda31e79a7ddd656d11af23e7a\",\"tokenname\":\"\",\"tokentype\":\"\",\"amount\":9}"
	err := d.InsertFChainTxAndCache(fctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_FChainTx(t *testing.T) {
	fChainTx, err := d.FChainTx("9765f91b8b1036a5ae29e9e872b4b937a201098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", fChainTx.Chain, fChainTx.TxHash)
}

func TestDao_InsertMChainTxAndCache(t *testing.T) {
	mctx := &model.MChainTx{}
	mctx.Chain = 0
	mctx.TxHash = "0665f93b8b1036a5ae51e9e871141911ad01098e1317ba7e36261ff7cdc49285"
	mctx.State = 1
	mctx.Fee = 0
	mctx.TT = uint32(1574307075)
	mctx.Height = 1
	mctx.FChain = 1
	mctx.FTxHash = "9765f91b8b1036a5ae29e9e872b4b937a201098e1317ba7e36261ff7cdc49285"
	mctx.TChain = 0
	err := d.InsertMChainTxAndCache(mctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_MChainTx(t *testing.T) {
	mChainTx, err := d.MChainTx("0665f93b8b1036a5ae51e9e871141911ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", mChainTx.Chain, mChainTx.TxHash)
}

func TestDao_MChainTxByFTx(t *testing.T) {
	mChainTx, err := d.MChainTxByFTx("9765f91b8b1036a5ae29e9e872b4b937a201098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", mChainTx.Chain, mChainTx.TxHash)
}

func TestDao_InsertTChainTxAndCache(t *testing.T) {
	chainContract := model.ChainContract{
		Id:       0,
		Contract: "1edC3Bce037B4F367c24e15B6772FCB191fE0f04",
	}
	tctx := &model.TChainTx{}
	tctx.Chain = 2
	tctx.TxHash = "3065f93b8b1036a5ae5919287614b937ad01098e1317ba7e36261ff7cdc49285"
	tctx.State = 1
	tctx.Fee = 0
	tctx.TT = uint32(1574307075)
	tctx.Height = 2
	tctx.FChain = 9
	tctx.Contract = chainContract.Contract
	tctx.RTxHash = "0665f93b8b1036a5ae51e9e871141911ad01098e1317ba7e36261ff7cdc49285"
	tctx.TokenAddress = "0x7dD16c0c71F71A123c4BDAF0a468aBC60Db41C0C"
	err := d.InsertTChainTxAndCache(tctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_TChainTx(t *testing.T) {
	tChainTx, err := d.TChainTx("3065f93b8b1036a5ae5919287614b937ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", tChainTx.Chain, tChainTx.TxHash)
}

func TestDao_TChainTxByMTx(t *testing.T) {
	tChainTx, err := d.TChainTxByMTx("0665f93b8b1036a5ae51e9e871141911ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", tChainTx.Chain, tChainTx.TxHash)
}
