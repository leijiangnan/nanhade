package main

import (
	"github.com/leijiangnan/nanhade/framework"
	"log"
)

func SubjectAddController(c *framework.Context) error {
	log.Println("SubjectAddController running...")
	c.Json(200, "ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	log.Println("SubjectListController running...")
	c.Json(200, "ok, SubjectListController")
	return nil
}
func SubjectDelController(c *framework.Context) error {
	log.Println("SubjectDelController running...")
	c.Json(200, "ok, SubjectDelController")
	return nil
}
func SubjectUpdateController(c *framework.Context) error {
	log.Println("SubjectUpdateController running...")
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}
func SubjectGetController(c *framework.Context) error {
	log.Println("SubjectDelController running...")
	c.Json(200, "ok, SubjectDelController")
	return nil
}
