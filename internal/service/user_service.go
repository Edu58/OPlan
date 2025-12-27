package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	if err := req.Validate(); err != nil {
		u.logger.Err(err)
		return nil, err
	}

	user, err := u.repo.GetByEmail(ctx, req.Email)

	if user != nil {
		return nil, fmt.Errorf("User with email '%s' already exists", req.Email)
	}

	if err != nil && err == pgx.ErrNoRows {
		newUser, err := u.repo.Create(ctx, req)

		if err != nil {
			u.logger.Err(fmt.Errorf("Error creating user: %v", err))
			return nil, errors.New("Error creating user. Verify details and try again")
		}

		return newUser, err
	}

	u.logger.Err(fmt.Errorf("Error during email lookup: %v", err))
	return nil, fmt.Errorf("Error verifying email. Verify details and try again")
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
