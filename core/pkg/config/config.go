package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataBase  DataBase
	Server    Server
	MsgBroker RabbitMQ
}

type RabbitMQ struct {
	Host      string
	Login     string
	Password  string
	Port      string
	QueueName string
}

type DataBase struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Server struct {
	Port string
}

func LoadCfg() (*Config, error) {
	// Только для dev-сервера
	if err := godotenv.Load(); err != nil {
		log.Println("No .env warning file found, using environment variables")
	}

	cfg := &Config{}

	cfg.DataBase.Host = os.Getenv("DB_HOST")
	if cfg.DataBase.Host == "" {
		return nil, fmt.Errorf("DB_HOST is required")
	}

	cfg.DataBase.Port = os.Getenv("DB_PORT")
	if cfg.DataBase.Port == "" {
		cfg.DataBase.Port = ""
	}

	cfg.DataBase.User = os.Getenv("DB_USER")
	if cfg.DataBase.User == "" {
		return nil, fmt.Errorf("DB_USER is required")
	}

	cfg.DataBase.Password = os.Getenv("DB_PASSWORD")
	if cfg.DataBase.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	cfg.DataBase.DBName = os.Getenv("DB_NAME")
	if cfg.DataBase.DBName == "" {
		return nil, fmt.Errorf("DB_NAME is required")
	}

	cfg.DataBase.SSLMode = os.Getenv("DB_SSLMODE")
	if cfg.DataBase.SSLMode == "" {
		cfg.DataBase.SSLMode = "disable"
	}

	cfg.Server.Port = os.Getenv("SERVER_PORT")
	if cfg.Server.Port == "" {
		cfg.Server.Port = ":8082"
	}

	cfg.MsgBroker.Host = os.Getenv("RMQ_HOST")
	if cfg.MsgBroker.Host == "" {
		return nil, fmt.Errorf("RMQ_HOST is required")
	}

	cfg.MsgBroker.Login = os.Getenv("RMQ_LOGIN")
	if cfg.MsgBroker.Login == "" {
		cfg.MsgBroker.Login = "guest"
	}

	cfg.MsgBroker.Password = os.Getenv("RMQ_PASSWORD")
	if cfg.MsgBroker.Password == "" {
		cfg.MsgBroker.Password = "guest"
	}

	cfg.MsgBroker.Port = os.Getenv("RMQ_PORT")
	if cfg.MsgBroker.Port == "" {
		cfg.MsgBroker.Port = "5672"
	}

	cfg.MsgBroker.QueueName = os.Getenv("RMQ_QUEUE_NAME")
	if cfg.MsgBroker.QueueName == "" {
		return nil, fmt.Errorf("RMQ_QUEUE_NAME is required")
	}
	return cfg, nil
}
