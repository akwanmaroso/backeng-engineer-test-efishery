package usecase

import (
	"context"
	"testing"

	commodityMock "github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity/mock"
	currencyMock "github.com/akwanmaroso/backend-efishery-test/core-service/internal/currency/mock"
	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
	"github.com/akwanmaroso/backend-efishery-test/core-service/pkg/cache"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_ListSuccess(t *testing.T) {
	cache := cache.NewCache()
	ctx := context.Background()
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommodityRepo := commodityMock.NewMockRepository(ctrl)
	mockCurrencyRepo := currencyMock.NewMockRepository(ctrl)
	commodityUC := NewCommodityUsecase(mockCommodityRepo, mockCurrencyRepo, cache)

	mockCommodities := []models.Commodity{{UUID: "abc", Komoditas: "test"}}
	mockCurrency := models.Currency{Base: "IDR"}

	mockCommodityRepo.EXPECT().List(ctx).Return(mockCommodities, nil)
	mockCurrencyRepo.EXPECT().GetCurrencyUSDToIDR(ctx).Return(mockCurrency, nil)

	commodities, err := commodityUC.List(ctx)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, commodities)
}

func TestUsecase_AggregateSuccess(t *testing.T) {
	cache := cache.NewCache()
	ctx := context.Background()
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommodityRepo := commodityMock.NewMockRepository(ctrl)
	mockCurrencyRepo := currencyMock.NewMockRepository(ctrl)
	commodityUC := NewCommodityUsecase(mockCommodityRepo, mockCurrencyRepo, cache)

	mockCommodities := []models.Commodity{{UUID: "abc", Komoditas: "test", Price: "24000", Size: "20", AreaProvinsi: "Palu"}}

	mockCommodityRepo.EXPECT().List(ctx).Return(mockCommodities, nil)

	aggregateCommodities, err := commodityUC.Aggregate(ctx)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, aggregateCommodities)
}
