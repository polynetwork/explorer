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
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	"math/big"
)

const (
	_insertTChainTx                 = "insert into tchain_tx(chain_id, txhash, state, tt, fee, height, fchain, contract, rtxhash) values (?,?,?,?,?,?,?,?,?)"
	_insertTChainTransfer          = "insert into tchain_transfer(txhash, chain_id, tt, asset, xfrom, xto, amount) values(?,?,?,?,?,?,?)"
	_insertFChainTx                 = "insert into fchain_tx(chain_id, txhash, state, tt, fee, height, xuser, tchain, contract, xkey, xparam) values (?,?,?,?,?,?,?,?,?,?,?)"
	_insertFChainTransfer          = "insert into fchain_transfer(txhash, chain_id, tt, asset, xfrom, xto, amount, tochainid, toasset, touser) values (?,?,?,?,?,?,?,?,?,?)"
	_insertMChainTx                 = "insert into mchain_tx(chain_id, txhash, state, tt, fee, height, fchain, ftxhash, tchain, xkey) values (?,?,?,?,?,?,?,?,?,?)"
	_selectMChainTxCount            = "select count(*) from mchain_tx"
	//_selectMChainTxByLimit          = "select A.chain_id, A.txhash, case when B.txhash is null OR C.txhash is null THEN 0 ELSE 1 END as state, A.tt, A.fee, A.height, A.fchain, A.tchain from mchain_tx A left join tchain_tx B on A.txhash = B.rtxhash left join fchain_tx C on A.ftxhash = C.txhash order by A.height desc limit ?,?;"
	_selectMChainTxByLimit          = "select C.chain_id, C.txhash, case when D.txhash is null THEN 0 ELSE 1 END as state, C.tt, C.fee, C.height, C.fchain, C.tchain from ((select B.chain_id, B.txhash, B.tt, B.fee, B.height, B.fchain, B.tchain from (select txhash, chain_Id, tt, height, tchain from fchain_tx order by tt desc limit ?, ?) A LEFT JOIN mchain_tx B ON B.ftxhash = A.txhash where B.txhash is not null) union all (select A.chain_id, case when A.chain_id = ? then A.xkey else A.txhash end as txhash, A.tt, A.fee, A.height, A.chain_id as fchain, A.tchain from (select txhash, xkey, chain_Id, tt, fee, height, tchain from fchain_tx order by tt desc limit ?, ?) A LEFT JOIN mchain_tx B ON B.ftxhash = A.txhash where B.txhash is null)) C LEFT JOIN tchain_tx D on C.txhash = D.rtxhash order by C.tt desc"
	_selectMChainTxByHash           = "select chain_id, txhash, state, tt, fee, height, fchain, ftxhash, tchain, xkey from mchain_tx where txhash = ?"
	_selectMChainTxByFHash          = "select chain_id, txhash, state, tt, fee, height, fchain, ftxhash, tchain, xkey from mchain_tx where ftxhash = ?"
	_selectFChainTxByHash           = "select chain_id, txhash, state, tt, fee, height, xuser, tchain, contract, xkey, xparam from fchain_tx where case when chain_id = ? then xkey = ? else txhash = ? end"
	_selectFChainTxByTime           = "select chain_id, unix_timestamp(FROM_UNIXTIME(tt,'%Y%m%d')) days, count(*) from fchain_tx where tt > ? and tt < ? group by chain_id, days order by chain_id, days desc"
	_selectFChainTransferByHash    = "select txhash, asset, xfrom, xto, amount, tochainid, toasset, touser from fchain_transfer where txhash = ?"
	_selectTChainTxByHash           = "select chain_id, txhash, state, tt, fee, height, fchain, contract,rtxhash from tchain_tx where txhash = ?"
	_selectTChainTxByMHash          = "select chain_id, txhash, state, tt, fee, height, fchain, contract,rtxhash from tchain_tx where rtxhash = ?"
	_selectTChainTxByTime           = "select chain_id, unix_timestamp(FROM_UNIXTIME(tt,'%Y%m%d')) days, count(*) from tchain_tx where tt > ? and tt < ? group by chain_id, days order by chain_id, days desc"
	_selectTChainTransferByHash    = "select txhash, asset, xfrom, xto, amount from tchain_transfer where txhash = ?"
	_selectChainAddresses           = "select chain_id, count(distinct addr) from (select chain_id,xfrom as addr from fchain_transfer union all select chain_id,xto as addr from tchain_transfer) c group by c.chain_id"
	_selectAllChainInfos            = "select xname, id, xtype, height, txin, txout from chain_info order by id"
	_selectContractById             = "select id, contract from chain_contract where id = ?"
	_selectTokenById                = "select id, xtoken, hash, xname, xtype, xprecision, xdesc from chain_token where id = ?"
	_selectTokenCount               = "select count(distinct xtoken) from chain_token"
	_updateChainInfoById            = "update chain_info set xname = ?, height = ?, txin = ?, txout = ? where id = ?"
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
	_insertAssetStatistic           = "insert into asset_statistic(xname, addressnum, amount,amount_btc, amount_usd, txnum, latestupdate) values (?,0,0,0,0,0,0) ON DUPLICATE KEY UPDATE txnum=txnum"
	_selectAssetAddressNum          =  "select token, count(distinct addr) as addrNum from (select B.xtoken as token, A.xfrom as addr from fchain_transfer as A inner join chain_token as B on A.asset = B.hash union all select D.xtoken as token, C.xto as addr from tchain_transfer as C inner join chain_token as D on C.asset = D.hash) E group by E.token"
	_selectAssetTxInfo              =  "select B.xtoken, amount from fchain_transfer as A inner join chain_token as B on A.asset = B.hash where A.tt >= ? and A.tt < ? order by B.xtoken"
	_updateStatistic                =  "update asset_statistic set addressnum = ?, amount = amount + ?, amount_btc = amount_btc + ?, amount_usd = amount_usd + ?, txnum = txnum + ?, latestupdate = ? where xname = ? and latestupdate = ?"
	_selectAssetStatistic           =  "select xname, addressnum, amount, amount_btc, amount_usd, txnum, latestupdate from asset_statistic where latestupdate < ? order by amount_usd desc"
	_selectAssetHistory             =  "select A.amount from fchain_transfer as A inner join chain_token as B on A.asset = B.hash where A.tt >= ? and A.tt < ? and B.xtoken = ?"
	_insertTransferStatistic        =  "insert into transfer_statistic(asset, amount, latestin, latestout) values(?, 0, 0, 0) ON DUPLICATE KEY UPDATE amount=amount"
	_selectTransferStatistic        =  "select asset, amount, latestin, latestout from transfer_statistic"
	_selectTransferOutHistory       =  "select amount, tt from fchain_transfer where asset = ? and tt > ?"
	_selectTransferInHistory        =  "select amount, tt from tchain_transfer where asset = ? and tt > ?"
	_updateTransferStatistic        =  "update transfer_statistic set amount = ?, latestin = ?, latestout = ? where asset = ?"
	_selectAllTransferStatistic     =  "select A.asset, A.amount, B.xtoken, B.id, B.xname, D.xname, D.id  from transfer_statistic A left join chain_token B on A.asset = B.hash left join chain_token_bind C on A.asset = C.hash_src left join chain_token D on C.hash_dest = D.hash"
	_selectAllTransferStatisticInChain     =  "select A.asset, A.amount, B.xtoken, B.id, B.xname, D.xname, D.id  from transfer_statistic A left join chain_token B on A.asset = B.hash left join chain_token_bind C on A.asset = C.hash_src left join chain_token D on C.hash_dest = D.hash where B.id = ?"
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
		if _, err = tx.Exec(_insertTChainTransfer, transfer.TxHash, t.Chain, t.TT, transfer.Asset, transfer.From, transfer.To, transfer.Amount.String()); err != nil {
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
		if _, err = tx.Exec(_insertFChainTransfer,transfer.TxHash,f.Chain, f.TT,transfer.Asset,transfer.From,transfer.To,transfer.Amount.String(), transfer.ToChain,transfer.ToAsset,transfer.ToUser); err != nil {
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


func (d *Dao) SelectAllChainInfos() (c []*model.ChainInfo, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAllChainInfos); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.ChainInfo)
		if err = rows.Scan(&r.Name, &r.Id, &r.XType, &r.Height, &r.In, &r.Out); err != nil {
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

func (d *Dao) SelectMChainTxByLimit(start int, limit int, chain uint32) (res []*model.MChainTx, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectMChainTxByLimit, start, limit, chain, start, limit); err != nil {
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
		row := d.db.QueryRow(_selectFChainTransferByHash, res.TxHash)
		var amount string
		if err = row.Scan(&res.Transfer.TxHash, &res.Transfer.Asset, &res.Transfer.From, &res.Transfer.To, &amount, &res.Transfer.ToChain, &res.Transfer.ToAsset, &res.Transfer.ToUser); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}
		res.Transfer.Amount, _ = new(big.Int).SetString(amount, 10)
	}
	return
}

func (d *Dao) SelectFChainTxByTime(start uint32, end uint32) (res []*model.CrossChainTxStatus, err error) {
	var rows *sql.Rows
	if rows, err  = d.db.Query(_selectFChainTxByTime, start, end); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		r := new(model.CrossChainTxStatus)
		if err = rows.Scan(&r.Id, &r.TT, &r.TxNumber); err != nil {
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
			var amount string
			if err = row.Scan(&res.Transfer.TxHash, &res.Transfer.Asset, &res.Transfer.From, &res.Transfer.To, &amount); err != nil {
				if err == sql.ErrNoRows {
					err = nil
				}
			}
			res.Transfer.Amount, _ = new(big.Int).SetString(amount, 10)
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
		var amount string
		if err = row.Scan(&res.Transfer.TxHash, &res.Transfer.Asset, &res.Transfer.From, &res.Transfer.To, &amount); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}
		res.Transfer.Amount, _ = new(big.Int).SetString(amount, 10)
	}
	return
}

