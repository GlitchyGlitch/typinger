package postgres

import (
	"github.com/go-pg/pg"
)

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
