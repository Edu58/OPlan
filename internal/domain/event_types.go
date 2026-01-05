package domain

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/google/uuid"
)

type EventTypesRepository interface {
	ListEventTypes(ctx context.Context) ([]sqlc.EventType, error)
	CreateEventType(ctx context.Context, arg sqlc.CreateEventTypeParams) (sqlc.EventType, error)
	GetEventTypeById(ctx context.Context, id uuid.UUID) (sqlc.EventType, error)
	GetEventTypeByName(ctx context.Context, name string) (sqlc.EventType, error)
	UpdateEventTypeById(ctx context.Context, arg sqlc.UpdateEventTypeByIdParams) (sqlc.EventType, error)
	DeleteEventType(ctx context.Context, id uuid.UUID) error
}

type EventTypesService interface {
	EventTypesRepository
}
