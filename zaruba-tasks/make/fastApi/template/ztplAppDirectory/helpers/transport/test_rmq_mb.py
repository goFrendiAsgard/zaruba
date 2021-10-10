from typing import Any
from helpers.transport.rmq_mb import RMQMessageBus
from helpers.transport.rmq_config import RMQEventMap
from helpers.transport.rmq_connection import create_rmq_connection_parameters

import os

def test_mb():
    host = os.getenv('TEST_RABBITMQ_HOST', 'localhost')
    user = os.getenv('TEST_RABBITMQ_USER', 'root')
    password = os.getenv('TEST_RABBITMQ_PASS', 'toor')
    vhost = os.getenv('TEST_RABBITMQ_VHOST', '/')
    rmq_connection_parameters = create_rmq_connection_parameters(host, user, password, vhost)
    rmq_event_map = RMQEventMap({})

    mb = RMQMessageBus(rmq_connection_parameters, rmq_event_map)

    @mb.handle('test_event')
    def handle(message: Any) -> Any:
        assert message == 'test_message'
        mb.shutdown()
    
    mb.publish('test_event', 'test_message')

