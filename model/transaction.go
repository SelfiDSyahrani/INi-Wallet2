package model

import "time"

type Transaction struct {
	Transaction_ID      string    `json:"transaction_ID"`
	Userwallet_id       string    `json:"userwallet_id"`
	Money_Changer_ID    string    `json:"money_Changer_ID"`
	Transaction_Type_ID string    `json:"transaction_Type_ID"`
	Payment_method_id   string    `json:"payment_method_id"`
	Amount              float64   `json:"amount"`
	Status              bool      `json:"status"`
	Date_Time           time.Time `json:"date_Time"`
}
