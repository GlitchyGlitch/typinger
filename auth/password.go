package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPasswd(passwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), 12) //TODO: Add to config.
	return string(bytes), err
}

func CheckPasswdHash(passwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	return err == nil
} // TODO: create separate modeule for this.
