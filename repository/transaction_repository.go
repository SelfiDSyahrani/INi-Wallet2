package repository

import (
	"INi-Wallet2/model"
	"INi-Wallet2/utils"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	InsertTransactionTransfer(trasaction model.Transaction) error
	InsertTransactionTopUp(trasaction model.Transaction) error
	InsertTransactionPayment(trasaction model.Transaction) error
	GetByID(transaction_ID string) (model.Transaction, error)
	GetAll() ([]model.Transaction, error)
	Delete(transaction_ID string) error
}

type transactionRepository struct {
	db *sqlx.DB
}

// insert transaksi transfer
func (tr *transactionRepository) InsertTransactionTransfer(trasaction model.Transaction) error {
	tx, err := tr.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	_, err = tx.Exec(utils.INSERT_RECORDS_TRANSFER)
	if err != nil {
		return err
	}
	return nil

	tx, err = tr.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	_, err = tx.Exec(utils.UPDATE_BALANCE_TRANSFER_USER)
	if err != nil {
		return err
	}
	return nil

	tx, err = tr.db.BeginTx(context.Background(), nil)
	_, err = tx.Query(utils.BALANCE_TRANSFER)
	if err != nil {
		return err

	}
	return nil
	tx.Commit()
	return nil
}

// insert transaction Top UP
func (tr *transactionRepository) InsertTransactionTopUp(trasaction model.Transaction) error {
	tx, err := tr.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()
	_, err = tx.Exec(utils.INSERT_RECORDS_TOPUP)
	if err != nil {
		return err
	}
	return nil
	tx, err = tr.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	_, err = tx.Exec(utils.UPDATE_BALANCE_TOPUP)
	if err != nil {
		return err
	}
	return nil

	tx, err = tr.db.BeginTx(context.Background(), nil)
	_, err = tx.Query(utils.BALANCE_TOPUP)
	if err != nil {
		return err

	}
	return nil
	tx.Commit()
	return nil

}

// insert transaction payment
func (tr *transactionRepository) InsertTransactionPayment(trasaction model.Transaction) error {
	tx, err := tr.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()
	_, err = tx.Exec(utils.INSERT_RECORDS_PAYMENT)
	if err != nil {
		return err
	}

	_, err = tx.Exec(utils.UPDATE_BALANCE_PAYMENT)
	if err != nil {
		return err
	}

	_, err = tx.Query(utils.BALANCE_PAYMENT)
	if err != nil {
		return err

	}

	tx.Commit()
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

// delete
func (tr *transactionRepository) Delete(transaction_ID string) error {
	_, err := tr.db.Exec(utils.DELETE_TRANSACTION, transaction_ID)
	if err != nil {
		return err
	}
	return nil
}

// object
func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
