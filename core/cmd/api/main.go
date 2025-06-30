package main

import (
	"IoTDevicesCore/internal/broker"
	"IoTDevicesCore/internal/usecase"
	"IoTDevicesCore/pkg/config"
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
)

func main() {

	cfg, err := config.LoadCfg()
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	// db, dbErr := db.New(cfg)
	// if dbErr != nil {
	// 	log.Fatal("Cannot conntect to DB: ", dbErr)
	// }
	// store := storage.NewPostgresStorage(db)

	brokerUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.MsgBroker.Login,
		cfg.MsgBroker.Password,
		cfg.MsgBroker.Host,
		cfg.MsgBroker.Port,
	)

	log.Println("Attemting to connect broker (RabbitMq) with: ", brokerUrl)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	rabbit, err := broker.NewConnection(brokerUrl, cfg.MsgBroker.QueueName)
	if err != nil {
		log.Fatal("Cannot subscribe to RabbitMQ: ", err)
	}

	consumer := usecase.NewDeviceDataConsumer(rabbit)

	msgChan, err := consumer.Consume(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for devices := range msgChan {
		log.Printf("Data recived: %v", devices)
	}

	<-ctx.Done()
	log.Println("Starting graceful shutdown...")
	rabbit.CloseConnection()
	log.Println("Graceful shutdown done.")

	// app := &application.Application{
	// 	Store:  store,
	// 	Config: cfg,
	// }
	// app.RunApp(gin.Default())

}
