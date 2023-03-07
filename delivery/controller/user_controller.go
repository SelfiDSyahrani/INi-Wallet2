package controller

import (
	// "INi-Wallet2/delivery/middleware"
	"INi-Wallet2/dto"
	"INi-Wallet2/model"
	"INi-Wallet2/usecase"
	"INi-Wallet2/utils"
	"INi-Wallet2/utils/authonticator"
	"fmt"
	"net/http"

	// "path/filepath"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUC  usecase.UserUseCase
	TransUC usecase.TransactionUscase
	jwt     authonticator.JWTService
}

func (uc *UserController) GenerateToken(ctx *gin.Context) {
	var request dto.LoginRequestBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	check, err := uc.userUC.Login(&request)
	if err != nil {
		utils.HandleInternalServerError(ctx, err.Error())
		return
	}
	tokenString, err := uc.jwt.GenerateToken(request)
	if err != nil {
		utils.HandleInternalServerError(ctx, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": check, "token": tokenString})
}

func (cc *UserController) loginUser(ctx *gin.Context) {
	input := dto.LoginRequestBody{}
	err := ctx.ShouldBindJSON(&input)
	fmt.Println(input)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		user, err := cc.userUC.Login(&input)
		if err != nil {
			utils.HandleInternalServerError(ctx, "Salah password")
		} else {
			userTampilan := &dto.UserResponseBody{}
			userTampilan.Name = user.Name
			userTampilan.Email = user.Email
			userTampilan.ID = user.ID
			userTampilan.Balance = user.Balance
			utils.HandleSuccessCreated(ctx, "Success log-in", userTampilan)
		}
	}
}

func (cc *UserController) daftarUser(ctx *gin.Context) {
	newuser := dto.RegisterReq{}
	err := ctx.ShouldBind(&newuser)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	}
	input := model.User{
		ID:       "",
		Name:     newuser.Name,
		Email:    newuser.Email,
		Phone:    newuser.Phone,
		Password: newuser.Password,
		Balance:  0,
	}
	file := newuser.Identitas
	dst := "identitas/" + file.Filename
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	err = cc.userUC.RegisterUser(&input)
	if err != nil {
		utils.HandleInternalServerError(ctx, err.Error())
	} else {
		userTampilan := &dto.UserResponseBody{}
		userTampilan.Name = newuser.Name
		userTampilan.Email = newuser.Email
		userTampilan.ID = input.ID
		utils.HandleSuccessCreated(ctx, "Success Create New User", userTampilan)
	}
}

func (cc *UserController) registerCustomer(ctx *gin.Context) {
	var newuser model.User
	err := ctx.ShouldBindJSON(&newuser)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := cc.userUC.RegisterUser(&newuser)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		} else {
			userTampilan := &dto.UserResponseBody{}
			userTampilan.Name = newuser.Name
			userTampilan.Email = newuser.Email
			userTampilan.ID = newuser.ID
			utils.HandleSuccessCreated(ctx, "Success Create New User", userTampilan)
		}
	}
}

func (cc *UserController) forgotPass(ctx *gin.Context) {
	input := dto.ForgotPasswordRequestBody{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := cc.userUC.ForgotPass(&input)
		if err != nil {
			utils.HandleInternalServerError(ctx, "gagal perbarui password")
		}
		utils.HandleSuccessCreated(ctx, "Success Reset Password", input)
	}
}

func (cc *UserController) getAllUser(ctx *gin.Context) {
	users, err := cc.userUC.GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (cc *UserController) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := cc.userUC.GetUserByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (cc *UserController) getByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := cc.userUC.GetByEmail(email)
	if err != nil {
		fmt.Println("error di get by email user controller")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (cc *UserController) userListTrans(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := cc.userUC.GetUserWithTrans(id)
	if err != nil {
		utils.HandleNotFound(ctx, "User Id is not Found")
	}
	ctx.JSON(http.StatusOK, user)
}

func NewUserController(router *gin.Engine, usecase usecase.UserUseCase) *UserController {
	newcontroller := UserController{
		userUC: usecase,
	}
	rG := router.Group("api/v1/INi-Wallet/user")
	rG.POST("/token", newcontroller.GenerateToken)
	// secure := rG.Group("/secure").Use(middleware.Auth())
	// {
	// 	secure.GET(":id/transaction", newcontroller.userListTrans)
	// 	secure.GET(":id", newcontroller.getUserById)
	// 	secure.GET("/users", newcontroller.getAllUser)
	// 	secure.GET("getEmail/:email", newcontroller.getByEmail)
	// }
	rG.GET(":id/transaction", newcontroller.userListTrans)
	rG.GET(":id", newcontroller.getUserById)
	rG.GET("/users", newcontroller.getAllUser)
	rG.GET("getEmail/:email", newcontroller.getByEmail)
	rG.POST("/register", newcontroller.registerCustomer)
	rG.POST("/registerForm", newcontroller.daftarUser)
	rG.POST("/login", newcontroller.loginUser)
	rG.POST("/forgetPassword", newcontroller.forgotPass)
	return &newcontroller
}
