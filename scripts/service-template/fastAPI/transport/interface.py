from typing import Any, Callable
import abc

class MessageBus(abc.ABC):

    @abc.abstractmethod
    def handle_rpc(self, queue: str, handler: Callable[..., Any]) -> Any:
        pass

    @abc.abstractmethod
    def call_rpc(self, queue: str, *args: Any) -> Any:
        pass

    @abc.abstractmethod
    def handle(self, queue: str, handler: Callable[[Any], Any]) -> Any:
        pass

    @abc.abstractmethod
    def publish(self, queue: str, msg: Any) -> Any:
        pass