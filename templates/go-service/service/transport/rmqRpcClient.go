package transport

import (
	"errors"
	"fmt"
	"log"
	"time"

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
	rmqMessages, err := rmqConsume(ch, replyTo, true, true, false, false, true, nil)
	if err != nil {
		c.logger.Printf("[INFO RmqRPCClient] Cannot Consume from %s: %#v", replyTo, err)
		return output, err
	}
	return c.call(ch, rmqMessages, functionName, replyTo, inputs)
}

func (c *RmqRPCClient) call(ch *amqp.Channel, rmqMessages <-chan amqp.Delivery, functionName, replyTo string, inputs []interface{}) (output interface{}, err error) {
	// consume from reply
	waitReply := make(chan bool)
	go func() {
		for rmqMessage := range rmqMessages {
			envelopedOutput, parseError := CreateEnvelopedMessageFromJSON(rmqMessage.Body)
			if parseError != nil {
				err = parseError
				c.logger.Printf("[INFO RmqRPCClient] Get Error %s %#v: %#v", functionName, inputs, err)
				waitReply <- true
				return
			}
			ok := true
			output, ok = envelopedOutput.Message["output"]
			if !ok {
				err = c.getErrorFromEnvelopedOutput(envelopedOutput)
				c.logger.Printf("[INFO RmqRPCClient] Get Error %s %#v: %#v", functionName, inputs, err)
			} else {
				c.logger.Printf("[INFO RmqRPCClient] Get Reply %s %#v: %#v", functionName, inputs, output)
			}
			waitReply <- true
			return
		}
	}()
	// send message
	exchangeName := c.eventMap.GetExchangeName(functionName)
	c.logger.Printf("[INFO RmqRPCClient] Call %s by Send to Exchange %s %#v", functionName, exchangeName, inputs)
	err = rmqRPCCall(ch, exchangeName, replyTo, inputs)
	if err != nil {
		c.logger.Printf("[ERROR RmqRPCClient] Cannot call %s by Sending to Exchange %s %#v: %#v", functionName, exchangeName, inputs, err)
		waitReply <- true
	}
	// timeout
	go func() {
		timeout := c.eventMap.GetRPCTimeout(functionName)
		ticker := time.NewTicker(timeout * time.Millisecond)
		<-ticker.C
		err = fmt.Errorf("Timeout %d ms", timeout)
		c.logger.Printf("[ERROR RmqRPCClient] Get timeout %s %#v: %d ms", functionName, inputs, timeout)
		waitReply <- true
		return
	}()
	// waiting for reply
	<-waitReply
	return output, err
}

func (c *RmqRPCClient) getErrorFromEnvelopedOutput(envelopedOutput *EnvelopedMessage) error {
	errorMessage := fmt.Sprintf("output not found in %#v", envelopedOutput.Message)
	return errors.New(errorMessage)
}
