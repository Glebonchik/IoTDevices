package storage

import (
	"IoTDevicesCore/internal/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type DataStorage struct {
	db *gorm.DB
}

func (d *DataStorage) Upload(ctx context.Context, data *models.Data) error {
	if err := d.db.WithContext(ctx).Save(data).Error; err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (d *DataStorage) Recive(ctx context.Context) ([]models.Data, error) {

	var data []models.Data
	if err := d.db.WithContext(ctx).Find(data).Error; err != nil {
		return data, fmt.Errorf("%w", err)
	}

	return data, nil
}
