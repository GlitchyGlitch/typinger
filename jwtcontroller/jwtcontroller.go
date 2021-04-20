package jwtcontroller

import (
	"fmt"
	"strings"
	"time"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/dgrijalva/jwt-go"
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
		return nil, fmt.Errorf("invalid authorization header")
	}
	rawToken := header[len(prefix):]

	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		return c.Config.JWTSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, err
}
