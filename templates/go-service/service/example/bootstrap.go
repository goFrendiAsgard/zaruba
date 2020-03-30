package example

import (
	"fmt"
	"net/http"

	"app/bootstrap"
	"app/example/components/greeting"

	"github.com/gin-gonic/gin"
)

// SetUp Example
func SetUp(s *bootstrap.Setting) {

	// HTTP EXAMPLE =====================================================================================================

	// Example: Simple HTTP Handler
	s.Router.Any("/", func(c *gin.Context) { c.String(http.StatusOK, "servicename") })

	// Example: More complex HTTP Handler, with side-effect
	s.Router.GET("/toggle-readiness", func(c *gin.Context) {
		s.Ctx.Status.IsReady = !s.Ctx.Status.IsReady
		c.String(http.StatusOK, fmt.Sprintf("Readiness %#v", s.Ctx.Status.IsReady))
	})

	// Example: Use HTTP Handler from greeting component
	s.Router.GET("/hello", greeting.GreetHTTPController)
	s.Router.POST("/hello", greeting.GreetHTTPController)
	s.Router.GET("/hello/:name", greeting.GreetHTTPController)

	GreetEveryoneHTTPController := greeting.CreateGreetEveryoneHTTPController(s.Ctx)
	s.Router.GET("/hello-all", GreetEveryoneHTTPController)
	s.Router.POST("/hello-all", GreetEveryoneHTTPController)
	s.Router.GET("/hello-all/:name", GreetEveryoneHTTPController)

	// RPC EXAMPLE ======================================================================================================

	// Example: RPC Handler  (Main)
	s.RPCServers.Main.RegisterHandler("greetRPC", greeting.GreetRPCController)

	// Example: HTTP handler to trigger RPC
	GreetRPCHTTPController := greeting.CreateGreetRPCHTTPController(s.RPCClients.MainLoopBack, "greetRPC")
	s.Router.GET("/hello-rpc", GreetRPCHTTPController)
	s.Router.POST("/hello-rpc", GreetRPCHTTPController)
	s.Router.GET("/hello-rpc/:name", GreetRPCHTTPController)

	// RPC EXAMPLE ======================================================================================================

	// Example: RPC (Secondary)
	s.RPCServers.Secondary.RegisterHandler("greetRPC", greeting.GreetRPCController)

	// Example: HTTP handler to trigger RPC
	SecondaryGreetRPCHTTPController := greeting.CreateGreetRPCHTTPController(s.RPCClients.SecondaryLoopBack, "greetRPC")
	s.Router.GET("/hello-secondary-rpc", SecondaryGreetRPCHTTPController)
	s.Router.POST("/hello-secondary-rpc", SecondaryGreetRPCHTTPController)
	s.Router.GET("/hello-secondary-rpc/:name", SecondaryGreetRPCHTTPController)

	// PUB SUB EXAMPLE ==================================================================================================

	// Example: Event Handler
	registerPersonEvenHandler := greeting.CreateRegisterPersonHandler(s.Ctx)
	s.Subscribers.Main.RegisterHandler("personRegistered", registerPersonEvenHandler)

	// Example: HTTP handler to publish event
	GreetPublishHTTPController := greeting.CreateGreetPublishHTTPController(s.Publishers.Main, "personRegistered")
	s.Router.GET("/hello-pub", GreetPublishHTTPController)
	s.Router.POST("/hello-pub", GreetPublishHTTPController)
	s.Router.GET("/hello-pub/:name", GreetPublishHTTPController)
}
