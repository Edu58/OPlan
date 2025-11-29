package service

import (
	"context"
	"fmt"

	db "github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
)

type AccountTypeService struct {
	repo   domain.AccountTypeRepository
	logger logger.Logger
}

func NewAccountTypesService(repo domain.AccountTypeRepository) *AccountTypeService {
	return &AccountTypeService{repo: repo}
}

func (a *AccountTypeService) Create(ctx context.Context, params db.CreateAccountTypeParams) (*db.AccountType, error) {
	acc_type, err := a.repo.CreateAccountType(ctx, params)

	if err != nil {
		a.logger.Error(fmt.Sprintf("Error creating account: %v", err))
		return nil, err
	}

	return acc_type, nil
}

func (a *AccountTypeService) List(ctx context.Context) ([]*db.AccountType, error) {
	acc_types, err := a.repo.ListAccountTypes(ctx)

	if err != nil {
		a.logger.Error(fmt.Sprintf("Error getting account types: %v", err))
		return nil, err
	}

	return acc_types, nil
}
