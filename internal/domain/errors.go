package domain

import "errors"

var (
	// Account Type errors
	ErrAccountTypeNotFound      = errors.New("account type not found")
	ErrAccountTypeAlreadyExists = errors.New("account type already exists")
	ErrInvalidAccountTypeName   = errors.New("invalid account type name")
	ErrAccountTypeNameTooShort  = errors.New("account type name too short")
	ErrAccountTypeNameTooLong   = errors.New("account type name too long")
	ErrAccountTypeInactive      = errors.New("account type is inactive")

	// Generic errors
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
)
