package transport

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// CreateRmqRPCClient create new RmqRPC
func CreateRmqRPCClient(connectionString string) *RmqRPCClient {
	return &RmqRPCClient{
		connectionString: connectionString,
		logger:           log.New(os.Stdout, "", log.LstdFlags),
	}
}

// RmqRPCClient implementation
type RmqRPCClient struct {
	connectionString string
	logger           *log.Logger
}

// SetLogger set custome logger
func (c *RmqRPCClient) SetLogger(logger *log.Logger) RPCClient {
	c.logger = logger
	return c
}

// Call remote function
func (c *RmqRPCClient) Call(functionName string, inputs ...interface{}) (output interface{}, err error) {
	replyTo, err := rmqRpcGenerateReplyQueueName(functionName)
	if err != nil {
		return output, err
	}
	// create connection and channel
	conn, ch, err := rmqCreateConnectionAndChannel(c.connectionString)
	if err != nil {
		return output, err
	}
	defer conn.Close()
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
			envelopedOutput, parseError := NewEnvelopedMessageFromJSON(rmqMessage.Body)
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
			waitReply <- true
			return
		}
	}()
	// send message
	err = rmqRpcCall(ch, functionName, replyTo, inputs)
	if err != nil {
		waitReply <- true
	}
	// return
	<-waitReply
	ch.QueueDelete(replyTo, false, false, true)
	return output, err
}
