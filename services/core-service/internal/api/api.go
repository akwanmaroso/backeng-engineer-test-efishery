package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Api struct {
	echo *echo.Echo
}

func NewApi() *Api {
	return &Api{
		echo: echo.New(),
	}
}

func (api *Api) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := api.MapHandlers(api.echo); err != nil {
		return err
	}

	srv := &http.Server{
		Addr:         api.getPort(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		if err := api.echo.StartServer(srv); err != nil {
			log.Panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatalf("signal notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}

	if err := api.echo.Server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (api *Api) getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	return port
}
