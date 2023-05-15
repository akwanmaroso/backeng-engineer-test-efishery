package models

// type Currency struct {
// 	Rate   float64 `json:"rate"`
// 	Amount float64 `json:"amount"`
// 	Result float64 `json:"result"`
// }

type BaseCurrencyResponse struct {
	Date string `json:"date"`
	Info struct {
		Rate      float64 `json:"rate"`
		Timestamp int     `json:"timestamp"`
	} `json:"info"`
	Query struct {
		Amount int    `json:"amount"`
		From   string `json:"from"`
		To     string `json:"to"`
	} `json:"query"`
	Result  float64 `json:"result"`
	Success bool    `json:"success"`
}

type Currency struct {
	Base  string
	Date  string `json:"date"`
	Rates struct {
		IDR float64
	}
}

// type CurrenyRequest struct {
// 	Amount float64 `json:"amount"`
// }
