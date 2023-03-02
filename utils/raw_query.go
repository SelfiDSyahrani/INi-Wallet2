package utils

const (
	//payment_method
	INSERT_PAYMENT_METHOD      = "INSERT INTO payment_method (payment_method_id,method_name) values ($1,$2)"
	SELECT_PAYMENT_METHOD_LIST = "SELECT * FROM payment_method"
	SELECT_PAYMENT_METHOD_ID   = "SELECT * FROM payment_method where payment_method_id = $1"
	//UPDATE_PAYMENT_METHOD = "UPDATE payment_method set payment_method_id=$1,method_name =$2 where id=$3"
	DELETE_PAYMENT_METHOD = "DELETE FROM payment_method where payment_method_id = $1"

	//transfer

	//user
	INSERT_USER       = "INSERT INTO users (userWallet_id, name, email, phone, password, balance) values ($!,$2, $3, $4, $5, $6)"
	DELETE_USER       = "DELETE FROM users where userWallet_id = $1"
	SELECT_USER_ID    = "SELECT * FROM users where userWallet_id = $1"
	SELECT_USER_LIST  = "SELECT * FROM users"
	SELECT_BY_EMAIL   = "SELECT * FROM users where email = $1"
	UPDATE_USER_BY_ID = "UPDATE users set name=$1, balance=$2 where userWallet_id=$3"
	UPDATE_USER_PASS  = "UPDATE users set password=$1 where email=$2"
)
