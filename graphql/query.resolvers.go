package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*models.User, error) {
	if !auth.Authorize(auth.FromContext(ctx)) {
		return nil, errs.Forbidden(ctx)
	}

	return r.Repos.GetUserByID(ctx, id)
}

func (r *queryResolver) Articles(ctx context.Context, filter *models.ArticleFilter, first *int, offset *int) ([]*models.Article, error) {
	return r.Repos.GetArticles(ctx, filter, *first, *offset)
}

func (r *queryResolver) Settings(ctx context.Context, id *string) ([]*models.Setting, error) {
	return r.Repos.GetSettings()
}

func (r *queryResolver) Images(ctx context.Context, id *string) ([]*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
