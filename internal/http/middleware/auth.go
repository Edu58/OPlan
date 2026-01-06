package middleware

import (
	"context"
	"net/http"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/google/uuid"
)

func RequireAuth(next http.Handler, sessionService domain.SessionService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("oplan_knob")

		if err != nil {
			http.Redirect(w, r, "/auth/signin", http.StatusForbidden)
			return
		}

		validUUID, err := uuid.Parse(authCookie.Value)

		if err != nil {
			http.Redirect(w, r, "/auth/signin", http.StatusForbidden)
			return
		}

		session, err := sessionService.GetSessionWithUserBySessionId(r.Context(), validUUID)

		if err != nil {
			http.Redirect(w, r, "/auth/signin", http.StatusForbidden)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "authenticated", true)
		ctx = context.WithValue(ctx, "userId", session.UserID)
		ctx = context.WithValue(ctx, "sessionId", session.SessionID)
		ctx = context.WithValue(ctx, "clientIP", session.ClientIp)
		ctx = context.WithValue(ctx, "userEmail", session.User.Email)
		ctx = context.WithValue(ctx, "userFirstName", session.User.FirstName)
		ctx = context.WithValue(ctx, "userLastName", session.User.LastName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
