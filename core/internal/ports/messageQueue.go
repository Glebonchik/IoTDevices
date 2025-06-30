package ports

import (
	"IoTDevicesCore/internal/models"
	"context"
)

type MessageQueue interface {
	Subscribe(ctx context.Context) (<-chan []models.Data, error)
	CloseConnection()
}