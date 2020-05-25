from .interfaces import App, SetupComponent
from flask import Flask
from logging import Logger
from typing import List
from transport import Publisher, Subscriber, RPCServer, RPCClient, RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient


class MainApp(App):

    def __init__(self, logger: Logger, router: Flask, subscribers: List[Subscriber], rpc_servers: List[RPCServer], http_port: int):
        self._http_port: int = http_port
        self._liveness: bool = False
        self._readiness: bool = False
        self._logger: Logger = logger
        self._router: Flask = router
        self._subscribers: List[Subscriber] = subscribers
        self._rpc_servers: List[RPCServer] = rpc_servers

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
            with self._router.app_context():
                try:
                    for subscriber in self._subscribers:
                        subscriber.subscribe()
                    for rpc_server in self._rpc_servers:
                        rpc_server.serve()
                    self.set_liveness(True)
                    self.set_readiness(True)
                except Exception as e:
                    self.set_liveness(False)
                    self.set_readiness(False)
                    raise e
            self._router.run("0.0.0.0", self._http_port)
        except Exception as e:
            self._logger.error(e)