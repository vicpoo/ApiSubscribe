// messaging_service.go
package infrastructure

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain/entities"
)

type MessagingService struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewMessagingService() (*MessagingService, error) {
	conn, err := amqp.Dial("amqp://reyhades:reyhades@44.223.218.9:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		return nil, err
	}

	// Declarar el intercambio
	err = ch.ExchangeDeclare(
		"orders_created", // nombre del intercambio
		"direct",         // tipo de intercambio
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
		return nil, err
	}

	return &MessagingService{conn: conn, ch: ch}, nil
}

func (ms *MessagingService) ConsumeOrderCreated() (<-chan []byte, error) {
	q, err := ms.ch.QueueDeclare(
		"cocina_queue", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ms.ch.QueueBind(
		q.Name,           // queue name
		"order_created",  // routing key
		"orders_created", // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ms.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	byteChan := make(chan []byte)
	go func() {
		for msg := range msgs {
			byteChan <- msg.Body
		}
	}()

	return byteChan, nil
}

func (ms *MessagingService) PublishOrderReady(pedido *entities.Orden) error {
	orderJSON, err := json.Marshal(pedido)
	if err != nil {
		return err
	}

	return ms.ch.Publish(
		"orders_ready", // exchange
		"order_ready",  // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        orderJSON,
		},
	)
}

func (ms *MessagingService) Close() {
	ms.ch.Close()
	ms.conn.Close()
}
