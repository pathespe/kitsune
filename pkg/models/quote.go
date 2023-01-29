package models

type Quote struct {
	Ask           float64 `json:"ask"`
	Name          string  `json:"name"`
	DividendYield float64 `json:"dividendYield"`
	YTDReturn     float64 `json:"yTDReturn"`
}
