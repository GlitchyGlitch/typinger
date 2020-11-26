package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/GlitchyGlitch/typinger/postgres"
)

func Handler(repos postgres.Repos) http.Handler {
	return handler.GraphQL(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repos: repos,
		},
	}))
}
