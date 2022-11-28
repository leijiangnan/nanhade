package main

import (
	"context"
	"github.com/leijiangnan/nanhade/framework"
	"github.com/leijiangnan/nanhade/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	registerRouter(core)

	server := http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
