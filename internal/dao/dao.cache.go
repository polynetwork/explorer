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
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
)

func (d *Dao) MChainTx(txHash string) (res *model.MChainTx, err error) {
	addCache := true
	res, err = d.CacheMChainTx(txHash)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		return
	}
	res, err = d.SelectMChainTxByHash(txHash)
	if err != nil {
		return
	}

	if res == nil || !addCache {
		return
	}
	err = d.AddMChainTx(res)
	return
}

func (d *Dao) MChainTxByFTx(fTxHash string, chain uint32) (res *model.MChainTx, err error) {
	addCache := true
	res, err = d.CacheMChainTxByFTx(fTxHash)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		return
	}
	res, err = d.SelectMChainTxByFHash(fTxHash, chain)
	if err != nil {
		return
	}
	if res == nil || !addCache {
		return
	}
	err = d.AddMChainTxByFTx(res)
	return
}

func (d *Dao) InsertMChainTxAndCache(m *model.MChainTx) (err error) {
	err = d.InsertMChainTx(m)
	if err != nil {
		return
	}
	err = d.AddMChainTx(m)
	if err != nil {
		log.Warnf("InsertMChainTxAndCache: AddMChainTx", err)
		err = nil
		return
	}
	err = d.AddMChainTxByFTx(m)
	if err != nil {
		log.Warnf("InsertMChainTxAndCache: AddMChainTxByFTx", err)
		err = nil
		return
	}
	return
}

func (d *Dao) TxInsertMChainTxAndCache(tx *sql.Tx, m *model.MChainTx) (err error) {
	err = d.TxInsertMChainTx(tx, m)
	if err != nil {
		return
	}
	err = d.AddMChainTx(m)
	if err != nil {
		log.Warnf("TxInsertMChainTxAndCache: AddMChainTx", err)
		err = nil
		return
	}
	err = d.AddMChainTxByFTx(m)
	if err != nil {
		log.Warnf("TxInsertMChainTxAndCache: AddMChainTxByFTx", err)
		err = nil
		return
	}
	return
}

func (d *Dao) FChainTx(txHash string) (res *model.FChainTx, err error) {
	addCache := true
	res, err = d.CacheFChainTx(txHash)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		return
	}
	res, err = d.SelectFChainTxByHash(txHash)
	if err != nil {
		return
	}
	if res == nil || !addCache {
		return
	}
	err = d.AddFChainTx(res)
	return
}

func (d *Dao) FChainTx1(txHash string, chain uint32) (res *model.FChainTx, err error) {
	addCache := true
	res, err = d.CacheFChainTx(txHash)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		return
	}
	res, err = d.SelectFChainTxByHash1(txHash, chain)
	if err != nil {
		return
	}
	if res == nil || !addCache {
		return
	}
	err = d.AddFChainTx(res)
	return
}

func (d *Dao) InsertFChainTxAndCache(f *model.FChainTx) (err error) {
	err = d.InsertFChainTx(f)
	if err != nil {
		return
	}
	err = d.AddFChainTx(f)
	if err != nil {
		log.Warnf("InsertFChainTxAndCache: AddFChainTx", err)
		err = nil
		return
	}
	return
}

func (d *Dao) TxInsertFChainTxAndCache(tx *sql.Tx, f *model.FChainTx) (err error) {
	err = d.TxInsertFChainTx(tx, f)
	if err != nil {
		return
	}
	err = d.AddFChainTx(f)
	if err != nil {
		log.Warnf("TxInsertFChainTxAndCache: AddFChainTx", err)
		err = nil
		return
	}
	return
}

func (d *Dao) TChainTx(txHash string) (res *model.TChainTx, err error) {
	addCache := true
	res, err = d.CacheTChainTx(txHash)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	if res != nil {
		return
	}
	res, err = d.SelectTChainTxByHash(txHash)
	if err != nil {
		return
	}

	if res == nil || !addCache {
		return
	}
	err = d.AddTChainTx(res)
	return
}

func (d *Dao) TChainTxByMTx(txHash string) (res *model.TChainTx, err error) {
	addCache := true
	res, err = d.CacheTChainTxByMTx(txHash)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		return
	}
	res, err = d.SelectTChainTxByMHash(txHash)
	if err != nil {
		return
	}

	if res == nil || !addCache {
		return
	}
	err = d.AddTChainTxByMTx(res)
	return
}

func (d *Dao) InsertTChainTxAndCache(t *model.TChainTx) (err error) {
	err = d.InsertTChainTx(t)
	if err != nil {
		return
	}
	err = d.AddTChainTx(t)
	if err != nil {
		log.Warnf("InsertTChainTxAndCache: AddTChainTx", err)
		err = nil
		return
	}
	err = d.AddTChainTxByMTx(t)
	if err != nil {
		log.Warnf("InsertTChainTxAndCache: AddTChainTxByMTx", err)
		err = nil
		return
	}
	return
}

func (d *Dao) TxInsertTChainTxAndCache(tx *sql.Tx, t *model.TChainTx) (err error) {
	err = d.TxInsertTChainTx(tx, t)
	if err != nil {
		return
	}
	err = d.AddTChainTx(t)
	if err != nil {
		log.Warnf("TxInsertTChainTxAndCache: AddTChainTx", err)
		err = nil
		return
	}
	err = d.AddTChainTxByMTx(t)
	if err != nil {
		log.Warnf("TxInsertTChainTxAndCache: AddTChainTxByMTx", err)
		err = nil
		return
	}
	return
}

// TODO
//func (d *Dao) UpdateBitcoinTxAndCache(txhash string, height uint32, tt uint32) (err error) {
//	if err = d.UpdateBitcoinTxConfirmed(txhash, height, tt); err != nil {
//		return
//	}
//	return
//}
