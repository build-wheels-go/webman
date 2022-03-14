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
			c.SetStatus(500).Json("panic")
			log.Println(p)
		case <-finish:
			c.SetOkStatus().Json("finish")
		case <-timeoutCtx.Done():
			c.HasTimeout()
			c.SetStatus(502).Json("timeout")
		}
		return nil
	}
}
