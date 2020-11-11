package postgres

import (
	"fmt"

	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type ArticleRepo struct {
	DB *pg.DB
}

func (a *ArticleRepo) GetArticles(filter *models.ArticleFilter, limit, offset int) ([]*models.Article, error) {
	var articles []*models.Article

	query := a.DB.Model(&articles).Order("id")

	if filter != nil {
		if filter.Title != nil {
			query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Title))
		}
	}
	if limit != 0 {
		query.Limit(limit)
	}
	if offset != 0 {
		query.Offset(offset)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *ArticleRepo) CreateArticle(input models.NewArticle) (*models.Article, error) {
	article := &models.Article{Title: input.Title, Content: input.Content, ThumbnailURL: input.ThumbnailURL, Author: "173057db-f127-4185-99df-dfa33787432d"} // TODO: Replace with author ID form auth module later
	_, err := a.DB.Model(article).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return article, nil
}
