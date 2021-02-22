package config

import (
	"fmt"
	"time"
)

type Config struct {
	DBURL        string
	Host         string
	Port         string
	JwtSecret    string
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
	return c
}
