package dataloaders

// go:generate go run github.com/vektah/dataloaden UserLoader string *github.com/GlitchyGlitch/typinger/models.User
// go:generate go run github.com/vektah/dataloaden ArticlesLoader string []*github.com/GlitchyGlitch/typinger/models.Article

type Loaders struct {
	UserByIDs         *UserLoader
	ArticlesByUserIDs *ArticlesLoader
}

func newLoaders(rep repos) *Loaders {
	return &Loaders{
		UserByIDs:         newUserByIDs(rep),
		ArticlesByUserIDs: newArticlesByUserIDs(rep),
	}
}
