package transport

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// CreateRmqRPCServer create new RmqRPC
func CreateRmqRPCServer(connectionString string) *RmqRPCServer {
	return &RmqRPCServer{
		connectionString: connectionString,
		handlers:         map[string]RPCHandler{},
		logger:           log.New(os.Stdout, "", log.LstdFlags),
	}
}

// RmqRPCServer implementation
type RmqRPCServer struct {
	connectionString string
	handlers         map[string]RPCHandler
	logger           *log.Logger
}

// SetLogger set custome logger
func (s *RmqRPCServer) SetLogger(logger *log.Logger) RPCServer {
	s.logger = logger
	return s
}

// RegisterHandler register servicemap for call
func (s *RmqRPCServer) RegisterHandler(functionName string, handler RPCHandler) RPCServer {
	s.handlers[functionName] = handler
	return s
}

// Serve serve RPC
func (s *RmqRPCServer) Serve(errChan chan error) {
	// create connection and channel
	conn, ch, err := rmqCreateConnectionAndChannel(s.connectionString)
	if err != nil {
		s.logger.Println("[ERROR RmqRPCServer]", err)
		errChan <- err
		return
	}
	defer conn.Close()
	defer ch.Close()
	for functionName, handler := range s.handlers {
		_, err = rmqDeclareQueueAndBindToDefaultExchange(ch, functionName)
		if err != nil {
			s.logger.Println("[ERROR RmqRPCServer]", err)
			errChan <- err
			return
		}
		// start consume
		s.logger.Printf("[INFO RmqRPCServer] Serve %s", functionName)
		rmqMessages, err := rmqConsume(ch, functionName)
		if err != nil {
			s.logger.Println("[ERROR RmqRPCServer]", err)
			errChan <- err
			return
		}
		// handle message
		messageHandler := handler
		thisFunctionName := functionName
		go func() {
			for rmqMessage := range rmqMessages {
				replyTo := rmqMessage.ReplyTo
				envelopedInput, err := NewEnvelopedMessageFromJSON(rmqMessage.Body)
				if err != nil {
					s.handleError(ch, thisFunctionName, replyTo, envelopedInput, err)
					continue
				}
				inputs, err := envelopedInput.Message.GetInterfaceArray("inputs")
				if err != nil {
					s.handleError(ch, thisFunctionName, replyTo, envelopedInput, err)
					continue
				}
				output, err := messageHandler(inputs...)
				if err != nil {
					s.handleError(ch, thisFunctionName, replyTo, envelopedInput, err)
					continue
				}
				s.logger.Printf("[INFO RmqRPCServer] Reply %s: %#v", thisFunctionName, output)
				rmqRpcReply(ch, replyTo, envelopedInput, output)
			}
		}()
	}
	forever := make(chan bool)
	<-forever
}

func (s *RmqRPCServer) handleError(ch *amqp.Channel, thisFunctionName, replyTo string, envelopedInput *EnvelopedMessage, err error) {
	s.logger.Printf(fmt.Sprintf("[ERROR RmqRPCServer] Reply %s:", thisFunctionName), err)
	rmqRpcReplyError(ch, replyTo, envelopedInput, err)
}
