from .interfaces import Message, Publisher
from .helpers import rmq_create_connection_and_channel, rmq_declare_queue_and_bind_to_default_exchange, rmq_publish, rmq_close_connection_and_channel
from .envelopedMessage import EnvelopedMessage
from logging import Logger, getLogger
import json


class RmqPublisher(Publisher):

    def __init__(self, connection_string: str):
        self.connection_string: str = connection_string
        self.logger: Logger = getLogger(__name__.split(".")[0])

    def set_logger(self, logger: Logger) -> Publisher:
        self.logger = logger
        return self

    def publish(self, event_name: str, msg: Message):
        self.logger.info(
            "[INFO RmqPublisher] Publish {} {}".format(event_name, json.dumps(msg)))
        conn, ch = rmq_create_connection_and_channel(self.connection_string)
        rmq_declare_queue_and_bind_to_default_exchange(ch, event_name)
        enveloped_message = EnvelopedMessage().set_correlation_id().set_message(msg)
        rmq_publish(ch, event_name, "", enveloped_message.to_json(), None)
        rmq_close_connection_and_channel(conn, ch)
