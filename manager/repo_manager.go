package manager

import "INi-Wallet2/repository"

type RepositoryManger interface {
	UserRepository() repository.UserRepository
	TransactionRepository() repository.TransactionRepository
	MoneyChangerRepsitory() repository.MoneyChangerRepository
	PaymentMethodRepository() repository.PaymentMethodRepository
	TransactionTypeRepository() repository.TransactionTypeRepository

}

type repositoryManger struct {
	infra InfraManager
}

func (r *repositoryManger) UserRepository() repository.UserRepository {
	return repository.NewUserRepository(r.infra.SqlDb())
}

func (r *repositoryManger) TransactionRepository() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.SqlDb())
}

func (r *repositoryManger) MoneyChangerRepsitory() repository.MoneyChangerRepository {
	return repository.NewMoneyChanger(r.infra.SqlDb())
}

func (r *repositoryManger) PaymentMethodRepository() repository.PaymentMethodRepository {
	return repository.NewPaymentMethodRepository(r.infra.SqlDb())
}

func (r *repositoryManger) TransactionTypeRepository() repository.TransactionTypeRepository {
	return repository.NewTransactionTypeRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManger {
	return &repositoryManger{
		infra: infraManager,
	}
}
