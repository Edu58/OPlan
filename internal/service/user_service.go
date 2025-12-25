package service

import (
	"context"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, req domain.CreateUserParams) (*domain.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByMSISDN(ctx context.Context, msisdn string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserService struct {
	repo   UserRepository
	logger logger.Logger
}

func NewUserService(repo UserRepository, logger logger.Logger) *UserService {
	return &UserService{repo, logger}
}

func (u *UserService) CreateUser(ctx context.Context, req domain.CreateUserParams) (*domain.User, error) {
	return u.repo.Create(ctx, req)
}

func (u *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return u.repo.GetByEmail(ctx, email)
}

func (u *UserService) GetUserByMSISDN(ctx context.Context, msisdn string) (*domain.User, error) {
	return u.repo.GetByMSISDN(ctx, msisdn)
}
