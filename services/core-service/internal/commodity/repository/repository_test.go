package repository

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"uuid": "d992453e-c26e-48d6-bedc-92bc55f8585e","komoditas": "BANDENG", "area_provinsi": "SULAWESI BARAT", "area_kota": "MAMUJU UTARA", "size": "180","price": "29000","timestamp": "1641064093344"}]`))
	}))

	defer mockServer.Close()

	client := &http.Client{}
	baseURL := mockServer.URL

	repo := NewCommodityRepository(client, baseURL)
	commodities, err := repo.List(context.Background())
	fmt.Println("=>>", commodities)
	assert.NoError(t, err)

	expectedCommodities := []models.Commodity{
		{UUID: "d992453e-c26e-48d6-bedc-92bc55f8585e", Komoditas: "BANDENG", AreaProvinsi: "SULAWESI BARAT", AreaKota: "MAMUJU UTARA", Size: "180", Price: "29000", TglParsed: time.Time{}, Timestamp: "1641064093344"},
	}
	assert.Equal(t, expectedCommodities, commodities)
}
