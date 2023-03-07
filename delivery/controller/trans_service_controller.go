package controller

import (
	"INi-Wallet2/dto"
	"INi-Wallet2/usecase"
	"INi-Wallet2/utils"
	"fmt"
	"time"

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
		TransId, err := tsc.usecase.Donasi(&input)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		} else {
			response := &dto.DonasiResponse{
				UserName:       input.Name,
				Amount:         input.Amount,
				ReceiverId:     input.ReceiverId,
				ReceiverName:   input.ReceiverName,
				Transaction_ID: TransId,
			}
			utils.HandleSuccessCreated(ctx, "Transaksi berhasil dilakukan", response)
		}
	}
}

func (tsc *TransServiceController) isiUlang(ctx *gin.Context) {
	input := dto.TopUpReq{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		TransId, err := tsc.usecase.IsiUlang(&input)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		} else {
			response := &dto.TopUpResponse{
				UserName:       input.Name,
				Amount:         input.Amount,
				Method:         TransId,
				Transaction_ID: input.Method_id,
			}
			utils.HandleSuccess(ctx, "Top-Up Success", response)
		}
	}
}

func (tsc *TransServiceController) bayar(ctx *gin.Context) {
	input := dto.BayarReq{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		TransId, err := tsc.usecase.Bayar(&input)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		} else {
			response := &dto.BayarResponse{
				UserName:       input.UserName,
				Amount:         input.Amount,
				Currency:       input.Currency,
				Transaction_ID: TransId,
				TimeofTrans:    time.Now(),
			}
			utils.HandleSuccess(ctx, "Payment Success", response)
		}
	}

}

func NewServiceController(router *gin.Engine, usecase usecase.TransService) *TransServiceController {
	newcontrollers := TransServiceController{
		router:  router,
		usecase: usecase,
	}
	router.Use()
	rGS := router.Group("api/v1/INi-Wallet/Transaction")
	rGS.POST("/transfer", newcontrollers.donasi)
	rGS.POST("/topUp", newcontrollers.isiUlang)
	rGS.POST("/payment", newcontrollers.bayar)
	return &newcontrollers
}
