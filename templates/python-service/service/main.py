import pika
import logging
from flask import Flask
from logging import getLogger
from config import Config
from core import MainApp
from transport import RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient

import components.defaultcomponent as defaultComponent
import components.monitoring as monitoring
import components.example as example


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
        defaultComponent.create_setup(config, router),
        monitoring.create_setup(config, app, router),
        example.Component(
            config, router, publisher, subscriber, rpc_server, rpc_client).setup
    ])

    # app execution
    app.run()


if __name__ == "__main__":
    main()
