package errs

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func Add(ctx context.Context, err error) {
	graphql.AddError(ctx, err)
}
