package errs

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ErrForbidden(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "Operation forbidden.",
		Extensions: map[string]interface{}{
			"code": "FORBIDDEN",
		},
	}
}

func ErrEmpty(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "No data found in this path.", // TODO: check ortography
		Extensions: map[string]interface{}{
			"code": "EMPTY",
		},
	}
}
