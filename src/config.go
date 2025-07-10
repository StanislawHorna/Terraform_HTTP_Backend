package src

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v10"

	"terraform_http_backend/src/store"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	LogLevel  string          `env:"LOG_LEVEL" envDefault:"error"`
	StoreType store.StoreType `env:"STORE_TYPE" envDefault:"file"`

	FileStore struct {
		Path          string `env:"FILE_STORE_PATH" envDefault:"./states"`
		FileExtension string `env:"FILE_STORE_EXTENSION" envDefault:".json"`
	}

	Loki struct {
		LokiURL     string `env:"LOKI_URL"`
		Environment string `env:"LOKI_ENV" envDefault:"dev"`
		GoVersion   string
		AppName     string
		AppVersion  string `env:"LOKI_APP_VERSION" envDefault:"0.0.0"`
	}
}

func GetConfig() *Config {
	once.Do(func() {
		var err error
		instance, err = loadConfig()
		if err != nil {
			panic("Failed to load configuration: " + err.Error())
		}
	})
	return instance
}

func loadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	cfg.Loki.AppName = AppName
	cfg.Loki.GoVersion = goVersion

	if !cfg.StoreType.Validate() {
		return nil, fmt.Errorf("invalid STORE_TYPE set")
	}
	return cfg, nil
}
