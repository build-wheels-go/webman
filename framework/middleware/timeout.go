package middleware

import (
	"context"
	"log"
	"time"
	"webman/framework"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
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
			c.Json(500, "panic")
			log.Println(p)
		case <-finish:
			c.Json(200, "finish")
		case <-timeoutCtx.Done():
			c.HasTimeout()
			c.Json(502, "timeout")
		}
		return nil
	}
}
