package services

import (
	"github.com/go-pg/pg"
)

type Repos struct {
	UserRepo
	ArticleRepo
	SettingRepo
	AuthRepo
}

func NewRepos(DB *pg.DB) *Repos {
	repos := &Repos{
		UserRepo:    UserRepo{DB: DB},
		ArticleRepo: ArticleRepo{DB: DB},
		SettingRepo: SettingRepo{DB: DB},
		AuthRepo:    AuthRepo{DB: DB},
	}
	return repos
}
