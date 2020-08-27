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

package ctx

import (
	"context"
	"time"
)

type Context struct {
	Context context.Context
	Cancel  context.CancelFunc
}

func New() *Context {
	return &Context{
		Context: context.Background(),
	}
}

func (c *Context) WithCancel() *Context {
	child, cancel := context.WithCancel(c.Context)
	return &Context{
		Context: child,
		Cancel:  cancel,
	}
}

func (c *Context) WithDeadline(sec int64, nsec int64) *Context {
	child, cancel := context.WithDeadline(c.Context, time.Unix(sec, nsec))
	return &Context{
		Context: child,
		Cancel:  cancel,
	}
}

func (c *Context) WithTimeout(nsec int64) *Context {
	child, cancel := context.WithTimeout(c.Context, time.Duration(nsec))
	return &Context{
		Context: child,
		Cancel:  cancel,
	}
}
