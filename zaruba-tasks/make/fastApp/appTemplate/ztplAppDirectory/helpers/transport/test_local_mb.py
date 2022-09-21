from typing import Any
from helpers.transport.localMessagebus import LocalMessageBus

import asyncio
import pytest


@pytest.mark.asyncio
async def test_local_mb():
    mb = LocalMessageBus()

    result = {}
    @mb.handle('test_event')
    def handle(message: Any) -> Any:
        result['message'] = message
    
    mb.publish('test_event', 'test_message')
    await asyncio.sleep(1)
    mb.shutdown()
    assert 'message' in result
    assert result['message'] == 'test_message'