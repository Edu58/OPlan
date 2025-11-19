package repository

import (
	"context"
	"log"

	db "github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountTypeRepository struct {
	db *db.Queries
}

func NewAccountTypeRepository(querier *db.Queries) domain.AccountTypeRepository {
	return &AccountTypeRepository{querier}
}

func (r *AccountTypeRepository) CreateAccountType(ctx context.Context, arg domain.AccountType) (domain.AccountType, error) {
	params := db.CreateAccountTypeParams{Name: arg.Name, Active: pgtype.Bool{Bool: arg.Active, Valid: false}}
	account_type, err := r.db.CreateAccountType(ctx, params)

	if err != nil {
		log.Printf("Error creating account type '%s' with error: %v", arg.Name, err)
		return domain.AccountType{}, nil
	}

	return domain.AccountType{
		ID:         account_type.ID.Bytes,
		Name:       account_type.Name,
		Active:     account_type.Active.Bool,
		InsertedAt: account_type.InsertedAt.Time,
		UpdatedAt:  account_type.UpdatedAt.Time,
	}, nil
}

func (r *AccountTypeRepository) DeleteAccountType(ctx context.Context, id uuid.UUID) (domain.AccountType, error) {
	account_type, err := r.db.DeleteAccountType(ctx, pgtype.UUID{Bytes: id, Valid: false})

	if err != nil {
		log.Printf("Error deleting account type with error: %v", err)
		return domain.AccountType{}, nil
	}

	return domain.AccountType{
		ID:         account_type.ID.Bytes,
		Name:       account_type.Name,
		Active:     account_type.Active.Bool,
		InsertedAt: account_type.InsertedAt.Time,
		UpdatedAt:  account_type.UpdatedAt.Time,
	}, nil
}

func (r *AccountTypeRepository) GetAccountTypeById(ctx context.Context, id uuid.UUID) (domain.AccountType, error) {
	account_type, err := r.db.GetAccountTypeById(ctx, pgtype.UUID{Bytes: id, Valid: false})

	if err != nil {
		log.Printf("Error getting account type by ID with error: %v", err)
		return domain.AccountType{}, nil
	}

	return domain.AccountType{
		ID:         account_type.ID.Bytes,
		Name:       account_type.Name,
		Active:     account_type.Active.Bool,
		InsertedAt: account_type.InsertedAt.Time,
		UpdatedAt:  account_type.UpdatedAt.Time,
	}, nil
}

func (r *AccountTypeRepository) GetAccountTypeByName(ctx context.Context, name string) (domain.AccountType, error) {
	account_type, err := r.db.GetAccountTypeByName(ctx, name)

	if err != nil {
		log.Printf("Error getting account type by name '%s' with error: %v", name, err)
		return domain.AccountType{}, nil
	}

	return domain.AccountType{
		ID:         account_type.ID.Bytes,
		Name:       account_type.Name,
		Active:     account_type.Active.Bool,
		InsertedAt: account_type.InsertedAt.Time,
		UpdatedAt:  account_type.UpdatedAt.Time,
	}, nil
}

func (r *AccountTypeRepository) ListAccountTypes(ctx context.Context) ([]domain.AccountType, error) {

	account_types, err := r.db.ListAccountTypes(ctx)

	if err != nil {
		log.Printf("Error getting getting account types with error: %v", err)
		return nil, nil
	}

	accountTypes := make([]domain.AccountType, 0, len(account_types))

	for _, t := range account_types {
		accountTypes = append(accountTypes, domain.AccountType{
			ID:         t.ID.Bytes,
			Name:       t.Name,
			Active:     t.Active.Bool,
			InsertedAt: t.InsertedAt.Time,
			UpdatedAt:  t.UpdatedAt.Time,
		})
	}

	return accountTypes, nil
}

func (r *AccountTypeRepository) UpdateAccountTypeByID(ctx context.Context, arg domain.AccountType) (domain.AccountType, error) {
	params := db.UpdateAccountTypeByIDParams{ID: pgtype.UUID{Bytes: arg.ID, Valid: false}, Name: arg.Name, Active: pgtype.Bool{Bool: arg.Active, Valid: false}}
	account_type, err := r.db.UpdateAccountTypeByID(ctx, params)

	if err != nil {
		log.Printf("Error updating account type by ID '%s' with error: %v", arg.Name, err)
		return domain.AccountType{}, nil
	}

	return domain.AccountType{
		ID:         account_type.ID.Bytes,
		Name:       account_type.Name,
		Active:     account_type.Active.Bool,
		InsertedAt: account_type.InsertedAt.Time,
		UpdatedAt:  account_type.UpdatedAt.Time,
	}, nil
}
