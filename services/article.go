package services

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type ArticleRepo struct {
	DB *pg.DB
}

func (a *ArticleRepo) GetArticleByID(ctx context.Context, id string) (*models.Article, error) {
	article := &models.Article{}
	err := a.DB.Model(article).Where("id = ?", id).First()
	if err == pg.ErrNoRows {
		return nil, errs.NotFound(ctx)
	}
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	return article, nil
}

func (a *ArticleRepo) GetArticles(ctx context.Context, filter *models.ArticleFilter, first, offset int) ([]*models.Article, error) {
	var articles []*models.Article

	query := a.DB.Model(&articles).Order("id")

	if filter != nil && filter.Title != "" {
		query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", filter.Title))
	}

	if first != 0 {
		query.Limit(first)
	}
	if offset != 0 {
		query.Offset(offset)
	}

	err := query.Select()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if len(articles) == 0 {
		return nil, nil
	}

	return articles, nil
}

func (a *ArticleRepo) GetArticlesByUserIDs(ids []string) ([][]*models.Article, []error) { //TODO: move it to new loader to get context here
	var articles []*models.Article
	result := make([][]*models.Article, len(ids))
	aMap := make(map[string][]*models.Article, len(ids))

	err := a.DB.Model(&articles).Where("author in (?)", pg.In(ids)).Order("author").Select()

	if err != nil {
		return nil, []error{} // TODO: internal error here
	}
	if len(articles) == 0 {
		return result, []error{}
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
	res, err := a.DB.Model(article).OnConflict("DO NOTHING").Insert()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		return nil, errs.Exists(ctx)
	}
	return article, nil
}

func (a *ArticleRepo) UpdateArticle(ctx context.Context, id string, input models.UpdateArticle) (*models.Article, error) { //TODO: It can't replace empty fields
	article, err := a.GetArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Title != "" {
		article.Title = input.Title
	}
	if input.Content != "" {
		article.Content = input.Content
	}
	if input.ThumbnailURL != "" {
		article.ThumbnailURL = input.ThumbnailURL
	}

	res, err := a.DB.Model(article).Where("id = ?", id).Update()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		return nil, errs.NotFound(ctx)
	}
	return article, nil
}

func (a *ArticleRepo) DeleteArticle(ctx context.Context, id string) (bool, error) {
	article := &models.Article{ID: id}
	res, err := a.DB.Model(article).Where("id = ?", id).Delete()
	if err != nil {
		return false, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		errs.Add(ctx, errs.NotFound(ctx))
		return false, nil
	}
	return true, nil
}
