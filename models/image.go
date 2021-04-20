package models

import "github.com/99designs/gqlgen/graphql"

type Image struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Img  []byte `json:"img"` // TODO: change field name to data
	MIME string `json:"mime"`
	Slug string `json:"slug"`
}

type NewImage struct {
	ID   int            `json:"id"`
	Name string         `json:"name"`
	Slug string         `json:"slug"`
	File graphql.Upload `json:"file"`
}

type UpdateImage struct {
	Name string         `json:"name"`
	Slug string         `json:"slug"`
	File graphql.Upload `json:"file"`
}
