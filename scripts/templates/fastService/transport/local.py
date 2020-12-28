from typing import Any, Callable, Mapping
from transport.interface import MessageBus

class LocalMessageBus(MessageBus):

    def __init__(self):
        self.event_handler: Mapping[str, Callable[[Any], Any]] = {}
        self.rpc_handler: Mapping[str, Callable[..., Any]] = {}


    def handle_rpc(self, event_name: str, handler: Callable[..., Any]) -> Any:
        self.rpc_handler[event_name] = handler


    def call_rpc(self, event_name: str, *args: Any) -> Any:
        if event_name not in self.rpc_handler:
            raise Exception('RPC handler for "{}" is not found'.format(event_name))
        return self.rpc_handler[event_name](*args)


    def handle(self, event_name: str, handler: Callable[[Any], Any]) -> Any:
        self.event_handler[event_name] = handler


    def publish(self, event_name: str, msg: Any) -> Any:
        if event_name not in self.event_handler:
            raise Exception('Event handler for "{}" is not found'.format(event_name))
        self.event_handler[event_name](msg)