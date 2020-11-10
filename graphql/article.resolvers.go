package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/models"
)

func (r *articleResolver) Author(ctx context.Context, obj *models.Article) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article returns ArticleResolver implementation.
func (r *Resolver) Article() ArticleResolver { return &articleResolver{r} }

type articleResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *articleResolver) ID(ctx context.Context, obj *models.Article) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
