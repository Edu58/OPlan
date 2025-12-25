package service

import (
	"context"
	"fmt"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
)

type AccountTypeRepository interface {
	GetByName(ctx context.Context, name string) (*domain.AccountType, error)
	Create(ctx context.Context, req domain.CreateAccountTypeParams) (*domain.AccountType, error)
	List(ctx context.Context) ([]*domain.AccountType, error)
}

type AccountTypeService struct {
	repo   AccountTypeRepository
	logger logger.Logger
}

func NewAccountTypesService(repo AccountTypeRepository, logger logger.Logger) *AccountTypeService {
	return &AccountTypeService{repo: repo, logger: logger}
}

func (a *AccountTypeService) Create(ctx context.Context, params domain.CreateAccountTypeParams) (*domain.AccountType, error) {
	acc_type, err := a.repo.Create(ctx, params)

	if err != nil {
		a.logger.Error(fmt.Sprintf("Error creating account: %v", err))
		return nil, err
	}

	return acc_type, nil
}

func (a *AccountTypeService) ListAll(ctx context.Context) ([]*domain.AccountType, error) {
	acc_types, err := a.repo.List(ctx)

	if err != nil {
		a.logger.Error(fmt.Sprintf("Error getting account types: %v", err))
		return nil, err
	}

	return acc_types, nil
}

func (a *AccountTypeService) GetByName(ctx context.Context, name string) (*domain.AccountType, error) {
	return a.repo.GetByName(ctx, name)
}
