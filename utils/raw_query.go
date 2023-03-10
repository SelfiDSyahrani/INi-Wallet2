package utils

const (

	//user
	INSERT_USER       = "INSERT INTO users (userWallet_id, name, email, phone, password, balance) values ($!,$2, $3, $4, $5, $6)"
	DELETE_USER       = "DELETE FROM users where userwallet_id = $1"
	SELECT_USER_ID    = "SELECT * FROM users where userWallet_id = $1"
	SELECT_USER_LIST  = "SELECT * FROM users"
	SELECT_BY_EMAIL   = "SELECT * FROM users where email = $1"
	UPDATE_USER_BY_ID = "UPDATE users set balance=$1 where userwallet_id=$2"
	UPDATE_USER_PASS  = "UPDATE users set password=$1 where email=$2"

	//=======================================================
	// transaction
	SELECT_TRANSACTION            = "SELECT * FROM transactions"
	SELECT_TRANSACTION_ID         = "SELECT * FROM transactions where transaction_id = $1"
	DELETE_TRANSACTION            = "DELETE FROM transactions where transaction_id = $1"
	SELECT_TRANSACTION_BY_USER_ID = "SELECT * FROM transactions where userWallet_id = $1"
	INSERT_TRANSACTION = "INSERT INTO transactions (transaction_ID, userwallet_id, money_Changer_ID, transaction_Type_ID, payment_method_id, amount, status, date_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	//transfer
	INSERT_RECORDS_TRANSFER      = "INSERT INTO transactions (transaction_id,userwallet_id,money_changer_id,trasaction_type_id,amount,status,date_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	UPDATE_BALANCE_TRANSFER_USER = "UPDATE users set balance = balance - $1"
	BALANCE_TRANSFER             = "SELECT *FROM users WHERE userWallet_ID"

	//topup
	INSERT_RECORDS_TOPUP = "INSERT INTO transactions (transaction_id,userwallet_id,money_changer_id,trasaction_type_id,amount,status,date_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	UPDATE_BALANCE_TOPUP = "UPDATE users SET balance= balance + (SELECT amount FROM transactions WHERE transactions_ID=$1) WHERE userwallet_id=$2"
	BALANCE_TOPUP        = "SELECT *FROM users WHERE userWallet_ID"

	//payment
	INSERT_RECORDS_PAYMENT = "INSERT INTO transactions (transaction_id,userwallet_id,money_changer_id,trasaction_type_id,amount,status,date_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	UPDATE_BALANCE_PAYMENT = "UPDATE users SET balance= balance - ((SELECT amount SELECT amount FROM transactions WHERE transactions_ID=$1) * (SELECT exchange_rate FROM money_changer WHERE money_changer_ID=$1))"
	BALANCE_PAYMENT        = "SELECT *FROM users WHERE userWallet_ID"
	//==========================================================
	//money changer
	SELECT_MONEY_CHANGER_ID = "SELECT * FROM money_changer where money_changer_ID = $1"
	SELECT_MONEY_CHANGER    = "SELECT * FROM money_changer "

	//payment_method
	INSERT_PAYMENT_METHOD      = "INSERT INTO payment_method (payment_method_id,method_name) values ($1,$2)"
	SELECT_PAYMENT_METHOD_LIST = "SELECT * FROM payment_method"
	SELECT_PAYMENT_METHOD_ID   = "SELECT * FROM payment_method where payment_method_id = $1"
	DELETE_PAYMENT_METHOD      = "DELETE FROM payment_method where payment_method_id = $1"

	//transaction_type
	SELECT_TRANSACTIONS_TYPE_ID = "SELECT * FROM transaction_type where transaction_type_ID = $1"
	SELECT_TRANSACTIONS_TYPE    = "SELECT * FROM transaction_type"
)
