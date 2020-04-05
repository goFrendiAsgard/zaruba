package transport

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// NewSimpleRPCServer create new SimpleRPC
func NewSimpleRPCServer(engine *gin.Engine) *SimpleRPCServer {
	return &SimpleRPCServer{
		engine:   engine,
		handlers: map[string]RPCHandler{},
		logger:   log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SimpleRPCServer implementation
type SimpleRPCServer struct {
	engine   *gin.Engine
	handlers map[string]RPCHandler
	logger   *log.Logger
}

// SetLogger set custome logger
func (s *SimpleRPCServer) SetLogger(logger *log.Logger) RPCServer {
	s.logger = logger
	return s
}

// RegisterHandler register servicemap for call
func (s *SimpleRPCServer) RegisterHandler(functionName string, handler RPCHandler) RPCServer {
	s.handlers[functionName] = handler
	return s
}

// Serve RPC request
func (s *SimpleRPCServer) Serve() {
	for functionName, handler := range s.handlers {
		endPoint := fmt.Sprintf("/api/%s", functionName)
		s.engine.POST(endPoint, func(c *gin.Context) {
			envelopedInput := &EnvelopedMessage{}
			err := c.BindJSON(envelopedInput)
			if err != nil {
				c.JSON(http.StatusBadRequest, rpcCreateEnvelopedErrorMessage(envelopedInput, err))
				return
			}
			inputs, err := envelopedInput.Message.GetInterfaceArray("inputs")
			if err != nil {
				c.JSON(http.StatusInternalServerError, rpcCreateEnvelopedErrorMessage(envelopedInput, err))
				return
			}
			output, err := handler(inputs...)
			if err != nil {
				c.JSON(http.StatusInternalServerError, rpcCreateEnvelopedErrorMessage(envelopedInput, err))
				return
			}
			c.JSON(http.StatusOK, rpcCreateEnvelopedOutputMessage(envelopedInput, output))
		})
	}
}
