package crypto

import (
	"crypto/rand"
)

const alphanum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return []byte{}, err
	}

	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}

	return bytes, nil
}
