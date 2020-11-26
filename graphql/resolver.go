package graphql

//go:generate go run github.com/99designs/

import (
	"github.com/GlitchyGlitch/typinger/postgres"
)

type Resolver struct {
	Repos postgres.Repos
}