func (d *Dao) SelectTChainTxByTime(start uint32, end uint32) (res []*model.CrossChainTxStatus, err error) {
	var rows *sql.Rows
	if rows, err  = d.db.Query(_selectTChainTxByTime, start, end); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		r := new(model.CrossChainTxStatus)
		if err = rows.Scan(&r.Id, &r.TT, &r.TxNumber); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}


func (d *Dao) TxUpdateChainInfoById(tx *sql.Tx, c *model.ChainInfo) (err error) {
	if _, err = tx.Exec(_updateChainInfoById, c.Name, c.Height, c.In, c.Out, c.Id); err != nil {
		return
	}
	return
}

func (d *Dao) SelectChainAddressNum() (res []*model.CrossChainAddressNum, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectChainAddresses); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.CrossChainAddressNum)
		if err = rows.Scan(&r.Id, &r.AddNum); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	return res, nil
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

func (d *Dao) UpdateBitcoinTxConfirmed(txhash string, height uint32, tt uint32, gas uint64, toaddress string, amount *big.Int) (err error) {
	if _, err = d.db.Exec(_updateBitcoinConfirmTx, tt, height, gas, txhash); err != nil {
		return
	}
	if _, err = d.db.Exec(_updateBitcoinConfirmTransfer, toaddress, amount.String(), txhash); err != nil {
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
		var amount string
		if err = rows.Scan(&r.TxHash, &r.From, &r.To, &amount, &r.Height, &r.TT, &r.Direct); err != nil {
			rows = nil
			return
		}
		r.Amount, _ = new(big.Int).SetString(amount, 10)
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
		var amount string
		if err = rows.Scan(&r.TxHash, &r.From, &r.To, &r.Asset, &amount, &r.Height, &r.TT, &r.Direct); err != nil {
			rows = nil
			return
		}
		r.Amount, _ = new(big.Int).SetString(amount, 10)
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

func (d *Dao) InsertAssetStatistic(name string) (err error) {
	if _, err = d.db.Exec(_insertAssetStatistic, name); err != nil {
		return
	}
	return
}

func (d *Dao) SelectAssetAddressNum()  (res []*model.AssetAddressNum, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAssetAddressNum); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.AssetAddressNum)
		if err = rows.Scan(&r.Name, &r.AddNum); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	return res, nil
}

