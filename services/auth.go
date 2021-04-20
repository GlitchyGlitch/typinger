package services

import (
	"context"
	"fmt"

	"github.com/GlitchyGlitch/typinger/crypto"
	"github.com/GlitchyGlitch/typinger/errs"
	"github.com/GlitchyGlitch/typinger/models"

	"github.com/go-pg/pg"
)

type AuthRepo struct {
	DB              *pg.DB
	TokenController tokenController
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
	tokenStr, err := a.TokenController.Token(user.ID) // TODO: add
	if err != nil {
		fmt.Println(err.Error())
		return "", errs.Internal(ctx)
	}
	return tokenStr, err
}
