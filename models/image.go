package models

import "github.com/99designs/gqlgen/graphql"

// TODO: Validate it

type Image struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Img  []byte `json:"img"` // TODO: change field name to data
	MIME string `json:"mime"`
	Slug string `json:"slug"`
}

type NewImage struct {
	Name string         `json:"name" validate:"required,max=128"`
	Slug string         `json:"slug" validate:"required,excludesall,max=256"`
	File graphql.Upload `json:"file" validate:"required"`
}

type UpdateImage struct {
	Name string         `json:"name" validate:"om,max=128"`
	Slug string         `json:"slug"`
	File graphql.Upload `json:"file"`
}

type ImageFilter struct {
	Name *string `json:"name" validate:"omitempty,max=128"`
	Slug *string `json:"slug" validate:"omitempty,excludesall,max=256"`
}
