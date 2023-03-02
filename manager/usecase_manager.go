package manager

import "INi-Wallet2/usecase"


type UsecaseManager interface {

	
	TransactionUscase() usecase.TransactionUscase
	
}
type useCaseManager struct {
	repoManager RepositoryManger
}


func (u *useCaseManager) TransactionUscase() usecase.TransactionUscase {
	return usecase.NewTransaction(u.repoManager.TransactionRepository())
}


func NewUseCaseManager(repoManager RepositoryManger) UsecaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