type AssetTxInfo2 struct {
	Name string
	Amount string
}

func (d *Dao) SelectAssetTxInfo1(start uint32, end uint32)  (res []*AssetTxInfo2, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAssetTxInfo, start, end); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(AssetTxInfo2)
		if err = rows.Scan(&r.Name, &r.Amount); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	return res, nil
}

func (d *Dao) SelectAssetTxInfo(start uint32, end uint32)  (res []*model.AssetTxInfo, err error) {
	oriRes, err := d.SelectAssetTxInfo1(start, end)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	r := new(model.AssetTxInfo)
	r.Name = ""
	r.Amount = big.NewInt(0)
	r.TxNum = 0
	for _, item := range oriRes {
		if item.Name != r.Name {
			if r.Name != "" && r.TxNum != 0 {
				res = append(res, r)
				r = new(model.AssetTxInfo)
				r.Name = item.Name
				r.Amount = big.NewInt(0)
				r.TxNum = 0
			}
		}
		r.Name = item.Name
		r.TxNum ++

		amount := big.NewInt(0)
		amount, ok := new(big.Int).SetString(item.Amount, 10)
		if !ok {
			amount = big.NewInt(0)
		}
		r.Amount = new(big.Int).Add(r.Amount, amount)
	}
	if r.Name != "" && r.TxNum != 0 {
		res = append(res, r)
	}
	return res, nil
}

