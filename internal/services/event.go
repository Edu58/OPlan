package services

import (
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
)

type EventService struct {
	_ domain.EventRepository
	_ logger.Logger
}

func NewEventService(repo domain.EventRepository, logger logger.Logger) *EventService {
	return &EventService{repo, logger}
}
