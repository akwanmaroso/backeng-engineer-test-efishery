package api

import (
	"net/http"
	"time"

	commodityRepository "github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity/repository"
	commodityUsecase "github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity/usecase"
	currencyRepository "github.com/akwanmaroso/backend-efishery-test/core-service/internal/currency/repository"
	"github.com/akwanmaroso/backend-efishery-test/core-service/pkg/cache"

	commodityHttp "github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity/delivery/http"

	"github.com/akwanmaroso/backend-efishery-test/core-service/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (api *Api) MapHandlers(e *echo.Echo) error {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cache := cache.NewCache()

	currencyRepo := currencyRepository.NewCurrencyRepository(httpClient)

	commodityRepo := commodityRepository.NewCommodityRepository(httpClient)
	commodityUC := commodityUsecase.NewCommodityUsecase(commodityRepo, currencyRepo, cache)
	commodityHandler := commodityHttp.NewCommodityHandler(commodityUC)

	middlewareManager := middlewares.NewMiddlewareManager(api.cfg)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	v1 := e.Group("/api/v1")
	healthRoute := e.Group("/health")
	commodityRoute := v1.Group("/commodities")

	commodityRoute.Use(middlewareManager.AuthMiddleware)

	commodityHttp.MapCommodityRoutes(commodityRoute, commodityHandler, *middlewareManager)

	healthRoute.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "Ok"})
	})

	return nil
}
