package services

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type ArticleRepo struct {
	DB         *pg.DB
	errHandler ErrHandler
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
		return nil, a.errHandler.Error("internal")
	}

	return articles, nil
}

func (a *ArticleRepo) GetArticlesByUserIDs(ids []string) ([][]*models.Article, []error) {
	var articles []*models.Article
	result := make([][]*models.Article, len(ids))
	aMap := make(map[string][]*models.Article, len(ids))

	err := a.DB.Model(&articles).Where("author in (?)", pg.In(ids)).Order("author").Select()
	if err != nil {
		return nil, []error{} // TODO: error here
	}

	for _, article := range articles {
		aMap[article.Author] = append(aMap[article.Author], article)
	}

	for i, id := range ids {
		result[i] = aMap[id]
	}
	return result, nil
}

func (a *ArticleRepo) CreateArticle(ctx context.Context, user *models.User, input *models.NewArticle) (*models.Article, error) {
	article := &models.Article{Title: input.Title, Content: input.Content, ThumbnailURL: input.ThumbnailURL, Author: user.ID}
	_, err := a.DB.Model(article).Returning("*").Insert()
	if err != nil {
		return nil, errs.Exists(ctx)
	}
	return article, nil
}
