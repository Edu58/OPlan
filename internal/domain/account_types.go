package domain

import (
	"context"

	db "github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewAccountType(name string, active bool) *db.AccountType {
	return &db.AccountType{Name: name, Active: pgtype.Bool{Bool: active}}
}

type AccountTypeService interface {
	Create(ctx context.Context, arg db.CreateAccountTypeParams) (*db.AccountType, error)
	// Delete(ctx context.Context, id uuid.UUID) (AccountType, error)
	// GetById(ctx context.Context, id uuid.UUID) (AccountType, error)
	// GetByName(ctx context.Context, name string) (AccountType, error)
	List(ctx context.Context) ([]*db.AccountType, error)
	// UpdateByID(ctx context.Context, arg AccountType) (AccountType, error)
}

type AccountTypeRepository interface {
	CreateAccountType(ctx context.Context, arg db.CreateAccountTypeParams) (*db.AccountType, error)
	DeleteAccountType(ctx context.Context, id pgtype.UUID) (*db.AccountType, error)
	GetAccountTypeById(ctx context.Context, id pgtype.UUID) (*db.AccountType, error)
	GetAccountTypeByName(ctx context.Context, name string) (*db.AccountType, error)
	ListAccountTypes(ctx context.Context) ([]*db.AccountType, error)
	UpdateAccountTypeByID(ctx context.Context, arg db.UpdateAccountTypeByIDParams) (*db.AccountType, error)
}
