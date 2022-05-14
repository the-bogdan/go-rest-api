package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsProd   bool   `env:"IS_PROD" env-default:"false"`
	LogLevel string `env:"LOG_LEVEL" env-default:"INFO"`

	Listen struct {
		Host string `env:"HOST" env-default:"127.0.0.1"`
		Port string `env:"PORT" env-default:"9090"`
	}
	Postgres struct{}
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig read envs and write them to Config instance
func GetConfig() *Config {
	// read envs only once
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatalf("connot get config from envs, %s. %s", err, help)
		}
	})
	return instance
}
