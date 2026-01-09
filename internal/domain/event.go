package domain

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/google/uuid"
)

type EventRepository interface {
	ListEvents(ctx context.Context, arg sqlc.ListEventsParams) ([]sqlc.ListEventsRow, error)
	GetEventById(ctx context.Context, id uuid.UUID) (sqlc.Event, error)
	GetEventByName(ctx context.Context, name string) (sqlc.Event, error)
	CreateEvent(ctx context.Context, arg sqlc.CreateEventParams) (sqlc.Event, error)
	UpdateEventById(ctx context.Context, arg sqlc.UpdateEventByIdParams) (sqlc.Event, error)
	DeleteEventById(ctx context.Context, id uuid.UUID) error
}

type EventService interface {
	EventsRepository
}
