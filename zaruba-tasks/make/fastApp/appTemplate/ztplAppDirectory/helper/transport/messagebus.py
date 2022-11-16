from typing import Any, Callable
import abc

class MessageBus(abc.ABC):

    @abc.abstractmethod
    def handle(self, event_name: str) -> Callable[..., Any]:
        pass

    @abc.abstractmethod
    def publish(self, event_name: str, message: Any) -> Any:
        pass

    @abc.abstractmethod
    def shutdown(self) -> Any:
        pass

    @abc.abstractclassmethod
    def get_error_count(self) -> int:
        pass

    @abc.abstractclassmethod
    def is_failing(self) -> bool:
        pass