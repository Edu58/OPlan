package domain

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/google/uuid"
)

type SessionRepository interface {
	GetSessionBySessionId(ctx context.Context, sessionID uuid.UUID) (sqlc.Session, error)
	GetSessionWithUserBySessionId(ctx context.Context, sessionID uuid.UUID) (sqlc.GetSessionWithUserBySessionIdRow, error)
	CreateSession(ctx context.Context, arg sqlc.CreateSessionParams) (sqlc.Session, error)
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
}

type SessionService interface {
	GetSessionWithUserBySessionId(ctx context.Context, sessionID uuid.UUID) (sqlc.GetSessionWithUserBySessionIdRow, error)
	GetSessionBySessionId(ctx context.Context, sessionId uuid.UUID) (sqlc.Session, error)
	CreateSession(ctx context.Context, arg sqlc.CreateSessionParams) (sqlc.Session, error)
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
}

type UserService interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (sqlc.User, error)
}

type OTPService interface {
	CreateOTP(ctx context.Context, arg sqlc.CreateOTPParams) (sqlc.OtpStore, error)
	GetOTP(ctx context.Context, identifier string) (sqlc.OtpStore, error)
	UpdateOTP(ctx context.Context, arg sqlc.UpdateOTPParams) (sqlc.OtpStore, error)
	DeleteOTP(ctx context.Context, identifier string) error
}
