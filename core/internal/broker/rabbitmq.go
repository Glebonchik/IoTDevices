package broker

import (
	"IoTDevicesCore/internal/models"
	"IoTDevicesCore/internal/ports"
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
	Conn    *amqp.Connection
}

func NewConnection(url string, name string) (ports.MessageQueue, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{q, ch, conn}, nil

}

func (r *RabbitMQ) CloseConnection() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Conn != nil {
		r.Conn.Close()
	}
}

func (r *RabbitMQ) Subscribe(ctx context.Context) (<-chan []models.Data, error) {
	msgs, err := r.Channel.ConsumeWithContext(ctx,
		r.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Subscribe error: %v", err)
		return nil, err
	}
	ch := make(chan []models.Data)
	go func() {
		defer close(ch)
		for d := range msgs {
			var devices []models.Data
			if err := json.Unmarshal(d.Body, &devices); err != nil {
				log.Printf("Unmarshal error: %v ", err)
				d.Nack(false, true)
				continue
			}
			select {
			case ch <- devices:
				d.Ack(false)
			case <-ctx.Done():
				d.Nack(false, true)
				return
			}
		}

	}()
	return ch, nil
}
