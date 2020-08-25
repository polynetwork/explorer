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

package dao

import (
	"database/sql"
	"encoding/json"
	"github.com/polynetwork/explorer/internal/common"
	"github.com/polynetwork/explorer/internal/model"
)

const (
	_insertTChainTx                 = "insert into tchain_tx(chain_id, txhash, state, tt, fee, height, fchain, contract, rtxhash) values (?,?,?,?,?,?,?,?,?)"
	_insertTChainTransfer          = "insert into tchain_transfer(txhash, asset, xfrom, xto, amount) values(?,?,?,?,?)"
	_insertFChainTx                 = "insert into fchain_tx(chain_id, txhash, state, tt, fee, height, xuser, tchain, contract, xkey, xparam) values (?,?,?,?,?,?,?,?,?,?,?)"
	_insertFChainTransfer          = "insert into fchain_transfer(txhash, asset, xfrom, xto, amount, tochainid, toasset, touser) values (?,?,?,?,?,?,?,?)"
	_insertMChainTx                 = "insert into mchain_tx(chain_id, txhash, state, tt, fee, height, fchain, ftxhash, tchain, xkey) values (?,?,?,?,?,?,?,?,?,?)"
	_selectMChainTxCount            = "select count(*) from mchain_tx"
	_selectMChainTxByLimit          = "select A.chain_id, A.txhash, case when B.txhash is null OR C.txhash is null THEN 0 ELSE 1 END as state, A.tt, A.fee, A.height, A.fchain, A.tchain from mchain_tx A left join tchain_tx B on A.txhash = B.rtxhash left join fchain_tx C on A.ftxhash = C.txhash order by A.height desc limit ?,?;"
	_selectMChainTxByHash           = "select chain_id, txhash, state, tt, fee, height, fchain, ftxhash, tchain, xkey from mchain_tx where txhash = ?"
	_selectMChainTxByFHash          = "select chain_id, txhash, state, tt, fee, height, fchain, ftxhash, tchain, xkey from mchain_tx where ftxhash = ?"
	_selectFChainTxByHash           = "select chain_id, txhash, state, tt, fee, height, xuser, tchain, contract, xkey, xparam from fchain_tx where case when chain_id = ? then xkey = ? else txhash = ? end"
	_selectFChainTxByTime           = "select unix_timestamp(FROM_UNIXTIME(tt,'%Y%m%d')) days, count(*) from fchain_tx where chain_id = ? and tt > ? and tt < ? group by chain_id,days order by days desc"
	_selectFChainTransferByHash    = "select txhash, asset, xfrom, xto, amount, tochainid, toasset, touser from fchain_transfer where txhash = ?" // TODO
	_selectTChainTxByHash           = "select chain_id, txhash, state, tt, fee, height, fchain, contract,rtxhash from tchain_tx where txhash = ?"
	_selectTChainTxByMHash          = "select chain_id, txhash, state, tt, fee, height, fchain, contract,rtxhash from tchain_tx where rtxhash = ?"
	_selectTChainTxByTime           = "select unix_timestamp(FROM_UNIXTIME(tt,'%Y%m%d')) days, count(*) from tchain_tx where chain_id = ? and tt > ? and tt < ? group by chain_id,days order by days desc"
	_selectTChainTransferByHash    = "select txhash, asset, xfrom, xto, amount from tchain_transfer where txhash = ?"
	_selectChainAddresses           = "select count(distinct a) from (select distinct xfrom as a from fchain_transfer aa left join fchain_tx bb on aa.txhash = bb.txhash where bb.chain_id = ? union all select distinct xto as a from tchain_transfer aa left join tchain_tx bb on aa.txhash = bb.txhash where bb.chain_id = ?) c"
	_selectChainInfoById            = "select xname, id, url, xtype, height, txin, txout from chain_info where id = ?"
	_selectAllChainInfos            = "select xname, id, url, xtype, height, txin, txout from chain_info order by id"
	_selectContractById             = "select id, contract from chain_contract where id = ?"
	_selectTokenById                = "select id, xtoken, hash, xname, xtype, xprecision, xdesc from chain_token where id = ?"
	_selectTokenCount               = "select count(distinct xtoken) from chain_token"
	_updateChainInfoById            = "update chain_info set xname = ?, url = ?, height = ?, txin = ?, txout = ? where id = ?"
	_selectAllianceTx               = "select chain_id, txhash, state, tt, fee, height, fchain, ftxhash, tchain,xkey from mchain_tx where (tchain = ? or fchain = ?) and height > ? order by height"
	_selectBitcoinUnconfirmTx       = "select txhash from tchain_tx where chain_id = ? and tt = ?"
	_updateBitcoinConfirmTx         = "update tchain_tx set tt = ?, height = ?, state = 1, fee = ? where txhash = ?"
	_updateBitcoinConfirmTransfer  = "update tchain_transfer set xto = ?, amount = ? where txhash = ?"
	_selectTokenTxList              = "select case when b.chain_id = ? then b.xkey else b.txhash end as txhash,a.xfrom,a.xto,a.amount,b.height,b.tt,1 as direct from fchain_transfer a left join fchain_tx b on a.txhash = b.txhash where a.asset = ? union all select d.txhash,c.xfrom,c.xto,c.amount,d.height,d.tt,2 as direct from tchain_transfer c left join tchain_tx d on c.txhash = d.txhash where c.asset = ? order by height desc limit ?,?;"
	_selectTokenTxTotal             = "select sum(cnt) from (select count(*) as cnt from fchain_transfer a left join fchain_tx b on a.txhash = b.txhash where a.asset = ? union all select count(*) as cnt from tchain_transfer c left join tchain_tx d on c.txhash = d.txhash where c.asset = ?) t"
	_selectAddressTxList            = "select case when b.chain_id = ? then b.xkey else b.txhash end as txhash,a.xfrom,a.xto,a.asset,a.amount,b.height,b.tt,1 as direct from fchain_transfer a left join fchain_tx b on a.txhash = b.txhash where a.xfrom = ? and b.chain_id = ? union all select d.txhash,c.xfrom,c.xto,c.asset,c.amount,d.height,d.tt,2 as direct from tchain_transfer c left join tchain_tx d on c.txhash = d.txhash where c.xto = ? and d.chain_id = ? order by height desc limit ?,?;"
	_selectAddressTxTotal           = "select sum(cnt) from (select count(*) as cnt from fchain_transfer a left join fchain_tx b on a.txhash = b.txhash where a.xfrom = ? and b.chain_id = ? union all select count(*) as cnt from tchain_transfer c left join tchain_tx d on c.txhash = d.txhash where c.xto = ? and d.chain_id = ?) t"
	_insertPolyValidator            = "insert into poly_validators(height, validators) values(?,?)"
	_selectPolyValidator            = "select height, validators from poly_validators order by height desc limit 1"
)

