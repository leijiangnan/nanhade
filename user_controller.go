package main

import (
	"github.com/leijiangnan/nanhade/framework"
	"log"
)

func UserLoginController(c *framework.Context) error {
	log.Printf("uri: %s,UserLoginController running...", c.GetRequest().URL.Path)
	c.Json(200, "ok,UserLoginController "+c.GetRequest().URL.Path)
	return nil
}
