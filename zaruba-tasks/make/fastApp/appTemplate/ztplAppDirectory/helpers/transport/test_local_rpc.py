from typing import Any
from helpers.transport.localRpc import LocalRPC

def test_local_rpc():
    rpc = LocalRPC()

    parameters = {}
    @rpc.handle('test_rpc')
    def handle(parameter_1: Any, parameter_2: str) -> Any:
        parameters['first'] = parameter_1
        parameters['second'] = parameter_2
        return 'hello world'
    
    result = rpc.call('test_rpc', 'hello', 'world')
    rpc.shutdown()
    assert parameters['first'] == 'hello'
    assert parameters['second'] == 'world'
    assert result == 'hello world'