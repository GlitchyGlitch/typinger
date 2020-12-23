package dataloaders

import (
	"time"
)

func newUserByIDs(rep repos) *UserLoader {
	return NewUserLoader(UserLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch:    rep.GetUsersByIDs,
	})
}
