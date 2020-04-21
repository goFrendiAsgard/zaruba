package transport

import (
	"log"
	"os"
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
		s.logger.Println("[ERROR RmqSubscriber]", err)
		errChan <- err
		return
	}
	defer conn.Close()
	defer ch.Close()
	for eventName, handler := range s.handlers {
		// start consume
		_, err = rmqDeclareQueueAndBindToDefaultExchange(ch, eventName)
		if err != nil {
			s.logger.Println("[ERROR RmqSubscriber]", err)
			errChan <- err
			return
		}
		s.logger.Printf("[INFO RmqRPCSubscriber] Subscribe %s", eventName)
		rmqMessages, err := rmqConsume(ch, eventName)
		if err != nil {
			s.logger.Println("[ERROR RmqSubscriber]", err)
			errChan <- err
			return
		}
		// handle message
		messageHandler := handler
		thisEventName := eventName
		go func() {
			for rmqMessage := range rmqMessages {
				envelopedMessage, err := NewEnvelopedMessageFromJSON(rmqMessage.Body)
				if err != nil {
					s.logger.Println("[ERROR RmqSubscriber]", err)
					continue
				}
				s.logger.Printf("[INFO RmqRPCSubscriber] Get %s: %#v", thisEventName, envelopedMessage.Message)
				err = messageHandler(envelopedMessage.Message)
				if err != nil {
					s.logger.Println("[ERROR RmqSubscriber]", err)
				}
			}
		}()
	}
	forever := make(chan bool)
	<-forever
}
