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

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	if !auth.Authorize(auth.FromContext(ctx)) {
		return nil, errs.Forbidden(ctx)
	}

	return r.Repos.CreateUser(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UpdateUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input models.NewArticle) (*models.Article, error) {
	user := auth.FromContext(ctx)
	if !auth.Authorize(user) {
		return nil, errs.Forbidden(ctx)
	}

	if ok := r.Validator.ValidateErrs(ctx, input); !ok {
		return nil, nil
	}

	return r.Repos.CreateArticle(ctx, user, &input)
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, id string, input models.UpdateArticle) (*models.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteArticle(ctx context.Context, id string) (bool, error) {
	user := auth.FromContext(ctx)
	if !auth.Authorize(user) {
		return false, errs.Forbidden(ctx)
	}

	if ok := r.Validator.CheckUUID(ctx, id); !ok {
		return false, nil
	}
	// TODO: implement delete
	return r.Repos.DeleteArticle(ctx, id)
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

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (string, error) {
	return r.Repos.Authenticate(input)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) RefreshToken(ctx context.Context, input models.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
