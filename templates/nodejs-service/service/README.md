# servicename

One Paragraph of project description goes here

# How to Start

```sh
export GLOBAL_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export LOCAL_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export HTTP_PORT=3000
... <other-env-setting>
npm start
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
[Other Same Domain App] ---    [Local Message Broker] --- [Other Same Domain App]
```

>__NOTE__ You can implement your own communication protocol (e.g: nats, kafka, etc)

# Project Structure 

```
.
├── dist
├── Dockerfile
├── Makefile
├── README.md
├── package-lock.json
├── package.json
├── src
│   ├── components                 # your components
│   │   ├── example
│   │   │   ├── helpers.ts
│   │   │   ├── index.ts
│   │   │   └── service.ts
│   │   └── monitoring
│   │       └── index.ts
│   ├── config
│   │   └── index.ts
│   ├── core
│   │   ├── application.ts         # core application
│   │   ├── index.ts
│   │   └── interfaces.ts          # application interface
│   ├── main.ts                    # container and application runner
│   └── transport
│       ├── envelopedMessage.ts
│       ├── expressMiddleware.ts
│       ├── helpers.ts
│       ├── index.ts
│       ├── interfaces.ts
│       ├── logExpressRoutes.ts
│       ├── rmqPublisher.ts
│       ├── rmqRpcClient.ts
│       ├── rmqRpcServer.ts
│       └── rmqSubscriber.ts
├── tsconfig.json
└── zaruba.config.yaml
```

# Inversion of Control

Eventhough we love the idea of dependency injection. We don't like these two things:

* Dependency-injection usually use `injector`. Those injector sometimes use reflection or even declarative language like `XML` or `YAML`. Thus, the flow of the program become less obvious. You need to read tons of documentation or dive into the framework code just to know what's going on.

* Most Dependency-injection framework heavily assume developer to only works with OOP paradigm. This is un-necessary. In fact we can even use `factory pattern` to inject things into our components.

By stripping those two things from our list, we get an explicit and less-assumption architecture.

The drawback is, this make your code more-verbose (which is actually a good thing because everything is under your control).

Below is the implementation of `container` in `main.ts`:

```typescript
function main() {
    // create config and app
    const config = new Config();
    console.log("CONFIG:", config);
    const app = new Application(
        config.httpPort,
        config.globalRmqConnectionString,
        config.localRmqConnectionString,
    );
    // setup components
    app.setup([
        monitoring.createSetup(app, config),                // setup monitoring
        () => (new example.Component(app, config)).setup(), // setup example
    ]);
    // run
    app.run();
}
```

## Swapping Components

Let's say you want to implement your own `Application` component. In order to do that, you need to modify the corresponding lines:

```ts
/*
const app = new Application(
	config.httpPort,
	config.globalRmqConnectionString,
	config.localRmqConnectionString,
);
*/
const app = MockApplication()
```

Of course, you have to make sure that `MockApplication` is comply with `App` interface. But aside from that, there is no magic here, it is just your day-to-day simple go code.

## Injecting Setup to App

Our `app` has a special method named `Setup`. You can use `Setup` to inject functions into app. This approach gives you a lot of freedom. For example you might use:

* constructor-based injection
* builder-based injection
* factory function
* anything that return a function

Let's focus on this part:

```ts
app.setup([
	monitoring.createSetup(app, config),                // setup monitoring
	() => (new example.Component(app, config)).setup(), // setup example
]);
```

> __NOTE:__ Javascript has somehow-strange behavior related to `this` context. Thus, for `setup example` you see that we wrap the component into an arrow function.

`monitoring.CreateSetup` is a factory function that produce anonymous function to change `app`'s behavior. Let's see how it looks like:

```ts
export function createSetup(app: App, config: Config): () => void {
    return () => {
        const serviceName = config.serviceName;
        const r = app.router();

        r.get("/liveness", (_, res) => {
            const liveness = app.liveness();
            const httpCode = liveness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_alive: liveness,
            });
        });

        r.get("/readiness", (_, res) => {
            const readiness = app.readiness();
            const httpCode = readiness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_alive: readiness,
            });
        });

    }
}
```

As you see, we have just inject `app` and `config` into the anonymous function. You might find that the approach above is similar to `decorator`.

On the other hand, `example.CreateComponent` is a class constructor:

