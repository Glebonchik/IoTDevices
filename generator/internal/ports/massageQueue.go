package ports

import (
	"IoTDevicesGenerator/internal/domain"
	"context"
)

type MassageQueue interface {
	Publish(ctx context.Context, d []domain.Data) error
	CloseConnection()
}
