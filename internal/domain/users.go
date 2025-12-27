package domain

import (
	"context"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Password       string    `json:"password"`
	MSISDN         string    `json:"msisdn"`
	EmailVerified  bool      `json:"email_verified"`
	MsisdnVerified bool      `json:"msisdn_verified"`
	Active         bool      `json:"active"`
	AccountTypeId  uuid.UUID `json:"account_type_id"`
	InsertedAt     time.Time `json:"inserted_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateUserParams struct {
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Password       string    `json:"password"`
	MSISDN         string    `json:"msisdn"`
	EmailVerified  bool      `json:"email_verified"`
	MsisdnVerified bool      `json:"msisdn_verified"`
	Active         bool      `json:"active"`
	AccountTypeId  uuid.UUID `json:"account_type_id"`
}

type UpdateUserParams struct {
	ID uuid.UUID `json:"id"`
	CreateUserParams
}

type GetByEmailorMSISDNParams struct {
	Email string `validate:"email"`
	MSIDN string
}

type ListUsersFilter struct {
	ActiveOnly    bool
	AccountTypeId uuid.UUID
	Limit         int
	Offset        int
}

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserParams) (*User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByMSISDN(ctx context.Context, msisdn string) (*User, error)
}

type UserRepository interface {
	Create(ctx context.Context, req CreateUserParams) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByMSISDN(ctx context.Context, msisdn string) (*User, error)
}

func (u CreateUserParams) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email, is.LowerCase),
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 255), is.Alphanumeric),
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 255), is.Alphanumeric),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.AccountTypeId, validation.Required, is.UUID),
		validation.Field(&u.MSISDN, validation.When(u.MSISDN != "", validation.Match(regexp.MustCompile(`^\+\d{1,3}\d{9}$`)))),
	)
}

func ValidateEmail(email string) error {
	return validation.Validate(email, is.Email)
}
