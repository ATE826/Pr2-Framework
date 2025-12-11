package config

import "os"

type Config struct {
	UserServiceURL  string
	OrderServiceURL string
	Port            string
	JWTSecret       string
}

func Load() *Config {
	cfg := &Config{
		UserServiceURL:  os.Getenv("USER_SERVICE_URL"),
		OrderServiceURL: os.Getenv("ORDER_SERVICE_URL"),
		Port:            os.Getenv("PORT"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.UserServiceURL == "" {
		cfg.UserServiceURL = "http://service_users:8001"
	}
	if cfg.OrderServiceURL == "" {
		cfg.OrderServiceURL = "http://service_orders:8002"
	}
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "secret123"
	}

	return cfg
}
