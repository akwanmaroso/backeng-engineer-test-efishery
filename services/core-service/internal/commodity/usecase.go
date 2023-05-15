package commodity

import (
	"context"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
)

// UseCase ...
type UseCase interface {
	List(ctx context.Context) ([]models.Commodity, error)
	Aggregate(ctx context.Context) ([]models.AggregateResult, error)
}
