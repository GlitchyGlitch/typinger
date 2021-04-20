package services

import (
	"github.com/go-pg/pg"
)

type Repos struct {
	UserRepo
	ArticleRepo
	AuthRepo
	ImageRepo
}

func NewRepos(db *pg.DB, tc tokenController) *Repos {
	repos := &Repos{
		UserRepo:    UserRepo{DB: db},
		ArticleRepo: ArticleRepo{DB: db},
		ImageRepo:   ImageRepo{DB: db},
		AuthRepo:    AuthRepo{DB: db, TokenController: tc},
	}
	return repos
}
