package repository

import (
	"context"

	db "github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountTypeRepository struct {
	repo db.Querier
}

func NewAccountTypeRepository(querier domain.AccountTypeRepository) domain.AccountTypeRepository {
	return &AccountTypeRepository{querier}
}

func (r *AccountTypeRepository) CreateAccountType(ctx context.Context, arg db.CreateAccountTypeParams) (*db.AccountType, error) {
	return r.repo.CreateAccountType(ctx, arg)
}

func (r *AccountTypeRepository) DeleteAccountType(ctx context.Context, id pgtype.UUID) (*db.AccountType, error) {
	return r.repo.DeleteAccountType(ctx, id)
}

func (r *AccountTypeRepository) GetAccountTypeById(ctx context.Context, id pgtype.UUID) (*db.AccountType, error) {
	return r.repo.GetAccountTypeById(ctx, id)
}

func (r *AccountTypeRepository) GetAccountTypeByName(ctx context.Context, name string) (*db.AccountType, error) {
	return r.repo.GetAccountTypeByName(ctx, name)
}

func (r *AccountTypeRepository) ListAccountTypes(ctx context.Context) ([]*db.AccountType, error) {
	return r.repo.ListAccountTypes(ctx)
}

func (r *AccountTypeRepository) UpdateAccountTypeByID(ctx context.Context, arg db.UpdateAccountTypeByIDParams) (*db.AccountType, error) {
	return r.repo.UpdateAccountTypeByID(ctx, arg)
}
