package msgqueue

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	SendMessage(message string) error
}

type service struct {
	amqpURL string
}

// NewService is used to create a single instance of the service
func NewService(amqpURL string) Service {
	return &service{
		amqpURL: amqpURL,
	}
}

func (s *service) SendMessage(message string) error {
	fmt.Println("Helo t1")
	conn, err := amqp.Dial(
		s.amqpURL,
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.Qos(1, 0, false)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = ch.PublishWithContext(
		ctx, "", q.Name, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message),
		})

	if err != nil {
		return err
	}
	return nil
}
