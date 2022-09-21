from typing import Any, Callable, Mapping
from helpers.transport.rpc import RPC

class LocalRPC(RPC):

    def __init__(self):
        self.rpc_handler: Mapping[str, Callable[..., Any]] = {}
        self.error_count = 0

    def get_error_count(self) -> int:
        return self.error_count

    def is_failing(self) -> bool:
        return False

    def shutdown(self):
        pass

    def handle(self, rpc_name: str) -> Callable[..., Any]:
        def register_rpc_handler(rpc_handler: Callable[..., Any]):
            self.rpc_handler[rpc_name] = rpc_handler
        return register_rpc_handler

    def call(self, rpc_name: str, *args: Any) -> Any:
        if rpc_name not in self.rpc_handler:
            self.error_count += 1
            raise Exception('RPC handler for "{}" is not found'.format(rpc_name))
        print({'action': 'call_local_rpc', 'rpc_name': rpc_name, 'args': args})
        try:
            print({'action': 'handle_local_rpc', 'rpc_name': rpc_name, 'args': args})
            result = self.rpc_handler[rpc_name](*args)
            print({'action': 'get_local_rpc_reply', 'rpc_name': rpc_name, 'args': args, 'result': result})
            return result
        except Exception as e:
            raise e