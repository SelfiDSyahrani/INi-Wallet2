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
	ID       string  `db:"userwallet_id" json:"userwallet_id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
