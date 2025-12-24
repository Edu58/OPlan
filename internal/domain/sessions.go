package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID    uuid.UUID
	SessionID string
	ClientIP  string
	IsBlocked bool
	ExpiresAt time.Time
}

type CreateSessionParams struct {
	UserID    uuid.UUID
	SessionID string
	ClientIP  string
	IsBlocked bool
	ExpiresAt time.Time
}

type UpdateSessionByIDParams struct {
	SessionID string
	ClientIP  string
	IsBlocked bool
	ExpiresAt time.Time
}

func (r *CreateSessionParams) Validate() error {
	return ValidateNotBeforeNow(r.ExpiresAt)
}

func (r *UpdateSessionByIDParams) Validate() error {
	return ValidateNotBeforeNow(r.ExpiresAt)
}

func ValidateNotBeforeNow(expiry time.Time) error {
	if time.Now().After(expiry) {
		return errors.New("Expiry cannot be before Now")
	}

	return nil
}

type SessionsService interface {
	GetSessionById(ctx context.Context, id uuid.UUID) (*Session, error)
	// GetBySessionId(ctx context.Context, session_id string) (*Session, error)
	CreateSession(ctx context.Context, params CreateSessionParams) (*Session, error)
	// UpdateSession(ctx context.Context, id uuid.UUID) (*Session, error)
	// DeleteSession(ctx context.Context, id uuid.UUID) error
}

type SessionsRepository interface {
	GetSessionById(ctx context.Context, id uuid.UUID) (*Session, error)
	GetBySessionId(ctx context.Context, session_id string) (*Session, error)
	CreateSession(ctx context.Context, params CreateSessionParams) (*Session, error)
	// UpdateSession(ctx context.Context, id uuid.UUID) (*Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID) error
}
