package repository

import (
	"INi-Wallet2/model"
	"INi-Wallet2/utils"

	"github.com/jmoiron/sqlx"
)

type TransactionTypeRepository interface {
	GetByID(transactionType_ID string) (model.TransactionType, error)
	GetAll() ([]model.TransactionType, error)
}

type trasactionTypeRepository struct {
	db *sqlx.DB
}

func (tp *trasactionTypeRepository) GetByID(transactionType_ID string) (model.TransactionType, error) {
	var transactionType model.TransactionType
	err := tp.db.QueryRow(utils.SELECT_TRANSACTIONS_TYPE_ID, transactionType_ID).Scan(
		&transactionType.Transaction_Type_ID,
		&transactionType.Type,
	)
	if err != nil {
		return model.TransactionType{}, err
	}
	return transactionType, nil
}

func (tp *trasactionTypeRepository) GetAll() ([]model.TransactionType, error) {
	var transactionTypes []model.TransactionType
	err := tp.db.Select(&transactionTypes, utils.SELECT_TRANSACTIONS_TYPE)
	if err != nil {
		return nil, err
	}
	return transactionTypes, nil

}

func NewTransactionTypeRepository(db * sqlx.DB) TransactionTypeRepository {
	return &trasactionTypeRepository{
		db: db,
	}
}