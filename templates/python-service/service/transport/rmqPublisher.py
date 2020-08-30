import json
from pika.adapters.blocking_connection import BlockingConnection
from .interfaces import Message, Publisher
from .rmqHelper import rmq_declare_and_bind_queue, rmq_declare_fanout_exchange, rmq_publish
from .rmqEventMap import RmqEventMap
from .envelopedMessage import EnvelopedMessage
from logging import Logger


class RmqPublisher(Publisher):

    def __init__(self, logger: Logger, connection: BlockingConnection, event_map: RmqEventMap):
        self.connection: BlockingConnection = connection
        self.logger: Logger = logger
        self.event_map: RmqEventMap = event_map

    def publish(self, event_name: str, msg: Message):
        self.logger.info(
            "[INFO RmqPublisher] Publish {} {}".format(event_name, json.dumps(msg)))
        ch = self.connection.channel()
        exchange_name = self.event_map.get_exchange_name(event_name)
        rmq_declare_fanout_exchange(ch, exchange_name)
        enveloped_message = EnvelopedMessage().set_correlation_id().set_message(msg)
        rmq_publish(ch, exchange_name, "", enveloped_message.to_json(), None)
        ch.close()
