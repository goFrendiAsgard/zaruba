package transport

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

// CreateRmqPublisher create new RmqPublisher
func CreateRmqPublisher(connectionString string) *RmqPublisher {
	return &RmqPublisher{
		connectionString: connectionString,
		logger:           log.New(os.Stdout, "", log.LstdFlags),
	}
}

// RmqPublisher for publish and subscribe
type RmqPublisher struct {
	connectionString string
	logger           *log.Logger
}

// SetLogger set custome logger
func (p *RmqPublisher) SetLogger(logger *log.Logger) Publisher {
	p.logger = logger
	return p
}

// Publish publish message to event
func (p *RmqPublisher) Publish(eventName string, message Message) (err error) {
	// create connection and channel
	conn, ch, err := rmqCreateConnectionAndChannel(p.connectionString)
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()
	// declare exchange
	err = rmqDeclareFanoutExchange(ch, eventName)
	if err != nil {
		return err
	}
	// create envelopedMessage
	envelopedMessage, err := NewEnvelopedMessage(message)
	if err != nil {
		return err
	}
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
