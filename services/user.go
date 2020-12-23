package services

import (
	"context"

	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.Model(user).Where("id = ?", id).First()
	if err != nil {
		return nil, errs.ErrEmpty(ctx)
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
