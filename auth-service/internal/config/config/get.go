package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	instance *Config
	// one      sync.Once
)

//Old function for read config
// func GetConfig() *Config {
// 	one.Do(func() {
// 		log.Printf("Read config")
// 		instance = &Config{}
// 		err := cleanenv.ReadConfig("internal/config/config.yml", instance)
// 		if err != nil {
// 			log.Fatalf("Config error: %s", err)
// 		}
// 	})
// 	return instance
// }

func GetConfig() *Config {
	log.Printf("Read config")
	v := viper.New()
	v.AddConfigPath("internal/config/")
	v.SetConfigName("config")
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found: %s", err)
		} else {
			log.Fatalf("Config error: %s", err)
		}
	}

	v.Unmarshal(&instance)
	return instance
}
