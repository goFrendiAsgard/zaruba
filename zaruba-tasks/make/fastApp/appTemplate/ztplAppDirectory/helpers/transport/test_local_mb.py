from typing import Any
from helpers.transport.local_mb import LocalMessageBus

import asyncio

def test_local_mb():
    asyncio.run(_test_local_mb())


async def _test_local_mb():
    mb = LocalMessageBus()

    result = {}
    @mb.handle('test_event')
    def handle(message: Any) -> Any:
        result['message'] = message
        mb.shutdown()
    
    mb.publish('test_event', 'test_message')
    await asyncio.sleep(1)
    assert 'message' in result
    assert result['message'] == 'test_message'