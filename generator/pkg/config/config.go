package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Broker RabbitMQ
}

type RabbitMQ struct {
	Host      string
	Login     string
	Password  string
	Port      string
	QueueName string
}

func LoadCfg() (*Config, error) {
	// Только для dev-сервера
	if err := godotenv.Load(); err != nil {
		log.Println("No .env warning file found, using environment variables")
	}

	cfg := &Config{}

	cfg.Broker.Host = os.Getenv("RMQ_HOST")
	if cfg.Broker.Host == "" {
		return nil, fmt.Errorf("RMQ_HOST is required")
	}

	cfg.Broker.Login = os.Getenv("RMQ_LOGIN")
	if cfg.Broker.Login == "" {
		cfg.Broker.Login = "guest"
	}

	cfg.Broker.Password = os.Getenv("RMQ_PASSWORD")
	if cfg.Broker.Password == "" {
		cfg.Broker.Password = "guest"
	}

	cfg.Broker.Port = os.Getenv("RMQ_PORT")
	if cfg.Broker.Port == "" {
		cfg.Broker.Port = "5672"
	}

	cfg.Broker.QueueName = os.Getenv("RMQ_QUEUE_NAME")
	if cfg.Broker.QueueName == "" {
		return nil, fmt.Errorf("RMQ_QUEUE_NAME is required")
	}
	return cfg, nil
}
