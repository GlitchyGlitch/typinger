package postgres

import (
	"log"

	"github.com/go-pg/pg"
)

type DBLogger struct {
}

func (d DBLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d DBLogger) AfterQuery(q *pg.QueryEvent) {
	log.Println(q.FormattedQuery())
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
