
// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.
 
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.
 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
 
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"context"
	"time"
)

//SendNoBlocking 无阻塞发送
func SendNoBlocking(m []byte, ch chan []byte) {
	select {
	case ch <- m:
	default:
	}
}

//IntermittentExec 间歇性执行
func IntermittentExec(ctx context.Context, f func(), t time.Duration) {
	tick := time.NewTicker(t)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			f()
		}
	}
}

//Exec 上下文执行
func Exec(ctx context.Context, f func() error, wait time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			re := f()
			if re != nil {
				return re
			}
		}
		//wait 1 second
		time.Sleep(wait)
	}
}
