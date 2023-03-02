package manager

import "INi-Wallet2/usecase"

type UsecaseManager interface {
	UserUseCase() usecase.UserUseCase
}
type useCaseManager struct {
	repoManager RepositoryManger
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepository())
}

func NewUseCaseManager(repoManager RepositoryManger) UsecaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
