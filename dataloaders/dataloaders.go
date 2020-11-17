package dataloaders

//go:generate go run github.com/vektah/dataloaden UserLoader string *github.com/GlitchyGlitch/typinger/models.User
//go:generate go run github.com/vektah/dataloaden ArticlesLoader string []*github.com/GlitchyGlitch/typinger/models.Article

import (
	"github.com/GlitchyGlitch/typinger/postgres" // update the username
)

type contextKey string

const key = contextKey("dataloaders")

type Loaders struct {
	UserByIDs         *UserLoader
	ArticlesByUserIDs *ArticlesLoader
}

func newLoaders(repos postgres.Repos) *Loaders {
	return &Loaders{
		UserByIDs:         newUserByIDs(repos),
		ArticlesByUserIDs: newArticlesByUserIDs(repos),
	}
}
