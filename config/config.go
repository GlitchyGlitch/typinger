package config

import (
	"fmt"
	"time"
)

type Config struct {
	Protocol     string
	Domain       string
	ImgDir       string
	DBURL        string
	Host         string
	Port         string
	JWTSecret    []byte
	IdleTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

func (c Config) Addr() string {
	if c.Port == "" {
		c.Port = "80" // Default value
	}
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func New() *Config {
	c := &Config{}
	c.DBURL = EnvDBURL()
	c.Host = EnvHost()
	c.Port = EnvPort()
	c.IdleTimeout = EnvIdleTimeout()
	c.WriteTimeout = EnvWriteTimeout()
	c.ReadTimeout = EnvReadTimeout()
	c.JWTSecret = EnvJWTSecret()
	c.Domain = EnvDomain(c.Host)
	c.ImgDir = EnvImgDir()
	c.Protocol = EnvProtocol()
	return c
}
