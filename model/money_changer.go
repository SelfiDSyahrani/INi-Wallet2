package model

type MoneyChanger struct {
	Money_changer_id string  `json:"money_changer_id"`
	Name             string  `json:"name"`
	Exchange_rate    float64 `json:"exchange_rate"`
	Currency_type    string  `json:"currency_type"`
	Country          string  `json:"country"`
}
