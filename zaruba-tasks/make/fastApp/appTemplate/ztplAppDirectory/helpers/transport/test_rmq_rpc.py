from typing import Any
from helpers.transport.rmqRpc import RMQRPC
from helpers.transport.rmqConfig import RMQEventMap
from helpers.transport.rmqConnection import create_rmq_connection_parameters

import os
import warnings
import asyncio
import pytest


@pytest.mark.asyncio
async def test_rmq_rpc():
    if os.getenv('TEST_INTEGRATION', '0') != '1':
        warnings.warn(UserWarning('TEST_INTEGRATION != 1, RMQRPC is not tested'))
        return None

    host = os.getenv('TEST_RABBITMQ_HOST', 'localhost')
    user = os.getenv('TEST_RABBITMQ_USER', '')
    password = os.getenv('TEST_RABBITMQ_PASS', '')
    vhost = os.getenv('TEST_RABBITMQ_VHOST', '/')
    rmq_connection_parameters = create_rmq_connection_parameters(host, user, password, vhost)
    rmq_event_map = RMQEventMap({})

    rpc = RMQRPC(rmq_connection_parameters, rmq_event_map)
    # await asyncio.sleep(3)

    result = ""
    try:
        parameters = {}
        @rpc.handle('test_rpc')
        def handle(parameter_1: Any, parameter_2: str) -> Any:
            parameters['first'] = parameter_1
            parameters['second'] = parameter_2
            return 'hello world'
        
        await asyncio.sleep(3)
        result = rpc.call('test_rpc', 'hello', 'world')
    finally:
        rpc.shutdown()
    assert parameters['first'] == 'hello'
    assert parameters['second'] == 'world'
    assert result == 'hello world'
    
