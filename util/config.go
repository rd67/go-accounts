package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APP_ENV string `mapstructure:"APP_ENV"`
	PORT    string `mapstructure:"PORT"`

	DB_DRIVER   string `mapstructure:"DB_DRIVER"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`

	DB_URL string `mapstructure:"DB_URL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(path)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)

	if len(config.DB_URL) == 0 {
		//	Updating DB URL if not present in the .env
		config.DB_URL = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
			config.DB_DRIVER,
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_HOST,
			config.DB_PORT,
			config.DB_NAME,
		)
	}

	return

	// viper.AddConfigPath(path)
	// viper.SetConfigFile(".env")

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatal("Can't find the file .env : ", err)
	// 	return
	// }

	// err = viper.Unmarshal(&config)
	// if err != nil {
	// 	log.Fatal("Environment can't be loaded: ", err)
	// 	return
	// }

	// log.Printf("The App is running in %s env", config.APP_ENV)

	// if config.DB_URL == "" {
	// 	//	Updating DB URL if not present in the .env
	// 	config.DB_URL = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
	// 		config.DB_DRIVER,
	// 		config.DB_USER,
	// 		config.DB_PASSWORD,
	// 		config.DB_HOST,
	// 		config.DB_PORT,
	// 		config.DB_NAME,
	// 	)
	// }

	// return config, err
}
