# servicename

One Paragraph of project description goes here

# Start

```sh
conda env create --file=environment.yml || conda env update --file=environment.yml  --prune 
export DEFAULT_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export HTTP_PORT=3000
# ... other-env-setting
python main.py
```

# Test

```sh
conda env create --file=environment.yml || conda env update --file=environment.yml  --prune 
export TEST_RMQ_CONNECTION_STRING=<your-rmq-connection-string>
export TEST_HTTP_PORT=3000
# ... other-env-setting
python -m pytest
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
├── __init__.py
├── components
│   ├── __pycache__
│   ├── defaultcomponent
│   ├── example
│   └── monitoring
├── config
│   └── __init__.py
├── core
│   ├── __init__.py
│   ├── interfaces.py
│   └── mainApp.py
├── environment.yml
├── main.py
├── transport
│   ├── __init__.py
│   ├── envelopedMessage.py
│   ├── helpers.py
│   ├── interfaces.py
│   ├── rmqPublisher.py
│   ├── rmqRpcClient.py
│   ├── rmqRpcServer.py
│   └── rmqSubscriber.py
└── servicename.zaruba.yaml
```

# Bootstrap

App initialization contains of 4 steps:

* __App component definitions:__ In this step we define all components that will be used by our app or other components (e.g: logger, config, router, publisher, subscriber, rpc-server, and rpc-client). The definition should be explicit so that other developers can easily understand which components depend on which component.
* __App creation:__ In this step, we initiate our app.
* __App setup:__ In this step, we setup our app, by registering your components. To register your component, you need to provide a function with no parameter and return nothing. You can surely wrap the function into object method or wrapper-function. If you need to inject dependencies, you can use [closure](https://en.wikipedia.org/wiki/Closure_(computer_programming)) (more about this latter).
* __App Execution:__ Finally, we run our app.

Below is how everything looks like in `main.py`:

```python

def main():

    # app component definitions
    logging.basicConfig(level="INFO")
    logger = getLogger()
    config = Config()
    logger.info("CONFIG: {}".format(config))
    router: Flask = Flask(__name__)

    rmq_connection_string = config.default_rmq_connection_string
    rmq_connection_url = pika.URLParameters(rmq_connection_string)

    rpc_server_connection = pika.BlockingConnection(rmq_connection_url)
    rpc_server = RmqRPCServer(logger, rpc_server_connection)

    rpc_client_connection = pika.BlockingConnection(rmq_connection_url)
    rpc_client = RmqRPCClient(logger, rpc_client_connection)

    subscriber_connection = pika.BlockingConnection(rmq_connection_url)
    subscriber = RmqSubscriber(logger, subscriber_connection)

    publisher_connection = pika.BlockingConnection(rmq_connection_url)
    publisher = RmqPublisher(logger, publisher_connection)

    # app creation
    app = MainApp(
        logger,
        router,
        [subscriber],
        [rpc_server],
        config.http_port,
    )

    # app setup
    app.setup([
        defaultComponent.Component(config, router),
        monitoring.Component(config, app, router),
        example.Component(
            config, router, publisher, subscriber, rpc_server, rpc_client)
    ])
    link_rmq_status_to_app(
        app, [rpc_server_connection, subscriber_connection, client_connection])

    # app execution
    app.run()
```

> __NOTE:__ Treat our app bootstrap as your own. You are free to define new components, replace one component with another ones, or wire the components differently.

> __NOTE:__ Pika blocking connection is not thread safe. THus we need to define different connection for publisher, subscriber, rpc-server, and rpc-client.

# Components

All components should be location on `component` directory. To expose a component and put them on `app setup`, you should provide a class with `setup` method.

For example, our `monitoring` component need `config`, `app`, and `router`. Thus we utilize closure to inject those components:

```python
class Component(Comp):

    def __init__(self, config: Config, app: App, router: Flask):
        self.config = config
        self.app = app
        self.router = router

    def setup(self):
        service_name = self.config.service_name

        @self.router.route("/liveness", methods=["GET"])
        def liveness():
            liveness = self.app.liveness()
            http_code = 200 if liveness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_alive": liveness})
            return Response(json_response, status=http_code, mimetype="application/json")

        @self.router.route("/readiness", methods=["GET"])
        def readiness():
            readiness = self.app.readiness()
            http_code = 200 if readiness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_ready": readiness})
            return Response(json_response, status=http_code, mimetype="application/json")

```

The returned function is then used on `app setup`:

```python
 app.setup([
	# ...
	monitoring.Component(config, app, router),
	# ...
])
```

We has already provide an `example` component that utilize OOP with some possible use-cases (including RPC and pub/sub communication). Feel free to explore.

# Component Convention

Typically a component should contains at least one file named `__init__.py`. Although you can write anything inside `__init__.py`, it is better to just put `router`, `rpc-server`, and `subscriber` handler in it.

Make your components as loosly coupled as possible, as it is going to help you in case of you want to split the app into several smaller apps.

If you need some business logic, please put them inside `service` directory or `service.py`.

If your components are depending on each other, please put the dependency in our `main.py` bootstrap, don't import the component directly as it will make your components tightly coupled to each other.

If you need database connection or cache mechanism, please also define them in our `main.py`.