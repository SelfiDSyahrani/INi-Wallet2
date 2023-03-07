package repository

import (
	"INi-Wallet2/model"
	"INi-Wallet2/utils"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	CreateTrans(transaction *model.Transaction) error
	GetByID(transaction_ID string) (model.Transaction, error)
	GetAll() ([]model.Transaction, error)
	GetByuserWalletID(userWallet_id string) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

// create new transaction
func (tr *transactionRepository) CreateTrans(trns *model.Transaction) error {
	_, err := tr.db.Query(utils.INSERT_TRANSACTION, trns.Transaction_ID, trns.Userwallet_id, trns.Money_Changer_ID, trns.Transaction_Type_ID, trns.Payment_method_id, trns.Amount, trns.Status, trns.Date_Time)
	if err != nil {
		return err
	}
	return nil
}

// GetByID
func (tr *transactionRepository) GetByID(transaction_ID string) (model.Transaction, error) {
	var transaction model.Transaction
	err := tr.db.QueryRow(utils.SELECT_TRANSACTION_ID, transaction_ID).Scan(
		&transaction.Transaction_ID,
		&transaction.Userwallet_id, // di ubah
		&transaction.Money_Changer_ID,
		&transaction.Transaction_Type_ID,
		&transaction.Payment_method_id,
		&transaction.Amount,
		&transaction.Status,
		&transaction.Date_Time,
	)
	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

// GetAll
func (tr *transactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := tr.db.Select(&transactions, utils.SELECT_TRANSACTION)
	if err != nil {
		log.Println("error repo =", err)
		return nil, err
	}
	return transactions, nil

}


// get all transaction for specific user //
func (tr *transactionRepository) GetByuserWalletID(userWallet_id string) ([]model.Transaction, error) {
	var transaction []model.Transaction
	err := tr.db.Select(&transaction, utils.SELECT_TRANSACTION_BY_USER_ID, userWallet_id)
	if err != nil {
		fmt.Println("erorr di sinii")
		return nil, err
	}
	return transaction, nil
}

// object
func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
