package transport

import (
	"errors"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// CreateRmqRPCClient create new RmqRPC
func CreateRmqRPCClient(logger *log.Logger, connection *amqp.Connection) *RmqRPCClient {
	return &RmqRPCClient{
		connection: connection,
		logger:     logger,
	}
}

// RmqRPCClient implementation
type RmqRPCClient struct {
	connection *amqp.Connection
	logger     *log.Logger
}

// Call remote function
func (c *RmqRPCClient) Call(functionName string, inputs ...interface{}) (output interface{}, err error) {
	replyTo, err := rmqRPCGenerateReplyQueueName(functionName)
	if err != nil {
		return output, err
	}
	// create connection and channel
	ch, err := c.connection.Channel()
	if err != nil {
		return output, err
	}
	defer ch.Close()
	// consume
	rmqMessages, err := rmqConsume(ch, replyTo)
	if err != nil {
		ch.QueueDelete(replyTo, false, false, true)
		return output, err
	}
	// create waiting channel
	waitReply := make(chan bool)
	go func() {
		for rmqMessage := range rmqMessages {
			envelopedOutput, parseError := CreateEnvelopedMessageFromJSON(rmqMessage.Body)
			if parseError != nil {
				err = parseError
				waitReply <- true
				return
			}
			ok := true
			output, ok = envelopedOutput.Message["output"]
			if !ok {
				errorMessage := fmt.Sprintf("output not found in %#v", envelopedOutput.Message)
				err = errors.New(errorMessage)
			}
			c.logger.Printf("[INFO RmqRPCClient] Get Reply %s %#v: %#v", functionName, inputs, output)
			waitReply <- true
			return
		}
	}()
	// send message
	c.logger.Printf("[INFO RmqRPCClient] Call %s %#v", functionName, inputs)
	err = rmqRPCCall(ch, functionName, replyTo, inputs)
	if err != nil {
		waitReply <- true
	}
	// return
	<-waitReply
	ch.QueueDelete(replyTo, false, false, true)
	return output, err
}
