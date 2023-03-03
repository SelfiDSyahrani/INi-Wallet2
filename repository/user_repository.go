package repository

import (
	"INi-Wallet2/model"
	"INi-Wallet2/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Delete(userWallet_ID string) error
	FindByEmail(email string) (model.User, error)
	GetByID(userWallet_ID string) (model.User, error)
	GetAll() ([]model.User, error)
	Insert(user *model.User) error
	UpdateById(user model.User) error
	UpdateByEmail(user model.User) error
}

type userRepository struct {
	db *sqlx.DB
}

// Creates a new user
func (r *userRepository) Insert(user *model.User) error {
	query := "INSERT INTO users (userWallet_id, name, email, phone, password, balance) VALUES ($1, $2, $3, $4, $5,$6)"
	_, err := r.db.Query(query, user.ID, user.Name, user.Email, user.Phone, user.Password, user.Balance)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes an existing user
func (r *userRepository) Delete(user_ID string) error {
	_, err := r.db.Exec(utils.DELETE_USER, user_ID)
	if err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a user by ID
// GetByID
func (r *userRepository) GetByID(user_ID string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(utils.SELECT_USER_ID, user_ID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Balance,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetAll retrieves all users
func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, utils.SELECT_USER_LIST)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateById(user model.User) error {
	res, err := r.db.Exec(utils.UPDATE_USER_BY_ID, user.Balance, user.ID)
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rows updated: %v\n", count)
	return nil
}

func (r *userRepository) UpdateByEmail(user model.User) error {
	_, err := r.db.Exec(utils.UPDATE_USER_PASS, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	fmt.Println(email)
	var user model.User
	err := r.db.QueryRow(utils.SELECT_BY_EMAIL, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Balance,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
