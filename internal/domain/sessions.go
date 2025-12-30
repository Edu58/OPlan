package domain

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/google/uuid"
)

type SessionsService interface {
	GetSessionById(ctx context.Context, id uuid.UUID) (*sqlc.Session, error)
	// GetBySessionId(ctx context.Context, session_id string) (*Session, error)
	CreateSession(ctx context.Context, params sqlc.CreateSessionParams) (*sqlc.Session, error)
	// UpdateSession(ctx context.Context, id uuid.UUID) (*Session, error)
	// DeleteSession(ctx context.Context, id uuid.UUID) error
}

type SessionsRepository interface {
	GetSessionById(ctx context.Context, id uuid.UUID) (*sqlc.Session, error)
	GetBySessionId(ctx context.Context, session_id string) (*sqlc.Session, error)
	CreateSession(ctx context.Context, params sqlc.CreateSessionParams) (*sqlc.Session, error)
	// UpdateSession(ctx context.Context, id uuid.UUID) (*Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID) error
}
