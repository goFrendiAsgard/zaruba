package transport

import (
	"log"

	"github.com/streadway/amqp"
)

// CreateRmqPublisher create new RmqPublisher
func CreateRmqPublisher(logger *log.Logger, connection *amqp.Connection) *RmqPublisher {
	return &RmqPublisher{
		connection: connection,
		logger:     logger,
	}
}

// RmqPublisher for publish and subscribe
type RmqPublisher struct {
	connection *amqp.Connection
	logger     *log.Logger
}

// Publish publish message to event
func (p *RmqPublisher) Publish(eventName string, message Message) (err error) {
	// create connection and channel
	ch, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	// declare exchange
	err = rmqDeclareFanoutExchange(ch, eventName)
	if err != nil {
		return err
	}
	// create envelopedMessage
	envelopedMessage := CreateEnvelopedMessage()
	if err = envelopedMessage.SetNewCorrelationID(); err != nil {
		return err
	}
	envelopedMessage.Message = message
	// make json representation of envelopedMessage
	jsonMessage, err := envelopedMessage.ToJSON()
	if err != nil {
		return err
	}
	// publish to exchange
	p.logger.Printf("[INFO RmqPublisher] Publish %s %#v", eventName, message)
	return rmqPublish(ch, eventName, "",
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: envelopedMessage.CorrelationID,
			Body:          jsonMessage,
		})
}
