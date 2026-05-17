package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUsernameExists    = errors.New("username already exists")
	ErrCustomerNotFound  = errors.New("customer not found")
	ErrEmailExists       = errors.New("email already exists")
	ErrProductNotFound   = errors.New("product not found")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrOrderNotFound     = errors.New("order not found")
	ErrInternalServer    = errors.New("internal server error")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidToken      = errors.New("invalid token")
	ErrForbidden         = errors.New("forbidden")
)
