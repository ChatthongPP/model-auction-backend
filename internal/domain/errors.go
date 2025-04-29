package domain

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrPhoneNumberExists  = errors.New("phone number already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden access")
	ErrEmailNotFound      = errors.New("email not found")
	ErrIncorrectPassword  = errors.New("incorrect password")
)
