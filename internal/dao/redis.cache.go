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
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/polynetwork/explorer/internal/conf"
	"github.com/polynetwork/explorer/internal/model"
	"github.com/pkg/errors"
	"time"
)

const (
	_multi    = "multi"
	_multiFtx = "multiftx"
	_from     = "from"
	_to       = "to"
	_toMtx    = "tomtx"
)

func keyMChainTx(hash string) string {
	return _multi + hash
}

func keyMChainTxByFTx(hash string) string {
	return _multiFtx + hash
}

func keyFChainTx(hash string) string {
	return _from + hash
}

func keyTChainTx(hash string) string {
	return _to + hash
}

func keyTChainTxByMTx(hash string) string {
	return _toMtx + hash
}

func (d *Dao) CacheMChainTx(txHash string) (m *model.MChainTx, err error) {

	key := keyMChainTx(txHash)
	var resp string
	resp, err = d.redis.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
			return
		}
		err = errors.Wrap(err, "dao cache mchaintx")
		return
	}
	m = new(model.MChainTx)
	err = json.Unmarshal([]byte(resp), &m)
	if err != nil {
		err = errors.Wrap(err, "dao cache mchaintx Unmarshal")
	}
	return
}

func (d *Dao) CacheMChainTxByFTx(fTxHash string) (m *model.MChainTx, err error) {
	key := keyMChainTxByFTx(fTxHash)
	var resp string

	resp, err = d.redis.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
			return
		}
		err = errors.Wrap(err, "dao cache mchaintxbyftx")
		return
	}
	m = new(model.MChainTx)
	err = json.Unmarshal([]byte(resp), &m)
	if err != nil {
		err = errors.Wrap(err, "dao cache mchaintxbyftx Unmarshal")
	}
	return
}

func (d *Dao) AddMChainTx(m *model.MChainTx) (err error) {

	key := keyMChainTx(m.TxHash)
	var mjson []byte
	mjson, err = json.Marshal(m)

	if err != nil {
		err = errors.Wrap(err, "dao add mchaintx marshal")
		return
	}
	if _, err = d.redis.Set(key, string(mjson), conf.DefConfig.Redis.Expiration*time.Minute).Result(); err != nil {
		err = errors.Wrap(err, "dao add mchaintx")
	}
	return
}

func (d *Dao) AddMChainTxByFTx(m *model.MChainTx) (err error) {
	key := keyMChainTxByFTx(m.FTxHash)
	var mjson []byte
	mjson, err = json.Marshal(m)

	if err != nil {
		err = errors.Wrap(err, "dao add mchaintx marshal by ftx")
		return
	}
	if _, err = d.redis.Set(key, string(mjson), conf.DefConfig.Redis.Expiration*time.Minute).Result(); err != nil {
		err = errors.Wrap(err, "dao add mchaintx by ftx")
	}
	return
}

func (d *Dao) CacheFChainTx(txHash string) (f *model.FChainTx, err error) {
	key := keyFChainTx(txHash)
	var resp string
	resp, err = d.redis.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
			return
		}
		err = errors.Wrap(err, "dao cache fchaintx")
		return
	}
	f = new(model.FChainTx)
	err = json.Unmarshal([]byte(resp), &f)
	if err != nil {
		err = errors.Wrap(err, "dao cache fchaintx unmarshal")
	}
	return
}

func (d *Dao) AddFChainTx(f *model.FChainTx) (err error) {
	key := keyFChainTx(f.TxHash)
	var fjson []byte
	fjson, err = json.Marshal(f)
	if err != nil {
		err = errors.Wrap(err, "dao add fchaintx marshal")
		return
	}
	if _, err = d.redis.Set(key, string(fjson), conf.DefConfig.Redis.Expiration*time.Minute).Result(); err != nil {
		err = errors.Wrap(err, "dao add fchaintx marshal")
	}
	return
}

func (d *Dao) CacheTChainTx(txHash string) (t *model.TChainTx, err error) {
	key := keyTChainTx(txHash)
	var resp string

	resp, err = d.redis.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
			return
		}
		err = errors.Wrap(err, "dao cache tchaintx")
		return
	}
	t = new(model.TChainTx)
	err = json.Unmarshal([]byte(resp), &t)
	if err != nil {
		err = errors.Wrap(err, "dao cache tchaintx Unmarshal")
	}
	return
}

func (d *Dao) CacheTChainTxByMTx(mTxHash string) (m *model.TChainTx, err error) {
	key := keyTChainTxByMTx(mTxHash)
	var resp string

	resp, err = d.redis.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
			return
		}
		err = errors.Wrap(err, "dao cache mchaintxbyftx")
		return
	}
	m = new(model.TChainTx)
	err = json.Unmarshal([]byte(resp), &m)
	if err != nil {
		err = errors.Wrap(err, "dao cache mchaintxbyftx Unmarshal")
	}
	return
}

func (d *Dao) AddTChainTx(t *model.TChainTx) (err error) {
	key := keyTChainTx(t.TxHash)
	var tjson []byte
	tjson, err = json.Marshal(t)
	if err != nil {
		err = errors.Wrap(err, "dao add tchaintx marshal")
		return
	}
	if _, err = d.redis.Set(key, string(tjson), conf.DefConfig.Redis.Expiration*time.Minute).Result(); err != nil {
		err = errors.Wrap(err, "dao add tchaintx")
	}
	return
}

func (d *Dao) AddTChainTxByMTx(t *model.TChainTx) (err error) {
	key := keyTChainTxByMTx(t.RTxHash)
	var tjson []byte
	tjson, err = json.Marshal(t)
	if err != nil {
		err = errors.Wrap(err, "dao add tchaintx marshal by mtx ")
		return
	}
	if _, err = d.redis.Set(key, string(tjson), conf.DefConfig.Redis.Expiration*time.Minute).Result(); err != nil {
		err = errors.Wrap(err, "dao add tchaintx by mtx ")
	}
	return
}
