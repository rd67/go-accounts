package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APP_ENV string `mapstructure:"APP_ENV"`
	PORT    string    `mapstructure:"PORT"`

	DB_DRIVER   string `mapstructure:"DB_DRIVER"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string    `mapstructure:"DB_PORT"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`

	DB_URL string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return
	}

	log.Printf("The App is running in %s env", config.APP_ENV)

	config.DB_URL = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB_DRIVER,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	return config, err
}
