package controller

import (
	"INi-Wallet2/usecase"

	"github.com/gin-gonic/gin"
)

type MoneyChangerController struct {
	router  *gin.Engine
	usecase usecase.MoneyChangerUsecase
}

func (mcc *MoneyChangerController) GetAllMoneyChangerController(c *gin.Context) {
	MoneyChanger, err := mcc.usecase.MoneyChangerAll()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"data":    MoneyChanger,
	})
}

func (mcc * MoneyChangerController)GetMoneyChangerControllerById(c *gin.Context)  {
	id := c.Param("id")
	MoneyChanger, err := mcc.usecase.MoneyChangerById(id) 
	if err != nil {
		c.JSON(400,gin.H{
			"messege" : err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"meesege" : "OK",
		"data" : MoneyChanger,
	})
} 

func NewMoneyChangerController(mcc *gin.Engine, usecase usecase.MoneyChangerUsecase) *MoneyChangerController{
	controller := MoneyChangerController{
		router:  mcc,
		usecase: usecase,
	}
	mcc.GET("/MoneyChanger", controller.GetAllMoneyChangerController)
	mcc.GET("/MoneyChanger/:id",controller.GetMoneyChangerControllerById)
	return &controller
}
