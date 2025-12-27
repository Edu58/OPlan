package httphandlers

import (
	"context"
	"net/http"

	"github.com/Edu58/Oplan/internal/domain"
	templates "github.com/Edu58/Oplan/internal/frontend/templates/auth"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type AuthAccountTypeService interface {
	GetByName(ctx context.Context, name string) (*domain.AccountType, error)
}

type SessionService interface {
	GetSessionById(ctx context.Context, id uuid.UUID) (*domain.Session, error)
	CreateSession(ctx context.Context, params domain.CreateSessionParams) (*domain.Session, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req domain.CreateUserParams) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByMSISDN(ctx context.Context, msisdn string) (*domain.User, error)
}

type SessionsHandler struct {
	account_type_service AuthAccountTypeService
	session_service      SessionService
	user_service         UserService
	logger               logger.Logger
}

func NewSessionHandler(session_service SessionService, user_service UserService, account_type_service AuthAccountTypeService, logger logger.Logger) *SessionsHandler {
	return &SessionsHandler{account_type_service, session_service, user_service, logger}
}

func (s *SessionsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/auth/signin", http.HandlerFunc(s.signin))
	mux.Handle("/auth/signup", http.HandlerFunc(s.signup))
	mux.Handle("/auth/verify-otp", http.HandlerFunc(s.verifyOTP))
}

func (s *SessionsHandler) signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			component := templates.ErrorMessage("Invalid")
			component.Render(context.Background(), w)
			return
		}

		fieldValue := r.FormValue("email")

		if err := domain.ValidateEmail(fieldValue); err != nil {
			w.WriteHeader(400)
			component := templates.ErrorMessage("Email is invalid. Verify and try again")
			component.Render(context.Background(), w)
			return
		}

		user, err := s.user_service.GetUserByEmail(r.Context(), fieldValue)

		if err != nil {
			w.WriteHeader(404)
			component := templates.ErrorMessage("User NOT found")
			component.Render(context.Background(), w)
			return
		}

		w.Header().Set("HX-Push-Url", "/auth/verify-otp")

		component := templates.OTPVerification(user.Email, "email")
		component.Render(context.Background(), w)
		return
	}

	component := templates.AuthPage("Sign in", "")
	component.Render(context.Background(), w)
}

func (s *SessionsHandler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			s.logger.Err(err)

			w.WriteHeader(400)
			component := templates.ErrorMessage("Invalid")
			component.Render(context.Background(), w)
			return
		}

		account_type, err := s.account_type_service.GetByName(r.Context(), "user")

		if err != nil {
			s.logger.Error(err.Error())
			w.WriteHeader(404)
			component := templates.ErrorMessage("Account type NOT found")
			component.Render(context.Background(), w)
			return
		}

		params := domain.CreateUserParams{
			Email:         r.PostForm.Get("email"),
			FirstName:     r.PostForm.Get("firstName"),
			LastName:      r.PostForm.Get("lastName"),
			Password:      r.PostForm.Get("password"),
			MSISDN:        r.PostForm.Get("msisdn"),
			AccountTypeId: account_type.ID,
		}

		user, err := s.user_service.CreateUser(r.Context(), params)

		if err != nil {
			s.logger.Error(err.Error())
			w.WriteHeader(400)
			component := templates.ErrorMessage(err.Error())
			component.Render(context.Background(), w)
			return
		}

		s.logger.WithField("Email", user.Email).Info("Account created successfully")

		w.Header().Set("HX-Push-Url", "/auth/verify-otp")

		component := templates.OTPVerification(user.Email, "email")
		component.Render(context.Background(), w)
		return
	}

	component := templates.AuthPage("Sign up", "")
	component.Render(context.Background(), w)
}

func (s *SessionsHandler) verifyOTP(w http.ResponseWriter, r *http.Request) {
	component := templates.AuthPage("OTP", "")
	component.Render(context.Background(), w)
}
