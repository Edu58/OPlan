package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Edu58/Oplan/internal/domain"
	templates "github.com/Edu58/Oplan/internal/frontend/templates/event"
	"github.com/Edu58/Oplan/pkg/logger"
)

type EventHandler struct {
	logger           logger.Logger
	eventService     domain.EventService
	eventTypeService domain.EventTypesService
	sessionService   domain.SessionService
}

func NewEventHandler(eventService domain.EventsService, eventTypeService domain.EventTypesService, sessionService domain.SessionService, logger logger.Logger) *EventHandler {
	return &EventHandler{logger, eventService, eventTypeService, sessionService}
}

// Registers all index page routes to the server mx
func (e *EventHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/event/{id}", http.HandlerFunc(e.index))
}

func (e *EventHandler) index(w http.ResponseWriter, r *http.Request) {
	email, _ := r.Context().Value("userEmail").(string)
	component := templates.EventDetail("Event", email)
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}
}
