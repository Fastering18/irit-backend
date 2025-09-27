// File: pkg/config/config.go

package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`
	Server struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"server"`
	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}