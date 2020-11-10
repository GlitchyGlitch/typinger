package postgres

import (
	"fmt"

	"github.com/go-pg/pg"
)

type DBLogger struct {
}

func (d DBLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d DBLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
