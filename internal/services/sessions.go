package services

import (
	"context"
	"fmt"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type SessionRepository interface {
	GetSessionBySessionId(ctx context.Context, sessionID uuid.UUID) (sqlc.Session, error)
	CreateSession(ctx context.Context, arg sqlc.CreateSessionParams) (sqlc.Session, error)
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
}

type SessionService struct {
	repo   SessionRepository
	logger logger.Logger
}

func NewSessionService(repo SessionRepository, logger logger.Logger) *SessionService {
	return &SessionService{repo: repo, logger: logger}
}

func (s *SessionService) CreateSession(ctx context.Context, params sqlc.CreateSessionParams) (sqlc.Session, error) {
	sess, err := s.repo.CreateSession(ctx, params)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error creating session: %v", err))
		return sqlc.Session{}, err
	}

	return sess, nil
}

func (s *SessionService) GetSessionBySessionId(ctx context.Context, sessionId uuid.UUID) (sqlc.Session, error) {
	sess, err := s.repo.GetSessionBySessionId(ctx, sessionId)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error getting account types: %v", err))
		return sqlc.Session{}, err
	}

	return sess, nil
}

func (s *SessionService) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	err := s.repo.DeleteSession(ctx, sessionID)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error deleting session: %v", err))
		return err
	}

	return nil
}
