package graphql

import (
	"context"

	"github.com/GlitchyGlitch/typinger/models"
)

type repos interface {
	GetArticles(context.Context, *models.ArticleFilter, int, int) ([]*models.Article, error)
	CreateArticle(context.Context, *models.User, *models.NewArticle) (*models.Article, error) // Pointer to NewArticle because of big Content field.

	GetArticlesByUserIDs([]string) ([][]*models.Article, []error)
	GetUsers(context.Context) ([]*models.User, error)
	GetUserByID(context.Context, *string) (*models.User, error)
	GetUsersByIDs([]string) ([]*models.User, []error)
	CreateUser(context.Context, models.NewUser) (*models.User, error)

	Authenticate(models.LoginInput) (string, error)
}
