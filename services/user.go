package services

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/crypto"
	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUsers(ctx context.Context, filter *models.UserFilter, first, offset *int) ([]*models.User, error) {
	var users []*models.User

	query := u.DB.Model(&users).Order("created_at DESC")
	if filter != nil {
		if filter.Name != nil {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
		if filter.Email != nil {
			query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Email))
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
	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func (u *UserRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}

	err := u.DB.Model(user).Where("id = ?", id).First()
	if err == pg.ErrNoRows {
		return nil, errs.NotFound(ctx)
	}
	if err != nil {
		return nil, errs.Internal(ctx)
	}

	return user, nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.Model(user).Where("email = ?", email).First()
	if user == nil {
		return nil, errs.NotFound(ctx)
	}
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	return user, err
}

func (u *UserRepo) GetUsersByIDs(ids []string) ([]*models.User, []error) {
	var users []*models.User

	err := u.DB.Model(&users).Where("id in (?)", pg.In(ids)).Order("created_at DESC").Select() // Check if order works
	if err != nil {
		return nil, []error{err}
	}

	uMap := make(map[string]*models.User, len(users))

	for _, user := range users {
		uMap[user.ID] = user
	}

	result := make([]*models.User, len(ids))
	for i, id := range ids {
		result[i] = uMap[id]
	}
	return result, nil
}

func (u *UserRepo) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	hash, err := crypto.HashPasswd(input.Password)
	if err != nil {
		return nil, errs.Internal(ctx)
	}

	user := &models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hash,
	}

	tx, err := u.DB.Begin()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	defer tx.Rollback()

	res, err := tx.Model(user).OnConflict("DO NOTHING").Insert()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		return nil, errs.Exists(ctx)
	}

	if err := tx.Commit(); err != nil {
		return nil, errs.Internal(ctx)
	}

	return user, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, id string, input models.UpdateUser) (*models.User, error) {
	user, err := u.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		passwd, err := crypto.HashPasswd(input.Password)
		if err != nil {
			return nil, errs.Internal(ctx)
		}
		user.PasswordHash = passwd
	}

	tx, err := u.DB.Begin()
	if err != nil {
		return nil, errs.Internal(ctx)
	}
	defer tx.Rollback()

	res, err := tx.Model(user).Where("id = ?", id).Update()
	if err != nil {
		return nil, errs.Internal(ctx)
	}

	if res.RowsAffected() <= 0 {
		return nil, errs.NotFound(ctx)
	}

	if err := tx.Commit(); err != nil {
		return nil, errs.Internal(ctx)
	}
	return user, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, id string) (bool, error) {
	user := &models.User{ID: id}
	tx, err := u.DB.Begin()
	if err != nil {
		return false, errs.Internal(ctx)
	}
	defer tx.Rollback()

	res, err := tx.Model(user).Where("id = ?", id).Delete()
	if err != nil {
		return false, errs.Internal(ctx)
	}
	if res.RowsAffected() <= 0 {
		return false, errs.NotFound(ctx)
	}

	if err := tx.Commit(); err != nil {
		return false, errs.Internal(ctx)
	}
	return true, nil
}
