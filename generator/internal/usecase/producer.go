package usecase

import (
	"IoTDevicesGenerator/internal/domain"
	"IoTDevicesGenerator/internal/ports"
	"context"
)

type DeviceDataProducer interface {
	Produce(ctx context.Context, devices []domain.Data) error
}

type deviceDataProducer struct {
	queue ports.MassageQueue
}

func NewDeviceDataProducer(queue ports.MassageQueue) DeviceDataProducer {
	return &deviceDataProducer{queue: queue}
}

func (p *deviceDataProducer) Produce(ctx context.Context, devices []domain.Data) error {
	return p.queue.Publish(ctx, devices)
}
