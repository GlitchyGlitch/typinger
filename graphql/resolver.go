package graphql

import "github.com/GlitchyGlitch/typinger/postgres"

type Resolver struct {
	UserRepo postgres.UserRepo
}
