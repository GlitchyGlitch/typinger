package postgres

import (
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.Model(user).Where("id = ?", id).First()

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
