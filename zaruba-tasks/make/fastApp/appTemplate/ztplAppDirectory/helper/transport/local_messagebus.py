from typing import Any, Callable, Mapping
from helper.transport.messagebus import MessageBus

import logging


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
            raise Exception('No event handler found for {}'.format(event_name))
        logging.info(
            'Publish event {} with message: {}'.format(event_name, message)
        )
        try:
            logging.info(
                'Handle event {} with message: {}'.format(event_name, message)
            )
            self._event_handler[event_name](message)
        except Exception as exception:
            logging.error(
                'Error publishing event {} with message: {}'.format(
                    event_name, message
                ),
                exc_info=True
            )
            raise exception
