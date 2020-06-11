from abc import ABC, abstractmethod
from flask import Flask
from logging import Logger
from typing import List, Callable
from transport import Publisher, Subscriber, RPCServer, RPCClient


class Comp(ABC):

    @abstractmethod
    def setup(self) -> None:
        pass


class App(ABC):

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
    def setup(self, setupComponents: List[Comp]) -> None:
        pass

    @abstractmethod
    def run(self) -> None:
        pass
