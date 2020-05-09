from .interfaces import App, SetupComponent
from flask import Flask
from logging import Logger, getLogger
from typing import List
from transport import Publisher, Subscriber, RPCServer, RPCClient, RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient
import logging


class MainApp(App):

    def __init__(self, http_port: int, global_rmq_connection_string: str, local_rmq_connection_string: str):
        self._http_port: int = http_port
        self._app: Flask = Flask(__name__.split(".")[0])
        logging.basicConfig(level="INFO")
        self._logger: Logger = getLogger()
        self._liveness: bool = False
        self._readiness: bool = False
        self._global_publisher: Publisher = RmqPublisher(
            global_rmq_connection_string).set_logger(self._logger)
        self._local_publisher: Publisher = RmqPublisher(
            local_rmq_connection_string).set_logger(self._logger)
        self._global_subscriber: Subscriber = RmqSubscriber(
            global_rmq_connection_string).set_logger(self._logger)
        self._local_subscriber: Subscriber = RmqSubscriber(
            local_rmq_connection_string).set_logger(self._logger)
        self._global_rpc_server: RPCServer = RmqRPCServer(
            global_rmq_connection_string).set_logger(self._logger)
        self._local_rpc_server: RPCServer = RmqRPCServer(
            local_rmq_connection_string).set_logger(self._logger)
        self._global_rpc_client: RPCClient = RmqRPCClient(
            global_rmq_connection_string).set_logger(self._logger)
        self._local_rpc_client: RPCClient = RmqRPCClient(
            local_rmq_connection_string).set_logger(self._logger)

    def logger(self) -> Logger:
        return self._logger

    def router(self) -> Flask:
        return self._app

    def global_publisher(self) -> Publisher:
        return self._global_publisher

    def local_publisher(self) -> Publisher:
        return self._local_publisher

    def global_subscriber(self) -> Subscriber:
        return self._global_subscriber

    def local_subscriber(self) -> Subscriber:
        return self._local_subscriber

    def global_rpc_server(self) -> RPCServer:
        return self._global_rpc_server

    def local_rpc_server(self) -> RPCServer:
        return self._local_rpc_server

    def global_rpc_client(self) -> RPCClient:
        return self._global_rpc_client

    def local_rpc_client(self) -> RPCClient:
        return self._local_rpc_client

    def liveness(self) -> bool:
        return self._liveness

    def readiness(self) -> bool:
        return self._readiness

    def set_liveness(self, liveness: bool) -> None:
        self._liveness = liveness

    def set_readiness(self, readiness: bool) -> None:
        self._readiness = readiness

    def setup(self, setupComponents: List[SetupComponent]) -> None:
        for setupComponent in setupComponents:
            setupComponent()

    def run(self) -> None:
        try:
            with self._app.app_context():
                try:
                    self._global_subscriber.subscribe()
                    self._local_subscriber.subscribe()
                    self._global_rpc_server.serve()
                    self._local_rpc_server.serve()
                    self.set_liveness(True)
                    self.set_readiness(True)
                except:
                    self.set_liveness(False)
                    self.set_readiness(False)
            self._app.run("0.0.0.0", self._http_port)
        except Exception as e:
            self._logger.error(e)
