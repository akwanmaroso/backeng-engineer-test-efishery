package utils

import (
	"encoding/json"
	"errors"
	"html"
	"strings"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractJWTFromRequest(e echo.Context, key string) (*models.UserClaims, error) {
	tokenStr, err := extractBearerToken(e)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature ")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	userClaims, err := convertToUserClaims(claims)
	if err != nil {
		return nil, err
	}

	return userClaims, nil
}

func convertToUserClaims(claims map[string]interface{}) (*models.UserClaims, error) {
	var userClaims models.UserClaims
	tokenByte, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tokenByte, &userClaims)
	if err != nil {
		return nil, err
	}

	return &userClaims, nil
}

func extractBearerToken(e echo.Context) (string, error) {
	headerAuth := e.Request().Header.Get("Authorization")
	bearerToken := strings.Split(headerAuth, " ")

	if len(bearerToken) < 2 {
		return "", errors.New("invalid format token")
	}

	return html.EscapeString(bearerToken[1]), nil

}
