package communication

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// NewSimpleRPC create new SimpleRPC
func NewSimpleRPC(engine *gin.Engine, serviceMap map[string]string) *SimpleRPC {
	return &SimpleRPC{
		engine:     engine,
		serviceMap: serviceMap,
		handlers:   map[string]RPCHandler{},
		logger:     log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SimpleRPC implementation
type SimpleRPC struct {
	engine     *gin.Engine
	serviceMap map[string]string
	handlers   map[string]RPCHandler
	logger     *log.Logger
}

// SetLogger set custome logger
func (rpc *SimpleRPC) SetLogger(logger *log.Logger) *SimpleRPC {
	rpc.logger = logger
	return rpc
}

// RegisterHandler register servicemap for call
func (rpc *SimpleRPC) RegisterHandler(functionName string, handler RPCHandler) {
	rpc.handlers[functionName] = handler
}

// Serve from remote client
func (rpc *SimpleRPC) Serve() {
	for functionName, handler := range rpc.handlers {
		endPoint := fmt.Sprintf("/api/%s", functionName)
		rpc.engine.POST(endPoint, func(c *gin.Context) {
			envelopedInput := EnvelopedMessage{}
			err := c.BindJSON(&envelopedInput)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"correlation_id": envelopedInput.CorrelationID, "error": fmt.Sprintf("%s", err)})
				return
			}
			output, err := handler(envelopedInput.Message)
			envelopedOutput := NewEnvelopedMessageWithCorrelationID(envelopedInput.CorrelationID, output)
			if err != nil {
				envelopedOutput.ErrorMessage = fmt.Sprintf("%s", err)
				c.JSON(http.StatusInternalServerError, envelopedOutput)
				return
			}
			c.JSON(http.StatusOK, envelopedOutput)
		})
	}
}

// Call remote function
func (rpc *SimpleRPC) Call(servicename, functionName string, input Message) (output Message, err error) {
	resp, err := rpc.sendRPCRequest(servicename, functionName, input)
	defer resp.Body.Close()
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
	output = envelopedOutput.Message
	// return
	return output, err
}

func (rpc *SimpleRPC) sendRPCRequest(servicename, functionName string, input Message) (resp *http.Response, err error) {
	req, err := rpc.createRPCRequest(servicename, functionName, input)
	if err != nil {
		return resp, err
	}
	// prepare HTTP Client and send
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (rpc *SimpleRPC) createRPCRequest(servicename, functionName string, input Message) (req *http.Request, err error) {
	envelopedJSON, err := rpc.inputMessageToJSON(input)
	if err != nil {
		return req, err
	}
	// create HTTP request
	remoteAddress := fmt.Sprintf("%s/api/%s", rpc.serviceMap[servicename], functionName)
	req, err = http.NewRequest("POST", remoteAddress, bytes.NewBuffer(envelopedJSON))
	if err != nil {
		return req, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, err

}

func (rpc *SimpleRPC) inputMessageToJSON(input Message) (envelopedJSON []byte, err error) {
	envelopedMessage, err := NewEnvelopedMessage(input)
	if err != nil {
		return envelopedJSON, err
	}
	return envelopedMessage.ToJSON()
}
