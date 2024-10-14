package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Name   string `mapstructure:"name"`
	Env    string `mapstructure:"env"`
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

type GRPCConfig struct {
	Address string `mapstructure:"address"`
}

type TMDBConfig struct {
	APIKey string `mapstructure:"api_key"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type Config struct {
	App  AppConfig  `mapstructure:"app"`
	GRPC GRPCConfig `mapstructure:"grpc"`
	TMDB TMDBConfig `mapstructure:"tmdb"`
	Log  LogConfig  `mapstructure:"log"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошика распаковки конфигурации: %w", err)
	}

	if cfg.TMDB.APIKey == "" {
		log.Println("Ключ Api не установлен")
	}

	return &cfg, nil
}
