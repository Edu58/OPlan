package service

import (
	"context"
	"fmt"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type SessionRepository interface {
	GetSessionById(ctx context.Context, id uuid.UUID) (*domain.Session, error)
	CreateSession(ctx context.Context, params domain.CreateSessionParams) (*domain.Session, error)
}

type SessionService struct {
	repo   SessionRepository
	logger logger.Logger
}

func NewSessionService(repo SessionRepository, logger logger.Logger) *SessionService {
	return &SessionService{repo: repo, logger: logger}
}

func (s *SessionService) CreateSession(ctx context.Context, params domain.CreateSessionParams) (*domain.Session, error) {
	sess, err := s.repo.CreateSession(ctx, params)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error creating session: %v", err))
		return nil, err
	}

	return sess, nil
}

func (s *SessionService) GetSessionById(ctx context.Context, session_id uuid.UUID) (*domain.Session, error) {
	sess, err := s.repo.GetSessionById(ctx, session_id)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error getting account types: %v", err))
		return nil, err
	}

	return sess, nil
}
