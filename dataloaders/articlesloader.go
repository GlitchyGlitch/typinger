package dataloaders

import (
	"time"
)

func newArticlesByUserIDs(rep repos) *ArticlesLoader {
	return NewArticlesLoader(ArticlesLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch:    rep.GetArticlesByUserIDs,
	})
}
