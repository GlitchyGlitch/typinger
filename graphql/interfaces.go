package graphql

import (
	"context"

	"github.com/GlitchyGlitch/typinger/models"
)

type repos interface {
	GetArticles(context.Context, *models.ArticleFilter, int, int) ([]*models.Article, error)
	CreateArticle(context.Context, *models.User, *models.NewArticle) (*models.Article, error)
	UpdateArticle(context.Context, string, models.UpdateArticle) (*models.Article, error)
	DeleteArticle(context.Context, string) (bool, error)
	GetArticlesByUserIDs([]string) ([][]*models.Article, []error)

	GetUsers(context.Context, *models.UserFilter, int, int) ([]*models.User, error)
	GetUsersByIDs([]string) ([]*models.User, []error)
	CreateUser(context.Context, models.NewUser) (*models.User, error)
	UpdateUser(context.Context, string, models.UpdateUser) (*models.User, error)
	DeleteUser(context.Context, string) (bool, error)

	Authenticate(context.Context, models.LoginInput) (string, error)
}
