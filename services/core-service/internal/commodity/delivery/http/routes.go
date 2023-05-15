package http

import (
	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity"
	"github.com/labstack/echo/v4"
)

// MapCommodityRoutes ...
func MapCommodityRoutes(group *echo.Group, h commodity.Handlers) {
	group.GET("", h.List())
	group.GET("/aggregate", h.Aggregate())
}
