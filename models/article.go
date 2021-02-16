package models

import "time"

type Article struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Author       string    `json:"author"`
	CreatedAt    time.Time `json:"created_at"`
}

type ArticleFilter struct {
	Title *string `json:"title"`
}

type NewArticle struct {
	Title        string `json:"title" validate:"required,max=255"`
	Content      string `json:"content" validate:"required"`
	ThumbnailURL string `json:"thumbnail_url" validate:"url,required,max=255"`
}

type UpdateArticle struct {
	Title        string `json:"title" validate:"required"`
	Content      string `json:"content" validate:"required"`
	ThumbnailURL string `json:"thumbnail_url" validate:"url,required"`
}
