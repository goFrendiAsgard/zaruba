# servicename

One Paragraph of project description goes here

# How to Start

```sh
export GLOBAL_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export LOCAL_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export HTTP_PORT=3000
... <other-env-setting>
go build && ./servicename
```

# Opinions

This application was built with the following opinions:

* You use message broker as main communication protocol.
* Even RPC call should also use message broker.
* The message broker you use is rabbitmq.
* Explicit is better than implicit. Everything should be accessible and readable without any hidden-magic.
* Your app has two contexts, global and local-domain (eventhough you can use one context only)

```
	[Other Domain App] --- [Global Message Broker] --- [Other Domain App]
	                        		  |
									  |
								 [Your App]
								      |
									  |
[Other Same Domain App] ---	[Local Message Broker] --- [Other Same Domain App]
```

>__NOTE__ You can implement your own communication protocol (e.g: nats, kafka, etc)

# Project Structure 

```
.
├── Dockerfile
├── Makefile
├── README.md
├── component                 # your components
│   ├── example
│   │   ├── component.go
│   │   ├── helpers.go
│   │   └── services.go
│   └── monitoring
│       ├── component.go
│       └── helpers.go
├── config
│   ├── config.go
│   └── helper.go
├── core
│   ├── application.go        # core application
│   └── interface.go          # application interface
├── go.mod
├── go.sum
├── main.go                   # container and application runner
├── transport
│   ├── envelopedMessage.go
│   ├── helpers.go
│   ├── interface.go
│   ├── message.go
│   ├── rmqPublisher.go
│   ├── rmqRpcClient.go
│   ├── rmqRpcServer.go
│   └── rmqSubscriber.go
└── zaruba.config.yaml
```

# Inversion of Control

Eventhough we love the idea of dependency injection. We don't like these two things:

* Dependency-injection usually use `injector`. Those injector sometimes use reflection or even declarative language like `XML` or `YAML`. Thus, the flow of the program become less obvious. You need to read tons of documentation or dive into the framework code just to know what's going on.

* Most Dependency-injection framework heavily assume developer to only works with OOP paradigm. This is un-necessary. In fact we can even use `factory pattern` to inject things into our components.

By stripping those two things from our list, we get an explicit and less-assumption architecture.

The drawback is, this make your code more-verbose (which is actually a good thing because everything is under your control).

Below is the implementation of `container` in `main.go`:

```go
func main() {
	// create config and app
	config := config.CreateConfig()
	fmt.Println(config)
	app := core.CreateApplication(
		config.HTTPPort,
		config.GlobalRmqConnectionString,
		config.LocalRmqConnectionString,
	)
	// setup components
	app.Setup([]core.SetupComponent{
		monitoring.CreateSetup(app, config),        // setup monitoring
		example.CreateComponent(app, config).Setup, // setup example
	})
	// run
	app.Run()
}
```

## Swapping Components

Let's say you want to implement your own `Application` component. In order to do that, you need to modify the corresponding lines:

```go
/*
app := core.CreateApplication(
	config.HTTPPort,
	config.GlobalRmqConnectionString,
	config.LocalRmqConnectionString,
)
*/
app := MockApplication()
```

Of course, you have to make sure that `MockApplication` is comply with `App` interface. But aside from that, there is no magic here, it is just your day-to-day simple go code.

## Injecting Setup to App

Our `app` has a special method named `Setup`. You can use `Setup` to inject functions into app. This approach gives you a lot of freedom. For example you might use:

* constructor-based injection
* builder-based injection
* factory function
* anything that return a function

Let's focus on this part:

```go
app.Setup([]core.SetupComponent{
	monitoring.CreateSetup(app, config),        // setup monitoring
	example.CreateComponent(app, config).Setup, // setup example
})
```

`monitoring.CreateSetup` is a factory function that produce anonymous function to change `app`'s behavior. Let's see how it looks like:

```go
func CreateSetup(app core.App, config *config.Config) core.SetupComponent {
	return func() {
		serviceName := config.ServiceName
		r := app.Router()

		r.GET("/liveness", func(c *gin.Context) {
			liveness := app.Liveness()
			// send response
			c.JSON(getHTTPCodeByStatus(liveness), gin.H{
				"service_name": serviceName,
				"is_alive":     liveness,
			})
		})

		r.GET("/readiness", func(c *gin.Context) {
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

As you see, we have just inject `app` and `config` into the anonymous function. You might find that the approach above is similar to `decorator`.

On the other hand, `example.CreateComponent` is a class constructor:

```go
// CreateComponent create component
func CreateComponent(app core.App, config *config.Config) *Component {
	return &Component{
		names:  []string{},
		app:    app,
		config: config,
	}
}

