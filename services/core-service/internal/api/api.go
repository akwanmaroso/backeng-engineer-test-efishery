package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akwanmaroso/backend-efishery-test/core-service/config"
	"github.com/labstack/echo/v4"
)

type Api struct {
	echo *echo.Echo
	cfg  *config.Config
}

func NewApi(cfg *config.Config) *Api {
	return &Api{
		cfg:  cfg,
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
		Addr:         api.cfg.Port,
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
