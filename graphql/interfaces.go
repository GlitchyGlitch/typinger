package graphql

import (
	"context"

	"github.com/GlitchyGlitch/typinger/models"
)

type repos interface {
	GetArticles(*models.ArticleFilter, int, int) ([]*models.Article, error)
	CreateArticle(*models.NewArticle) (*models.Article, error) // Pointer to NewArticle because of big Content field.
	GetArticlesByUserIDs([]string) ([][]*models.Article, []error)

	GetUserByID(context.Context, string) (*models.User, error)
	GetUsersByIDs([]string) ([]*models.User, []error)

	GetSettings() ([]*models.Setting, error)

	Authenticate(models.LoginInput) (string, error)
}
