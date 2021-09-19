from typing import Any, Callable, Mapping
from helpers.transport.interface import MessageBus

class LocalMessageBus(MessageBus):

    def __init__(self):
        self.event_handler: Mapping[str, Callable[[Any], Any]] = {}
    
    def shutdown(self):
        pass

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            self.event_handler[event_name] = event_handler
        return register_event_handler

    def publish(self, event_name: str, message: Any) -> Any:
        if event_name not in self.event_handler:
            raise Exception('Event handler for "{}" is not found'.format(event_name))
        print({'action': 'publish_local_event', 'event_name': event_name, 'message': message})
        self.event_handler[event_name](message)