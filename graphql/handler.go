package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
)

func Server(rep repos, ep graphql.ErrorPresenterFunc) http.Handler {
	s := handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repos: rep,
		},
	}))
	s.SetErrorPresenter(ErrorPresenter()) // TODO: move it to server.go somehow
	return s
}
