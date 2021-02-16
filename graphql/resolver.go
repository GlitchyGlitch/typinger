package graphql

//go:generate go run github.com/99designs/gqlgen

import "github.com/GlitchyGlitch/typinger/validator"

type Resolver struct {
	Repos     repos
	Validator *validator.Validator
}
