package jwtcontroller

import (
	"errors"
	"strings"
	"time"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidHeader = errors.New("invalid authorization header")
	ErrInvalidToken  = errors.New("invalid token")
)

const prefix = "Bearer "

type JWTController struct {
	Config *config.Config
}

func New(config *config.Config) *JWTController {
	c := &JWTController{Config: config}
	return c
}

func (c JWTController) Token(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // TODO: Add as config to config struct
	rawToken, err := token.SignedString(c.Config.JWTSecret)
	if err != nil {
		return "", err
	}
	return rawToken, nil
}

func (c JWTController) ParseAuthorization(header string) (map[string]interface{}, error) {
	if !strings.HasPrefix(header, prefix) {
		return nil, ErrInvalidHeader
	}
	rawToken := header[len(prefix):]

	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		return c.Config.JWTSecret, nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, err
}
