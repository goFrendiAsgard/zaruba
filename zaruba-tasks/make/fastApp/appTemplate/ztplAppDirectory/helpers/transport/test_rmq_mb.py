from typing import Any
from helpers.transport.rmqMessagebus import RMQMessageBus
from helpers.transport.rmqConfig import RMQEventMap
from helpers.transport.rmqConnection import create_rmq_connection_parameters

import os
import warnings
import asyncio
import pytest


@pytest.mark.asyncio
async def test_rmq_mb():
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
    # await asyncio.sleep(3)

    result = {}
    try:
        @mb.handle('test_event')
        def handle(message: Any) -> Any:
            result['message'] = message
        
        await asyncio.sleep(3)
        mb.publish('test_event', 'test_message')

        trial: int = 10
        while trial > 0 and not 'message' in result:
            await asyncio.sleep(1)
            trial -= 1
    finally:
        mb.shutdown()
    assert 'message' in result
    assert result['message'] == 'test_message'
    