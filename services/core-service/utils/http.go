package utils

import (
	"context"
	"errors"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
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

type UserCtxKey struct{}

func GetUserFromCtx(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(UserCtxKey{}).(*models.User)
	if !ok {
		return nil, errors.New("invalid get role from ctx")
	}

	return user, nil
}
