package middleware

import "github.com/leijiangnan/nanhade/framework"

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
		}()
		c.Next()
		return nil
	}
}
