package broker

import (
	"IoTDevicesGenerator/internal/domain"
	"IoTDevicesGenerator/internal/ports"
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
	Conn    *amqp.Connection
}

func NewBrokerConn(url string, name string) (ports.MassageQueue, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
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
		return nil, err
	}

	return &RabbitMQ{q, ch, conn}, nil
}

func (r *RabbitMQ) CloseConnection() {
	r.Channel.Close()
	r.Conn.Close()
}

func (r *RabbitMQ) Publish(ctx context.Context, data []domain.Data) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return r.Channel.PublishWithContext(ctx,
		"",
		r.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
