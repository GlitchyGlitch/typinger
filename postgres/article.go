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

func (a *ArticleRepo) CreateArticle(article *models.Article) (*models.Article, error) {
	_, err := a.DB.Model(article).Returning("*").Insert()
	return article, err
}
