from typing import Any, Callable
import abc


class RPC(abc.ABC):

    @abc.abstractmethod
    def handle(self, rpc_name: str) -> Callable[..., Any]:
        '''
        Decorator to handle an RPC request

        Keyword arguments:
        - rpc_name -- Name of the RPC you want to handle.
        '''
        pass

    @abc.abstractmethod
    def call(self, rpc_name: str, *args: Any) -> Any:
        '''
        Publish an RPC with *args as arguments.

        Keyword arguments:
        - rpc_name -- Name of the RPC you want to call.
        '''
        pass

    @abc.abstractmethod
    def shutdown(self) -> Any:
        '''
        Shutdown the RPC
        '''
        pass

    @abc.abstractclassmethod
    def get_error_count(self) -> int:
        '''
        Get how many errors has been occurred while call/handle RPC.
        '''
        pass

    @abc.abstractclassmethod
    def is_failing(self) -> bool:
        '''
        Get whether RPC is failing (and should be terminated) or not.
        '''
        pass
