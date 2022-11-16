from typing import Any, Callable
from helper.transport import RPC


class AppRPC(RPC):
    '''
    RPC with special methods to support app use case.
    Feel free to add methods as necessary.
    '''

    def __init__(self, rpc: RPC):
        self.rpc = rpc


    def handle(self, rpc_name: str) -> Callable[..., Any]:
        return self.rpc.handle(rpc_name)


    def call(self, rpc_name: str, *args: Any) -> Any:
        return self.rpc.call(rpc_name, *args)


    def shutdown(self) -> Any:
        return self.rpc.shutdown()


    def get_error_count(self) -> int:
        return self.rpc.get_error_count()


    def is_failing(self) -> bool:
        return self.rpc.is_failing()

