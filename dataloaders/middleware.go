package dataloaders

import (
	"context"
	"net/http"
)

type contextKey string

const key = contextKey("dataloaders")

//Middleware is a intermediary function that allows fix graphql's n+1 problem
func Middleware(rep repos) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := newLoaders(rep)
			augmentedCtx := context.WithValue(ctx, key, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}

// FromContext is function that returns a bunch of dataloaders as pointer to Loaders struct
func FromContext(ctx context.Context) *Loaders {
	loaders, ok := ctx.Value(key).(*Loaders)
	if !ok {
		return nil
	}
	return loaders
}