func (d *Dao) UpdateAssetStatistics(assetStatistics []*model.AssetStatistic, tt uint32) (err error) {
	for _, statistic := range assetStatistics {
		if statistic.LatestUpdate == 1 {
			continue
		}
		err := d.UpdateAssetStatistic(statistic, tt)
		if err != nil {
			log.Errorf("UpdateAssetStatistic err: %s", err.Error())
		}
	}
	return nil
}

func (d *Dao) UpdateAssetStatistic(assetStatistic *model.AssetStatistic, tt uint32) (err error) {
	if _, err = d.db.Exec(_updateStatistic, assetStatistic.Addressnum, assetStatistic.Amount.Int64(), assetStatistic.Amount_btc.Int64(), assetStatistic.Amount_usd.Int64(),assetStatistic.TxNum, tt, assetStatistic.Name, assetStatistic.LatestUpdate); err != nil {
		return err
	}
	return nil
}

func (d *Dao) SelectAssetStatistic(tt uint32) (res []*model.AssetStatistic, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAssetStatistic, tt); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.AssetStatistic)
		Amount := int64(0)
		Amount_btc := int64(0)
		Amount_usd := int64(0)
		if err = rows.Scan(&r.Name, &r.Addressnum, &Amount, &Amount_btc, &Amount_usd, &r.TxNum, &r.LatestUpdate); err != nil {
			res = nil
			return
		}
		r.Amount = new(big.Int).SetInt64(Amount)
		r.Amount_btc = new(big.Int).SetInt64(Amount_btc)
		r.Amount_usd = new(big.Int).SetInt64(Amount_usd)
		res = append(res, r)
	}
	return res, nil
}

type AssetTxInfo1 struct {
	Amount  string
}

func (d *Dao) SelectAssetHistory1(start uint32, end uint32, name string) (res []*AssetTxInfo1, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectAssetHistory, start, end, name); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(AssetTxInfo1)
		if err = rows.Scan(&r.Amount); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	return res, nil
}

