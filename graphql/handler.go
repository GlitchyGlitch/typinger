package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/GlitchyGlitch/typinger/dataloaders"
	"github.com/GlitchyGlitch/typinger/postgres"
)

func Handler(repos postgres.Repos, dl dataloaders.Retriever) http.Handler {
	return handler.GraphQL(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repos:       repos,
			DataLoaders: dl,
		},
	}))
}
