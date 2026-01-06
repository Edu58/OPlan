package middleware

import (
	"context"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// Iterate backwards:
// 1. Put shirt on you: Shirt(You)
// 2. Put sweater over shirt: Sweater(Shirt(You))
// 3. Put jacket over sweater: Jacket(Sweater(Shirt(You)))

// Result: Jacket is outermost (executes first)
//
//	Shirt is innermost (executes last, closest to you)
func Chain(middlewares ...Middleware) Middleware {
	// Does not execute the func. Just prepares it
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}

		return final
	}
}

func WithValue(key, value any) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add value to context
			ctx := context.WithValue(r.Context(), key, value)

			// Call next handler with modified request
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
