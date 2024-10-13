package config

import (
	"log"
	"os"
	"path/filepath"
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
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting working directory: %s", err)
		}
		configPath := filepath.Join(wd, "..", "..", "internal/config/config.yml")
		err = cleanenv.ReadConfig(configPath, instance)
		if err != nil {
			log.Fatalf("Config error: %s", err)
		}
	})
	return instance
}
