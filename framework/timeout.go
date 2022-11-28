package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			fun(c)

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
}
