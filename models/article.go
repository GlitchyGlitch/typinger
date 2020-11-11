package models

type Article struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Author       string `json:"author"`
}

type ArticleFilter struct {
	Title *string `json:"title"`
}

type NewArticle struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type UpdateArticle struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
