package services

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type EventsService struct {
	repo   domain.EventRepository
	logger logger.Logger
}

func NewEventsService(repo domain.EventRepository, logger logger.Logger) *EventsService {
	return &EventsService{repo, logger}
}

func (e *EventsService) ListEvents(ctx context.Context, arg sqlc.ListEventsParams) ([]sqlc.ListEventsRow, error) {
	return e.repo.ListEvents(ctx, arg)
}

func (e *EventsService) CreateEvent(ctx context.Context, arg sqlc.CreateEventParams) (sqlc.Event, error) {
	return e.repo.CreateEvent(ctx, arg)
}

func (e *EventsService) GetEventByName(ctx context.Context, name string) (sqlc.Event, error) {
	return e.repo.GetEventByName(ctx, name)
}

func (e *EventsService) GetEventById(ctx context.Context, id uuid.UUID) (sqlc.Event, error) {
	return e.repo.GetEventById(ctx, id)
}

func (e *EventsService) UpdateEventById(ctx context.Context, arg sqlc.UpdateEventByIdParams) (sqlc.Event, error) {
	return e.repo.UpdateEventById(ctx, arg)
}

func (e *EventsService) DeleteEventById(ctx context.Context, id uuid.UUID) error {
	return e.repo.DeleteEventById(ctx, id)
}
