package config

import (
	"fmt"
	"os"
	"strconv"
)

const envPortKey = "SNITCH_PORT"
const defaultPort = 1323

type Config struct {
	Port int
}

func (c *Config) GetPort() int {
	return c.Port
}

func New() (*Config, error) {
	port := defaultPort

	if portEnv, found := os.LookupEnv(envPortKey); found {
		var err error
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			return nil, fmt.Errorf("invalid port value from environment %s: %v", envPortKey, portEnv)
		}
	}

	return &Config{
		Port: port,
	}, nil
}
