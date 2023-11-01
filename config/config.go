package config

import (
	"fmt"
	"github.com/gookit/slog"
	"os"
	"redishop/util/helper"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	DBDriver string
}

type ApiConfig struct {
	Host string
	Port string
}

type RedisConfig struct {
	Host     string
	Password string
}

type Config struct {
	*DbConfig
	*ApiConfig
	*RedisConfig
}

// init db
func (c *Config) ReadDbConfig() error {
	//load env
	if err := helper.Loadenv(); err != nil {
		return err
	}
	//assign struct dbconfig
	c.DbConfig = &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		DBDriver: os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = &ApiConfig{
		Host: os.Getenv("API_HOST"),
		Port: os.Getenv("API_PORT"),
	}

	//if env is missing
	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.DBName == "" ||
		c.DbConfig.Username == "" || c.DbConfig.Password == "" || c.DbConfig.DBDriver == "" {
		return fmt.Errorf("missing required environment variable ")
	}

	slog.Infof("Database connected to %v", c.DbConfig.Host)
	return nil
}

func (c *Config) ReadRedisConfig() error {
	if err := helper.Loadenv(); err != nil {
		return err
	}

	c.RedisConfig = &RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cfg.ReadDbConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read db config %v", err.Error())
	}

	if err := cfg.ReadRedisConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read redis config %v", err.Error())
	}

	return cfg, nil
}
