package storage

import (
	"IoTDevicesCore/internal/models"
	"context"

	"gorm.io/gorm"
)

type Storage struct {
	Data interface {
		Upload(context.Context, *models.Data) error
		Recive(context.Context) ([]models.Data, error)
	}
}

func NewPostgresStorage(db *gorm.DB) Storage {
	return Storage{
		Data: &DataStorage{db},
	}
}
