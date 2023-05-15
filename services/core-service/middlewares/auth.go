package middlewares

import (
	"context"
	"net/http"

	"github.com/akwanmaroso/backend-efishery-test/core-service/utils"
	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := utils.ExtractJWTFromRequest(c, mw.cfg.JwtSecretKey)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "invalid"})
		}

		c.Set("user", claims.User)
		ctx := context.WithValue(c.Request().Context(), utils.UserCtxKey{}, claims.User)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
