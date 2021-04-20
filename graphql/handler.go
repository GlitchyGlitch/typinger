package graphql

import (
	"net/http"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/GlitchyGlitch/typinger/validator"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
)

func newSchema(rep repos, conf *config.Config, valid *validator.Validator, ep graphql.ErrorPresenterFunc) graphql.ExecutableSchema { //TODO: replace validator with interface maybe
	c := Config{
		Resolvers: &Resolver{
			Repos:     rep,
			Config:    conf,
			Validator: valid,
		},
	}
	return NewExecutableSchema(c)
}

func Handler(rep repos, conf *config.Config, validate *validator.Validator, ep graphql.ErrorPresenterFunc) http.Handler {
	s := handler.NewDefaultServer(newSchema(rep, conf, validate, ep))
	// s.SetErrorPresenter(ErrorPresenter()) // TODO: move it to server.go somehow
	return s
}
