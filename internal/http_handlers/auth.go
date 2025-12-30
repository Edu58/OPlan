package httphandlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	templates "github.com/Edu58/Oplan/internal/frontend/templates/auth"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type SessionService interface {
	GetSessionBySessionId(ctx context.Context, sessionId uuid.UUID) (sqlc.Session, error)
	CreateSession(ctx context.Context, arg sqlc.CreateSessionParams) (sqlc.Session, error)
}

type UserService interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (sqlc.User, error)
}

type SessionsHandler struct {
	sessionService SessionService
	userService    UserService
	logger         logger.Logger
}

func NewSessionHandler(sessionService SessionService, userService UserService, logger logger.Logger) *SessionsHandler {
	return &SessionsHandler{sessionService, userService, logger}
}

func (s *SessionsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/auth/signin", http.HandlerFunc(s.signIn))
	mux.Handle("/auth/signup", http.HandlerFunc(s.signup))
	mux.Handle("/auth/verify-otp", http.HandlerFunc(s.verifyOTP))
}

func (s *SessionsHandler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			component := templates.ErrorMessage("Invalid")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
			}
			return
		}

		fieldValue := r.FormValue("email")

		user, err := s.userService.GetUserByEmail(r.Context(), fieldValue)

		if err != nil {
			w.WriteHeader(404)
			component := templates.ErrorMessage("User NOT found")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		w.Header().Set("HX-Push-Url", "/auth/verify-otp")

		component := templates.OTPVerification(user.Email, "email")
		err = component.Render(context.Background(), w)

		if err != nil {
			http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
			return
		}
		return
	}

	component := templates.AuthPage("Sign in", "")
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
		return
	}
	return
}

func (s *SessionsHandler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			component := templates.ErrorMessage("error processing request")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
				return
			}
		}

		lastName := r.PostForm.Get("lastName")

		params := sqlc.CreateUserParams{
			Email:     r.PostForm.Get("email"),
			FirstName: r.PostForm.Get("firstName"),
			LastName:  &lastName,
			Password:  r.PostForm.Get("password"),
		}

		user, err := s.userService.CreateUser(r.Context(), params)

		if err != nil {
			w.WriteHeader(400)
			component := templates.ErrorMessage("error processing request")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
				return
			}
		}

		s.logger.WithField("Email", user.Email).Info("Account created successfully")

		w.Header().Set("HX-Push-Url", "/auth/verify-otp")

		component := templates.OTPVerification(user.Email, "email")
		err = component.Render(context.Background(), w)

		if err != nil {
			http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
			return
		}
		return
	}

	component := templates.AuthPage("Sign up", "")
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
		return
	}
	return
}

func (s *SessionsHandler) verifyOTP(w http.ResponseWriter, r *http.Request) {
	component := templates.AuthPage("OTP", "")
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprint("error processing request"), http.StatusInternalServerError)
		return
	}
	return
}
