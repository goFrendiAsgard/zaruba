from typing import Any, Callable, Mapping
from transport.interface import MessageBus
from transport.logger import log_info

class LocalMessageBus(MessageBus):

    def __init__(self):
        self.event_handler: Mapping[str, Callable[[Any], Any]] = {}
        self.rpc_handler: Mapping[str, Callable[..., Any]] = {}

    
    def shutdown(self):
        pass


    def handle_rpc(self, event_name: str, handler: Callable[..., Any]) -> Any:
        self.rpc_handler[event_name] = handler


    def call_rpc(self, event_name: str, *args: Any) -> Any:
        if event_name not in self.rpc_handler:
            raise Exception('RPC handler for "{}" is not found'.format(event_name))
        log_info('CALL RPC', event_name=event_name, args=args)
        result = self.rpc_handler[event_name](*args)
        log_info('GET RPC REPLY', event_name=event_name, args=args, result=result)
        return result


    def handle(self, event_name: str, handler: Callable[[Any], Any]) -> Any:
        self.event_handler[event_name] = handler


    def publish(self, event_name: str, msg: Any) -> Any:
        if event_name not in self.event_handler:
            raise Exception('Event handler for "{}" is not found'.format(event_name))
        log_info('PUBLISH EVENT', event_name=event_name, msg=msg)
        self.event_handler[event_name](msg)