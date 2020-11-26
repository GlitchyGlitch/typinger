package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/GlitchyGlitch/typinger/dataloaders"
	"github.com/GlitchyGlitch/typinger/models"
)

func (r *articleResolver) Author(ctx context.Context, obj *models.Article) (*models.User, error) {
	return dataloaders.ForContext(ctx).UserByIDs.Load(obj.Author)
}

// Article returns ArticleResolver implementation.
func (r *Resolver) Article() ArticleResolver { return &articleResolver{r} }

type articleResolver struct{ *Resolver }
