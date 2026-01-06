package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	templates "github.com/Edu58/Oplan/internal/frontend/templates/index"
	"github.com/Edu58/Oplan/pkg/logger"
)

const (
	DefaultEventsPageNum  = 1
	DefaultEventsPageSize = 15
)

type IndexHandler struct {
	logger           logger.Logger
	eventService     domain.EventService
	eventTypeService domain.EventTypesService
	sessionService   domain.SessionService
}

func NewIndexHandler(eventService domain.EventService, eventTypeService domain.EventTypesService, sessionService domain.SessionService, logger logger.Logger) *IndexHandler {
	return &IndexHandler{logger, eventService, eventTypeService, sessionService}
}

// Registers all index page routes to the server mx
func (i *IndexHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(i.index))
}

func (s *IndexHandler) index(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters with defaults and validation
	page := parseIntQuery(r, "page", DefaultEventsPageNum, 1)
	pageSize := parseIntQuery(r, "page_size", DefaultEventsPageSize, 1)

	event_types, _ := s.eventTypeService.ListEventTypes(r.Context())

	events, _ := s.eventService.ListEvents(r.Context(), sqlc.ListEventsParams{Limit: int32(page), Offset: int32(pageSize)})

	email, _ := r.Context().Value("userEmail").(string)

	component := templates.Index("Oplan", email, &events, &event_types)
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}
}

// Helper function (add to your handler struct or utils package)
func parseIntQuery(r *http.Request, key string, defaultVal, min int) int {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultVal
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed < min {
		return defaultVal
	}

	return parsed
}
