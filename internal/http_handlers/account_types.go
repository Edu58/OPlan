package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/Edu58/Oplan/internal/domain"
)

type AccountTypesHandler struct {
	service domain.AccountTypeService
}

func NewAccountTypesHandler(service domain.AccountTypeService) *AccountTypesHandler {
	return &AccountTypesHandler{service: service}
}

func (a *AccountTypesHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(a.listAccountTypes))
}

func (a *AccountTypesHandler) listAccountTypes(w http.ResponseWriter, r *http.Request) {
	account_types, err := a.service.ListAll(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(account_types)

}
