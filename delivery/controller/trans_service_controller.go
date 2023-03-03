package controller

import (
	"INi-Wallet2/dto"
	"INi-Wallet2/usecase"
	"INi-Wallet2/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TransServiceController struct {
	router  *gin.Engine
	usecase usecase.TransService
}

func (tsc *TransServiceController) donasi(ctx *gin.Context) {
	input := dto.DonasiReq{}
	err := ctx.ShouldBindJSON(&input)
	fmt.Println(input)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := tsc.usecase.Donasi(&input)
		if err != nil {
			utils.HandleInternalServerError(ctx, "gagal transaksi")
			fmt.Println(err)
		} else {
			response := &dto.DonasiResponse{
				UserName:     input.Name,
				Amount:       input.Amount,
				ReceiverId:   input.ReceiverId,
				ReceiverName: input.ReceiverName}
			utils.HandleSuccessCreated(ctx, "Transaksi berhasil dilakukan", response)
		}
	}
}

func NewServiceController(router *gin.Engine, usecase usecase.TransService) *TransServiceController {
	newcontrollers := TransServiceController{
		router:  router,
		usecase: usecase,
	}
	rGS := router.Group("api/v1/INi-Wallet")
	rGS.POST("/transfer", newcontrollers.donasi)
	return &newcontrollers
}
