package domain

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/google/uuid"
)

type SessionRepository interface {
	GetSessionBySessionId(ctx context.Context, sessionID uuid.UUID) (sqlc.Session, error)
	CreateSession(ctx context.Context, arg sqlc.CreateSessionParams) (sqlc.Session, error)
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
}
