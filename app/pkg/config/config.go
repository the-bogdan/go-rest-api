package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// You need to create Config struct and fill it with envs that you need. There is example below.
// For more info reed https://github.com/ilyakaznacheev/cleanenv#model-format
//
//
//type DefaultConfig struct {
//	IsProd   bool   `env:"IS_PROD" env-required:"" env-description:"is prod env"`
//	LogLevel string `env:"LOG_LEVEL" env-required:"" env-description:"logging level"`
//
//	Listen struct {
//		Host string `env:"HOST" env-required:"" env-description:"server host"`
//		Port string `env:"PORT" env-required:"" env-description:"server port"`
//	}
//}

// GetConfig read envs and write them to the received config instance
func GetConfig(config interface{}) {
	// read envs only once
	if err := cleanenv.ReadEnv(config); err != nil {
		help, _ := cleanenv.GetDescription(config, nil)
		log.Fatalf("connot get config from envs, %s. %s", err, help)
	}
}
