package graphql

import (
	"context"

	"github.com/GlitchyGlitch/typinger/models"
)

type repos interface {
	GetArticles(context.Context, *models.ArticleFilter, int, int) ([]*models.Article, error)
	CreateArticle(context.Context, *models.User, *models.NewArticle) (*models.Article, error) // Pointer to NewArticle because of big Content field.
	DeleteArticle(context.Context, string) (bool, error)
	GetArticlesByUserIDs([]string) ([][]*models.Article, []error)

	GetUserByID(context.Context, *string) (*models.User, error)
	GetUsersByIDs([]string) ([]*models.User, []error)
	CreateUser(context.Context, models.NewUser) (*models.User, error)

	Authenticate(context.Context, models.LoginInput) (string, error)
}
