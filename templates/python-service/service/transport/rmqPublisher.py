import pika
import json
from .interfaces import Message, Publisher
from .helpers import rmq_declare_queue_and_bind_to_default_exchange, rmq_publish
from .envelopedMessage import EnvelopedMessage
from logging import Logger


class RmqPublisher(Publisher):

    def __init__(self, logger: Logger, connection: pika.BlockingConnection):
        self.connection: pika.BlockingConnection = connection
        self.logger: Logger = logger

    def publish(self, event_name: str, msg: Message):
        self.logger.info(
            "[INFO RmqPublisher] Publish {} {}".format(event_name, json.dumps(msg)))
        ch = self.connection.channel()
        rmq_declare_queue_and_bind_to_default_exchange(ch, event_name)
        enveloped_message = EnvelopedMessage().set_correlation_id().set_message(msg)
        rmq_publish(ch, event_name, "", enveloped_message.to_json(), None)
        ch.close()
