package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/GlitchyGlitch/typinger/models"
)

func (r *userResolver) Articles(ctx context.Context, obj *models.User) ([]*models.Article, error) {
	return r.DataLoaders.Retrieve(ctx).ArticlesByUserIDs.Load(obj.ID)
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
