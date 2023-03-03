package utils

type NotValidEmailError struct {
}

func (e *NotValidEmailError) Error() string {
	return "not a valid email"
}

type NotEnoughBalance struct{
}

func (b *NotEnoughBalance) Error() string{
	return "not enough balance"
}
