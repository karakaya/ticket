package rabbit

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TicketRabbit struct {
	ID        uuid.UUID `bson:"_id" json:"id"`
	Title     string    `bson:"title" json:"title"`
	Body      string    `bson:"body" json:"body"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// var conn *amqp.Connection
var channel *amqp.Channel

func ConnRabbitMQ() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:7001/")
	if err != nil {
		return err
	}

	channel, err = conn.Channel()
	if err != nil {
		return err
	}

	channel.QueueDeclare(
		"email", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	return nil
}

func Publish(message []byte) error {
	var ticket TicketRabbit
	json.Unmarshal(message, &ticket)
	err := channel.PublishWithContext(
		context.TODO(),
		"",      // exchange
		"email", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	if err != nil {
		return err
	}
	return nil
}
