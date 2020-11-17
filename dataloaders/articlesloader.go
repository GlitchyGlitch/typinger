package dataloaders

import (
	"time"

	"github.com/GlitchyGlitch/typinger/postgres"
)

func newArticlesByUserIDs(repos postgres.Repos) *ArticlesLoader {
	return NewArticlesLoader(ArticlesLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch:    repos.GetArticlesByUserIDs,
	})
}
