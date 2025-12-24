package repository

import (
	"context"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepository struct {
	db *pgxpool.Pool
}

func NewSessionRepository(db *pgxpool.Pool) *SessionRepository {
	return &SessionRepository{db}
}

func (s *SessionRepository) GetSessionById(ctx context.Context, id uuid.UUID) (*domain.Session, error) {
	query := `
	SELECT * WHERE id = $1 LIMIT 1
	`
	row := s.db.QueryRow(ctx, query, id)

	var i domain.Session

	err := row.Scan(
		&i.UserID,
		&i.SessionID,
		&i.ClientIP,
		&i.IsBlocked,
		&i.ExpiresAt,
	)
	return &i, err
}

func (s *SessionRepository) GetBySessionId(ctx context.Context, session_id string) (*domain.Session, error) {
	query := `
	SELECT * WHERE session_id = $1 LIMIT 1
	`
	row := s.db.QueryRow(ctx, query, session_id)

	var i domain.Session

	err := row.Scan(
		&i.UserID,
		&i.SessionID,
		&i.ClientIP,
		&i.IsBlocked,
		&i.ExpiresAt,
	)
	return &i, err
}

func (s *SessionRepository) CreateSession(ctx context.Context, params domain.CreateSessionParams) (*domain.Session, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	query := `
	INSERT INTO sessions
		(user_id, session_id, client_ip, is_blocked, expires_at)
		VALUES($1, $2, $3, $4, $5)
	RETURNING *
	`

	row := s.db.QueryRow(ctx, query, params.UserID, params.SessionID, params.ClientIP, params.IsBlocked, params.ExpiresAt)

	var i domain.Session

	err := row.Scan(
		&i.UserID,
		&i.SessionID,
		&i.ClientIP,
		&i.IsBlocked,
		&i.ExpiresAt,
	)
	return &i, err
}

func (s *SessionRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {

	query := `
	DELETE FROM sessions
	WHERE id = $1
	RETURNING *
	`

	row := s.db.QueryRow(ctx, query, id)

	var i domain.Session

	err := row.Scan(
		&i.UserID,
		&i.SessionID,
		&i.ClientIP,
		&i.IsBlocked,
		&i.ExpiresAt,
	)
	return err
}
