package communication

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// NewRmqRPC create RPC based on rmq
func NewRmqRPC(connectionString string) *RmqRPC {
	return &RmqRPC{
		connectionString: connectionString,
		handlers:         map[string]RPCHandler{},
		logger:           log.New(os.Stdout, "", log.LstdFlags),
	}
}

// RmqRPC rmq rpc
type RmqRPC struct {
	connectionString string
	handlers         map[string]RPCHandler
	logger           *log.Logger
}

// SetLogger set custome logger
func (rpc *RmqRPC) SetLogger(logger *log.Logger) *RmqRPC {
	rpc.logger = logger
	return rpc
}

// RegisterHandler register servicemap for call
func (rpc *RmqRPC) RegisterHandler(functionName string, handler RPCHandler) {
	rpc.handlers[functionName] = handler
}

// Serve from remote client
func (rpc *RmqRPC) Serve() {
	for eventName, handler := range rpc.handlers {
		// create connection and channel
		conn, ch, err := createRmqConnectionAndChannel(rpc.connectionString)
		if err != nil {
			rpc.logger.Println("[ERROR]", err)
			return
		}
		defer conn.Close()
		defer ch.Close()
		// declare queue and bind
		q, err := declareAndBindRmqQueueToExchange(ch, eventName)
		if err != nil {
			rpc.logger.Println("[ERROR]", err)
			return
		}
		// start consume
		rmqMessages, err := rmqConsume(ch, q)
		if err != nil {
			rpc.logger.Println("[ERROR]", err)
			return
		}
		// handle message
		messageHandler := handler
		go func() {
			for rmqMessage := range rmqMessages {
				envelopedInput, err := NewEnvelopedMessageFromJSON(rmqMessage.Body)
				if err != nil {
					rpc.replyError(ch, rmqMessage, err)
					continue
				}
				output, err := messageHandler(envelopedInput.Message)
				if err != nil {
					rpc.replyError(ch, rmqMessage, err)
					continue
				}
				rpc.replySuccess(ch, rmqMessage, output)
			}
		}()
	}
	forever := make(chan bool)
	<-forever
}

func (rpc *RmqRPC) replySuccess(ch *amqp.Channel, rmqMessage amqp.Delivery, output Message) error {
	replyTo := rmqMessage.ReplyTo
	correlationID := rmqMessage.CorrelationId
	envelopedOutput := NewEnvelopedMessageWithCorrelationID(correlationID, output)
	jsonMessage, err := envelopedOutput.ToJSON()
	if err != nil {
		return rpc.replyError(ch, rmqMessage, err)
	}
	err = rmqPublish(ch, "", replyTo,
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: correlationID,
			Body:          jsonMessage,
		})
	if err != nil {
		return rpc.replyError(ch, rmqMessage, err)
	}
	return err
}

func (rpc *RmqRPC) replyError(ch *amqp.Channel, rmqMessage amqp.Delivery, err error) error {
	replyTo := rmqMessage.ReplyTo
	correlationID := rmqMessage.CorrelationId
	envelopedOutput := NewEnvelopedMessageWithCorrelationID(correlationID, Message{})
	envelopedOutput.ErrorMessage = fmt.Sprintf("%s", err)
	jsonMessage, err := envelopedOutput.ToJSON()
	if err != nil {
		rpc.logger.Println("[ERROR]", err)
		return err
	}
	err = rmqPublish(ch, "", replyTo,
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: correlationID,
			Body:          jsonMessage,
		})
	if err != nil {
		rpc.logger.Println("[ERROR]", err)
	}
	return err
}

// Call remote function
func (rpc *RmqRPC) Call(servicename, functionName string, input Message) (output Message, err error) {
	replyTo, err := rpc.generateReplyQueueName(functionName)
	if err != nil {
		return output, err
	}
	// create connection and channel
	conn, ch, err := createRmqConnectionAndChannel(rpc.connectionString)
	if err != nil {
		return output, err
	}
	defer conn.Close()
	defer ch.Close()
	// declare queue
	q, err := declareRmqQueue(ch, replyTo)
	if err != nil {
		return output, err
	}
	// start consume
	rmqMessages, err := rmqConsume(ch, q)
	if err != nil {
		return output, err
	}
	// create waiting channel
	waitReply := make(chan bool)
	go func() {
		for rmqMessage := range rmqMessages {
			envelopedOutput, parseError := NewEnvelopedMessageFromJSON(rmqMessage.Body)
			if parseError != nil {
				err = parseError
				waitReply <- true
				return
			}
			output = envelopedOutput.Message
			waitReply <- true
			return
		}
	}()
	// create envelopedMessage
	envelopedInput, err := NewEnvelopedMessage(input)
	if err != nil {
		return output, err
	}
	// make json representation of envelopedMessage
	jsonMessage, err := envelopedInput.ToJSON()
	if err != nil {
		return output, err
	}
	// publish
	err = rmqPublish(ch, functionName, "",
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: envelopedInput.CorrelationID,
			ReplyTo:       replyTo,
			Body:          jsonMessage,
		})
	if err != nil {
		return output, err
	}
	// return
	<-waitReply
	ch.QueueDelete(replyTo, false, false, true)
	return output, err
}

func (rpc *RmqRPC) generateReplyQueueName(functionName string) (queueName string, err error) {
	randomID, err := uuid.NewUUID()
	if err != nil {
		return queueName, err
	}
	queueName = fmt.Sprintf("%s.reply.%s", functionName, strings.ReplaceAll(randomID.String(), "-", ""))
	return queueName, err
}