```ts
export class Component {
    private names: string[];
    private app: App;
    private config: Config;

    constructor(app: App, config: Config) {
        this.names = [];
        this.app = app;
        this.config = config;
    }

    setup() {
        const r = this.app.router();
        const rpcServer = this.app.globalRPCServer();
        const subscriber = this.app.globalSubscriber();

        // Simple HTTP Handler
        r.all("/", (_, res) => res.send("servicename"));

        // More complex HTTP Handler, with side-effect
        r.get("/toggle-readiness", (_, res) => {
            this.app.setReadiness(!this.app.readiness());
            const readiness = this.app.readiness();
            res.send(`Readiness: ${readiness}`);
        });

        // Use the same HTTP Handler for multiple URLS
        r.get("/hello", this.handleHTTPHello.bind(this));
        r.get("/hello/:name", this.handleHTTPHello.bind(this));
        r.post("/hello", this.handleHTTPHello.bind(this));

        // Use HTTP Handler that take state from component
        r.get("/hello-all", this.handleHTTPHelloAll.bind(this));

        // Trigger RPC Call
        r.get("/hello-rpc", this.handleHTTPHelloRPC.bind(this));
        r.get("/hello-rpc/:name", this.handleHTTPHelloRPC.bind(this));
        r.post("/hello-rpc", this.handleHTTPHelloRPC.bind(this));

        // Trigger RPC Call
        r.get("/hello-pub", this.handleHTTPHelloPub.bind(this));
        r.get("/hello-pub/:name", this.handleHTTPHelloPub.bind(this));
        r.post("/hello-pub", this.handleHTTPHelloPub.bind(this));

        // Serve RPC
        rpcServer.registerHandler("servicename.helloRPC", this.handleRPCHello.bind(this));

        // Event
        subscriber.registerHandler("servicename.helloEvent", this.handleEventHello.bind(this));

	}

	// etc ...
}
```

This approach is better if you want to encapsulate state and share it among your handlers. For example, let's take a look on `handleHTTPHelloAll` and `handleEventHello`. Both method are sharing `names` property as state.

`handleHTTPHelloAll` simply use `names` to show HTTP response to user:

```ts
handleHTTPHelloAll(req: any, res: any) {
	res.send(greetEveryone(this.names));
}
```

On the other hand, `handleEventHello` listen to `helloEvent` and add user's name into `names`:

```ts
handleEventHello(msg: Message) {
	const { name } = msg;
	this.names.push(name);
}
```

# Publish and Handle Event

Pub-sub is an asyncrhonous communication pattern. To put it simple, you just fire an event without waiting for response. This pattern is quite common for long-running process, such as Machine-learning model training (i.e: You don't want to wait for days until your training process complete. Instead, you want to receive email once the process finished or failed).

By default, you can use `app.globalPublisher` and `app.localPublisher` to publish message. Here is an example:

```ts
const pub = this.app.localPublisher();
try {
	publisher.publish("helloEvent", { name });
} catch(err) {
	// do something
}
```

To handle the event, you can use `app.globalSubscriber` and `app.localSubscriber`. Here is an example:

```ts
sub := app.localSubscriber()
sub.RegisterHandler("helloEvent", (msg: Message) => {
	const { name } = msg;
    this.names.push(name);
)
```

# Call and Handle RPC

RPC is stands for Remote Procecure Call. It let you call procedure/function from other application. Unlike pub-sub, you usually use RPC if you need to wait for the result.

You can use `app.globalRPCClient` and `app.localRPCClient` to send RPC Call:

```ts
const client = app.localRPCClient();
try {
	const result = await client.call("helloRPC", name);
	console.log(result);
} catch(err) {
	// do something
}
```

To handle RPC, you can use `app.globalRPCServer` and `app.localRPCServer`:

```ts
const server := app.localRPCServer()
server.registerhandler("helloRPC", (...inputs: any[]): any => {
	if (inputs.length === 0) {
		throw new Error("Message accepted but input is invalid");
	}
	const name = inputs[0] as string
	return greet(name);
}
```

# Naming Convention

A component usually contains several files:

* `index.ts`: This file contains code to produce your `setup`.
* `helpers.ts`: This file contains any general-purpose code that are used by `index.ts` or `services.ts`.
* `services.ts`: This file contains your main business logic. It should not care about how it is called.