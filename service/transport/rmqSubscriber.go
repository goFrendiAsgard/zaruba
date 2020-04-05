package transport

import (
	"log"
	"os"
)

// NewRmqSubscriber create new RmqSubscriber
func NewRmqSubscriber(connectionString string) *RmqSubscriber {
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
func (s *RmqSubscriber) Subscribe() {
	// create connection and channel
	conn, ch, err := rmqCreateConnectionAndChannel(s.connectionString)
	if err != nil {
		s.logger.Println("[ERROR]", err)
		return
	}
	defer conn.Close()
	defer ch.Close()
	for eventName, handler := range s.handlers {
		// start consume
		_, err = rmqDeclareQueueAndBindToDefaultExchange(ch, eventName)
		if err != nil {
			s.logger.Println("[ERROR]", err)
			return
		}
		rmqMessages, err := rmqConsume(ch, eventName)
		if err != nil {
			s.logger.Println("[ERROR]", err)
			return
		}
		// handle message
		messageHandler := handler
		go func() {
			for rmqMessage := range rmqMessages {
				envelopedMessage, err := NewEnvelopedMessageFromJSON(rmqMessage.Body)
				if err != nil {
					s.logger.Println("[ERROR]", err)
					continue
				}
				err = messageHandler(envelopedMessage.Message)
				if err != nil {
					s.logger.Println("[ERROR]", err)
				}
			}
		}()
	}
	forever := make(chan bool)
	<-forever
}
