package services

import (
	"github.com/go-pg/pg"
)

type repos struct {
	UserRepo
	ArticleRepo
	SettingRepo
	AuthRepo
}

func NewRepos(DB *pg.DB) *repos {
	repos := &repos{
		UserRepo:    UserRepo{DB: DB},
		ArticleRepo: ArticleRepo{DB: DB},
		SettingRepo: SettingRepo{DB: DB},
		AuthRepo:    AuthRepo{DB: DB},
	}
	return repos
}
