package dataloaders

import (
	"context"
	"net/http"

	"github.com/GlitchyGlitch/typinger/postgres"
)

type contextKey string

const key = contextKey("dataloaders")

//Middleware is a intermediary function that allows fix graphql's n+1 problem
func Middleware(repos postgres.Repos) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := newLoaders(repos)
			augmentedCtx := context.WithValue(ctx, key, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}

//ForContext is function that returns a bunch of dataloaders as pointer to Loaders struct
func ForContext(ctx context.Context) *Loaders {
	return ctx.Value(key).(*Loaders)
}
