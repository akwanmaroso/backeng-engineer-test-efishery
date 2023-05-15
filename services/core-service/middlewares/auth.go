package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/akwanmaroso/backend-efishery-test/core-service/utils"
	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := utils.ExtractJWTFromRequest(c, mw.cfg.JwtSecretKey)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "invalid"})
		}

		c.Set("user", claims)
		ctx := context.WithValue(c.Request().Context(), utils.UserCtxKey{}, claims)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func (mw *MiddlewareManager) RoleMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := utils.GetUserFromCtx(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "invalid"})
			}
			if user.Role != role {
				return c.JSON(http.StatusForbidden, map[string]interface{}{"message": "invalid role"})
			}

			return next(c)
		}
	}

}
