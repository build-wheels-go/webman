package middleware

import (
	"fmt"
	"time"
	"webman/framework"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		//开始时间
		start := time.Now()
		c.Next()

		//结束时间
		end := time.Now()
		cost := end.Sub(start)
		fmt.Printf("api uri: %v,cost: %v", c.GetRequest().RequestURI, cost.Seconds())
		return nil
	}
}
