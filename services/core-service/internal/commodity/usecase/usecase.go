package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity"
	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/currency"
	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
	"github.com/akwanmaroso/backend-efishery-test/core-service/pkg/cache"
)

type commodityUsecaseImpl struct {
	commodityRepo commodity.Repository
	currencyRepo  currency.Repository
	cache         *cache.Cache
}

func NewCommodityUsecase(commodityRepo commodity.Repository, currencyRepo currency.Repository, cache *cache.Cache) commodity.UseCase {
	return &commodityUsecaseImpl{
		commodityRepo: commodityRepo,
		currencyRepo:  currencyRepo,
		cache:         cache,
	}
}

func (uc *commodityUsecaseImpl) List(ctx context.Context) ([]models.Commodity, error) {
	commodities, err := uc.commodityRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var idrCurrency models.Currency
	cachedValue, exist := uc.cache.Get("currency_usd_idr")
	if exist {
		cachedByte, err := json.Marshal(cachedValue)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(cachedByte, &idrCurrency); err != nil {
			return nil, err
		}
	} else {
		idrCurrency, err = uc.currencyRepo.GetCurrencyUSDToIDR(ctx)
		if err != nil {
			return nil, err
		}
		uc.cache.Set("currency_usd_idr", idrCurrency, 24*time.Hour)
	}

	for i := 0; i < len(commodities); i++ {
		convertedCurrency := uc.calculateCurrency(commodities[i].Price, idrCurrency)
		commodities[i].PriceUSD = convertedCurrency
	}

	return commodities, nil
}

func (uc *commodityUsecaseImpl) Aggregate(ctx context.Context) ([]models.AggregateResult, error) {
	commodities, err := uc.commodityRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	aggregateCommodities := make(map[string]*models.AggregateCommodities)
	for _, v := range commodities {
		if v.UUID == "" {
			continue
		}

		key := v.AreaProvinsi + "/" + uc.getWeekStartDate(v.TglParsed)
		price, err := strconv.ParseFloat(v.Price, 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		size, err := strconv.ParseFloat(v.Size, 64)
		if err != nil {
			return nil, err
		}

		aggregate, exists := aggregateCommodities[key]
		if !exists {
			aggregate = &models.AggregateCommodities{
				Price: models.AggregateField{
					Min:    price,
					Max:    price,
					Avg:    price,
					Median: price,
					Sum:    price,
					Values: []float64{price},
				},
				Size: models.AggregateField{
					Min:    size,
					Max:    size,
					Avg:    size,
					Median: size,
					Sum:    size,
					Values: []float64{size},
				},
				Count: 1,
			}
		} else {
			if price < aggregate.Price.Min {
				aggregate.Price.Min = price
			}
			if price > aggregate.Price.Max {
				aggregate.Price.Max = price
			}
			if size < aggregate.Size.Min {
				aggregate.Size.Min = size
			}
			if size > aggregate.Size.Max {
				aggregate.Size.Max = size
			}
			aggregate.Price.Sum += price
			aggregate.Size.Sum += size

			aggregate.Count++

			aggregate.Price.Avg = aggregate.Price.Sum / float64(aggregate.Count)
			aggregate.Size.Avg = aggregate.Size.Sum / float64(aggregate.Count)

			aggregate.Price.Values = append(aggregate.Price.Values, price)
			aggregate.Size.Values = append(aggregate.Size.Values, size)

			aggregate.Price.Median = uc.calculateMedian(aggregate.Price.Values)
			aggregate.Size.Median = uc.calculateMedian(aggregate.Size.Values)
		}
		aggregateCommodities[key] = aggregate
	}

	var aggregateResult []models.AggregateResult
	for key, v := range aggregateCommodities {
		keys := strings.Split(key, "/")
		aggregateResult = append(aggregateResult, models.AggregateResult{
			Area:            keys[0],
			StartOfWeekDate: keys[1],
			Aggregate:       v,
		})
	}

	return aggregateResult, nil
}

func (uc *commodityUsecaseImpl) calculateMedian(data []float64) float64 {
	sort.Float64s(data)
	length := len(data)
	if length%2 == 0 {
		return (data[length/2-1] + data[length/2]/2)

	}

	return data[length/2]
}

func (uc *commodityUsecaseImpl) getWeekStartDate(t time.Time) string {
	year, week := t.ISOWeek()
	startOfWeek := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC).AddDate(0, 0, (week-1)*7)

	return startOfWeek.Format("2006-01-02")
}

func (uc *commodityUsecaseImpl) calculateCurrency(amount string, currency models.Currency) string {
	amountInt, err := strconv.ParseFloat(amount, 10)
	if err != nil {
		return ""

	}

	result := amountInt / currency.Rates.IDR

	return fmt.Sprintf("%.2f", result)
}
