package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webman/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	serve := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	go func() {
		serve.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serve.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
