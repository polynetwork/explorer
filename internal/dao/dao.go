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
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/polynetwork/explorer/internal/conf"
	"time"
)

type Dao struct {
	redis *redis.Client
	db    *sql.DB
}

func NEW(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		redis: GetRedisClient(c),
		db:    GetDBConn(c),
	}
	return
}

func (d *Dao) Close() {
	d.db.Close()
	d.redis.Close()
}

func GetRedisClient(c *conf.Config) *redis.Client {
	if c.Redis.DialTimeout <= 0 || c.Redis.ReadTimeout <= 0 || c.Redis.WriteTimeout <= 0 {
		panic("must config redis timeout")
	}
	options := &redis.Options{
		Network:      c.Redis.Proto,
		Addr:         c.Redis.Addr,
		DialTimeout:  c.Redis.DialTimeout * time.Second,
		ReadTimeout:  c.Redis.ReadTimeout * time.Second,
		WriteTimeout: c.Redis.WriteTimeout * time.Second,
		PoolSize:     c.Redis.PoolSize,
		IdleTimeout:  c.Redis.IdleTimeout * time.Second,
	}
	return redis.NewClient(options)
}

func GetDBConn(c *conf.Config) *sql.DB {
	db, err := sql.Open("mysql",
		c.Mysql.User+
			":"+c.Mysql.Password+
			"@tcp("+c.Mysql.Url+
			")/"+c.Mysql.DbName+
			"?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic("connect mysql failed")
	}
	return db
}

func (d *Dao) BeginTran() (*sql.Tx, error) {
	return d.db.Begin()
}

func (d *Dao) Ping() (err error) {
	if err = d.db.Ping(); err != nil {
		return
	}
	if _, err = d.redis.Ping().Result(); err != nil {
		return
	}
	return
}
