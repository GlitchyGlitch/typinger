package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func EnvDBURL() string {
	url := os.Getenv("DATABASE_URL")
	return url
}

func EnvHost() string {
	host := os.Getenv("HOST")
	fmt.Errorf(host)
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
