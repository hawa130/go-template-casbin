package reqerr

import (
	"errors"
)

var (
	ErrInvalidLoginInput = errors.New("invalid phone or password")

	ErrForbidden = errors.New("forbidden")
)
