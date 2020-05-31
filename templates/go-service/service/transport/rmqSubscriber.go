package transport

import (
	"log"

	"github.com/streadway/amqp"
)

// CreateRmqSubscriber create new RmqSubscriber
func CreateRmqSubscriber(logger *log.Logger, connection *amqp.Connection) *RmqSubscriber {
	return &RmqSubscriber{
		connection: connection,
		handlers:   map[string]EventHandler{},
		logger:     logger,
	}
}

// RmqSubscriber for publish and subscribe
type RmqSubscriber struct {
	connection *amqp.Connection
	handlers   map[string]EventHandler
	logger     *log.Logger
}

// RegisterHandler register servicemap for call
func (s *RmqSubscriber) RegisterHandler(eventName string, handler EventHandler) Subscriber {
	s.handlers[eventName] = handler
	return s
}

// Subscribe consuming message from all event
func (s *RmqSubscriber) Subscribe(errChan chan error) {
	// create connection and channel
	ch, err := s.connection.Channel()
	if err != nil {
		s.handleRmqError(err, errChan)
		return
	}
	amqpErrChan := s.connection.NotifyClose(make(chan *amqp.Error))
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
