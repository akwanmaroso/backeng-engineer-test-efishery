package currency

import (
	"context"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
)

type Repository interface {
	GetCurrencyUSDToIDR(context.Context) (models.Currency, error)
}