func (d *Dao) InsertTChainTx(t *model.TChainTx) (err error) {
	if _, err = d.db.Exec(_insertTChainTx, t.Chain, t.TxHash, t.State, t.TT, t.Fee, t.Height, t.FChain, t.Contract, t.RTxHash); err != nil {
		return
	}
	return
}

func (d *Dao) InsertFChainTx(f *model.FChainTx) (err error) {
	if _, err = d.db.Exec(_insertFChainTx, f.Chain, f.TxHash, f.State, f.TT, f.Fee, f.Height, f.User, f.TChain, f.Contract, f.Key, f.Param); err != nil {
		return
	}
	return
}

func (d *Dao) InsertMChainTx(m *model.MChainTx) (err error) {
	if _, err = d.db.Exec(_insertMChainTx, m.Chain, m.TxHash, m.State, m.TT, m.Fee, m.Height, m.FChain, m.FTxHash, m.TChain, m.Key); err != nil {
		return
	}
	return
}

func (d *Dao) TxInsertTChainTx(tx *sql.Tx, t *model.TChainTx) (err error) {
	if _, err = tx.Exec(_insertTChainTx, t.Chain, t.TxHash, t.State, t.TT, t.Fee, t.Height, t.FChain, t.Contract, t.RTxHash); err != nil {
		return
	}
	if t.Transfer != nil && t.Transfer.TxHash != "" {
		transfer := t.Transfer
		if _, err = tx.Exec(_insertTChainTransfer, transfer.TxHash, transfer.Asset, transfer.From, transfer.To, transfer.Amount); err != nil {
			return
		}
	}
	return
}

