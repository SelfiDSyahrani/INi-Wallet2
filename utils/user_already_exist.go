package utils

type UserAlreadyExistsError struct {
}

func (e *UserAlreadyExistsError) Error() string {
	return "email already in use"
}
