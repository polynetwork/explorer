package dao

import (
	"encoding/hex"
	"github.com/polynetwork/explorer/internal/model"
	"testing"
)

func TestDao_InsertFChainTx(t *testing.T) {
	chainContract := model.ChainContract{
		Id:       0,
		Contract: "1edC3Bce037B4F367c24e15B6772FCB191fE0f04",
	}
	fctx := &model.FChainTx{}
	fctx.Chain = 6
	fctx.TxHash = "0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285"
	fctx.State = 1
	fctx.Fee = 0
	fctx.TT = uint32(1574307075)
	fctx.Height = 1
	fctx.Key = ""
	fctx.TokenAddress = "0x7dD16c0c71F71A123c4BDAF0a468aBC60Db41C0C"
	fctx.Contract = chainContract.Contract
	fctx.Value = hex.EncodeToString([]byte("123"))
	fctx.TChain = 7
	fctx.Info = "fsdfssdffsd"
	err := d.InsertFChainTx(fctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_InsertTChainTx(t *testing.T) {
	chainContract := model.ChainContract{
		Id:       0,
		Contract: "1edC3Bce037B4F367c24e15B6772FCB191fE0f04",
	}
	tctx := &model.TChainTx{}
	tctx.Chain = 2
	tctx.TxHash = "0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285"
	tctx.State = 1
	tctx.Fee = 0
	tctx.TT = uint32(1574307075)
	tctx.Height = 2
	tctx.FChain = 9
	tctx.Contract = chainContract.Contract
	tctx.RTxHash = "0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285"
	tctx.TokenAddress = "0x7dD16c0c71F71A123c4BDAF0a468aBC60Db41C0C"
	err := d.InsertTChainTx(tctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_InsertMChainTx(t *testing.T) {
	mctx := &model.MChainTx{}
	mctx.Chain = 0
	mctx.TxHash = "0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285"
	mctx.State = 1
	mctx.Fee = 0
	mctx.TT = uint32(1574307075)
	mctx.Height = 1
	mctx.FChain = 1
	mctx.FTxHash = "0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285"
	mctx.TChain = 0
	err := d.InsertMChainTx(mctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_SelectChainTokenByIdAndAddress(t *testing.T) {
	chainToken, err := d.SelectChainTokenByIdAndAddress(1, "0x7dD16c0c71F71A123c4BDAF0a468aBC60Db41C0C")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chainToken %s, name %s", chainToken.Address, chainToken.Name)
}