// Component implementation
type Component struct {
	names  []string
	app    core.App
	config *config.Config
}

// Setup component
func (comp *Component) Setup() {
	r := comp.app.Router()
	rpcServer := comp.app.GlobalRPCServer()
	subscriber := comp.app.GlobalSubscriber()

	// Simple HTTP Handler
	r.Any("/", func(c *gin.Context) { c.String(http.StatusOK, "servicename") })

	// More complex HTTP Handler, with side-effect
	r.GET("/toggle-readiness", func(c *gin.Context) {
		comp.app.SetReadiness(!comp.app.Readiness())
		c.String(http.StatusOK, fmt.Sprintf("Readiness: %#v", comp.app.Readiness()))
	})

	// Use the same HTTP Handler for multiple URLS
	r.GET("/hello", comp.handleHTTPHello)
	r.POST("/hello", comp.handleHTTPHello)
	r.GET("/hello/:name", comp.handleHTTPHello)

	// Use HTTP Handler that take state from component
	r.GET("/hello-all", comp.handleHTTPHelloAll)

	// Trigger RPC Call
	r.GET("/hello-rpc", comp.handleHTTPHelloRPC)
	r.GET("/hello-rpc/:name", comp.handleHTTPHelloRPC)
	r.POST("/hello-rpc", comp.handleHTTPHelloRPC)

	// Trigger Publisher Call
	r.GET("/hello-pub", comp.handleHTTPHelloPub)
	r.GET("/hello-pub/:name", comp.handleHTTPHelloPub)
	r.POST("/hello-pub", comp.handleHTTPHelloPub)

	// Serve RPC
	rpcServer.RegisterHandler("helloRPC", comp.handleRPCHello)

	// Event
	subscriber.RegisterHandler("helloEvent", comp.handleEventHello)

}

// etc ...
```

This approach is better if you want to encapsulate state and share it among your handlers. For example, let's take a look on `comp.handleHTTPHelloAll` and `comp.handleEventHello`. Both method are sharing `comp.names` property as state.

`comp.handleHTTPHelloAll` simply use `comp.names` to show HTTP response to user:

```go
func (comp *Component) handleHTTPHelloAll(c *gin.Context) {
	c.String(http.StatusOK, GreetEveryone(comp.names))
}
```

On the other hand, `comp.handleEventHello` listen to `helloEvent` and add user's name into `comp.names`:

```go
func (comp *Component) handleEventHello(msg transport.Message) (err error) {
	name, err := msg.GetString("name")
	if err != nil {
		return err
	}
	comp.names = append(comp.names, name)
	return err
}
```

# Publish and Handle Event

Pub-sub is an asyncrhonous communication pattern. To put it simple, you just fire an event without waiting for response. This pattern is quite common for long-running process, such as Machine-learning model training (i.e: You don't want to wait for days until your training process complete. Instead, you want to receive email once the process finished or failed).

By default, you can use `app.GlobalPublisher` and `app.LocalPublisher` to publish message. Here is an example:

```go
pub := app.LocalPublisher()
err := pub.Publish("helloEvent", transport.Message{"name": name})
```

To handle the event, you can use `app.GlobalSubscriber` and `app.LocalSubscriber`. Here is an example:

```go
sub := app.LocalSubscriber()
sub.RegisterHandler("helloEvent", func (msg transport.Message) (err error) {
		name, err := msg.GetString("name")
		if err != nil {
			return err
		}
		comp.names = append(comp.names, name)
		return err
	}
)
```

# Call and Handle RPC

RPC is stands for Remote Procecure Call. It let you call procedure/function from other application. Unlike pub-sub, you usually use RPC if you need to wait for the result.

You can use `app.GlobalRPCClient` and `app.LocalRPCClient` to send RPC Call:

```go
client := app.LocalRPCClient()
result, err := client.Call("helloRPC", name)
fmt.Println(result.(string))
```

To handle RPC, you can use `app.GlobalRPCServer` and `app.LocalRPCServer`:

```go
server := app.LocalRPCServer()
server.RegisterHandler("helloRPC", func(inputs ...interface{}) (greeting interface{}, err error) {
	if len(inputs) == 0 {
		return greeting, errors.New("Message accepted but input is invalid")
	}
	name, success := inputs[0].(string)
	if !success {
		errorMessage := fmt.Sprintf("Cannot convert %#v to string", inputs[0])
		return greeting, errors.New(errorMessage)
	}
	return Greet(name), err
})
```

# Naming Convention

A component usually contains several files:

* `component.go`: This file contains code to produce your `setup`.
* `helpers.go`: This file contains any general-purpose code that are used by `component.go` or `services.go`.
* `services.go`: This file contains your main business logic. It should not care about how it is called.