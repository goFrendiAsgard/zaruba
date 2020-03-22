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
		conn, ch, err := pubSub.createConnectionAndChannel()
		if err != nil {
			pubSub.logger.Println("[ERROR]", err)
			return
		}
		defer conn.Close()
		defer ch.Close()
		// declare queue and bind
		q, err := pubSub.declareAndBindQueueToExchange(ch, eventName)
		if err != nil {
			pubSub.logger.Println("[ERROR]", err)
			return
		}
		// start consume
		rmqMessages, err := pubSub.consume(ch, q)
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
					return
				}
				err = messageHandler(envelopedMessage.Message)
				if err != nil {
					pubSub.logger.Println("[ERROR]", err)
					return
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
	conn, ch, err := pubSub.createConnectionAndChannel()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()
	// declare exchange
	err = pubSub.declareFanoutExchange(ch, eventName)
	if err != nil {
		return err
	}
	// create jsonMessage
	jsonMessage, err := pubSub.messageToJSON(message)
	if err != nil {
		return err
	}
	// publish to exchange
	return ch.Publish(
		eventName, // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonMessage,
		})
}

func (pubSub *RmqPubSub) createConnectionAndChannel() (conn *amqp.Connection, ch *amqp.Channel, err error) {
	conn, err = amqp.Dial(pubSub.connectionString)
	if err != nil {
		return conn, ch, err
	}
	ch, err = conn.Channel()
	return conn, ch, err
}

func (pubSub *RmqPubSub) declareAndBindQueueToExchange(ch *amqp.Channel, queueName string) (q amqp.Queue, err error) {
	// declare exchange
	err = pubSub.declareFanoutExchange(ch, queueName)
	if err != nil {
		return q, err
	}
	// declare queue
	q, err = pubSub.declareQueue(ch, queueName)
	if err != nil {
		return q, err
	}
	// bind queue to exchange
	err = ch.QueueBind(
		q.Name,    // queue name
		"",        // routing key
		queueName, // exchange
		false,
		nil,
	)
	return q, err
}

func (pubSub *RmqPubSub) messageToJSON(message Message) (envelopedJSON []byte, err error) {
	envelopedMessage, err := NewEnvelopedMessage(message)
	if err != nil {
		return envelopedJSON, err
	}
	return envelopedMessage.ToJSON()
}

func (pubSub *RmqPubSub) declareQueue(ch *amqp.Channel, queueName string) (q amqp.Queue, err error) {
	return ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
}

func (pubSub *RmqPubSub) declareFanoutExchange(ch *amqp.Channel, exchangeName string) (err error) {
	return ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func (pubSub *RmqPubSub) consume(ch *amqp.Channel, q amqp.Queue) (rmqMessages <-chan amqp.Delivery, err error) {
	// start consume
	return ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}
