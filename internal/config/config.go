package config

import (
	"fmt"
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	DB DBConfig
	JWT JWTConfig
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
}

type JWTConfig struct {
	Secret string
	AccessExpire time.Duration
	RefreshExpire time.Duration
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	cfg := &Config{
		ServerPort: viper.GetString("SERVER_PORT"),
		DB: DBConfig{
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetString("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName: viper.GetString("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("JWT_SECRET"),
			AccessExpire: viper.GetDuration("JWT_ACCESS_EXPIRE"),
			RefreshExpire: viper.GetDuration("JWT_REFRESH_EXPIRE"),
		},
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}

	return cfg, nil
}