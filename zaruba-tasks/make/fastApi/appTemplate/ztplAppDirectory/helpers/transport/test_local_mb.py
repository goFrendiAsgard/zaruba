from typing import Any
from helpers.transport.local_mb import LocalMessageBus

def test_mb():
    mb = LocalMessageBus()

    @mb.handle('test_event')
    def handle(message: Any) -> Any:
        assert message == 'test_message'
        mb.shutdown()
    
    mb.publish('test_event', 'test_message')