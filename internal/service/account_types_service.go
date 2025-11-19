package service

import (
	"context"

	"github.com/Edu58/Oplan/internal/domain"
)

type AccountTypeService struct {
	repo domain.AccountTypeRepository
}

func NewAccountTypesService(repo domain.AccountTypeRepository) *AccountTypeService {
	return &AccountTypeService{repo: repo}
}

func (a *AccountTypeService) Create(ctx context.Context, params domain.AccountType) (domain.AccountType, error) {
	return a.repo.CreateAccountType(ctx, params)
}

func (a *AccountTypeService) List(ctx context.Context) ([]domain.AccountType, error) {
	return a.repo.ListAccountTypes(ctx)
}
