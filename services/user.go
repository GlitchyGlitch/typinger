package services

import (
	"context"

	"github.com/GlitchyGlitch/typinger/crypto"
	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Model(&users).Order("id").Select()
	if err != nil {
		return []*models.User{}, errs.Internal(ctx)
	}
	return users, nil
}

func (u *UserRepo) GetUserByID(ctx context.Context, id *string) (*models.User, error) {
	user := &models.User{}

	if id == nil {
		return nil, errs.InvalidInput(ctx) // TODO: move it to data validation module
	}

	err := u.DB.Model(user).Where("id = ?", id).First()
	if err != nil {
		return nil, errs.NotFound(ctx)
	}

	return user, nil
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.Model(user).Where("email = ?", email).First()

	return user, err
}

func (u *UserRepo) GetUsersByIDs(ids []string) ([]*models.User, []error) {
	var users []*models.User

	err := u.DB.Model(&users).Where("id in (?)", pg.In(ids)).Select()
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
	_, err := u.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errs.Exists(ctx)
	}

	hash, err := crypto.HashPasswd(input.Password)
	if err != nil {
		return nil, errs.InvalidInput(ctx)
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
	if _, err := tx.Model(user).Returning("*").Insert(); err != nil {
		return nil, errs.Internal(ctx)
	}
	if err := tx.Commit(); err != nil {
		return nil, errs.Internal(ctx)
	}

	return user, err
}
