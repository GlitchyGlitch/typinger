package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/models"
)

func (r *imageResolver) URL(ctx context.Context, obj *models.Image) (*string, error) {
	url := fmt.Sprintf("%s://%s/%s/%s", r.Config.Protocol, r.Config.Domain, r.Config.ImgEndpoint, obj.Slug)
	return &url, nil
}

// Image returns ImageResolver implementation.
func (r *Resolver) Image() ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
