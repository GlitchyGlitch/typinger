package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.NewUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UpdateUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input models.NewArticle) (*models.Article, error) {
	return r.Repos.CreateArticle(&input)
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, id string, input models.UpdateArticle) (*models.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteArticle(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSetting(ctx context.Context, input *models.NewSetting) (*models.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSetting(ctx context.Context, id string, input models.UpdateSetting) (*models.Setting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSetting(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateImages(ctx context.Context, input []*models.NewImage) ([]*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateImage(ctx context.Context, id string, input models.UpdateImage) (*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteImages(ctx context.Context, ids []string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
