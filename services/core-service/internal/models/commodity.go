package models

import "time"

// Commodity ...
type Commodity struct {
	UUID         string    `json:"uuid"`
	Komoditas    string    `json:"komoditas"`
	AreaProvinsi string    `json:"area_provinsi"`
	AreaKota     string    `json:"area_kota"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	PriceUSD     string    `json:"price_usd"`
	TglParsed    time.Time `json:"tgl_parsed"`
	Timestamp    string    `json:"timestamp"`
}

type AggregateResult struct {
	Area            string                `json:"area"`
	StartOfWeekDate string                `json:"start_of_week_date"`
	Aggregate       *AggregateCommodities `json:"aggregate"`
}

type AggregateCommodities struct {
	Price AggregateField `json:"price"`
	Size  AggregateField `json:"size"`
	Count int            `json:"count"`
}

type AggregateField struct {
	Min    float64   `json:"min"`
	Max    float64   `json:"max"`
	Avg    float64   `json:"avg"`
	Median float64   `json:"median"`
	Sum    float64   `json:"-"`
	Values []float64 `json:"-"`
}
