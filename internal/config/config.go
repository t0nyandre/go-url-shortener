package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gookit/validate"
)

const (
	defaultAppPort = 4000
	defaultAppHost = "localhost"
	defaultAppEnv  = "development"
	defaultAppName = "go-url-shortener"
)

type Config struct {
	AppPort int    `json:"app_port" env:"APP_PORT"`
	AppHost string `json:"app_host" env:"APP_HOST"`
	AppEnv  string `json:"app_env" env:"APP_ENV"`
	AppName string `json:"app_name" env:"APP_NAME"`
}

func (c *Config) Validate() error {
	v := validate.Struct(c)
	if v.Validate() {
		return nil
	}
	return v.Errors
}

func Load(file string) (*Config, error) {
	c := &Config{
		AppPort: defaultAppPort,
		AppHost: defaultAppHost,
		AppEnv:  defaultAppEnv,
		AppName: defaultAppName,
	}

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
