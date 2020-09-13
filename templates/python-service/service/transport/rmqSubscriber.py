from typing import Dict, cast
import threading
from logging import Logger
from pika.adapters.blocking_connection import BlockingConnection, BlockingChannel
from pika.spec import Basic, BasicProperties
from .interfaces import EventHandler, Subscriber
from .envelopedMessage import EnvelopedMessage
from .rmqHelper import rmq_consume, rmq_declare_and_bind_queue, OnMessageCallback
from .rmqEventMap import RmqEventMap


class RmqSubscriber(Subscriber):

    def __init__(self, logger: Logger, connection: BlockingConnection, event_map: RmqEventMap):
        self.connection: BlockingConnection = connection
        self.logger: Logger = logger
        self.handlers: Dict[str, EventHandler] = cast(
            Dict[str, EventHandler], {})
        self.event_map: RmqEventMap = event_map

    def register_handler(self, event_name: str, handler: EventHandler) -> Subscriber:
        self.handlers[event_name] = handler
        return self

    def subscribe(self):
        ch = self.connection.channel()
        for key in self.handlers:
            event_name = key
            # declare dlx
            args = {}
            if self.event_map.get_ttl(event_name) > 0:
                dead_letter_exchange = self.event_map.get_dead_letter_exchange(event_name)
                dead_letter_queue = self.event_map.get_dead_letter_queue(event_name)
                rmq_declare_and_bind_queue(ch, dead_letter_exchange, dead_letter_queue, True, False, {})
                args = self.event_map.get_queue_args(event_name)
            # declare queue
            exchange_name = self.event_map.get_exchange_name(event_name)
            queue_name = self.event_map.get_queue_name(event_name)
            handler = self.handlers[event_name]
            rmq_declare_and_bind_queue(ch, exchange_name, queue_name, True, False, args)
            self.logger.info("[INFO RmqSubscriber] Subscribe {}".format(event_name))
            # consume
            auto_ack = self.event_map.get_auto_ack(event_name)
            rmq_consume(ch, queue_name, True, False, auto_ack, args, self._create_rmq_handler(event_name, auto_ack, handler))
        thread = threading.Thread(target=ch.start_consuming)
        thread.start()

    def _create_rmq_handler(self, event_name: str, auto_ack: bool, handler: EventHandler) -> OnMessageCallback:
        def on_message(ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, body: str):
            try:
                json_enveloped_input = body
                enveloped_input = EnvelopedMessage(json_enveloped_input)
                self.logger.info("[INFO RmqSubscriber] Get Event {}".format(event_name))
                try:
                    handler(enveloped_input.message)
                    if not auto_ack:
                        ch.basic_ack(delivery_tag=method.delivery_tag)
                except Exception as e:
                    if not auto_ack:
                        ch.basic_nack(delivery_tag=method.delivery_tag)
            except Exception as e:
                self.logger.info("[ERROR RmqSubscriber] Get Event {}: {}".format(event_name, e))
                self.logger.error(e)
                if not auto_ack:
                    ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_message
