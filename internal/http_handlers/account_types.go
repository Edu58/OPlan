package httphandlers

import (
	"context"
	"net/http"

	"github.com/Edu58/Oplan/internal/domain"
	templates "github.com/Edu58/Oplan/internal/frontend/templates/auth"
)

type AccountTypeService interface {
	Create(ctx context.Context, params domain.CreateAccountTypeParams) (*domain.AccountType, error)
	ListAll(ctx context.Context) ([]*domain.AccountType, error)
}

type AccountTypesHandler struct {
	service AccountTypeService
}

func NewAccountTypesHandler(service AccountTypeService) *AccountTypesHandler {
	return &AccountTypesHandler{service: service}
}

func (a *AccountTypesHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(a.listAccountTypes))
}

func (a *AccountTypesHandler) listAccountTypes(w http.ResponseWriter, r *http.Request) {
	// account_types, err := a.service.ListAll(r.Context())

	component := templates.AuthPage("Sign In")
	component.Render(context.Background(), w)
}
