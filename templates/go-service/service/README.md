# servicename

One Paragraph of project description goes here

# How to Start

```sh
go build && ./servicename
```

# Project Structure

```
servicename
├── Dockerfile
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── main.go 							# (Customizable) Application entry point
├── bootstrap                           # (Customizable) Code to set-up and run components
│   ├── run.go                          # (Customizable) Code to run components
│   ├── setting.go                      # (Customizable) Global Dependency Definition
│   └── setup.go                        # (Customizable) Code to set-up components
├── components                          # (Customizable) You should put your components here
│   └── monitoring
│       └── livenessAndReadiness.go
├── config                              # (Customizable) Config related Code
│   ├── config.go
│   └── helper.go
├── context                             # (Customizable) Wrapper of Config, status, and other properties
│   └── context.go
├── example                             # (Example) Example. Delete this if not necessary
│   ├── bootstrap.go                    # (Example) Code to link local components with the global ones
│   └── components
│       └── greeting                    # (Example) Component example
│           ├── eventHandler.go         # (Example) Event handler functions and factories
│           ├── httpController.go       # (Example) HTTP handler functions and factories
│           ├── rpcHandler.go           # (Example) RPC handler functions and factories
│           └── serviceGreeting.go      # (Example) Service (business logic), avoid side-effects here
├── transport                           # (Shared-Lib) Low-level transport implementation
│   ├── envelopedMessage.go
│   ├── helpers.go
│   ├── interface.go
│   ├── message.go
│   ├── rmqPublisher.go
│   ├── rmqRpcClient.go
│   ├── rmqRpcServer.go
│   ├── rmqSubscriber.go
│   ├── simpleRpcClient.go
│   └── simpleRpcServer.go
└── zaruba.config.yaml
```

# Component

We try to make components as loosely coupled as possible. Business logic should not tightly bound to transport/communication layer. To serve this purpose we divide component into several parts:

* __Service (business logic)__
* __HTTP Handler__
* __RPC Handler__
* __Event Handler__

## Service

Service is the highest level abstraction of your business logic. It should be technology and transport agnostic.

For example, to calculate average numbers of user visiting your website, you should only care about the logic:

```go
func GetAverageVisit(visitPerDays []int) float64 {
	totalVisit := 0
	for visitInADay := range(visitPerDays) {
		totalVisit += visitInADay
	}
	return totalVisit/len(vistPerDays)
}
```

You don't need to care about how to get the data (e.g: message-queue, database, etc)

By convention you should put your services in `service` prefixed files.

## HTTP Handler

HTTP Handlers are functions or factories to handle HTTP request. By convention, you should put your HTTP Handlers in `httpController.go`

Here is how a simple HTTP Handler looks like:

```go
func SayHi (c *gin.Context) { 
	c.String(http.StatusOK, "Hi") 
}
```

In case of you need something more than `gin.context`, you should write factory instead:

```go
func CreateRegisterHandler(publisher transport.Publisher, registerEvent string) func(c *gin.Context) {
	return func(c *gin.Context) {
		registrant := c.Query("registrant")
		err := publisher.Publish(registerEvent, transport.Message{"registrant": registrant})
		if err != nil {
			c.String(http.StatusInternalServerError, "Registration failed, please try again")
		}
		c.String(http.StatusOK, "Registration Submitted")
	}
}
```

## RPC Handler

RPC Handlers are functions or factories to handle RPC Request. By convention, you should put your RPC Controller in `rpcHandler.go`.

RPC Handler function signature is as follow:

```go
func RPCHandler(inputs ...[]interface{]}) (output interface{}, err error) {
	return output, err
}
```

Just like HTTP Handlers, you can wrap RPC Handler functions into factories.

## Event Handler

Event Handlers are functions or factories to handle Event Request. By convention, you should put your Event Controller in `eventHandler.go`.

Event Handler function signature is as follow:

```go
func EventHandler(message transport.Message) (err error) {
	return err
}
```

Just like HTTP Handlers, you can wrap Event Handler functions into factories.

## Bootstrap

Bootstrap is a function to connect your component to the application. You can think of it as a place to wire up component's dependencies.

```go
func SetUp(s *bootstrap.Setting) {
	// TODO: setup your HTTP handlers etc here...
}
```

The `bootstrap.Setting` is free to modified to fit your need.  By default, it contains `gin.Engine`, `context.Context`, and bunch of interfaces:

```go
type Setting struct {
	Ctx        *context.Context
	Router     *gin.Engine
	Publishers struct {
		Main transport.Publisher
	}
	Subscribers struct {
		Main transport.Subscriber
	}
	RPCServers struct {
		Main      transport.RPCServer
		Secondary transport.RPCServer
	}
	RPCClients struct {
		MainLoopBack      transport.RPCClient
		SecondaryLoopBack transport.RPCClient
	}
}
```

We don't try to give you any hidden magic, so you should register your bootstrap in `main.go`.

For example we register example bootstrap as follow:

```go
func main() {

	s := bootstrap.NewSetting() // <-- "container" for our dependency injection

	s.Ctx = context.NewContext()
	s.Router = gin.Default()

	logger := s.Ctx.Config.Logger
	rmqConnectionString := s.Ctx.Config.RmqConnectionString

	s.Publishers.Main = transport.NewRmqPublisher(rmqConnectionString).SetLogger(logger)
	s.Subscribers.Main = transport.NewRmqSubscriber(rmqConnectionString).SetLogger(logger)

	s.RPCServers.Main = transport.NewRmqRPCServer(rmqConnectionString).SetLogger(logger)
	s.RPCClients.MainLoopBack = transport.NewRmqRPCClient(rmqConnectionString).SetLogger(logger)

	s.RPCServers.Secondary = transport.NewSimpleRPCServer(s.Router).SetLogger(logger)
	s.RPCClients.SecondaryLoopBack = transport.NewSimpleRPCClient(s.Ctx.Config.LocalServiceAddress).SetLogger(logger)

	// TODO: remove the example, and implement your own
	example.SetUp(s) // <-- No magic, just function call

	bootstrap.Setup(s)
	bootstrap.Run(s)
}
```

Please visit `example` folder as it already show you almost every possible use-case.
