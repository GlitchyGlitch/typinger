package services

import (
	"context"
	"fmt"
	"io"

	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type ImageRepo struct {
	DB *pg.DB
}

func (i *ImageRepo) GetImages(ctx context.Context, filter *models.ImageFilter, first, offset *int) ([]*models.Image, error) {
	var images []*models.Image

	query := i.DB.Model(&images).Order("created_at DESC")

	if filter != nil {
		if filter.Name != nil {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
		if filter.Slug != nil {
			query.Where("slug ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Slug))
		}
	}

	if first != nil {
		query.Limit(*first)
	}
	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if len(images) == 0 {
		return nil, nil
	}
	return images, nil
}

func (i *ImageRepo) GetImageByID(ctx context.Context, id string) (*models.Image, error) {
	image := &models.Image{}

	err := i.DB.Model(image).Where("id = ?", id).First()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if image == nil {
		return nil, nil
	}

	return image, nil
}

func (i *ImageRepo) GetImageBySlug(ctx context.Context, slug string) (*models.Image, error) {
	image := &models.Image{}

	err := i.DB.Model(image).Where("slug = ?", slug).First()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if image == nil {
		return nil, nil
	}

	return image, nil
}

func (i *ImageRepo) CreateImage(ctx context.Context, input models.NewImage) (*models.Image, error) {
	fileBuf := make([]byte, 5120)
	for {
		_, err := input.File.File.Read(fileBuf)
		if err == io.EOF {
			break
		}
	}
	mime := input.File.ContentType
	//TODO: check content type
	image := &models.Image{Name: input.Name, Img: fileBuf, MIME: mime, Slug: input.Slug}
	res, err := i.DB.Model(image).OnConflict("DO NOTHING").Insert()

	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		return nil, errs.Exists(ctx)
	}
	return image, nil
}

func (i *ImageRepo) UpdateImage(ctx context.Context, id string, input models.UpdateImage) (*models.Image, error) {
	image, err := i.GetImageByID(ctx, id)
	if err != nil {
		return nil, err
	}

	fileBuf := make([]byte, 5120)
	for {
		_, err := input.File.File.Read(fileBuf)
		if err == io.EOF {
			break
		}
	}

	if len(fileBuf) != 0 {
		image.Img = fileBuf
		image.MIME = input.File.ContentType
	}
	if input.Name != "" {
		image.Name = input.Name
	}
	if input.Slug != "" {
		image.Slug = input.Slug
	}

	res, err := i.DB.Model(input).Where("id = ?", id).Update()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		return nil, errs.NotFound(ctx)
	}
	return image, nil
}

func (i *ImageRepo) DeleteImage(ctx context.Context, id string) (bool, error) {
	image := &models.Image{ID: id}
	res, err := i.DB.Model(image).Where("id = ?", id).Delete()
	if err != nil {
		return false, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		errs.Add(ctx, errs.NotFound(ctx))
		return false, nil
	}

	return true, nil
}
