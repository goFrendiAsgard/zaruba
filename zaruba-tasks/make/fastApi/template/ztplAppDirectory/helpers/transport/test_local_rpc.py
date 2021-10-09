from typing import Any
from helpers.transport.local_rpc import LocalRPC

def test_mb():
    rpc = LocalRPC()

    @rpc.handle('test_rpc')
    def handle(parameter_1: Any, parameter_2: str) -> Any:
        assert parameter_1 == 'hello'
        assert parameter_2 == 'world'
        return 'hello world'
    
    result = rpc.call('test_rpc', 'hello', 'world')
    assert result == 'hello world'
    rpc.shutdown()