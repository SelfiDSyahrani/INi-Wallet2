package usecase

import (
	"INi-Wallet2/model"
	"INi-Wallet2/repository"
	"log"
)

type MoneyChangerUsecase interface {
	MoneyChangerById(MoneyChanger_ID string) (*model.MoneyChanger, error)
	MoneyChangerAll() ([]model.MoneyChanger, error)
}

type moneyChangerUsecase struct {
	moneyChangerRepo repository.MoneyChangerRepository
}

func (mcu *moneyChangerUsecase) MoneyChangerById(MoneyChanger_ID string) (*model.MoneyChanger, error) {
	return mcu.moneyChangerRepo.GetByID(MoneyChanger_ID)
}

func (mcu *moneyChangerUsecase) MoneyChangerAll() ([]model.MoneyChanger, error) {
	var moneyChangerList []model.MoneyChanger
	moneyChangerList,err := mcu.moneyChangerRepo.GetAll()
	if err != nil {
		log.Println("error use case ", err.Error())
		
	}
	return moneyChangerList,err
}


func NewMoneyChanger(moneyChangerRepo repository.MoneyChangerRepository) MoneyChangerUsecase {
	return &moneyChangerUsecase{
		moneyChangerRepo: moneyChangerRepo,
	}
}
