package middlewares

import "github.com/akwanmaroso/backend-efishery-test/core-service/config"

type MiddlewareManager struct {
	cfg *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg}

}
