package manager

import "INi-Wallet2/repository"

type RepositoryManger interface {
	TransactionRepository() repository.TransactionRepository
}

type repositoryManger struct {
	infra InfraManager
}


func (r *repositoryManger) TransactionRepository() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.SqlDb())
}


func NewRepositoryManager(infraManager InfraManager) RepositoryManger {
	return &repositoryManger{
		infra: infraManager,
	}
}
