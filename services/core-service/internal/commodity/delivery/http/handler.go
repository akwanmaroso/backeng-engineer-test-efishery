package http

import (
	"net/http"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity"
	"github.com/akwanmaroso/backend-efishery-test/core-service/utils"
	"github.com/labstack/echo/v4"
)

type commodityHandlerImpl struct {
	commodityUC commodity.UseCase
}

func NewCommodityHandler(commodityUC commodity.UseCase) commodity.Handlers {
	return &commodityHandlerImpl{
		commodityUC: commodityUC,
	}
}

func (handler *commodityHandlerImpl) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetContextFromRequest(c)
		commodities, err := handler.commodityUC.List(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, commodities)
	}
}

func (handler *commodityHandlerImpl) Aggregate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetContextFromRequest(c)
		commodities, err := handler.commodityUC.Aggregate(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, commodities)
	}
}
