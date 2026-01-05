package services

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type EventService struct {
	repo   domain.EventRepository
	logger logger.Logger
}

func NewEventService(repo domain.EventRepository, logger logger.Logger) *EventService {
	return &EventService{repo, logger}
}

func (e *EventService) ListEvents(ctx context.Context, arg sqlc.ListEventsParams) ([]sqlc.ListEventsRow, error) {
	return e.repo.ListEvents(ctx, arg)
}

func (e *EventService) CreateEvent(ctx context.Context, arg sqlc.CreateEventParams) (sqlc.Event, error) {
	return e.repo.CreateEvent(ctx, arg)
}

func (e *EventService) GetEventByName(ctx context.Context, name string) (sqlc.Event, error) {
	return e.repo.GetEventByName(ctx, name)
}

func (e *EventService) GetEventById(ctx context.Context, id uuid.UUID) (sqlc.Event, error) {
	return e.repo.GetEventById(ctx, id)
}

func (e *EventService) UpdateEventById(ctx context.Context, arg sqlc.UpdateEventByIdParams) (sqlc.Event, error) {
	return e.repo.UpdateEventById(ctx, arg)
}

func (e *EventService) DeleteEventById(ctx context.Context, id uuid.UUID) error {
	return e.repo.DeleteEventById(ctx, id)
}
