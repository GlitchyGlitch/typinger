package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

func Server(rep repos) http.Handler {
	s := handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repos: rep,
		},
	}))
	s.SetErrorPresenter(ErrorPresenter()) // TODO: fove it ro server.go somehow
	return s
}
