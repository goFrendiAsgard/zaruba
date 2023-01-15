from typing import Any
from helper.transport.rmq_messagebus import RMQMessageBus
from helper.transport.rmq_config import RMQEventMap
from helper.transport.rmq_connection import create_rmq_connection_parameters
from helper.config.boolean import get_boolean_env

import os
import logging
import asyncio
import pytest


@pytest.mark.asyncio
async def test_rmq_mb():
    if not get_boolean_env('TEST_INTEGRATION', False):
        logging.warn('RMQMessageBus is not tested')
        return None

    host = os.getenv('TEST_RABBITMQ_HOST', 'localhost')
    user = os.getenv('TEST_RABBITMQ_USER', '')
    password = os.getenv('TEST_RABBITMQ_PASS', '')
    vhost = os.getenv('TEST_RABBITMQ_VHOST', '/')
    rmq_connection_parameters = create_rmq_connection_parameters(
        host, user, password, vhost
    )
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
        while trial > 0 and 'message' not in result:
            await asyncio.sleep(1)
            trial -= 1
    finally:
        mb.shutdown()
    assert 'message' in result
    assert result['message'] == 'test_message'
