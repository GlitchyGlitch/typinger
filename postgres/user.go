package postgres

import (
	"github.com/GlitchyGlitch/typinger/models"
	"github.com/go-pg/pg"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetByUUID(uuid string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.Model(user).Where("user_uuid = ?", uuid).First()

	return user, err
}
