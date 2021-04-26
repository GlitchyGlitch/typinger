package config

import (
	"errors"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/GlitchyGlitch/typinger/crypto"
)

var ErrInvalidStatic = errors.New("invalid static path configuration")

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
	if port == "" {
		return "80"
	}
	return port
}

func EnvWriteTimeout() time.Duration {
	tStr := os.Getenv("WRITE_TIMEOUT")
	tInt, err := strconv.Atoi(tStr)
	if err != nil {
		tInt = 10
	}
	t := time.Duration(tInt) * time.Second
	return t
}

func EnvReadTimeout() time.Duration {
	tStr := os.Getenv("READ_TIMEOUT")
	tInt, err := strconv.Atoi(tStr)
	if err != nil {
		tInt = 5
	}
	t := time.Duration(tInt) * time.Second
	return t
}

func EnvIdleTimeout() time.Duration {
	tStr := os.Getenv("IDLE_TIMEOUT")
	tInt, err := strconv.Atoi(tStr)
	if err != nil {
		tInt = 120
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

func EnvImgEndpoint() string {
	path := os.Getenv("IMG_ENDPOINT")
	if path == "" {
		return "img"
	}
	return path
}

func EnvStaticPath() string {
	path := os.Getenv("STATIC_PATH")
	if path == "" {
		path = "./static"
	}
	path, err := filepath.Abs(path)
	if err != nil {
		panic(ErrInvalidStatic)
	}
	return path
}

func EnvStaticDashPath() string {
	path := os.Getenv("STATIC_DASH_PATH")
	if path == "" {
		path = "./static_dash"
	}
	path, err := filepath.Abs(path)
	if err != nil {
		panic(ErrInvalidStatic)
	}
	return path
}

func EnvProtocol() string {
	proto := os.Getenv("PROTOCOL")
	if proto == "" {
		return "http"
	}
	return proto
}

// TODO: Check where to panic
