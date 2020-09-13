package transport

import (
	"log"

	"github.com/streadway/amqp"
)

// CreateRmqSubscriber create new RmqSubscriber
func CreateRmqSubscriber(logger *log.Logger, connection *amqp.Connection, eventMap *RmqEventMap) *RmqSubscriber {
	return &RmqSubscriber{
		connection: connection,
		handlers:   map[string]EventHandler{},
		logger:     logger,
		eventMap:   eventMap,
	}
}

// RmqSubscriber for publish and subscribe
type RmqSubscriber struct {
	connection *amqp.Connection
	handlers   map[string]EventHandler
	logger     *log.Logger
	eventMap   *RmqEventMap
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
		// declare dlx
		var args amqp.Table = nil
		if s.eventMap.GetTTL(eventName) > 0 {
			deadLetterExchange := s.eventMap.GetDeadLetterExchange(eventName)
			deadLetterQueue := s.eventMap.GetDeadLetterQueue(eventName)
			if _, err = rmqDeclareAndBindQueue(ch, deadLetterExchange, deadLetterQueue, true, false, false, false, nil); err != nil {
				s.handleRmqError(err, errChan)
				return
			}
			args = s.eventMap.GetQueueArgs(eventName)
		}
		// declare queue
		exchangeName, queueName := s.eventMap.GetExchangeName(eventName), s.eventMap.GetQueueName(eventName)
		if _, err = rmqDeclareAndBindQueue(ch, exchangeName, queueName, true, false, false, false, args); err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		s.logger.Printf("[INFO RmqRPCSubscriber] Subscribe %s", eventName)
		autoAck := s.eventMap.GetAutoAck(eventName)
		rmqMessages, err := rmqConsume(ch, queueName, true, false, false, false, autoAck, args)
		if err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		// handle message
		thisHandler, thisEventName := handler, eventName
		go s.handleRmqMessages(thisEventName, thisHandler, rmqMessages, autoAck)
	}
	err = <-amqpErrChan
	errChan <- err
}

func (s *RmqSubscriber) handleRmqMessages(eventName string, handler EventHandler, rmqMessages <-chan amqp.Delivery, autoAck bool) {
	for rmqMessage := range rmqMessages {
		envelopedMessage, err := CreateEnvelopedMessageFromJSON(rmqMessage.Body)
		if err != nil {
			s.logger.Println("[ERROR RmqSubscriber]", err)
			if !autoAck {
				rmqMessage.Nack(false, true)
			}
			continue
		}
		s.logger.Printf("[INFO RmqSubscriber] Get %s: %#v", eventName, envelopedMessage.Message)
		if err = handler(envelopedMessage.Message); err != nil {
			s.logger.Println("[ERROR RmqSubscriber]", err)
			if !autoAck {
				rmqMessage.Nack(false, true)
			}
		} else if !autoAck {
			rmqMessage.Ack(false)
		}
	}
}

func (s *RmqSubscriber) handleRmqError(err error, errChan chan error) {
	s.logger.Println("[ERROR RmqSubscriber]", err)
	errChan <- err
}
