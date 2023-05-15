package api

import (
	"net/http"
	"time"

	currencyRepository "github.com/akwanmaroso/backend-efishery-test/core-service/internal/currency/repository"

	commodityRepository "github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity/repository"
	commodityUsecase "github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity/usecase"
	"github.com/akwanmaroso/backend-efishery-test/core-service/pkg/cache"

	commodityHttp "github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity/delivery/http"

	"github.com/labstack/echo/v4"
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

	v1 := e.Group("/api/v1")
	healthRoute := e.Group("/health")
	commodityRoute := v1.Group("/commodities")

	commodityHttp.MapCommodityRoutes(commodityRoute, commodityHandler)

	healthRoute.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "Ok"})
	})

	return nil
}
