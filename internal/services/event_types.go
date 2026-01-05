package services

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type EventTypeService struct {
	repo   domain.EventTypesRepository
	logger logger.Logger
}

func NewEventTypeService(repo domain.EventTypesRepository, logger logger.Logger) *EventTypeService {
	return &EventTypeService{repo, logger}
}

func (e *EventTypeService) ListEventTypes(ctx context.Context) ([]sqlc.EventType, error) {
	return e.repo.ListEventTypes(ctx)
}

func (e *EventTypeService) CreateEventType(ctx context.Context, arg sqlc.CreateEventTypeParams) (sqlc.EventType, error) {
	return e.repo.CreateEventType(ctx, arg)
}

func (e *EventTypeService) GetEventTypeByName(ctx context.Context, name string) (sqlc.EventType, error) {
	return e.repo.GetEventTypeByName(ctx, name)
}

func (e *EventTypeService) GetEventTypeById(ctx context.Context, id uuid.UUID) (sqlc.EventType, error) {
	return e.repo.GetEventTypeById(ctx, id)
}

func (e *EventTypeService) UpdateEventTypeById(ctx context.Context, arg sqlc.UpdateEventTypeByIdParams) (sqlc.EventType, error) {
	return e.repo.UpdateEventTypeById(ctx, arg)
}

func (e *EventTypeService) DeleteEventType(ctx context.Context, id uuid.UUID) error {
	return e.repo.DeleteEventType(ctx, id)
}
