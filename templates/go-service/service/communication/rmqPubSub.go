package communication

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

// NewRmqPubSub create new RmqPubSub
func NewRmqPubSub(connectionString string) *RmqPubSub {
	return &RmqPubSub{
		connectionString: connectionString,
		handlers:         map[string]PubSubHandler{},
		logger:           log.New(os.Stdout, "", log.LstdFlags),
	}
}

// RmqPubSub for publish and subscribe
type RmqPubSub struct {
	connectionString string
	handlers         map[string]PubSubHandler
	logger           *log.Logger
}

// SetLogger set custome logger
func (pubSub *RmqPubSub) SetLogger(logger *log.Logger) *RmqPubSub {
	pubSub.logger = logger
	return pubSub
}

// RegisterHandler register servicemap for call
func (pubSub *RmqPubSub) RegisterHandler(eventName string, handler PubSubHandler) {
	pubSub.handlers[eventName] = handler
}

// Start consuming message from all event
func (pubSub *RmqPubSub) Start() {
	for eventName, handler := range pubSub.handlers {
		// create connection and channel
		conn, ch, err := createRmqConnectionAndChannel(pubSub.connectionString)
		if err != nil {
			pubSub.logger.Println("[ERROR]", err)
			return
		}
		defer conn.Close()
		defer ch.Close()
		// declare queue and bind
		q, err := declareAndBindRmqQueueToExchange(ch, eventName)
		if err != nil {
			pubSub.logger.Println("[ERROR]", err)
			return
		}
		// start consume
		rmqMessages, err := rmqConsume(ch, q)
		if err != nil {
			pubSub.logger.Println("[ERROR]", err)
			return
		}
		// handle message
		messageHandler := handler
		go func() {
			for rmqMessage := range rmqMessages {
				envelopedMessage, err := NewEnvelopedMessageFromJSON(rmqMessage.Body)
				if err != nil {
					pubSub.logger.Println("[ERROR]", err)
					continue
				}
				err = messageHandler(envelopedMessage.Message)
				if err != nil {
					pubSub.logger.Println("[ERROR]", err)
				}
			}
		}()
	}
	forever := make(chan bool)
	<-forever
}

// Publish publish message to event
func (pubSub *RmqPubSub) Publish(eventName string, message Message) (err error) {
	// create connection and channel
	conn, ch, err := createRmqConnectionAndChannel(pubSub.connectionString)
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()
	// declare exchange
	err = declareRmqFanoutExchange(ch, eventName)
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
	return rmqPublish(ch, eventName, "",
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: envelopedMessage.CorrelationID,
			Body:          jsonMessage,
		})
}
