package models

import "github.com/99designs/gqlgen/graphql"

type Image struct {
	ID      string `json:"uuid"`
	Name    string `json:"name"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

type NewImage struct {
	ID   int            `json:"id"`
	Name string         `json:"name"`
	URL  string         `json:"url"`
	File graphql.Upload `json:"file"`
}

type UpdateImage struct {
	Name *string         `json:"name"`
	URL  *string         `json:"url"`
	File *graphql.Upload `json:"file"`
}
