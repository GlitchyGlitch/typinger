package auth

import (
	"context"

	"github.com/GlitchyGlitch/typinger/models"
)

type repos interface {
	GetUserByID(context.Context, string) (*models.User, error)
}
type tokenController interface { //TODO: implement it
	ParseAuthorization(string) (map[string]interface{}, error)
	Token(string) (string, error)
}
