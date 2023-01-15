from typing import Any, Callable
from helper.transport import RPC


class AppRPC(RPC):
    '''
    Wrapper for RPC with special methods to support app use case.

    Feel free to add methods as necessary.
    '''

    def __init__(self, rpc: RPC):
        '''
        Initiate a new AppRPC based on specified rpc_type.

        Keyword arguments:
        - rpc_type -- RPC type (e.g., rmq, local)
        '''
        self.rpc = rpc

    def handle(self, rpc_name: str) -> Callable[..., Any]:
        '''
        Decorator to handle an RPC request.

        Keyword arguments:
        - rpc_name -- Name of the RPC you want to handle.
        '''
        return self.rpc.handle(rpc_name)

    def call(self, rpc_name: str, *args: Any) -> Any:
        '''
        Publish an RPC with *args as arguments.

        Keyword arguments:
        - rpc_name -- Name of the RPC you want to call.
        '''
        return self.rpc.call(rpc_name, *args)

    def shutdown(self) -> Any:
        '''
        Shutdown the AppRPC
        '''
        return self.rpc.shutdown()

    def get_error_count(self) -> int:
        '''
        Get how many errors has been occurred while call/handle RPC.
        '''
        return self.rpc.get_error_count()

    def is_failing(self) -> bool:
        '''
        Get whether AppRPC is failing (and should be terminated) or not.
        '''
        return self.rpc.is_failing()

