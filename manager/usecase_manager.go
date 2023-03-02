package manager

import "INi-Wallet2/usecase"


type UsecaseManager interface {

	UserUseCase() usecase.UserUseCase
	TransactionUscase() usecase.TransactionUscase
	MoneyChangerUsecase() usecase.MoneyChangerUsecase
	PaymentMethodUsecase() usecase.PaymentMethodUsecase
	TransactionTypeUsecase() usecase.TransactionTypeUsecase

}

type useCaseManager struct {
	repoManager RepositoryManger
}


func (u *useCaseManager) TransactionUscase() usecase.TransactionUscase {
	return usecase.NewTransaction(u.repoManager.TransactionRepository())
}


func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepository())
}

func (u *useCaseManager) MoneyChangerUsecase() usecase.MoneyChangerUsecase {
	return usecase.NewMoneyChanger(u.repoManager.MoneyChangerRepsitory())
}

func (u *useCaseManager) PaymentMethodUsecase() usecase.PaymentMethodUsecase {
	return usecase.NewPaymentMethod(u.repoManager.PaymentMethodRepository())
}

func (u *useCaseManager) TransactionTypeUsecase() usecase.TransactionTypeUsecase {
	return usecase.NewTransactionType(u.repoManager.TransactionTypeRepository())
}

func NewUseCaseManager(repoManager RepositoryManger) UsecaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
