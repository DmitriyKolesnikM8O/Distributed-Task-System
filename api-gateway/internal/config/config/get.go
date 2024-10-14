package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	instance *Config
	one      sync.Once
)

func GetConfig() *Config {
	one.Do(func() {
		log.Printf("Read config")
		instance = &Config{}
		err := cleanenv.ReadConfig("internal/config/config.yml", instance)
		if err != nil {
			log.Fatalf("Config error: %s", err)
		}
	})
	return instance
}
