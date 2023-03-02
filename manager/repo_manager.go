package manager

import "INi-Wallet2/repository"

type RepositoryManger interface {
	UserRepository() repository.UserRepository
	TransactionRepository() repository.TransactionRepository
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

func NewRepositoryManager(infraManager InfraManager) RepositoryManger {
	return &repositoryManger{
		infra: infraManager,
	}
}
