package transport

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// NewSimpleRPCClient create new SimpleRPC
func NewSimpleRPCClient(serverAddress string) *SimpleRPCClient {
	return &SimpleRPCClient{
		serverAddress: serverAddress,
		logger:        log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SimpleRPCClient implementation
type SimpleRPCClient struct {
	serverAddress string
	logger        *log.Logger
}

// SetLogger set custome logger
func (c *SimpleRPCClient) SetLogger(logger *log.Logger) RPCClient {
	c.logger = logger
	return c
}

// Call remote function
func (c *SimpleRPCClient) Call(functionName string, inputs ...interface{}) (output interface{}, err error) {
	// prepare JSON message to be send
	envelopedMsgJSON, err := rpcInputsToJSON(inputs)
	if err != nil {
		return output, err
	}
	// prepare request
	remoteAddr := fmt.Sprintf("%s/api/%s", c.serverAddress, functionName)
	req, err := http.NewRequest("POST", remoteAddr, bytes.NewBuffer(envelopedMsgJSON))
	if err != nil {
		return output, err
	}
	req.Header.Set("Content-Type", "application/json")
	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return output, err
	}
	// handle reponse error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errorMessage := fmt.Sprintf("Response HTTP Status: %d %s", resp.StatusCode, resp.Status)
		return output, errors.New(errorMessage)
	}
	// get response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}
	// enveloped Output
	envelopedOutput, err := NewEnvelopedMessageFromJSON(body)
	if err != nil {
		return output, err
	}
	return envelopedOutput.Message.GetInterface("output")
}
