package usecase

import (
	"INi-Wallet2/model"
	"INi-Wallet2/repository"
	"log"
)

type TransactionUscase interface {
	Transfer(transaction model.Transaction) error
	TopuP(transaction model.Transaction) error
	Payment(transaction model.Transaction) error
	TransactionGetByID(transaction_ID string) (model.Transaction, error)
	TransactionGetAll() ([]model.Transaction, error)
}

type transactionUscase struct {
	transactionRepo repository.TransactionRepository
}

func (t *transactionUscase) Transfer(transaction model.Transaction) error {
	return t.transactionRepo.InsertTransactionTransfer(transaction)
	/* if t.transactionRepo.InsertTransactionPayment <= t.Transfer {
		return errors.New("saldo anda tidak cukup")
	} */

	//Proses dengan transactions
	//1. transctionRepo.InsertIntoTabelTransction
	// bila sukses ambil amount ransaction dan Update saldo user menjadi saldo + amount_transfer

	//3. select data saldo terakhir setelah di update where ID_USER = id

	return nil
}

func (t *transactionUscase) TopuP(transaction model.Transaction) error {
	return t.transactionRepo.InsertTransactionTopUp(transaction)
	return nil
}

func (t *transactionUscase) Payment(transaction model.Transaction) error {
	return t.transactionRepo.InsertTransactionPayment(transaction)
	return nil
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

func NewTransaction(transactionRepo repository.TransactionRepository) TransactionUscase {
	return &transactionUscase{
		transactionRepo: transactionRepo,
	}
}
