package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/crypto"
	"github.com/Edu58/Oplan/shared/generators"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	components "github.com/Edu58/Oplan/internal/frontend/components/shared"
	templates "github.com/Edu58/Oplan/internal/frontend/templates/auth"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/google/uuid"
)

type SessionsHandler struct {
	sessionService domain.SessionService
	userService    domain.UserService
	otpService     domain.OTPService
	logger         logger.Logger
}

func NewSessionHandler(sessionService domain.SessionService, userService domain.UserService, otpService domain.OTPService, logger logger.Logger) *SessionsHandler {
	return &SessionsHandler{sessionService, userService, otpService, logger}
}

func (s *SessionsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/auth/signin", http.HandlerFunc(s.signIn))
	mux.Handle("/auth/signup", http.HandlerFunc(s.signup))
	mux.Handle("/auth/verify-otp", http.HandlerFunc(s.verifyOTP))
	mux.Handle("/auth/signout", http.HandlerFunc(s.signout))
}

func (s *SessionsHandler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			component := components.ErrorMessage("Invalid")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
			}

			return
		}

		fieldValue := r.FormValue("email")

		err := domain.ValidateEmail(fieldValue)

		if err != nil {
			w.WriteHeader(404)
			component := components.ErrorMessage(err.Error())
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		user, err := s.userService.GetUserByEmail(r.Context(), fieldValue)

		if err != nil {
			w.WriteHeader(404)
			component := components.ErrorMessage("User NOT found")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		otpExpiry := time.Now().Add(3 * time.Minute)

		err = generateOTP(r.Context(), s.otpService, user.Email, &otpExpiry)

		if err != nil {
			s.logger.Err(err)
			http.Error(w, fmt.Sprintln("error sending otp"), http.StatusInternalServerError)
			return
		}

		cookie := &http.Cookie{
			Name:     "auth",
			Value:    user.Email,
			Path:     "/auth/verify-otp",
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   300,
			Expires:  time.Time.Add(time.Now(), 5*time.Minute),
		}

		http.SetCookie(w, cookie)

		w.Header().Set("HX-Push-Url", "/auth/verify-otp")

		component := templates.OTPVerification(user.Email, "email")
		err = component.Render(context.Background(), w)

		if err != nil {
			http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
			return
		}

		return
	}

	component := templates.AuthPage("Sign in", "")
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}

}

func (s *SessionsHandler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			component := components.ErrorMessage("error processing request")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		lastName := r.PostForm.Get("lastName")

		params := sqlc.CreateUserParams{
			Email:     r.PostForm.Get("email"),
			FirstName: r.PostForm.Get("firstName"),
			LastName:  &lastName,
			Password:  r.PostForm.Get("password"),
		}

		err := domain.ValidateCreateUser(params)

		user, err := s.userService.CreateUser(r.Context(), params)

		if err != nil {
			var errMsg = "error processing request"
			var validationErr validation.Errors

			if errors.As(err, &validationErr) {
				errMsg = err.Error()
			}

			w.WriteHeader(400)
			component := components.ErrorMessage(errMsg)
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		s.logger.WithField("Email", user.Email).Info("Account created successfully")

		otpExpiry := time.Now().Add(3 * time.Minute)

		err = generateOTP(r.Context(), s.otpService, user.Email, &otpExpiry)

		if err != nil {
			http.Error(w, fmt.Sprintln("error sending otp"), http.StatusInternalServerError)
			return
		}

		cookie := &http.Cookie{
			Name:     "auth",
			Value:    user.Email,
			Path:     "/auth/verify-otp",
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   300,
			Expires:  time.Time.Add(time.Now(), 5*time.Minute),
		}

		http.SetCookie(w, cookie)

		w.Header().Set("HX-Push-Url", "/auth/verify-otp")

		component := templates.OTPVerification(user.Email, "email")
		err = component.Render(context.Background(), w)

		if err != nil {
			http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
			return
		}

		return
	}

	component := templates.AuthPage("Sign up", "")
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}
}

func (s *SessionsHandler) verifyOTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			component := components.ErrorMessage("Invalid OTP")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		authCookie, err := r.Cookie("auth")

		if err != nil {
			http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
			return
		}

		otp, err := s.otpService.GetOTP(r.Context(), authCookie.Value)

		if err != nil {
			http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
			return
		}

		rawOTP := domain.UserOTP{
			OTP1: r.FormValue("otp1"),
			OTP2: r.FormValue("otp2"),
			OTP3: r.FormValue("otp3"),
			OTP4: r.FormValue("otp4"),
			OTP5: r.FormValue("otp5"),
			OTP6: r.FormValue("otp6"),
		}

		err = rawOTP.ValidateUserOTP()

		if err != nil {
			w.WriteHeader(400)
			component := components.ErrorMessage(err.Error())
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		userOTP := strings.Join([]string{rawOTP.OTP1, rawOTP.OTP2, rawOTP.OTP3, rawOTP.OTP4, rawOTP.OTP5, rawOTP.OTP6}, "")

		otpHash := crypto.HashStringSHA512(userOTP)

		if time.Now().After(*otp.ExpiresAt) || otpHash != otp.Value {
			w.WriteHeader(400)
			component := components.ErrorMessage("Invalid OTP")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		user, err := s.userService.GetUserByEmail(r.Context(), authCookie.Value)

		if err != nil {
			w.WriteHeader(404)
			component := components.ErrorMessage("User NOT found")
			err = component.Render(context.Background(), w)

			if err != nil {
				http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
				return
			}

			return
		}

		sessExpiry := time.Time.Add(time.Now(), time.Hour)

		userIP := ReadUserIP(r)

		session, err := s.sessionService.CreateSession(r.Context(), sqlc.CreateSessionParams{
			UserID:    user.ID,
			ClientIp:  &userIP,
			IsBlocked: false,
			ExpiresAt: &sessExpiry,
		})

		if err != nil {
			s.logger.Err(err)
			http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth",
			Value:    "",
			MaxAge:   -1,
			Path:     "/",
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "oplan_knob",
			Value:    session.SessionID.String(),
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   3600,
			Expires:  sessExpiry,
		})

		s.logger.WithField("Email", user.Email).
			Info("User logged in successfully")

		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
		return
	}

	component := templates.AuthPage("OTP", "")
	err := component.Render(context.Background(), w)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}

}

func (s *SessionsHandler) signout(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("oplan_knob")

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}

	sess, err := s.sessionService.GetSessionBySessionId(r.Context(), uuid.MustParse(authCookie.Value))

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}

	err = s.sessionService.DeleteSession(r.Context(), sess.SessionID)

	if err != nil {
		http.Error(w, fmt.Sprintln("error processing request"), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})

	s.logger.WithField("Session ID", sess.UserID).
		Info("User logged out successfully")

	http.Redirect(w, r, "/", http.StatusFound)
}

func generateOTP(ctx context.Context, otpService domain.OTPService, email string, expiry *time.Time) (err error) {
	otp, err := generators.GenerateCode(6)
	fmt.Printf("GENERATED OTP %s", otp)

	_, err = otpService.CreateOTP(ctx, sqlc.CreateOTPParams{
		Identifier: email,
		Value:      crypto.HashStringSHA512(otp),
		ExpiresAt:  expiry,
	})

	return
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
