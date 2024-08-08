package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emmiter struct {
	connection *amqp.Connection
}

func (e *Emmiter) setup() error {
	//create a channel
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emmiter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Println("Pushing to channel")

	err = channel.Publish(
		"logs_topic", //exchange
		severity,     //routing key
		false,        //mandatory
		false,        //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	if err != nil {	
		return err
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emmiter, error) {
	emmiter := Emmiter{
		connection: conn,
	}

	err := emmiter.setup()
	if err != nil {
		return Emmiter{}, err
	}

	return emmiter, nil
}