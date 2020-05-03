package transport

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

// CreateRmqSubscriber create new RmqSubscriber
func CreateRmqSubscriber(connectionString string) *RmqSubscriber {
	return &RmqSubscriber{
		connectionString: connectionString,
		handlers:         map[string]EventHandler{},
		logger:           log.New(os.Stdout, "", log.LstdFlags),
	}
}

// RmqSubscriber for publish and subscribe
type RmqSubscriber struct {
	connectionString string
	handlers         map[string]EventHandler
	logger           *log.Logger
}

// SetLogger set custome logger
func (s *RmqSubscriber) SetLogger(logger *log.Logger) Subscriber {
	s.logger = logger
	return s
}

// RegisterHandler register servicemap for call
func (s *RmqSubscriber) RegisterHandler(eventName string, handler EventHandler) Subscriber {
	s.handlers[eventName] = handler
	return s
}

// Subscribe consuming message from all event
func (s *RmqSubscriber) Subscribe(errChan chan error) {
	// create connection and channel
	conn, ch, err := rmqCreateConnectionAndChannel(s.connectionString)
	if err != nil {
		s.handleRmqError(err, errChan)
		return
	}
	amqpErrChan := conn.NotifyClose(make(chan *amqp.Error))
	defer conn.Close()
	defer ch.Close()
	for eventName, handler := range s.handlers {
		// start consume
		_, err = rmqDeclareQueueAndBindToDefaultExchange(ch, eventName)
		if err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		s.logger.Printf("[INFO RmqRPCSubscriber] Subscribe %s", eventName)
		rmqMessages, err := rmqConsume(ch, eventName)
		if err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		// handle message
		thisHandler := handler
		thisEventName := eventName
		go s.handleRmqMessages(thisEventName, thisHandler, rmqMessages)
	}
	err = <-amqpErrChan
	errChan <- err
}

func (s *RmqSubscriber) handleRmqMessages(eventName string, handler EventHandler, rmqMessages <-chan amqp.Delivery) {
	for rmqMessage := range rmqMessages {
		envelopedMessage, err := CreateEnvelopedMessageFromJSON(rmqMessage.Body)
		if err != nil {
			s.logger.Println("[ERROR RmqSubscriber]", err)
			continue
		}
		s.logger.Printf("[INFO RmqSubscriber] Get %s: %#v", eventName, envelopedMessage.Message)
		err = handler(envelopedMessage.Message)
		if err != nil {
			s.logger.Println("[ERROR RmqSubscriber]", err)
		}
	}
}

func (s *RmqSubscriber) handleRmqError(err error, errChan chan error) {
	s.logger.Println("[ERROR RmqSubscriber]", err)
	errChan <- err
}
