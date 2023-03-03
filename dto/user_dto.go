package dto

import (
	// "INi-Wallet2/model"
	"time"
)

type UserRequestParams struct {
	UserWallet_ID string `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name  string `form:"name"`
	Email string `form:"email"`
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

type DonasiReq struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Amount       uint   `json:"Amount"`
	ReceiverId   string `json:"receiver Id"`
	ReceiverName string `json:"receiver Name"`
}

type DonasiResponse struct {
	UserName     string `json:"name"`
	Amount       uint   `json:"Amount"`
	ReceiverId   string `json:"receiverId"`
	ReceiverName string `json:"receiver Name"`
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
