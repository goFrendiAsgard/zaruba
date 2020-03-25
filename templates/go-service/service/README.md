# servicename

One Paragraph of project description goes here

# Project Structure

```
servicename
├── README.md
├── communication/                 # (shared-lib)    Any code related to communication protocols (RPC, Pubsub, etc).
├── env/                           # (shared-lib)    Any code related to environment manipulation
├── config/                        # (customizeable) Your service cofiguration.
├── context/                       # (customizeable) Wrapper for config and global states.
├── httphandlers/                  # (customizeable) Functions and factories to handle HTTP requests.
├── pubsubhandlers/                # (customizeable) Functions and factories to handle Pub-sub messages.
├── rpchandlers/                   # (customizeable) Functions and factories to handle RPC requests.
├── registerHttpHandlers.go        # (customizeable) Function to register HTTP handlers.
├── registerPubSubHandlers.go      # (customizeable) Function to register pubsub handlers. 
├── registerRpcHandlers.go         # (customizeable) Function to register RPC handlers. 
├── main.go                        # (customizaable) Service entry point
├── go.mod
├── go.sum
├── Dockerfile
├── Makefile
└── zaruba.config.yaml
```

# Convention

## Reusing instead of initializing components

If your components is going to be used over and over, you should put it in `main.go`. By default, `main.go` contains some components (`context`, `router`, `pubSub`, `rpc`):

```go
context := context.NewContext()
config := context.Config
rmqConnectionString := config.DefaultRmq.CreateConnectionString()
logger := config.Logger

router := gin.Default()
pubSub := communication.NewRmqPubSub(rmqConnectionString).SetLogger(logger)
rpc := communication.NewRmqRPC(rmqConnectionString).SetLogger(logger)
```

In some cases, you might need second pubSub connection, redis, or even database-connection. Feel free to implement and add more components as you need.

## Component naming

Component naming should represent usage/purpose, not technology. For example, `externalPubSub` is a better name than `rabbitmqPubSub` or `natsPubSub`. We believe it is a better practice since technologies can be easily replaced, while architecture might not be affected.

Component naming should be consistent within the application or among the applications if possible.

## Swapping components

It is recommended to create component interfaces, so that you can swap components without affecting business logic.

For example, since `RmqRPC` and `SimpleRPC` have the same interface, you can easily swap `rpc` component with HTTP instead of rabbitMq:

```go
// rpc := communication.NewRmqRPC(rmqConnectionString).SetLogger(logger)
rpc := communication.NewSimpleRPC(router, config.ServiceURLMap).SetLogger(logger)
```

## Injecting components

Rather than using `constructor`, we recommend you to use `closure` (factories and function parameters) instead.

For example, you need `context` and `router` to serve "liveness". Thus you pass both component to `registerHTTPHandlers`:

```go
// main.go
registerHTTPHandlers(context, router)
```

In `registerHTTPHandlers`, we use `router` to link `/liveness` URL with corresponding HTTP handler. Since the handler itself only has `*gin.Context` input parameter, we need to wrap the handler inside `CreateLivenessHandler` factory.

```go
// registerHttpHandlers.go
func registerHTTPHandlers(context *context.Context, router *gin.Engine) {
    router.GET("/liveness", httphandlers.CreateLivenessHandler(context))
}
```

`CreateLivenessHandler` return a handler function and pass `context` to it.

```go
// httphandlers/livenesHandler.go
func CreateLivenessHandler(context *context.Context) (handler func(c *gin.Context)) {
	return func(c *gin.Context) {
		// get http status
		httpCode := http.StatusOK
		if !context.Status.IsAlive {
			httpCode = http.StatusInternalServerError
		}
		// send response
		c.JSON(httpCode, gin.H{
			"servicename": context.Config.ServiceName,
			"isAlive":     context.Status.IsAlive,
		})
	}
}


```

## Handlers

You should put your handlers in either `httphandlers`, `pubsubhandlers`, or `rpchandlers` depending on the message/request you want to handle. 

The handle could be a simple function or a factory (a function that return another function). You should use factory if your handler need external components (`context`, `pubSub`, `rpc`, `router`, etc).
