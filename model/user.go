package model

type User struct {
	ID       string  `db:"userwallet_id" json:"userwallet_id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

type UserForm struct {
	ID       string  `db:"userwallet_id" form:"userwallet_id"`
	Name     string  `form:"name"`
	Email    string  `form:"email"`
	Phone    string  `form:"phone"`
	Password string  `form:"password"`
	Balance  float64 `form:"balance"`
}
