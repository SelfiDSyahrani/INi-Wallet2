package manager

import "INi-Wallet2/usecase"

type UsecaseManager interface {
	UserUseCase() usecase.UserUseCase
	TransactionUscase() usecase.TransactionUscase
	TransService() usecase.TransService
}

type useCaseManager struct {
	repoManager RepositoryManger
}

// TransService implements UsecaseManager
func (u *useCaseManager) TransService() usecase.TransService {
	return usecase.NewServiceTrans(
		u.repoManager.UserRepository(),
		u.UserUseCase(),
		u.repoManager.TransactionRepository(),
	)
}

func (u *useCaseManager) TransactionUscase() usecase.TransactionUscase {
	return usecase.NewTransaction(u.repoManager.TransactionRepository())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(
		u.repoManager.UserRepository(),
		u.TransactionUscase(),
	)
}

func NewUseCaseManager(repoManager RepositoryManger) UsecaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
