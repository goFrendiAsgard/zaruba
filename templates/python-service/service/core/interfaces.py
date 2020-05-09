from abc import ABC, abstractmethod
from flask import Flask
from logging import Logger
from typing import List, Callable
from transport import Publisher, Subscriber, RPCServer, RPCClient

SetupComponent = Callable[[], None]


class App(ABC):

    @abstractmethod
    def logger(self) -> Logger:
        pass

    @abstractmethod
    def router(self) -> Flask:
        pass

    @abstractmethod
    def global_publisher(self) -> Publisher:
        pass

    @abstractmethod
    def local_publisher(self) -> Publisher:
        pass

    @abstractmethod
    def global_subscriber(self) -> Subscriber:
        pass

    @abstractmethod
    def local_subscriber(self) -> Subscriber:
        pass

    @abstractmethod
    def global_rpc_server(self) -> RPCServer:
        pass

    @abstractmethod
    def local_rpc_server(self) -> RPCServer:
        pass

    @abstractmethod
    def global_rpc_client(self) -> RPCClient:
        pass

    @abstractmethod
    def local_rpc_client(self) -> RPCClient:
        pass

    @abstractmethod
    def liveness(self) -> bool:
        pass

    @abstractmethod
    def readiness(self) -> bool:
        pass

    @abstractmethod
    def set_liveness(self, liveness: bool) -> None:
        pass

    @abstractmethod
    def set_readiness(self, readiness: bool) -> None:
        pass

    @abstractmethod
    def setup(self, setupComponents: List[SetupComponent]) -> None:
        pass

    @abstractmethod
    def run(self) -> None:
        pass
