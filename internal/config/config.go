package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gookit/validate"
)

func Default() *Config {
	return &Config{
		Port:        4000,
		Host:        "localhost",
		Environment: "development",
		Name:        "go-url-shortener",
	}
}

type Config struct {
	Port        int       `json:"port" env:"APP_PORT" validate:"required|numeric"`
	Host        string    `json:"host" env:"APP_HOST" validate:"required|string"`
	Environment string    `json:"environment" env:"APP_ENV" validate:"required|string"`
	Name        string    `json:"name" env:"APP_NAME" validate:"required|string"`
	Database    *Database `json:"database" validate:"required"`
}

type Database struct {
	Driver string `json:"driver" env:"DATABASE_DRIVER" validate:"required|string"`
	DSN    string `json:"dsn" env:"DATABASE_DSN" validate:"required|string"`
}

func (c *Config) Validate() error {
	v := validate.Struct(c)
	if v.Validate() {
		return nil
	}
	return v.Errors
}

func Load(file string) (*Config, error) {
	c := Default()

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, c); err != nil {
		return nil, err
	}

	// TODO: Read env variables

	if err = c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}
