package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ReqIDCtxKey struct {
}

func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

func GetContextFromRequest(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}
