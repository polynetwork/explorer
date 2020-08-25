package dao

import (
	"encoding/hex"
	"github.com/polynetwork/explorer/internal/model"
	"testing"
)

func TestDao_AddFChainTx(t *testing.T) {
	chainContract := model.ChainContract{
		Id:       0,
		Contract: "1edC3Bce037B4F367c24e15B6772FCB191fE0f04",
	}
	fctx := &model.FChainTx{}
	fctx.Chain = 6
	fctx.TxHash = "0x8865f93b8b1036a5ae59e9e476b4b937ad01098e1317ba7e36261ff7cdc49285"
	fctx.State = 1
	fctx.Fee = 0
	fctx.TT = uint32(1574307075)
	fctx.Height = 1
	fctx.Key = ""
	fctx.TokenAddress = "0x7dD16c0c71F71A223c4BDAF0a468aBC60Db41C0C"
	fctx.Contract = chainContract.Contract
	fctx.Value = hex.EncodeToString([]byte("123"))
	fctx.TChain = 7
	fctx.Info = "fsdfssdffsd"

	err := d.AddFChainTx(fctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_CacheFChainTx(t *testing.T) {
	fChainTx, err := d.CacheFChainTx("0x8865f93b8b1036a5ae59e9e476b4b937ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", fChainTx.Chain, fChainTx.TxHash)
}

func TestDao_AddMChainTx(t *testing.T) {
	mctx := &model.MChainTx{}
	mctx.Chain = 0
	mctx.TxHash = "0x8865f93b8b1036a5ae59e9e873b4b937ad01098e1317ba7e36261ff7cdc49285"
	mctx.State = 1
	mctx.Fee = 0
	mctx.TT = uint32(1574307075)
	mctx.Height = 1
	mctx.FChain = 1
	mctx.FTxHash = "0x8865f93b8b1036a5ae59e9e876b41937ad01098e1317ba7e36261ff7cdc49285"
	mctx.TChain = 0
	err := d.AddMChainTx(mctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_AddMChainTxByFTx(t *testing.T) {
	mctx := &model.MChainTx{}
	mctx.Chain = 0
	mctx.TxHash = "0x8865f93b8b1036a5ae59e9e873b4b937ad01098e1317ba7e36261ff7cdc49285"
	mctx.State = 1
	mctx.Fee = 0
	mctx.TT = uint32(1574307075)
	mctx.Height = 1
	mctx.FChain = 1
	mctx.FTxHash = "0x8865f93b8b1036a5ae59e9e876b41937ad01098e1317ba7e36261ff7cdc49285"
	mctx.TChain = 0
	err := d.AddMChainTxByFTx(mctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_CacheMChainTx(t *testing.T) {
	mChainTx, err := d.CacheMChainTx("0x8865f93b8b1036a5ae59e9e873b4b937ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", mChainTx.Chain, mChainTx.TxHash)
}

func TestDao_CacheMChainTxByFTx(t *testing.T) {
	mChainTx, err := d.CacheMChainTxByFTx("0x8865f93b8b1036a5ae59e9e876b41937ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", mChainTx.Chain, mChainTx.TxHash)
}

func TestDao_AddTChainTx(t *testing.T) {
	chainContract := model.ChainContract{
		Id:       0,
		Contract: "1edC3Bce037B4F367c24e15B6772FCB191fE0f04",
	}
	tctx := &model.TChainTx{}
	tctx.Chain = 2
	tctx.TxHash = "0x8865f93b8b1036a5ae59e91876b4b937ad01098e1317ba7e36261ff7cdc49285"
	tctx.State = 1
	tctx.Fee = 0
	tctx.TT = uint32(1574307075)
	tctx.Height = 2
	tctx.FChain = 9
	tctx.Contract = chainContract.Contract
	tctx.RTxHash = "0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285"
	tctx.TokenAddress = "0x7dD16c0c71F71A123c4BDAF0a468aBC60Db41C0C"

	err := d.AddTChainTx(tctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_AddTChainTxByMTx(t *testing.T) {
	chainContract := model.ChainContract{
		Id:       0,
		Contract: "1edC3Bce037B4F367c24e15B6772FCB191fE0f04",
	}
	tctx := &model.TChainTx{}
	tctx.Chain = 2
	tctx.TxHash = "0x8865f93b8b1036a5ae59e91876b4b937ad01098e1317ba7e36261ff7cdc49285"
	tctx.State = 1
	tctx.Fee = 0
	tctx.TT = uint32(1574307075)
	tctx.Height = 2
	tctx.FChain = 9
	tctx.Contract = chainContract.Contract
	tctx.RTxHash = "0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285"
	tctx.TokenAddress = "0x7dD16c0c71F71A123c4BDAF0a468aBC60Db41C0C"
	err := d.AddTChainTxByMTx(tctx)
	if err != nil {
		t.FailNow()
	}
}

func TestDao_CacheTChainTx(t *testing.T) {
	tChainTx, err := d.CacheTChainTx("0x8865f93b8b1036a5ae59e91876b4b937ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", tChainTx.Chain, tChainTx.TxHash)
}

func TestDao_CacheTChainTxByMTx(t *testing.T) {
	tChainTx, err := d.CacheTChainTxByMTx("0x8865f93b8b1036a5ae59e9e876b4b937ad01098e1317ba7e36261ff7cdc49285")
	if err != nil {
		t.FailNow()
	}
	t.Logf("chain id %d, txHash %s", tChainTx.Chain, tChainTx.TxHash)
}
