package communication

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewDefaultSimpleRPC create new SimpleRPC
func NewDefaultSimpleRPC(engine *gin.Engine) *SimpleRPC {
	return NewSimpleRPC(engine, map[string]string{})
}

// NewSimpleRPC create new SimpleRPC
func NewSimpleRPC(engine *gin.Engine, serviceMap map[string]string) *SimpleRPC {
	return &SimpleRPC{
		engine:     engine,
		serviceMap: serviceMap,
		handlers:   map[string]RPCHandler{},
	}
}

// SimpleRPC implementation
type SimpleRPC struct {
	engine     *gin.Engine
	serviceMap map[string]string
	handlers   map[string]RPCHandler
}

// RegisterService register servicemap for call
func (rpc *SimpleRPC) RegisterService(serviceName, serviceURL string) {
	rpc.serviceMap[serviceName] = serviceURL
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
				c.JSON(http.StatusBadRequest, gin.H{"correlation_id": envelopedInput.CorrelationID})
				return
			}
			correlationID := envelopedInput.CorrelationID
			input := envelopedInput.Message
			output, err := handler(input)
			envelopedOutput := EnvelopedMessage{
				CorrelationID: correlationID,
				Message:       output,
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, envelopedOutput)
				return
			}
			c.JSON(http.StatusOK, envelopedOutput)
		})
	}
}

// Call remote function
func (rpc *SimpleRPC) Call(serviceName, functionName string, input Message) (output Message, err error) {
	remoteAddress := fmt.Sprintf("%s/api/%s", rpc.serviceMap[serviceName], functionName)
	envelopedJSON, err := inputMessageToJSON(input)
	if err != nil {
		return output, err
	}
	// create HTTP request
	req, err := http.NewRequest("POST", remoteAddress, bytes.NewBuffer(envelopedJSON))
	if err != nil {
		return output, err
	}
	req.Header.Set("Content-Type", "application/json")
	// prepare HTTP Client and send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()
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
	envelopedOutput := EnvelopedMessage{}
	if err = json.Unmarshal(body, &envelopedOutput); err != nil {
		return output, err
	}
	output = envelopedOutput.Message
	// return
	return output, err
}

func inputMessageToJSON(input Message) (envelopedJSON []byte, err error) {
	correlationID, err := uuid.NewUUID()
	if err != nil {
		return envelopedJSON, err
	}
	envelopedInput := EnvelopedMessage{
		CorrelationID: correlationID.String(),
		Message:       input,
	}
	return json.Marshal(envelopedInput)
}

// EnvelopedMessage Message structure while transporting
type EnvelopedMessage struct {
	CorrelationID string  `json:"correlation_id"`
	Message       Message `json:"message"`
}
