package postgres

import (
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type Repos interface {
	GetArticles(*models.ArticleFilter, int, int) ([]*models.Article, error)
	CreateArticle(*models.NewArticle) (*models.Article, error) // Pointer to NewArticle because of big Content field.
	GetArticlesByUserIDs([]string) ([][]*models.Article, []error)

	GetUserByID(string) (*models.User, error)
	GetUsersByIDs([]string) ([]*models.User, []error)

	GetSettings() ([]*models.Setting, error)
}

type repos struct {
	UserRepo
	ArticleRepo
	SettingRepo
}

func NewRepos(DB *pg.DB) Repos {
	repos := &repos{
		UserRepo:    UserRepo{DB: DB},
		ArticleRepo: ArticleRepo{DB: DB},
		SettingRepo: SettingRepo{DB: DB},
	}
	return repos
}
