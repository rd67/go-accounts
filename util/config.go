package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT string `mapstructure:"APP_ENV"`
	HTTP_SERVER_ADDRESS    string `mapstructure:"HTTP_SERVER_ADDRESS"`

	DB_DRIVER   string `mapstructure:"DB_DRIVER"`
	DB_SOURCE   string `mapstructure:"DB_SOURCE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
