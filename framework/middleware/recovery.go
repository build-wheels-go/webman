package middleware

import "webman/framework"

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if p := recover(); p != nil {
				c.Json(500, p)
			}
		}()

		if err := c.Next(); err != nil {
			return err
		}

		return nil
	}
}
