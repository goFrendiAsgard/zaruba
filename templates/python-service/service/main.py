import logging
from flask import Flask
from logging import getLogger
from config import Config
from core import MainApp, App, create_rmq_connection
from transport import RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient

import components.main.component as mainComponent
import components.monitoring.component as monitoring
import components.example.component as example

def main():

    # app component definitions
    logging.basicConfig(level="INFO")
    logger = getLogger()
    config = Config()
    logger.info("CONFIG: {}".format(config))
    router: Flask = Flask(__name__)
    rmq_connection_string = config.default_rmq_connection_string
    rmq_event_map = config.rmq_event_map

    rpc_server_connection = create_rmq_connection(rmq_connection_string, 10)
    rpc_server = RmqRPCServer(logger, rpc_server_connection, rmq_event_map)

    subscriber_connection = create_rmq_connection(rmq_connection_string, 10)
    subscriber = RmqSubscriber(logger, subscriber_connection, rmq_event_map)

    client_connection = create_rmq_connection(rmq_connection_string, 10)
    rpc_client = RmqRPCClient(logger, client_connection, rmq_event_map)
    publisher = RmqPublisher(logger, client_connection, rmq_event_map)

    # app creation
    app: App = MainApp(
        logger=logger,
        router=router,
        subscribers=[subscriber],
        rpc_servers=[rpc_server],
        http_port=config.http_port,
        rmq_connection_list=[rpc_server_connection, subscriber_connection, client_connection]
    )

    # app setup
    app.setup([
        mainComponent.Component(config, router),
        monitoring.Component(config, app, router),
        example.Component(config, app, router, publisher, subscriber, rpc_server, rpc_client)
    ])

    # app execution
    app.run()


if __name__ == "__main__":
    main()
