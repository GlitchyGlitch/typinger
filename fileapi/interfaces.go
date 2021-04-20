package fileapi

import (
	"context"

	"github.com/GlitchyGlitch/typinger/models"
)

type repos interface {
	GetImageBySlug(context.Context, string) (*models.Image, error)
}
