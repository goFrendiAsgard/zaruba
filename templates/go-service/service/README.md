# servicename

One Paragraph of project description goes here

# Start

```sh
export DEFAULT_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export HTTP_PORT=3000
# ... other-env-setting
go build && ./servicename
```

# Test

```sh
export TEST_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export TEST_HTTP_PORT=3000
# ... other-env-setting
go test ./...
```

# Assumptions

This application was built with the following assumptions:

* Explicit is better than implicit. We try to make everything as accessible, as readable, and as replace-able as possible. 
* We don't try to hide anything using any magic. No reflection, automatic dependency resolver, etc.
* We use `rabbit-mq` as inter-service communication protocol.
	- Pub/Sub
	- RPC Call

>__NOTE__ You can implement your own communication protocol (e.g: nats, kafka, etc)

# Project Structure 

```
.
├── Dockerfile
├── Makefile
├── README.md
├── app
├── component
│   ├── defaultcomponent
│   │   └── index.go
│   ├── example
│   │   ├── helpers.go
│   │   ├── helpers_test.go
│   │   ├── index.go
│   │   ├── services.go
│   │   └── services_test.go
│   └── monitoring
│       ├── helpers.go
│       ├── helpers_test.go
│       └── index.go
├── config
│   ├── config.go
│   └── helper.go
├── core
│   ├── interfaces.go
│   └── mainApp.go
├── go.mod
├── go.sum
├── main.go
├── transport
│   ├── envelopedMessage.go
│   ├── helpers.go
│   ├── interfaces.go
│   ├── message.go
│   ├── rmqPublisher.go
│   ├── rmqRpcClient.go
│   ├── rmqRpcServer.go
│   └── rmqSubscriber.go
└── servicename.zaruba.yaml
```

# Bootstrap

App initialization contains of 4 steps:

* __App component definitions:__ In this step we define all components that will be used by our app or other components (e.g: logger, config, router, publisher, subscriber, rpc-server, and rpc-client). The definition should be explicit so that other developers can easily understand which components depend on which component.
* __App creation:__ In this step, we initiate our app.
* __App setup:__ In this step, we setup our app, by registering your components. To register your component, you need to provide a function with no parameter and return nothing. You can surely wrap the function into object method or wrapper-function. If you need to inject dependencies, you can use [closure](https://en.wikipedia.org/wiki/Closure_(computer_programming)) (more about this latter).
* __App Execution:__ Finally, we run our app.

Below is how everything looks like in `main.go`:

```go

func main() {

	// app component definitions
	logger := log.New(os.Stdout, "", log.LstdFlags)
	config := config.CreateConfig()
	logger.Println("CONFIG:", config.ToString())
	router := gin.Default()
	defaultRmqConnection, err := amqp.Dial(config.DefaultRmqConnectionString)
	if err != nil {
		logger.Fatal("[RmqConnection]", err)
	}
	rpcServer := transport.CreateRmqRPCServer(logger, defaultRmqConnection)
	rpcClient := transport.CreateRmqRPCClient(logger, defaultRmqConnection)
	subscriber := transport.CreateRmqSubscriber(logger, defaultRmqConnection)
	publisher := transport.CreateRmqPublisher(logger, defaultRmqConnection)

	// app creation
	app := core.CreateMainApp(
		logger,
		router,
		[]transport.Subscriber{subscriber},
		[]transport.RPCServer{rpcServer},
		config.HTTPPort,
	)

	// app setup
	app.Setup([]core.SetupComponent{
		defaultcomponent.CreateSetup(config, router), // setup landingPage
		monitoring.CreateSetup(config, app, router),  // setup monitoring
		example.CreateComponent(
			config, router, publisher, subscriber, rpcServer, rpcClient,
		).Setup, // setup example
	})

	// app execution
	app.Run()

}
```

> __NOTE:__ Treat our app bootstrap as your own. You are free to define new components, replace one component with another ones, or wire the components differently.

# Components

All components should be location on `component` directory. To expose a component and put them on `app setup`, you should provide a struct with `Setup` function receiver.

For example, our `monitoring` component need `config`, `app`, and `router`. Thus we utilize closure to inject those components:

```go
// CreateSetup factory to create SetupComponent
func CreateSetup(config *config.Config, app core.App, router *gin.Engine) core.SetupComponent {
	return func() {
		serviceName := config.ServiceName

		router.GET("/liveness", func(c *gin.Context) {
			liveness := app.Liveness()
			// send response
			c.JSON(getHTTPCodeByStatus(liveness), gin.H{
				"service_name": serviceName,
				"is_alive":     liveness,
			})
		})

		router.GET("/readiness", func(c *gin.Context) {
			readiness := app.Readiness()
			// send response
			c.JSON(getHTTPCodeByStatus(readiness), gin.H{
				"service_name": serviceName,
				"is_ready":     readiness,
			})
		})

	}
}
```

The returned function is then used on `app setup`:

```go
app.Setup([]core.SetupComponent{
	// ...
	monitoring.CreateSetup(config, app, router),  // setup monitoring
	// ...
})
```

We has already provide an `example` component that utilize OOP with some possible use-cases (including RPC and pub/sub communication). Feel free to explore.

# Component Convention

Typically a component should contains at least one file named `index.go`. Although you can write anything inside `index.go`, it is better to just put `router`, `rpc-server`, and `subscriber` handler in it.

Make your components as loosly coupled as possible, as it is going to help you in case of you want to split the app into several smaller apps.

If you need some business logic, please put them inside `service` directory or `service.go`.

If your components are depending on each other, please put the dependency in our `main.go` bootstrap, don't import the component directly as it will make your components tightly coupled to each other.

If you need database connection or cache mechanism, please also define them in our `main.go`.