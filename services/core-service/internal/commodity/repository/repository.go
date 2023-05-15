package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/commodity"
	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
)

const BASE_URL = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"

type commodityRepositoryImpl struct {
	client  *http.Client
	baseURL string
}

// NewCommodityRepository ...
func NewCommodityRepository(httpClient *http.Client, baseURL string) commodity.Repository {
	return &commodityRepositoryImpl{client: httpClient, baseURL: baseURL}
}

func (repo *commodityRepositoryImpl) List(ctx context.Context) ([]models.Commodity, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, repo.baseURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := repo.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var commodities []models.Commodity
	if err := json.NewDecoder(resp.Body).Decode(&commodities); err != nil {
		return nil, err
	}

	return commodities, nil
}
