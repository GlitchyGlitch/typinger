package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/GlitchyGlitch/typinger/config"
	"github.com/GlitchyGlitch/typinger/validator"
)

type Resolver struct {
	Repos     repos
	Config    *config.Config
	Validator *validator.Validator
}
