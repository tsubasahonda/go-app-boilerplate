package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config def
type Config struct {
	DB     DBConfig
	Server Server
}

// DBConfig def
type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

// Server def
type Server struct {
	Port int
}

// ReadConfig get config from env
func ReadConfig() (*Config, error) {
	sPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'SERVER_PORT': %w", err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'DB_PORT': %w", err)
	}

	return &Config{
		DBConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
		},
		Server{
			Port: sPort,
		},
	}, nil
}
