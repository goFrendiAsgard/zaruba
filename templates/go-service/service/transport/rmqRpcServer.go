package transport

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

// NewRmqRPCServer create new RmqRPC
func NewRmqRPCServer(connectionString string) *RmqRPCServer {
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

// Serve
func (s *RmqRPCServer) Serve() {
	// create connection and channel
	conn, ch, err := rmqCreateConnectionAndChannel(s.connectionString)
	if err != nil {
		s.logger.Println("[ERROR]", err)
		return
	}
	defer conn.Close()
	defer ch.Close()
	for functionName, handler := range s.handlers {
		_, err = rmqDeclareQueueAndBindToDefaultExchange(ch, functionName)
		if err != nil {
			s.logger.Println("[ERROR]", err)
			return
		}
		// start consume
		rmqMessages, err := rmqConsume(ch, functionName)
		if err != nil {
			s.logger.Println("[ERROR]", err)
			return
		}
		// handle message
		messageHandler := handler
		go func() {
			for rmqMessage := range rmqMessages {
				replyTo := rmqMessage.ReplyTo
				envelopedInput, err := NewEnvelopedMessageFromJSON(rmqMessage.Body)
				if err != nil {
					s.handleError(ch, replyTo, envelopedInput, err)
					continue
				}
				inputs, err := envelopedInput.Message.GetInterfaceArray("inputs")
				if err != nil {
					s.handleError(ch, replyTo, envelopedInput, err)
					continue
				}
				output, err := messageHandler(inputs...)
				if err != nil {
					s.handleError(ch, replyTo, envelopedInput, err)
					continue
				}
				rmqRpcReply(ch, replyTo, envelopedInput, output)
			}
		}()
	}
	forever := make(chan bool)
	<-forever
}

func (s *RmqRPCServer) handleError(ch *amqp.Channel, replyTo string, envelopedInput *EnvelopedMessage, err error) {
	s.logger.Println("[ERROR]", err)
	rmqRpcReplyError(ch, replyTo, envelopedInput, err)
}
