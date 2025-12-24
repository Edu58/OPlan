package httphandlers

import (
	"context"
	"net/http"

	"github.com/Edu58/Oplan/internal/domain"
	templates "github.com/Edu58/Oplan/internal/frontend/templates/auth"
	"github.com/google/uuid"
)

type SessionService interface {
	GetSessionById(ctx context.Context, id uuid.UUID) (*domain.Session, error)
	CreateSession(ctx context.Context, params domain.CreateSessionParams) (*domain.Session, error)
}

type SessionsHandler struct {
	service SessionService
}

func NewSessionHandler(service SessionService) *SessionsHandler {
	return &SessionsHandler{service}
}

func (s *SessionsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/auth/signin", http.HandlerFunc(s.signin))
	mux.Handle("/auth/signup", http.HandlerFunc(s.signup))
	mux.Handle("/auth/signin-form", http.HandlerFunc(s.signinForm))
	mux.Handle("/auth/signup-form", http.HandlerFunc(s.signupForm))
}

func (s *SessionsHandler) signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(400)
		component := templates.ErrorMessage("Not valid")
		component.Render(context.Background(), w)
		return
	} else {
		component := templates.AuthPage("Sign in")
		component.Render(context.Background(), w)
		return
	}
}

func (s *SessionsHandler) signup(w http.ResponseWriter, r *http.Request) {
	component := templates.AuthPage("Sign up")
	component.Render(context.Background(), w)
}

func (s *SessionsHandler) signinForm(w http.ResponseWriter, r *http.Request) {
	component := templates.SignInForm("Sign in")
	component.Render(context.Background(), w)
}

func (s *SessionsHandler) signupForm(w http.ResponseWriter, r *http.Request) {
	component := templates.SignUpForm("Sign up")
	component.Render(context.Background(), w)
}
