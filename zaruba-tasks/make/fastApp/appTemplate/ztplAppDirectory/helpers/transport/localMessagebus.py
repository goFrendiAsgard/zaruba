from typing import Any, Callable, Mapping
from helpers.transport.messagebus import MessageBus

class LocalMessageBus(MessageBus):

    def __init__(self):
        self._event_handler: Mapping[str, Callable[[Any], Any]] = {}
        self._error_count = 0

    def get_error_count(self) -> int:
        return self._error_count

    def is_failing(self) -> bool:
        return False

    def shutdown(self):
        pass

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            self._event_handler[event_name] = event_handler
        return register_event_handler

    def publish(self, event_name: str, message: Any) -> Any:
        if event_name not in self._event_handler:
            self._error_count += 1
            raise Exception('Event handler for "{}" is not found'.format(event_name))
        print({'action': 'publish_local_event', 'event_name': event_name, 'message': message})
        try:
            print({'action': 'handle_local_event', 'event_name': event_name, 'message': message})
            self._event_handler[event_name](message)
        except Exception as e:
            raise e