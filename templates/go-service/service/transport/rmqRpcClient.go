package transport

import (
	"errors"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// CreateRmqRPCClient create new RmqRPC
func CreateRmqRPCClient(logger *log.Logger, connection *amqp.Connection, eventMap *RmqEventMap) *RmqRPCClient {
	return &RmqRPCClient{
		connection: connection,
		logger:     logger,
		eventMap:   eventMap,
	}
}

// RmqRPCClient implementation
type RmqRPCClient struct {
	connection *amqp.Connection
	logger     *log.Logger
	eventMap   *RmqEventMap
}

// Call remote function
func (c *RmqRPCClient) Call(functionName string, inputs ...interface{}) (output interface{}, err error) {
	queueName := c.eventMap.GetQueueName(functionName)
	replyTo, err := rmqRPCGenerateReplyQueueName(queueName)
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
	defer ch.QueueDelete(replyTo, false, false, true)
	if err != nil {
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
				err = c.getErrorFromEnvelopedOutput(envelopedOutput)
			}
			c.logger.Printf("[INFO RmqRPCClient] Get Reply %s %#v: %#v", functionName, inputs, output)
			waitReply <- true
			return
		}
	}()
	// send message
	exchangeName := c.eventMap.GetExchangeName(functionName)
	c.logger.Printf("[INFO RmqRPCClient] Call %s %#v", exchangeName, inputs)
	err = rmqRPCCall(ch, exchangeName, replyTo, inputs)
	if err != nil {
		waitReply <- true
	}
	<-waitReply
	return output, err
}

func (c *RmqRPCClient) getErrorFromEnvelopedOutput(envelopedOutput *EnvelopedMessage) error {
	errorMessage := fmt.Sprintf("output not found in %#v", envelopedOutput.Message)
	return errors.New(errorMessage)
}
