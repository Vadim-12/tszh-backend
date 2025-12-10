package errors

import "errors"

var (
	ErrInvalidCredentials               = errors.New("invalid credentials")
	ErrUserNotFound                     = errors.New("user not found")
	ErrUserWithPhoneNumberAlreadyExists = errors.New("user with this phone number already exists")
	ErrUserWithEmailAlreadyExists       = errors.New("user with this email already exists")
	ErrRefreshInvalid                   = errors.New("invalid refresh token")
	ErrTokenExpired                     = errors.New("token expired")
)
