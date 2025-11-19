package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AccountType struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Active     bool      `json:"active"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewAccountType(name string, active bool) *AccountType {
	return &AccountType{Name: name, Active: active}
}

type AccountTypeService interface {
	Create(ctx context.Context, arg AccountType) (AccountType, error)
	// Delete(ctx context.Context, id uuid.UUID) (AccountType, error)
	// GetById(ctx context.Context, id uuid.UUID) (AccountType, error)
	// GetByName(ctx context.Context, name string) (AccountType, error)
	List(ctx context.Context) ([]AccountType, error)
	// UpdateByID(ctx context.Context, arg AccountType) (AccountType, error)
}

type AccountTypeRepository interface {
	CreateAccountType(ctx context.Context, arg AccountType) (AccountType, error)
	DeleteAccountType(ctx context.Context, id uuid.UUID) (AccountType, error)
	GetAccountTypeById(ctx context.Context, id uuid.UUID) (AccountType, error)
	GetAccountTypeByName(ctx context.Context, name string) (AccountType, error)
	ListAccountTypes(ctx context.Context) ([]AccountType, error)
	UpdateAccountTypeByID(ctx context.Context, arg AccountType) (AccountType, error)
}
