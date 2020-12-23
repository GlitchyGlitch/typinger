package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ErrorPresenter() graphql.ErrorPresenterFunc {
	return graphql.ErrorPresenterFunc(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})
}
