package delivery

import (
	"INi-Wallet2/config"
	"INi-Wallet2/delivery/controller"
	"INi-Wallet2/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine         *gin.Engine
	useCaseManager manager.UsecaseManager
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		engine:         ginEngine,
		useCaseManager: usecase,
	}
}

func (a *appServer) initHandlers() {

	controller.NewUserController(a.engine, a.useCaseManager.UserUseCase())

	// controller.NewStoreController(a.engine, a.useCaseManager.StoreUseCase(), a.useCaseManager.ProductUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(":8085")
	if err != nil {
		panic(err.Error())
	}
}
