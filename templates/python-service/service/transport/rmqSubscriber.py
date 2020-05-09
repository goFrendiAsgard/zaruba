from .interfaces import EventHandler, Subscriber
from .envelopedMessage import EnvelopedMessage
from .helpers import rmq_create_connection_and_channel, rmq_consume, rmq_declare_queue_and_bind_to_default_exchange, OnMessageCallback
from typing import Dict, cast
from logging import Logger, getLogger
import threading
from pika.adapters.blocking_connection import BlockingChannel
from pika.spec import Basic, BasicProperties


class RmqSubscriber(Subscriber):

    def __init__(self, connection_string: str):
        self.connection_string: str = connection_string
        self.logger: Logger = getLogger()
        self.handlers: Dict[str, EventHandler] = cast(
            Dict[str, EventHandler], {})

    def register_handler(self, event_name: str, handler: EventHandler) -> Subscriber:
        self.handlers[event_name] = handler
        return self

    def set_logger(self, logger: Logger) -> Subscriber:
        self.logger = logger
        return self

    def subscribe(self):
        _, ch = rmq_create_connection_and_channel(self.connection_string)
        for key in self.handlers:
            event_name = key
            handler = self.handlers[event_name]
            rmq_declare_queue_and_bind_to_default_exchange(ch, event_name)
            self.logger.info(
                "[INFO RmqSubscriber] Subscribe {}".format(event_name))
            rmq_consume(ch, event_name, self.create_rmq_handler(
                event_name, handler))
        thread = threading.Thread(target=ch.start_consuming)
        thread.start()

    def create_rmq_handler(self, event_name: str, handler: EventHandler) -> OnMessageCallback:
        def on_message(ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, body: str):
            try:
                json_enveloped_input = body
                enveloped_input = EnvelopedMessage(json_enveloped_input)
                self.logger.info(
                    "[INFO RmqSubscriber] Get Event {}".format(event_name))
                handler(enveloped_input.message)
            except Exception as e:
                self.logger.info(
                    "[ERROR RmqSubscriber] Get Event {}: {}".format(event_name, e))
                self.logger.error(e)
        return on_message
