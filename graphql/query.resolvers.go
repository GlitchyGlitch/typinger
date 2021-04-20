package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
)

func (r *queryResolver) Users(ctx context.Context, filter *models.UserFilter, first *int, offset *int) ([]*models.User, error) {
	if !auth.Authorize(auth.FromContext(ctx)) {
		return nil, errs.Forbidden(ctx)
	}
	if ok := r.Validator.CheckStruct(ctx, filter, true); !ok {
		return nil, nil
	}
	if ok := r.Validator.CheckPagination(ctx, first, offset); !ok {
		return nil, nil
	}
	return r.Repos.GetUsers(ctx, filter, first, offset)
}

func (r *queryResolver) Articles(ctx context.Context, filter *models.ArticleFilter, first *int, offset *int) ([]*models.Article, error) {
	if ok := r.Validator.CheckStruct(ctx, filter, true); !ok {
		return nil, nil
	}
	if ok := r.Validator.CheckPagination(ctx, first, offset); !ok {
		return nil, nil
	}
	return r.Repos.GetArticles(ctx, filter, first, offset)
}

func (r *queryResolver) Images(ctx context.Context, filter *models.ImageFilter, first *int, offset *int) ([]*models.Image, error) {
	if !auth.Authorize(auth.FromContext(ctx)) {
		return nil, errs.Forbidden(ctx)
	}
	if ok := r.Validator.CheckStruct(ctx, filter, true); !ok {
		return nil, nil
	}
	if ok := r.Validator.CheckPagination(ctx, first, offset); !ok {
		return nil, nil
	}

	return r.Repos.GetImages(ctx, filter, first, offset)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
