package auth

import (
	"context"
	"net/http"

	"github.com/GlitchyGlitch/typinger/models"
)

type contextKey string

const key = contextKey("user")

func Middleware(tc tokenController, rep repos) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			claims, err := tc.ParseAuthorization(header)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			id, ok := claims["sub"].(string)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			user, err := rep.GetUserByID(r.Context(), id)
			if err != nil || user == nil {
				next.ServeHTTP(w, r)
				return
			}

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
