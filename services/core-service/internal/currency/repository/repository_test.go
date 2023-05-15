package repository

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencyUSDToIDR(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"base":"IDR","date":"2023-03-15","rates":{"idr": 14000}}`))
	}))

	defer mockServer.Close()

	client := &http.Client{}
	baseURL := mockServer.URL

	repo := NewCurrencyRepository(client, baseURL)
	commodities, err := repo.GetCurrencyUSDToIDR(context.Background())
	assert.NoError(t, err)

	expectedCommodities := models.Currency{Base: "IDR", Date: "2023-03-15", Rates: struct{ IDR float64 }{IDR: 14000}}

	assert.Equal(t, expectedCommodities, commodities)
}
