package commodity

import (
	"context"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
)

// Repository ...
type Repository interface {
	List(context.Context) ([]models.Commodity, error)
}
