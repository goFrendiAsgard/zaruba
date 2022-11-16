from typing import Any, Callable
import abc

class RPC(abc.ABC):

    @abc.abstractmethod
    def handle(self, rpc_name: str) -> Callable[..., Any]:
        pass

    @abc.abstractmethod
    def call(self, rpc_name: str, *args: Any) -> Any:
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

