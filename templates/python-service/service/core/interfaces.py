from abc import ABC, abstractmethod
from flask import Flask
from logging import Logger
from typing import List, Callable, NoReturn

SetupComponent = Callable[[], NoReturn]


class App(ABC):

    @abstractmethod
    def logger(self) -> Logger:
        pass

    @abstractmethod
    def router(self) -> Flask:
        pass

    @abstractmethod
    def liveness(self) -> bool:
        pass

    @abstractmethod
    def readiness(self) -> bool:
        pass

    @abstractmethod
    def set_liveness(self, liveness: bool) -> NoReturn:
        pass

    @abstractmethod
    def set_readiness(self, readiness: bool) -> NoReturn:
        pass

    @abstractmethod
    def setup(self, setupComponents: List[SetupComponent]) -> NoReturn:
        pass

    @abstractmethod
    def run(self) -> NoReturn:
        pass
