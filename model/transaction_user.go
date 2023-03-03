package model

type UserTransaction struct {
	UserWallet_id string
	UserName      string
	Transaction   []Transaction
}
