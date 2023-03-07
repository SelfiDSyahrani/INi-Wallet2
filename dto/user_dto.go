package dto

import (
	"mime/multipart"
	"time"
)

type UserRequestParams struct {
	UserWallet_ID string `uri:"id" binding:"required"`
}

type RegisterReq struct {
	Name      string                `form:"name" binding:"required"`
	Email     string                `form:"email" binding:"required"`
	Phone     string                `form:"phone" binding:"required"`
	Password  string                `form:"password" binding:"required"`
	Identitas *multipart.FileHeader `form:"identitas" binding:"required"`
}

type UserResponseBody struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}
type ForgotPasswordRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPasswordResponseBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type PasswordReset struct {
	Email     string
	Token     string
	ExpiredAt time.Time
}

// input json di postman
type DonasiReq struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Amount       uint   `json:"Amount"`
	ReceiverId   string `json:"receiver Id"`
	ReceiverName string `json:"receiver Name"`
}

// response di postman
type DonasiResponse struct {
	UserName       string `json:"name"`
	Amount         uint   `json:"Amount"`
	ReceiverId     string `json:"receiverId"`
	ReceiverName   string `json:"receiver Name"`
	Transaction_ID string `json:"Transaction Id"`
}

// input json di postman
type TopUpReq struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Amount    uint   `json:"Amount" binding:"required,min=50000,max=10000000"`
	Method_id string `json:"Method Id"`
}

type TopUpResponse struct {
	UserName       string    `json:"name"`
	Amount         uint      `json:"Amount"`
	Method         string    `json:"Method"`
	Transaction_ID string    `json:"Transaction Id"`
	TimeofTrans    time.Time `json:"Time"`
}

type BayarReq struct {
	UserName string `json:"name"`
	Email    string `json:"Email"`
	Password string `json:"password"`
	Amount   uint   `json:"Amount"`
	Currency string `json:"Currency" binding:"required"`
}

type BayarResponse struct {
	UserName       string    `json:"name"`
	Amount         uint      `json:"Amount"`
	Currency       string    `json:"Currency"`
	Transaction_ID string    `json:"Transaction Id"`
	TimeofTrans    time.Time `json:"Time"`
}

type RegisterRequest struct {
}

// func FormatUser(user *model.User) UserResponseBody {
// 	formattedUser := UserResponseBody{}
// 	formattedUser.ID = user.ID
// 	formattedUser.Name = user.Name
// 	formattedUser.Email = user.Email
// 	return formattedUser
// }

// func FormatUsers(authors []*model.User) []UserResponseBody {
// 	formattedUsers := []UserResponseBody{}
// 	for _, user := range authors {
// 		formattedUser := FormatUser(user)
// 		formattedUsers = append(formattedUsers, formattedUser)
// 	}
// 	return formattedUsers
// }

// func FormatForgotPassword(passwordReset PasswordReset) ForgotPasswordResponseBody {
// 	return ForgotPasswordResponseBody{
// 		Email: passwordReset.Email,
// 		Token: passwordReset.Token,
// 	}
// }
