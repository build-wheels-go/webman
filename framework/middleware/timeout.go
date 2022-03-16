// Copyright 2021 build-wheels-go.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"context"
	"log"
	"time"

	"webman/framework/gin"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		timeoutCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			//使用next执行具体业务
			c.Next()

			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			c.ISetStatus(500).IJson("panic")
			log.Println(p)
		case <-finish:
			c.ISetOkStatus().IJson("finish")
		case <-timeoutCtx.Done():
			c.ISetStatus(502).IJson("timeout")
		}
	}
}
