package repository

import (
	"INi-Wallet2/model"
	"INi-Wallet2/utils"

	"github.com/jmoiron/sqlx"
)

type PaymentMethodRepository interface {
	GetByID(PaymentMethod_ID string) (model.PaymentMethod, error)
	GetAll() ([]model.PaymentMethod, error)
	//Update(paymentMethod model.PaymentMethod) error
}

type paymentMethodRepository struct {
	db *sqlx.DB
}


//GetByID
func (pmr *paymentMethodRepository) GetByID(PaymentMethod_ID string) (model.PaymentMethod, error) {
	var paymentMethod model.PaymentMethod
	err := pmr.db.QueryRow(utils.SELECT_PAYMENT_METHOD_ID, PaymentMethod_ID).Scan(
		&paymentMethod.Payment_Method_ID,
		&paymentMethod.Method_name,
	)
	if err != nil {
		return model.PaymentMethod{},err
	}
	return paymentMethod, nil
}

//GetAll
func (pmr *paymentMethodRepository) GetAll() ([]model.PaymentMethod, error) {
	var paymentMethods []model.PaymentMethod
	err := pmr.db.Select(&paymentMethods, utils.SELECT_PAYMENT_METHOD_LIST)
	if err != nil {
		return nil, err
	}
	return paymentMethods, nil
}


//object
func NewPaymentMethodRepository(db *sqlx.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		db: db,
	}
}