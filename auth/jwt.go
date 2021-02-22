package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte("secret") // TODO: move to config struct
)

// Token generates a jwt token and assign a id and exp to it's claims and return it.
func Token(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // TODO: Add as config to config struct
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// parseToken parses a jwt token and returns the id in it's claims
func parseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["sub"].(string)
		return id, nil
	}
	return "", err
}
