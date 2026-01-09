package handlers

import (
	"context"
	"fmt"
	"net/http"

	components "github.com/Edu58/Oplan/internal/frontend/components/shared"
)

func renderHXError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	newComponent := components.ErrorMessage(err.Error())
	err = newComponent.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), status)
		return
	}
}
