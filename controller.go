package main

import (
	"context"
	"fmt"
	"github.com/leijiangnan/nanhade/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Second)
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		time.Sleep(2 * time.Second)
		log.Println("2 second later")
		c.SetOkStatus().Json("ok")
		sonCtx := context.WithValue(c.BaseContext(), "aaa", "bbb")
		select {
		case <-sonCtx.Done():
			log.Println("durationCtx is Done")
		default:
			log.Println("durationCtx is not Done")
		}
		log.Println("other work")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}
	return nil
}
