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

package error

const (
	SUCCESS                 int64 = 1
	DB_CONNECTTION_FAILED   int64 = 10000
	DB_LOADDATA_FAILED      int64 = 10001
	REST_PARAM_INVALID      int64 = 20000
	REST_METHOD_INVALID     int64 = 20001
	REST_ILLEGAL_DATAFORMAT int64 = 20002
)

var ErrMap = map[int64]string{
	SUCCESS:                 "success",
	DB_CONNECTTION_FAILED:   "connect db error",
	DB_LOADDATA_FAILED:      "load db data error",
	REST_PARAM_INVALID:      "invalid rest parameter",
	REST_METHOD_INVALID:     "invalid rest method",
	REST_ILLEGAL_DATAFORMAT: "rest illegal data format",
}
