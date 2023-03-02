package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWrapper struct {
	Message  string
	Response int
	Result   interface{}
}

func HandleSuccess(context *gin.Context, message string, data interface{}) {
	response := ResponseWrapper{
		Message:  message,
		Response: 200,
		Result:   data,
	}
	context.JSON(http.StatusOK, response)
}

func HandleSuccessCreated(context *gin.Context, message string, data interface{}) {
	response := ResponseWrapper{
		Message:  message,
		Response: 202,
		Result:   data,
	}
	context.JSON(http.StatusCreated, response)
}

func HandleBadRequest(context *gin.Context, message string) {
	response := ResponseWrapper{
		Message:  message,
		Response: 400,
		Result:   nil,
	}
	context.JSON(http.StatusBadRequest, response)
}

func HandleNotFound(context *gin.Context, message string) {
	response := ResponseWrapper{
		Message:  message,
		Response: 404,
		Result:   nil,
	}
	context.JSON(http.StatusNotFound, response)
}

func HandleInternalServerError(context *gin.Context, message string) {
	response := ResponseWrapper{
		Message:  message,
		Response: 500,
		Result:   nil,
	}
	context.JSON(http.StatusInternalServerError, response)
}
