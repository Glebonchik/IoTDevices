package main

import (
	"IoTDevicesGenerator/internal/broker"
	"IoTDevicesGenerator/internal/domain"
	"IoTDevicesGenerator/internal/usecase"
	"IoTDevicesGenerator/pkg/config"
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.LoadCfg()
	if err != nil {
		log.Fatalf("Unable to load cfg file: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	//потом сюда конфиг
	brokerUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.Broker.Login,
		cfg.Broker.Password,
		cfg.Broker.Host,
		cfg.Broker.Port,
	)

	rabbitmq, err := broker.NewBrokerConn(brokerUrl, cfg.Broker.QueueName)

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
				log.Printf("Producer error: %v", err)
			} else {
				log.Printf("Data sent: %v", data)
			}
		}
	}
}
