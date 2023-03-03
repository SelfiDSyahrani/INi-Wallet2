package controller

import (
	"INi-Wallet2/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	router  *gin.Engine
	usecase usecase.TransactionUscase
}

func (tc *TransactionController) GetAllTransaction(c *gin.Context) {
	transaction, err := tc.usecase.TransactionGetAll()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"data":    transaction,
	})

}
func (tc *TransactionController) GetTransactionById(c *gin.Context) {
	id := c.Param("id")
	transaction, err := tc.usecase.TransactionGetByID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    transaction,
		})
	}
}

/* func (tc * TransactionController) TransferControler(c *gin.Context)  {
	transactionn, err := tc.usecase.Transfer()
	if err != nil {
		c.JSON(400,gin.H{
			"message" : err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"message" : "OK",
			"data" : transactionn,
		})
	}
} */

func NewControllerTransaksi(tc *gin.Engine, usecase usecase.TransactionUscase) *TransactionController {
	controller := TransactionController{
		router:  tc,
		usecase: usecase,
	}
	tc.GET("/transaction", controller.GetAllTransaction)
	tc.GET("/transaction/:id", controller.GetTransactionById)
	//tc.POST("/transfer",controller.TransferControler)
	return &controller
}