func (d *Dao) TxInsertFChainTx(tx *sql.Tx, f *model.FChainTx) (err error) {
	if _, err = tx.Exec(_insertFChainTx, f.Chain, f.TxHash, f.State, f.TT, f.Fee, f.Height, f.User, f.TChain, f.Contract, f.Key, f.Param); err != nil {
		return
	}
	if f.Transfer != nil && f.Transfer.TxHash != "" {
		transfer := f.Transfer
		if _, err = tx.Exec(_insertFChainTransfer, transfer.TxHash,transfer.Asset,transfer.From,transfer.To,transfer.Amount,transfer.ToChain,transfer.ToAsset,transfer.ToUser); err != nil {
			return
		}
	}
	return
}

func (d *Dao) TxInsertMChainTx(tx *sql.Tx, m *model.MChainTx) (err error) {
	if _, err = tx.Exec(_insertMChainTx, m.Chain, m.TxHash, m.State, m.TT, m.Fee, m.Height, m.FChain, m.FTxHash, m.TChain, m.Key); err != nil {
		return
	}
	return
}

func (d *Dao) SelectChainInfoById(id uint32) (res *model.ChainInfo, err error) {
	res = new(model.ChainInfo)
	row := d.db.QueryRow(_selectChainInfoById, id)
	if err = row.Scan(&res.Name, &res.Id, &res.Url, &res.XType, &res.Height, &res.In, &res.Out); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		res = nil
	}
	return
}

func (d *Dao) SelectAllChainInfos() (c []*model.ChainInfo, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAllChainInfos); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.ChainInfo)
		if err = rows.Scan(&r.Name, &r.Id, &r.Url, &r.XType, &r.Height, &r.In, &r.Out); err != nil {
			c = nil
			return
		}
		c = append(c, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectContractById(id uint32) (c []*model.ChainContract, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectContractById, id); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.ChainContract)
		if err = rows.Scan(&r.Id, &r.Contract); err != nil {
			c = nil
			return
		}
		c = append(c, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectTokenById(id uint32) (c []*model.ChainToken, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectTokenById, id); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.ChainToken)
		if err = rows.Scan(&r.Id, &r.Token, &r.Hash, &r.Name, &r.Type, &r.Precision, &r.Desc); err != nil {
			c = nil
			return
		}
		c = append(c, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectTokenCount() (count *uint32, err error) {
	count = new(uint32)
	row := d.db.QueryRow(_selectTokenCount)
	if err = row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		*count = uint32(0)
	}
	return
}

func (d *Dao) SelectMChainTxCount() (count *uint32, err error) {
	count = new(uint32)
	row := d.db.QueryRow(_selectMChainTxCount)
	if err = row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		*count = uint32(0)
	}
	return
}

func (d *Dao) SelectMChainTxByHash(hash string) (res *model.MChainTx, err error) {
	res = new(model.MChainTx)
	row := d.db.QueryRow(_selectMChainTxByHash, hash)
	if err = row.Scan(&res.Chain, &res.TxHash, &res.State, &res.TT, &res.Fee, &res.Height, &res.FChain, &res.FTxHash, &res.TChain, &res.Key); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		res = nil
	}
	return
}

func (d *Dao) SelectMChainTxByFHash(hash string) (res *model.MChainTx, err error) {
	res = new(model.MChainTx)
	row := d.db.QueryRow(_selectMChainTxByFHash, hash)
	if err = row.Scan(&res.Chain, &res.TxHash, &res.State, &res.TT, &res.Fee, &res.Height, &res.FChain, &res.FTxHash, &res.TChain, &res.Key); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		res = nil
	}
	return
}

