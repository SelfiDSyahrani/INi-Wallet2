package controller

import (
	"INi-Wallet2/usecase"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	router *gin.Engine
	usecase usecase.PaymentMethodUsecase
}

func (pc *PaymentController) GetAllPayment(c *gin.Context)  {
	payments ,err := pc.usecase.PaymentMethodGetAll()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"data":    payments,
	})
}

func (pc *PaymentController) GetPaymentById(c *gin.Context) {
	id := c.Param("id")
	payment, err := pc.usecase.PaymentMethodGetByID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    payment,
		})
	}
}

func NewControllerPayment(pc *gin.Engine,usecase usecase.PaymentMethodUsecase) *PaymentController {
	controller := PaymentController{
		router:  pc,
		usecase: usecase,
	}
	pc.GET("/payments", controller.GetAllPayment)
	pc.GET("/payment/:id", controller.GetPaymentById)

	return &controller
}