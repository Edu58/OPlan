package httphandlers

import (
	"context"
	"fmt"
	"net/http"

	templates "github.com/Edu58/Oplan/internal/frontend/templates/index"
	"github.com/Edu58/Oplan/pkg/logger"
)

type IndexHandler struct {
	logger logger.Logger
}

func NewIndexHandler(logger logger.Logger) *IndexHandler {
	return &IndexHandler{logger}
}

func (i *IndexHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(i.index))
}

func (s *IndexHandler) index(w http.ResponseWriter, r *http.Request) {
	component := templates.Index("Oplan", false, "edu@mail.com")
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}
}