func (d *Dao) SelectMChainTxByLimit(start int, limit int) (res []*model.MChainTx, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectMChainTxByLimit, start, limit); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.MChainTx)
		if err = rows.Scan(&r.Chain, &r.TxHash, &r.State, &r.TT, &r.Fee, &r.Height, &r.FChain, &r.TChain); err != nil {
			rows = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectFChainTxByHash(hash string, chain uint32) (res *model.FChainTx, err error) {
	res = new(model.FChainTx)
	res.Transfer = new(model.FChainTransfer)
	{
		row := d.db.QueryRow(_selectFChainTxByHash, chain, hash, hash)
		if err = row.Scan(&res.Chain, &res.TxHash, &res.State, &res.TT, &res.Fee, &res.Height, &res.User, &res.TChain, &res.Contract, &res.Key, &res.Param); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			res = nil
			return
		}
	}
	{
		row := d.db.QueryRow(_selectFChainTransferByHash, hash)
		if err = row.Scan(&res.Transfer.TxHash, &res.Transfer.Asset, &res.Transfer.From, &res.Transfer.To, &res.Transfer.Amount, &res.Transfer.ToChain, &res.Transfer.ToAsset, &res.Transfer.ToUser); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}
	}
	return
}

func (d *Dao) SelectFChainTxByTime(chainId uint32, start uint32, end uint32) (res []*model.CrossChainTxStatus, err error) {
	var rows *sql.Rows
	if rows, err  = d.db.Query(_selectFChainTxByTime, chainId, start, end); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		r := new(model.CrossChainTxStatus)
		if err = rows.Scan(&r.TT, &r.TxNumber); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectTChainTxByHash(hash string) (res *model.TChainTx, err error) {
	res = new(model.TChainTx)
	res.Transfer = new(model.TChainTransfer)
	{
		row := d.db.QueryRow(_selectTChainTxByHash, hash)
		if err = row.Scan(&res.Chain, &res.TxHash, &res.State, &res.TT, &res.Fee, &res.Height, &res.FChain, &res.Contract, &res.RTxHash); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			res = nil
			return
		}
		{
			row := d.db.QueryRow(_selectTChainTransferByHash, hash)
			if err = row.Scan(&res.Transfer.TxHash, &res.Transfer.Asset, &res.Transfer.From, &res.Transfer.To, &res.Transfer.Amount); err != nil {
				if err == sql.ErrNoRows {
					err = nil
				}
			}
		}
	}
	return
}

func (d *Dao) SelectTChainTxByMHash(hash string) (res *model.TChainTx, err error) {
	res = new(model.TChainTx)
	res.Transfer = new(model.TChainTransfer)
	{
		row := d.db.QueryRow(_selectTChainTxByMHash, hash)
		if err = row.Scan(&res.Chain, &res.TxHash, &res.State, &res.TT, &res.Fee, &res.Height, &res.FChain, &res.Contract, &res.RTxHash); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			res = nil
			return
		}
	}
	{
		row := d.db.QueryRow(_selectTChainTransferByHash, res.TxHash)
		if err = row.Scan(&res.Transfer.TxHash, &res.Transfer.Asset, &res.Transfer.From, &res.Transfer.To, &res.Transfer.Amount); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}
	}
	return
}

func (d *Dao) SelectTChainTxByTime(chainId uint32, start uint32, end uint32) (res []*model.CrossChainTxStatus, err error) {
	var rows *sql.Rows
	if rows, err  = d.db.Query(_selectTChainTxByTime, chainId, start, end); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		r := new(model.CrossChainTxStatus)
		if err = rows.Scan(&r.TT, &r.TxNumber); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}

/*
func (d *Dao) UpdateChainInfoById(c *model.ChainInfo) (err error) {
	if _, err = d.db.Exec(_updateChainInfoById, c.Name, c.Url, c.Height, c.In, c.Out, c.Id); err != nil {
		return
	}
	return
}
*/

func (d *Dao) TxUpdateChainInfoById(tx *sql.Tx, c *model.ChainInfo) (err error) {
	if _, err = tx.Exec(_updateChainInfoById, c.Name, c.Url, c.Height, c.In, c.Out, c.Id); err != nil {
		return
	}
	return
}

