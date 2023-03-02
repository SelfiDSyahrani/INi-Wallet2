package utils

type IncorrectCredentialsError struct {
}

func (e *IncorrectCredentialsError) Error() string {
	return "incorrect email or password"
}
