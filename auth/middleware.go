package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/GlitchyGlitch/typinger/models"
)

type repos interface {
	GetUserByID(context.Context, *string) (*models.User, error)
}

type contextKey string

const key = contextKey("user")

func Middleware(rep repos) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			if strings.ToLower(header[0:6]) != "bearer" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate jwt token and check if expired
			tokenStr := header[7:]
			id, err := parseToken(tokenStr)
			if err != nil {
				next.ServeHTTP(w, r) // TODO: Handle forbidden status properly here
				return
			}

			// Check if user exists
			user, err := rep.GetUserByID(r.Context(), &id)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// Put it in context
			ctx := context.WithValue(r.Context(), key, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func FromContext(ctx context.Context) *models.User {
	user, ok := ctx.Value(key).(*models.User)
	if !ok {
		return nil
	}
	return user
}
