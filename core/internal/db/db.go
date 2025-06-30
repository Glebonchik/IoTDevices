package db

import (
	"IoTDevicesCore/internal/models"
	"IoTDevicesCore/pkg/config"
	"fmt"
	"log"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	dbConf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DataBase.Host,
		cfg.DataBase.Port,
		cfg.DataBase.User,
		cfg.DataBase.Password,
		cfg.DataBase.DBName,
		cfg.DataBase.SSLMode,
	)
	log.Println("Attempting to connect postgressDB:", dbConf)
	db, err := gorm.Open(postgres.Open(dbConf), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	err = db.AutoMigrate(&models.Data{})

	if err != nil {
		log.Fatal("failed to auto migrate", slog.String("error", err.Error()))
		return nil, nil
	}
	log.Println("migrations applied successfully")

	return db, nil
}
