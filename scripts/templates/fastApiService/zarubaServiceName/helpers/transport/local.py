from typing import Any, Callable, Mapping
from helpers.transport.interface import MessageBus

class LocalMessageBus(MessageBus):

    def __init__(self):
        self.event_handler: Mapping[str, Callable[[Any], Any]] = {}
        self.rpc_handler: Mapping[str, Callable[..., Any]] = {}
    
    def shutdown(self):
        pass

    def handle_rpc(self, event_name: str) -> Callable[..., Any]:
        def register_rpc_handler(rpc_handler: Callable[..., Any]):
            self.rpc_handler[event_name] = rpc_handler
        return register_rpc_handler

    def call_rpc(self, event_name: str, *args: Any) -> Any:
        if event_name not in self.rpc_handler:
            raise Exception('RPC handler for "{}" is not found'.format(event_name))
        print({'action': 'call_local_rpc', 'event_name': event_name, 'args': args})
        result = self.rpc_handler[event_name](*args)
        print({'action': 'get_local_rpc_reply', 'event_name': event_name, 'args': args, 'result': result})
        return result

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            self.event_handler[event_name] = event_handler
        return register_event_handler

    def publish(self, event_name: str, message: Any) -> Any:
        if event_name not in self.event_handler:
            raise Exception('Event handler for "{}" is not found'.format(event_name))
        print({'action': 'publish_local_event', 'event_name': event_name, 'message': message})
        self.event_handler[event_name](message)