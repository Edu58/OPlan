package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email" validate:"required,email"`
	Username       string    `json:"username" validate:"required,gte=3,lte=130"`
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	Password       string    `json:"password" validate:"required"`
	MSISDN         string    `json:"msisdn" validate:"required,mobile"`
	Dob            time.Time `json:"dob" validate:"required,date"`
	EmailVerified  bool      `json:"email_verified"`
	MsisdnVerified bool      `json:"msisdn_verified"`
	Active         bool      `json:"active"`
	AccountTypeId  uuid.UUID `json:"account_type_id"`
	InsertedAt     time.Time `json:"inserted_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateUserParams struct {
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Password       string    `json:"password"`
	MSISDN         string    `json:"msisdn"`
	Dob            time.Time `json:"dob"`
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
