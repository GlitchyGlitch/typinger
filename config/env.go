package config

import (
	"net"
	"os"
	"strconv"
	"time"

	"github.com/GlitchyGlitch/typinger/crypto"
)

func EnvDBURL() string {
	url := os.Getenv("DATABASE_URL")
	return url
}

func EnvHost() string {
	host := os.Getenv("HOST")
	return host
}

func EnvPort() string {
	port := os.Getenv("PORT")
	return port
}

func EnvWriteTimeout() time.Duration {
	tStr := os.Getenv("WRITE_TIMEOUT")
	tInt, err := strconv.Atoi(tStr)
	if err != nil {
		tInt = 10 // Default value
	}
	t := time.Duration(tInt) * time.Second
	return t
}

func EnvReadTimeout() time.Duration {
	tStr := os.Getenv("READ_TIMEOUT")
	tInt, err := strconv.Atoi(tStr)
	if err != nil {
		tInt = 5 // Default value
	}
	t := time.Duration(tInt) * time.Second
	return t
}

func EnvIdleTimeout() time.Duration {
	tStr := os.Getenv("IDLE_TIMEOUT")
	tInt, err := strconv.Atoi(tStr)
	if err != nil {
		tInt = 120 // Default value
	}
	t := time.Duration(tInt) * time.Second
	return t
}

func EnvJWTSecret() []byte {
	jwtStr := os.Getenv("JWT_SECRET")
	jwt := []byte(jwtStr)
	if len(jwt) == 0 {
		return crypto.GenJWTSecret()
	}

	return jwt
}

func EnvDomain(host string) string {
	if host == "" {
		return "0.0.0.0"
	}
	domain := os.Getenv("DOMAIN")
	if domain != "" {
		return domain
	}
	domains, err := net.LookupAddr(host)
	if err != nil {
		return host
	}
	return domains[0]
}

func EnvImgDir() string {
	dir := os.Getenv("IMG_DIR")
	if dir == "" {
		return "img"
	}
	return dir
}

func EnvProtocol() string {
	dir := os.Getenv("PROTOCOL")
	if dir == "" {
		return "http"
	}
	return dir
}
