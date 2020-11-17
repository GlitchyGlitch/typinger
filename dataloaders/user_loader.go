package dataloaders

import (
	"time"

	"github.com/GlitchyGlitch/typinger/postgres"
)

func newUserByIDs(repos postgres.Repos) *UserLoader {
	return NewUserLoader(UserLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch:    repos.GetUsersByIDs,
	})
}
