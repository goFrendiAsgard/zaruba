package transport

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// CreateRmqRPCServer create new RmqRPC
func CreateRmqRPCServer(logger *log.Logger, connection *amqp.Connection, eventMap *RmqEventMap) *RmqRPCServer {
	return &RmqRPCServer{
		connection: connection,
		handlers:   map[string]RPCHandler{},
		logger:     logger,
		eventMap:   eventMap,
	}
}

// RmqRPCServer implementation
type RmqRPCServer struct {
	connection *amqp.Connection
	handlers   map[string]RPCHandler
	logger     *log.Logger
	eventMap   *RmqEventMap
}

// RegisterHandler register servicemap for call
func (s *RmqRPCServer) RegisterHandler(functionName string, handler RPCHandler) RPCServer {
	s.handlers[functionName] = handler
	return s
}

// Serve serve RPC
func (s *RmqRPCServer) Serve(errChan chan error) {
	// create connection and channel
	ch, err := s.connection.Channel()
	if err != nil {
		s.handleRmqError(err, errChan)
		return
	}
	amqpErrChan := s.connection.NotifyClose(make(chan *amqp.Error))
	defer ch.Close()
	for functionName, handler := range s.handlers {
		// declare dlx
		var args amqp.Table = nil
		if s.eventMap.GetTTL(functionName) > 0 {
			deadLetterExchange := s.eventMap.GetDeadLetterExchange(functionName)
			deadLetterQueue := s.eventMap.GetDeadLetterQueue(functionName)
			if _, err = rmqDeclareAndBindQueue(ch, deadLetterExchange, deadLetterQueue, true, false, false, false, nil); err != nil {
				s.handleRmqError(err, errChan)
				return
			}
			args = s.eventMap.GetQueueArgs(functionName)
		}
		// declare queue
		exchangeName := s.eventMap.GetExchangeName(functionName)
		queueName := s.eventMap.GetQueueName(functionName)
		_, err = rmqDeclareAndBindQueue(ch, exchangeName, queueName, true, false, false, false, args)
		if err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		// start consume
		s.logger.Printf("[INFO RmqRPCServer] Serve %s", functionName)
		autoAck := s.eventMap.GetAutoAck(functionName)
		rmqMessages, err := rmqConsume(ch, queueName, true, false, false, false, autoAck, args)
		if err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		// handle message
		thisHandler, thisFunctionName := handler, functionName
		go s.handleRmqMessages(ch, thisFunctionName, thisHandler, rmqMessages, autoAck)
	}
	err = <-amqpErrChan
	errChan <- err
}

func (s *RmqRPCServer) handleRmqMessages(ch *amqp.Channel, functionName string, handler RPCHandler, rmqMessages <-chan amqp.Delivery, autoAck bool) {
	for rmqMessage := range rmqMessages {
		replyTo := rmqMessage.ReplyTo
		envelopedInput, err := CreateEnvelopedMessageFromJSON(rmqMessage.Body)
		if err != nil {
			s.handleRmqMessageError(ch, functionName, replyTo, envelopedInput, err)
			if !autoAck {
				rmqMessage.Nack(false, true)
			}
			continue
		}
		inputs, err := envelopedInput.Message.GetInterfaceArray("inputs")
		if err != nil {
			s.handleRmqMessageError(ch, functionName, replyTo, envelopedInput, err)
			if !autoAck {
				rmqMessage.Nack(false, true)
			}
			continue
		}
		output, err := handler(inputs...)
		if err != nil {
			s.handleRmqMessageError(ch, functionName, replyTo, envelopedInput, err)
			if !autoAck {
				rmqMessage.Nack(false, true)
			}
			continue
		} else if !autoAck {
			rmqMessage.Ack(false)
		}
		s.logger.Printf("[INFO RmqRPCServer] Reply %s by Sending to Queue %s: %#v", functionName, replyTo, output)
		if err = rmqRPCReplyOutput(ch, replyTo, envelopedInput, output); err != nil {
			s.logger.Printf("[INFO RmqRPCServer] Failed to Reply %s by Sending to Queue %s: %#v", functionName, replyTo, err)
		}
	}
}

func (s *RmqRPCServer) handleRmqError(err error, errChan chan error) {
	s.logger.Println("[ERROR RmqRPCServer]", err)
	errChan <- err
}

func (s *RmqRPCServer) handleRmqMessageError(ch *amqp.Channel, thisFunctionName, replyTo string, envelopedInput *EnvelopedMessage, err error) {
	s.logger.Printf(fmt.Sprintf("[ERROR RmqRPCServer] Reply Error %s by Sending to Queue %s:", thisFunctionName, replyTo), err)
	if err = rmqRPCReplyError(ch, replyTo, envelopedInput, err); err != nil {
		s.logger.Printf("[ERROR RmqRPCServer] Failed to Reply Error %s by Sending to Queue %s: %#v", thisFunctionName, replyTo, err)
	}
}
