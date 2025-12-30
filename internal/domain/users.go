package domain

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

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

func Validate(u sqlc.CreateUserParams) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email, is.LowerCase),
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 255), is.Alphanumeric),
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 255), is.Alphanumeric),
		validation.Field(&u.Password, validation.Required),
	)
}

func ValidateEmail(email string) error {
	return validation.Validate(email, is.Email)
}
