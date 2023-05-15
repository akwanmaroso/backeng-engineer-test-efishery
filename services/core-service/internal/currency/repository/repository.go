package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/currency"
	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
)

type currencyRepositoryImpl struct {
	client  *http.Client
	baseURL string
}

// NewCurrencyRepository ...
func NewCurrencyRepository(httpClient *http.Client, baseURL string) currency.Repository {
	return &currencyRepositoryImpl{client: httpClient, baseURL: baseURL}
}

func (repo *currencyRepositoryImpl) GetCurrencyUSDToIDR(ctx context.Context) (models.Currency, error) {
	var convertedCurrency models.Currency

	uri := fmt.Sprintf("%s/latest?symbols=IDR&base=USD", repo.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	req.Header.Set("apiKey", "X26IJdrfDFLnD8NszOASpB3YFUN9BtHJ")
	if err != nil {
		return convertedCurrency, err
	}

	resp, err := repo.client.Do(req)
	if err != nil {
		return convertedCurrency, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&convertedCurrency); err != nil {
		return convertedCurrency, err
	}

	return convertedCurrency, nil
}
