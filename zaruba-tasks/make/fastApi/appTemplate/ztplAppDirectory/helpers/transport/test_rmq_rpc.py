from typing import Any
from helpers.transport.rmq_rpc import RMQRPC
from helpers.transport.rmq_config import RMQEventMap
from helpers.transport.rmq_connection import create_rmq_connection_parameters

import os
import warnings

def test_rmq_rpc():
    if os.getenv('TEST_INTEGRATION', '0') != '1':
        warnings.warn(UserWarning('TEST_INTEGRATION != 1, RMQRPC is not tested'))
        return None

    host = os.getenv('TEST_RABBITMQ_HOST', 'localhost')
    user = os.getenv('TEST_RABBITMQ_USER', 'root')
    password = os.getenv('TEST_RABBITMQ_PASS', 'Alch3mist')
    vhost = os.getenv('TEST_RABBITMQ_VHOST', '/')
    rmq_connection_parameters = create_rmq_connection_parameters(host, user, password, vhost)
    rmq_event_map = RMQEventMap({})

    rpc = RMQRPC(rmq_connection_parameters, rmq_event_map)

    @rpc.handle('test_rpc')
    def handle(parameter_1: Any, parameter_2: str) -> Any:
        assert parameter_1 == 'hello'
        assert parameter_2 == 'world'
        return 'hello world'
    
    result = rpc.call('test_rpc', 'hello', 'world')
    assert result == 'hello world'
    rpc.shutdown()