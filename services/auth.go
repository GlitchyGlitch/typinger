package services

import (
	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/crypto"
	"github.com/GlitchyGlitch/typinger/models"

	"github.com/go-pg/pg"
)

type AuthRepo struct {
	DB *pg.DB
}

func (a *AuthRepo) Authenticate(login models.LoginInput) (string, error) {
	user := &models.User{}
	err := a.DB.Model(user).Where("email = ?", login.Email).First()
	if err != nil {
		return "", err // TODO: handle this error properly
	}

	if ok := crypto.CheckPasswdHash(login.Password, user.PasswordHash); !ok {
		return "", err // TODO: handle this error properly
	}
	jwtStr, err := auth.Token(user.ID)
	if err != nil {
		return "", err // TODO: handle this error properly
	}
	return jwtStr, err
}
