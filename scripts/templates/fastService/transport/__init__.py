from typing import Mapping
from transport.interface import MessageBus
from transport.rmq import RMQMessageBus
from transport.local import LocalMessageBus
from transport.helper import handle, handle_rpc


def init_mb(config: Mapping[str, str]) -> MessageBus:
    message_bus_type = config.get('message_bus_type', 'local')
    if message_bus_type == 'rmq':
        return RMQMessageBus(
            rmq_host = config.get('rabbitmq_host', 'localhost'),
            rmq_user = config.get('rabbitmq_user', 'root'),
            rmq_pass = config.get('rabbitmq_pass', 'toor'),
            rmq_vhost = config.get('rabbitmq_vhost', '/')
        )
    return LocalMessageBus()