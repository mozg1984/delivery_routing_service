package config

import (
	"github.com/spf13/viper"
)

type Server struct {
	Host string
	Port string
}

type Redis struct {
	Auth        string
	DB          int
	Host        string
	Port        string
	Type        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type Config struct {
	Server Server
	Redis  Redis
}

func NewConfig() (*Config, error) {
	config := Config{
		Server: Server{},
		Redis:  Redis{},
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./pkg/config/")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
