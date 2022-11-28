package main

import (
	"github.com/leijiangnan/nanhade/framework"
	"log"
	"time"
)

func UserLoginController(c *framework.Context) error {
	log.Printf("uri: %s,UserLoginController running...", c.GetRequest().URL.Path)
	time.Sleep(5 * time.Second)
	c.SetOkStatus().Json("ok,UserLoginController " + c.GetRequest().URL.Path)
	return nil
}
