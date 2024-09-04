package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port     string   `mapstructure:"PORT"`
	Backends []string `mapstructure:"BACKENDS"`
}

func LoadConfig(env string) (*Config, error) {
	viper.SetConfigFile(env)

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
