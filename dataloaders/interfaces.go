package dataloaders

import "github.com/GlitchyGlitch/typinger/models"

type repos interface {
	GetArticlesByUserIDs([]string) ([][]*models.Article, []error)
	GetUsersByIDs([]string) ([]*models.User, []error)
}
