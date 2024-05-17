package errorx

import (
	"errors"
)

type InvalidLoginInputError struct{}

func NewInvalidLoginInputError() *InvalidLoginInputError {
	return &InvalidLoginInputError{}
}

func (e *InvalidLoginInputError) Error() string {
	return "invalid phone or password"
}

func IsInvalidLoginInputError(err error) bool {
	if err == nil {
		return false
	}
	var e *InvalidLoginInputError
	ok := errors.As(err, &e)
	return ok
}
