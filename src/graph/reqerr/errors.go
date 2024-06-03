package reqerr

import (
	"errors"
)

var (
	ErrInvalidLoginInput = errors.New("invalid phone or password")
	ErrPhoneAlreadyExist = errors.New("phone already exist")

	ErrForbidden        = errors.New("forbidden")
	ErrPasswordNotMatch = errors.New("password not match")
)
