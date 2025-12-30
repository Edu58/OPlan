package domain

import (
	"context"
	"github.com/Edu58/Oplan/pkg/crypto"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type UserOTP struct {
	OTP1 string
	OTP2 string
	OTP3 string
	OTP4 string
	OTP5 string
	OTP6 string
}

type UserService interface {
	CreateUser(ctx context.Context, req sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (sqlc.User, error)
	GetUserByMSISDN(ctx context.Context, msisdn string) (sqlc.User, error)
}

type UserRepository interface {
	Create(ctx context.Context, req sqlc.CreateUserParams) (sqlc.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (sqlc.User, error)
	GetByEmail(ctx context.Context, email string) (sqlc.User, error)
	GetByMSISDN(ctx context.Context, msisdn string) (sqlc.User, error)
}

func ValidateCreateUser(u sqlc.CreateUserParams) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email, is.LowerCase),
		validation.Field(&u.FirstName, validation.Required, validation.Length(2, 255), is.Alphanumeric),
		validation.Field(&u.LastName, validation.Required, validation.Length(2, 255), is.Alphanumeric),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 30), is.Alphanumeric),
	)
}

func ValidateEmail(email string) error {
	return validation.Validate(email, validation.Required, is.Email, is.LowerCase)
}

func HashPassword(password string) (string, error) {
	return crypto.HashPassword(password)
}

func VerifyPassword(password string, hash string) (bool, error) {
	return crypto.VerifyPassword(password, hash)
}

func (otp UserOTP) ValidateUserOTP() error {
	return validation.ValidateStruct(&otp,
		validation.Field(&otp.OTP1, validation.Required, is.Alphanumeric, validation.Length(1, 1)),
		validation.Field(&otp.OTP2, validation.Required, is.Alphanumeric, validation.Length(1, 1)),
		validation.Field(&otp.OTP3, validation.Required, is.Alphanumeric, validation.Length(1, 1)),
		validation.Field(&otp.OTP4, validation.Required, is.Alphanumeric, validation.Length(1, 1)),
		validation.Field(&otp.OTP5, validation.Required, is.Alphanumeric, validation.Length(1, 1)),
		validation.Field(&otp.OTP6, validation.Required, is.Alphanumeric, validation.Length(1, 1)),
	)
}
