package jwtutils

import (
	"errors"
)

var (
	ErrGenToken             = errors.New("Token generation error")
	ErrTokenAlreadyBeenUsed = errors.New("Token already been used")
	ErrTokenNotFound        = errors.New("Token not found in database")
	ErrTokenDisabled     = errors.New("Token is disabled")
)
