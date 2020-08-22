import logging
from flask import Flask
from logging import getLogger
from config import Config
from core import MainApp, App, create_rmq_connection
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

    rpc_server_connection = create_rmq_connection(rmq_connection_string, 10)
    rpc_server = RmqRPCServer(logger, rpc_server_connection)

    subscriber_connection = create_rmq_connection(rmq_connection_string, 10)
    subscriber = RmqSubscriber(logger, subscriber_connection)

    client_connection = create_rmq_connection(rmq_connection_string, 10)
    rpc_client = RmqRPCClient(logger, client_connection)
    publisher = RmqPublisher(logger, client_connection)

    # app creation
    app: App = MainApp(
        logger=logger,
        router=router,
        subscribers=[subscriber],
        rpc_servers=[rpc_server],
        http_port=config.http_port,
        rmq_connection_list=[
            rpc_server_connection, subscriber_connection, client_connection]
    )

    # app setup
    app.setup([
        defaultComponent.Component(config, router),
        monitoring.Component(config, app, router),
        example.Component(
            config, router, publisher, subscriber, rpc_server, rpc_client)
    ])

    # app execution
    app.run()


if __name__ == "__main__":
    main()
