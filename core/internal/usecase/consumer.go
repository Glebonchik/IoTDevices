package usecase

import (
	"IoTDevicesCore/internal/models"
	"IoTDevicesCore/internal/ports"
	"context"
)

type DeviceDataConsumer interface {
	Consume(ctx context.Context) (<-chan []models.Data, error)
}

type deviceDataConsumer struct {
	queue ports.MessageQueue
}

func NewDeviceDataConsumer(queue ports.MessageQueue) DeviceDataConsumer {
	return &deviceDataConsumer{queue: queue}
}

func (c *deviceDataConsumer) Consume(ctx context.Context) (<-chan []models.Data, error) {
	ch, err := c.queue.Subscribe(ctx)
	if err != nil {
		c.queue.CloseConnection()
		return nil, err
	}
	out := make(chan []models.Data, 100)
	go func() {
		defer close(out)
		for devices := range ch {
			select {
			case out <- devices:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, nil
}
