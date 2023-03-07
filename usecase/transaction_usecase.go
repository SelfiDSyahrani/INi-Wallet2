package usecase

import (
	"INi-Wallet2/model"
	"INi-Wallet2/repository"
	// "fmt"
	"log"
)

type TransactionUscase interface {
	TransactionGetByID(transaction_ID string) (model.Transaction, error)
	TransactionGetAll() ([]model.Transaction, error)
	TransactionByUserId(userWallet_id string) ([]model.Transaction, error)
}

type transactionUscase struct {
	transactionRepo repository.TransactionRepository
}


func (t *transactionUscase) TransactionGetByID(transaction_ID string) (model.Transaction, error) {
	return t.transactionRepo.GetByID(transaction_ID)
}

func (t *transactionUscase) TransactionGetAll() ([]model.Transaction, error) {
	var trxList []model.Transaction
	trxList, err := t.transactionRepo.GetAll()
	if err != nil {
		log.Println("error use case ", err.Error())

	}
	return trxList, err
}

func (t *transactionUscase) TransactionByUserId(userWallet_id string) ([]model.Transaction, error) {
	return t.transactionRepo.GetByuserWalletID(userWallet_id)
}

func NewTransaction(transactionRepo repository.TransactionRepository) TransactionUscase {
	return &transactionUscase{
		transactionRepo: transactionRepo,
	}
}
