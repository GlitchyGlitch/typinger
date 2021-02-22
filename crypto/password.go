package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPasswd(passwd string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), 12) //TODO: Add to config.
	return string(bytes), err
}

func CheckPasswdHash(passwd, hash string) (ok bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	return err == nil
}
