from .interfaces import App, Comp
import threading, pika
from flask import Flask
from logging import Logger
from typing import List
from transport import Subscriber, RPCServer

class MainApp(App):

    def __init__(self, logger: Logger, router: Flask, subscribers: List[Subscriber], rpc_servers: List[RPCServer], http_port: int, rmq_connection_list: List[pika.BlockingConnection]):
        self._http_port: int = http_port
        self._liveness: bool = False
        self._readiness: bool = False
        self._logger: Logger = logger
        self._router: Flask = router
        self._subscribers: List[Subscriber] = subscribers
        self._rpc_servers: List[RPCServer] = rpc_servers
        self._rmq_connection_list: List[pika.BlockingConnection] = rmq_connection_list

    def liveness(self) -> bool:
        return self._liveness

    def readiness(self) -> bool:
        return self._readiness

    def set_liveness(self, liveness: bool) -> None:
        self._liveness = liveness

    def set_readiness(self, readiness: bool) -> None:
        self._readiness = readiness

    def setup(self, components: List[Comp]) -> None:
        for component in components:
            component.setup()

    def run(self) -> None:
        try:
            with self._router.app_context():
                try:
                    self._serve_and_subscribe()
                    self._set_liveness_and_readiness(True)
                except Exception as e:
                    self._set_liveness_and_readiness(False)
                    raise e
            self._link_rmq_status_to_app(self._rmq_connection_list)
            self._router.run("0.0.0.0", self._http_port)
        except Exception as e:
            self._set_liveness_and_readiness(False)
            self._logger.error(e)
    
    def _link_rmq_status_to_app(self, rmq_connection_list: List[pika.BlockingConnection]):
        def check_closed():
            is_ok = True
            while is_ok:
                for rmq_connection in rmq_connection_list:
                    if rmq_connection.is_closed:
                        self._set_liveness_and_readiness(False)
                        is_ok = False
                        break
        thread = threading.Thread(target=check_closed)
        thread.start()

    def _set_liveness_and_readiness(self, liveness_and_readiness: bool) -> None:
        self.set_liveness(liveness_and_readiness)
        self.set_readiness(liveness_and_readiness)

    def _serve_and_subscribe(self) -> None:
        for subscriber in self._subscribers:
            subscriber.subscribe()
        for rpc_server in self._rpc_servers:
            rpc_server.serve()