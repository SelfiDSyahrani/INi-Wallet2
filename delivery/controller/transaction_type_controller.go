package controller

import (
	"INi-Wallet2/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionTypeController struct {
	router *gin.Engine
	usecase usecase.TransactionTypeUsecase
}

func (ttc *TransactionTypeController) GetAllTransactionType(c *gin.Context)  {
	transactionTypes ,err := ttc.usecase.TransactionTypeGetAll()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"data":    transactionTypes,
	})
}

func (ttc *TransactionTypeController) GetTransactionTyprById(c *gin.Context) {
	id := c.Param("id")
	transactionType, err := ttc.usecase.TransactionTypeGetByID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    transactionType,
		})
	}
}

func NewControllerTransactionType(ttc *gin.Engine,usecase usecase.TransactionTypeUsecase) *TransactionTypeController {
	controller := TransactionTypeController{
		router:  ttc,
		usecase: usecase,
	}
	ttc.GET("/trasactionTypes", controller.GetAllTransactionType)
	ttc.GET("/trasactionType/:id", controller.GetTransactionTyprById)
	return &controller
}