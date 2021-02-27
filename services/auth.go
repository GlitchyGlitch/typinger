package services

import (
	"context"

	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/crypto"
	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"

	"github.com/go-pg/pg"
)

type AuthRepo struct {
	DB *pg.DB
}

func (a *AuthRepo) Authenticate(ctx context.Context, login models.LoginInput) (string, error) {
	user := &models.User{}
	err := a.DB.Model(user).Where("email = ?", login.Email).First()
	if err == pg.ErrNoRows {
		errs.Add(ctx, errs.BadCredencials(ctx))
		return "", nil
	}
	if err != nil {
		return "", errs.Internal(ctx)
	}

	if ok := crypto.CheckPasswdHash(login.Password, user.PasswordHash); !ok {
		errs.Add(ctx, errs.BadCredencials(ctx))
		return "", nil
	}
	jwtStr, err := auth.Token(user.ID)
	if err != nil {
		return "", errs.Internal(ctx) // TODO: handle this error properly
	}
	return jwtStr, err
}
