package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/models"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*models.User, error) {
	return r.UserRepo.GetByID(id)
}

func (r *queryResolver) Articles(ctx context.Context, filter *models.ArticleFilter, limit *int, offset *int) ([]*models.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Settings(ctx context.Context, id *string) ([]*models.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Images(ctx context.Context, id *string) ([]*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
