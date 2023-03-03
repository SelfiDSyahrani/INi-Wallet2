package usecase

import (
	"INi-Wallet2/model"
	"INi-Wallet2/repository"
)

type PaymentMethodUsecase interface {
	PaymentMethodGetByID(PaymentMethod_ID string) (model.PaymentMethod, error)
	PaymentMethodGetAll() ([]model.PaymentMethod, error)
}

type paymentMethodUsecase struct {
	paymentMethodRepo repository.PaymentMethodRepository
} 

func (pmu * paymentMethodUsecase) PaymentMethodGetByID(PaymentMethod_ID string) (model.PaymentMethod, error){
	return pmu.paymentMethodRepo.GetByID(PaymentMethod_ID)
}
func (pmu * paymentMethodUsecase) PaymentMethodGetAll() ([]model.PaymentMethod, error) {
	return pmu.paymentMethodRepo.GetAll()
}
func NewPaymentMethod(paymentMethodRepo repository.PaymentMethodRepository) PaymentMethodUsecase  {
	return &paymentMethodUsecase{
	paymentMethodRepo :paymentMethodRepo,
	}
}