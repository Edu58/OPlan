package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (sqlc.User, error)
}

type UserService struct {
	repo   UserRepository
	logger logger.Logger
}

func NewUserService(repo UserRepository, logger logger.Logger) *UserService {
	return &UserService{repo, logger}
}

func (u *UserService) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error) {
	// if err := domain.Validate(req); err != nil {
	// 	u.logger.Err(err)
	// 	return nil, err
	// }

	user, err := u.repo.GetUserByEmail(ctx, params.Email)

	if user.Active != nil {
		return sqlc.User{}, fmt.Errorf("user with email '%s' already exists", params.Email)
	}

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		newUser, err := u.repo.CreateUser(ctx, params)

		if err != nil {
			u.logger.Err(fmt.Errorf("error creating user: %v", err))
			return sqlc.User{}, errors.New("error creating user. Verify details and try again")
		}

		return newUser, err
	}

	u.logger.Err(fmt.Errorf("error during email lookup: %v", err))
	return sqlc.User{}, fmt.Errorf("error verifying email. Verify details and try again")
}

func (u *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (sqlc.User, error) {
	return u.repo.GetUserById(ctx, id)
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}
