package main

import (
	"IoTDevicesGenerator/internal/broker"
	"IoTDevicesGenerator/internal/domain"
	"IoTDevicesGenerator/internal/usecase"
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	//потом сюда конфиг
	brokerUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		"guest",
		"guest",
		"localhost",
		"5672",
	)

	rabbitmq, err := broker.NewBrokerConn(brokerUrl, "io_devices")

	if err != nil {
		log.Fatal("Unable to reach broker: ", err)
		rabbitmq.CloseConnection()
	}

	go func() {
		<-ctx.Done()
		log.Println("Graceful shutdown start...")
		rabbitmq.CloseConnection()
	}()

	producer := usecase.NewDeviceDataProducer(rabbitmq)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			data := domain.GenerateData()
			if err := producer.Produce(ctx, data); err != nil {
				log.Printf("Ошибка публикации: %v", err)
			} else {
				log.Printf("Отправлены данные: %v", data)
			}
		}
	}
}
