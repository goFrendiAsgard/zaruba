package transport

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// CreateRmqRPCServer create new RmqRPC
func CreateRmqRPCServer(logger *log.Logger, connection *amqp.Connection) *RmqRPCServer {
	return &RmqRPCServer{
		connection: connection,
		handlers:   map[string]RPCHandler{},
		logger:     logger,
	}
}

// RmqRPCServer implementation
type RmqRPCServer struct {
	connection *amqp.Connection
	handlers   map[string]RPCHandler
	logger     *log.Logger
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
		_, err = rmqDeclareQueueAndBindToDefaultExchange(ch, functionName)
		if err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		// start consume
		s.logger.Printf("[INFO RmqRPCServer] Serve %s", functionName)
		rmqMessages, err := rmqConsume(ch, functionName)
		if err != nil {
			s.handleRmqError(err, errChan)
			return
		}
		// handle message
		thisHandler := handler
		thisFunctionName := functionName
		go s.handleRmqMessages(ch, thisFunctionName, thisHandler, rmqMessages)
	}
	err = <-amqpErrChan
	errChan <- err
}

func (s *RmqRPCServer) handleRmqMessages(ch *amqp.Channel, functionName string, handler RPCHandler, rmqMessages <-chan amqp.Delivery) {
	for rmqMessage := range rmqMessages {
		replyTo := rmqMessage.ReplyTo
		envelopedInput, err := CreateEnvelopedMessageFromJSON(rmqMessage.Body)
		if err != nil {
			s.handleRmqMessageError(ch, functionName, replyTo, envelopedInput, err)
			continue
		}
		inputs, err := envelopedInput.Message.GetInterfaceArray("inputs")
		if err != nil {
			s.handleRmqMessageError(ch, functionName, replyTo, envelopedInput, err)
			continue
		}
		output, err := handler(inputs...)
		if err != nil {
			s.handleRmqMessageError(ch, functionName, replyTo, envelopedInput, err)
			continue
		}
		s.logger.Printf("[INFO RmqRPCServer] Reply %s: %#v", functionName, output)
		rmqRPCReplyOutput(ch, replyTo, envelopedInput, output)
	}
}

func (s *RmqRPCServer) handleRmqError(err error, errChan chan error) {
	s.logger.Println("[ERROR RmqRPCServer]", err)
	errChan <- err
}

func (s *RmqRPCServer) handleRmqMessageError(ch *amqp.Channel, thisFunctionName, replyTo string, envelopedInput *EnvelopedMessage, err error) {
	s.logger.Printf(fmt.Sprintf("[ERROR RmqRPCServer] Reply %s:", thisFunctionName), err)
	rmqRPCReplyError(ch, replyTo, envelopedInput, err)
}
