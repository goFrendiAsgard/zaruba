from .interfaces import EventHandler, Subscriber
from .envelopedMessage import EnvelopedMessage
from .helpers import rmq_create_connection_and_channel, rmq_consume, rmq_declare_queue_and_bind_to_default_exchange
from typing import Dict, Mapping
from logging import Logger, getLogger


class RmqSubscriber(Subscriber):

    def __init__(self, connection_string: str):
        self.connection_string: str = connection_string
        self.logger: Logger = getLogger(__name__.split(".")[0])
        self.handlers: Mapping[str, EventHandler] = {}

    def register_handler(self, event_name: str, handler: EventHandler) -> Subscriber:
        self.handlers[event_name] = handler
        return self

    def set_logger(self, logger: Logger) -> Subscriber:
        self.logger = logger
        return self
