package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HOST            string `mapstructure:"host"`
	PORT            string `mapstructure:"port"`
	DsnUrl          string `mapstructure:"dsn_url"`
	LoggerLevel     string `mapstructure:"logger_level"`
	LoggerPath      string `mapstructure:"logger_path"`
	DsnOptions      string `mapstructure:"dsn_options"`
	MigrationsUrl   string `mapstructure:"migrations_url"`
	FileStoragePath string `mapstructure:"file_storage_path"`
}

func LoadConfig(dir string, configName string, configType string) (config Config, err error) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(dir)
	viper.SetConfigType(configType)

	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "4000")

	viper.AutomaticEnv()

	// Find and read the config file
	err = viper.ReadInConfig()

	// Handle errors
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	// Handle errors
	if err != nil {
		return Config{}, err
	}

	return
}
