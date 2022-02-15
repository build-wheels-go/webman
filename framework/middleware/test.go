package middleware

import (
	"fmt"
	"webman/framework"
)

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware test1 pre")
		if err := c.Next(); err != nil {
			return err
		}
		fmt.Println("middleware test1 post")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware test2 pre")
		if err := c.Next(); err != nil {
			return err
		}
		fmt.Println("middleware test2 post")
		return nil
	}
}
