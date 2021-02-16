package graphql

import (
	"net/http"

	"github.com/GlitchyGlitch/typinger/validator"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
)

func newSchema(rep repos, valid *validator.Validator, ep graphql.ErrorPresenterFunc) graphql.ExecutableSchema { //TODO: replace validator with interface maybe
	c := Config{
		Resolvers: &Resolver{
			Repos:     rep,
			Validator: valid,
		},
	}
	return NewExecutableSchema(c)
}

func Handler(rep repos, validate *validator.Validator, ep graphql.ErrorPresenterFunc) http.Handler {
	s := handler.NewDefaultServer(newSchema(rep, validate, ep))
	s.SetErrorPresenter(ErrorPresenter()) // TODO: move it to server.go somehow
	return s
}