func (d *Dao) SelectAssetHistory(start uint32, end uint32, name string)  (res *model.AssetTxInfo, err error) {
	res = new(model.AssetTxInfo)
	res.Name = name
	res.Amount = big.NewInt(0)
	oriRes, err := d.SelectAssetHistory1(start, end, name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for _, item := range oriRes {
		amount := big.NewInt(0)
		amount, ok := new(big.Int).SetString(item.Amount, 10)
		if !ok {
			amount = big.NewInt(0)
		}
		res.Amount = new(big.Int).Add(res.Amount, amount)
		res.TxNum ++
	}
	return
}


func (d *Dao) InsertTransferStatistic(hash string) (err error) {
	if _, err = d.db.Exec(_insertTransferStatistic, hash); err != nil {
		return
	}
	return
}

func (d *Dao) SelectTransferStatistic() (res []*model.TransferStatistic, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectTransferStatistic); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.TransferStatistic)
		Amount := int64(0)
		if err = rows.Scan(&r.Hash, &Amount, &r.LatestIn, &r.LatestOut); err != nil {
			res = nil
			return
		}
		r.Amount = new(big.Int).SetInt64(Amount)
		res = append(res, r)
	}
	return res, nil
}


type TransferTxInfo1 struct {
	Amount  string
	TT   uint32
}

func (d *Dao) SelectTransferOutHistory1(start uint32, hash string) (res []*TransferTxInfo1, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectTransferOutHistory, hash, start); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(TransferTxInfo1)
		if err = rows.Scan(&r.Amount, &r.TT); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	return res, nil
}

func (d *Dao) SelectTransferOutHistory(start uint32, hash string)  (res *model.TransferTxInfo, err error) {
	res = new(model.TransferTxInfo)
	res.Hash = hash
	res.Amount = big.NewInt(0)
	res.TT = start
	oriRes, err := d.SelectTransferOutHistory1(start, hash)
	if err == sql.ErrNoRows {
		return res, nil
	}
	if err != nil {
		return nil, err
	}
	tt := start
	for _, item := range oriRes {
		amount := big.NewInt(0)
		amount, ok := new(big.Int).SetString(item.Amount, 10)
		if !ok {
			amount = big.NewInt(0)
		}
		res.Amount = new(big.Int).Add(res.Amount, amount)
		if item.TT > tt {
			tt = item.TT
		}
	}
	res.TT = tt
	return
}

func (d *Dao) SelectTransferInHistory1(start uint32, hash string) (res []*TransferTxInfo1, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(_selectTransferInHistory, hash, start); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(TransferTxInfo1)
		if err = rows.Scan(&r.Amount, &r.TT); err != nil {
			res = nil
			return
		}
		res = append(res, r)
	}
	return res, nil
}

func (d *Dao) SelectTransferInHistory(start uint32, hash string)  (res *model.TransferTxInfo, err error) {
	res = new(model.TransferTxInfo)
	res.Hash = hash
	res.Amount = big.NewInt(0)
	res.TT = start
	oriRes, err := d.SelectTransferInHistory1(start, hash)
	if err == sql.ErrNoRows {
		return res, nil
	}
	if err != nil {
		return nil, err
	}
	tt := start
	for _, item := range oriRes {
		amount := big.NewInt(0)
		amount, ok := new(big.Int).SetString(item.Amount, 10)
		if !ok {
			amount = big.NewInt(0)
		}
		res.Amount = new(big.Int).Add(res.Amount, amount)
		if item.TT > tt {
			tt = item.TT
		}
	}
	res.TT = tt
	return
}

func (d *Dao) UpdateTransferStatistic(transferStatistic *model.TransferStatistic) (err error) {
	if _, err = d.db.Exec(_updateTransferStatistic, transferStatistic.Amount.Int64(), transferStatistic.LatestIn, transferStatistic.LatestOut, transferStatistic.Hash); err != nil {
		return err
	}
	return nil
}


func (d *Dao) SelectAllTransferStatistic(chainid int) (res []*model.AllTransferStatistic, err error) {
	var rows *sql.Rows
	if chainid == 0 {
		if rows, err = d.db.Query(_selectAllTransferStatistic); err != nil {
			return
		}
	} else {
		if rows, err = d.db.Query(_selectAllTransferStatisticInChain, chainid); err != nil {
			return
		}
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.AllTransferStatistic)
		Amount := int64(0)
		if err = rows.Scan(&r.Hash, &Amount, &r.Token, &r.Chain, &r.Name, &r.SourceName, &r.SourceChain); err != nil {
			res = nil
			return
		}
		r.Amount = new(big.Int).SetInt64(Amount)
		res = append(res, r)
	}
	return res, nil
}