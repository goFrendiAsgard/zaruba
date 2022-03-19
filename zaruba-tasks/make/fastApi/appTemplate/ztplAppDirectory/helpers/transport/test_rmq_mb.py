from typing import Any
from helpers.transport.rmq_mb import RMQMessageBus
from helpers.transport.rmq_config import RMQEventMap
from helpers.transport.rmq_connection import create_rmq_connection_parameters

import os
import warnings
import asyncio

def test_rmq_mb():
    asyncio.run(_test_rmq_mb())


async def _test_rmq_mb():
    if os.getenv('TEST_INTEGRATION', '0') != '1':
        warnings.warn(UserWarning('TEST_INTEGRATION != 1, RMQMessageBus is not tested'))
        return None

    host = os.getenv('TEST_RABBITMQ_HOST', 'localhost')
    user = os.getenv('TEST_RABBITMQ_USER', '')
    password = os.getenv('TEST_RABBITMQ_PASS', '')
    vhost = os.getenv('TEST_RABBITMQ_VHOST', '/')
    rmq_connection_parameters = create_rmq_connection_parameters(host, user, password, vhost)
    rmq_event_map = RMQEventMap({})

    mb = RMQMessageBus(rmq_connection_parameters, rmq_event_map)

    result = {}
    @mb.handle('test_event')
    def handle(message: Any) -> Any:
        result['message'] = message
        mb.shutdown()
    
    mb.publish('test_event', 'test_message')
    await asyncio.sleep(5)
    assert 'message' in result
    assert result['message'] == 'test_message'