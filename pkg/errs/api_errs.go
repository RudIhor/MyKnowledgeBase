package errs

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid credentials.")
	ErrSomethingWentWrong = errors.New("Something went wrong.")
)
