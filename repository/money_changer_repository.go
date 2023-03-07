package repository

import (
	"INi-Wallet2/model"
	"INi-Wallet2/utils"
	"log"

	"github.com/jmoiron/sqlx"
)

type MoneyChangerRepository interface {
	GetByID(MoneyChanger_ID string) (*model.MoneyChanger, error)
	GetAll() ([]model.MoneyChanger, error)
}

type moneyChangerRepository struct {
	db *sqlx.DB
}

func (mc *moneyChangerRepository) GetByID(MoneyChanger_ID string) (*model.MoneyChanger, error) {
	var MoneyChanger model.MoneyChanger
	err := mc.db.QueryRow(utils.SELECT_MONEY_CHANGER_ID, MoneyChanger_ID).Scan(
		&MoneyChanger.Money_changer_id,
		&MoneyChanger.Name,
		&MoneyChanger.Exchange_rate,
		&MoneyChanger.Currency_type,
		&MoneyChanger.Country,
	)
	if err != nil {
		return &model.MoneyChanger{}, err
	}
	return &MoneyChanger, nil
}

func (mc *moneyChangerRepository) GetAll() ([]model.MoneyChanger, error) {
	var moneyChanger []model.MoneyChanger
	err := mc.db.Select(&moneyChanger, utils.SELECT_MONEY_CHANGER)
	if err != nil {
		log.Println("error repo =", err)
		return nil, err
	}
	return moneyChanger, nil

}

func NewMoneyChanger(db *sqlx.DB) MoneyChangerRepository {
	return &moneyChangerRepository{
		db: db,
	}

}
