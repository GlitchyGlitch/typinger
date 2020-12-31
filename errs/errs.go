package errs

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Forbidden(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "Operation forbidden.",
		Extensions: map[string]interface{}{
			"code": "FORBIDDEN",
		},
	}
}

func NotFound(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "No data found.", // TODO: check ortography
		Extensions: map[string]interface{}{
			"code": "NOT_FOUND",
		},
	}
}

func BadInput(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "Input is invalid", // TODO: check ortography
		Extensions: map[string]interface{}{
			"code": "BAD_INPUT",
		},
	}
}

func Exists(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "Resource already exists.", // TODO: check ortography
		Extensions: map[string]interface{}{
			"code": "EXISTS",
		},
	}
}

func Internal(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "Internal server error.", // TODO: check ortography
		Extensions: map[string]interface{}{
			"code": "INTERNAL_ERROR",
		},
	}
}
