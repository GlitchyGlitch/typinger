package crypto

func GenJWTSecret() []byte {
	secret, err := randomBytes(512)
	if err != nil {
		return []byte{} // TODO: Figure out what should happen here.
	}
	return secret
}