func (d *Dao) SelectChainAddresses(chainId uint32) (uint32, error) {
	row := d.db.QueryRow(_selectChainAddresses, chainId, chainId)
	var counter uint32
	if err := row.Scan(&counter); err != nil {
		return 0, err
	}
	return counter, nil
}

/*
func (d *Dao) SelectChainTokenByIdAndAddress(chainId uint32, address string) (res *model.ChainToken, err error) {
	res = new(model.ChainToken)
	row := d.db.QueryRow(_selectChainTokenByIdAndAddress, chainId, address)
	if err = row.Scan(&res.Id, &res.Hash, &res.Name, &res.Type, &res.Precision, &res.Desc); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		res = nil
	}
	return
}
*/

func (d *Dao) SelectAllianceTx(height uint32, chain uint32) (res []*model.MChainTx, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAllianceTx, chain, chain, height); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.MChainTx)
		if err = rows.Scan(&r.Chain, &r.TxHash, &r.State, &r.TT, &r.Fee, &r.Height, &r.FChain, &r.FTxHash, &r.TChain, &r.Key); err != nil {
			rows = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectBitcoinTxUnConfirm(id uint32) (res []string, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectBitcoinUnconfirmTx, id, 0); err != nil {
		return
	}
	defer rows.Close()

	var txhash string
	txhashs := make([]string, 0)
	for rows.Next() {
		if err = rows.Scan(&txhash); err != nil {
			rows = nil
			return nil, err
		}
		txhashs = append(txhashs, txhash)
	}
	return txhashs, nil
}

func (d *Dao) UpdateBitcoinTxConfirmed(txhash string, height uint32, tt uint32, gas uint64, toaddress string, amount uint64) (err error) {
	if _, err = d.db.Exec(_updateBitcoinConfirmTx, tt, height, gas, txhash); err != nil {
		return
	}
	if _, err = d.db.Exec(_updateBitcoinConfirmTransfer, toaddress, amount, txhash); err != nil {
		return
	}
	return
}

func (d *Dao) SelectTokenTxList(token string, start uint32, end uint32) (res []*model.TokenTx, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectTokenTxList, common.CHAIN_ETH, token, token, start, end); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.TokenTx)
		if err = rows.Scan(&r.TxHash, &r.From, &r.To, &r.Amount, &r.Height, &r.TT, &r.Direct); err != nil {
			rows = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectTokenTxTotal(token string) (count *uint32, err error) {
	count = new(uint32)
	row := d.db.QueryRow(_selectTokenTxTotal, token, token)
	if err = row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		*count = uint32(0)
	}
	return
}

func (d *Dao) SelectAddressTxList(chainId uint32, addr string, start uint32, end uint32) (res []*model.AddressTx, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAddressTxList, common.CHAIN_ETH, addr, chainId, addr, chainId, start, end); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.AddressTx)
		if err = rows.Scan(&r.TxHash, &r.From, &r.To, &r.Asset, &r.Amount, &r.Height, &r.TT, &r.Direct); err != nil {
			rows = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}

func (d *Dao) SelectAddressTxTotal(chainId uint32, addr string) (count *uint32, err error) {
	count = new(uint32)
	row := d.db.QueryRow(_selectAddressTxTotal, addr, chainId, addr, chainId)
	if err = row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		*count = uint32(0)
	}
	return
}

func (d *Dao) InsertPolyValidator(height uint32, validators []string) (err error) {
	validators_json, err := json.Marshal(validators)
	if err != nil {
		return
	}
	if _, err = d.db.Exec(_insertPolyValidator, height, string(validators_json)); err != nil {
		return
	}
	return
}

func (d *Dao) SelectPolyValidator() (validator []string, err error) {
	row := d.db.QueryRow(_selectPolyValidator)
	var height uint32
	var validator_json string
	if err = row.Scan(&height, &validator_json); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return nil, err
	}

	validator = make([]string, 0)
	err = json.Unmarshal([]byte(validator_json), &validator)
	if err != nil {
		return nil, err
	}
	return validator, nil
}
