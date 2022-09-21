from typing import Any
from helpers.transport.kafkaMessagebus import KafkaMessageBus, create_kafka_connection_parameters
from helpers.transport.kafkaConfig import KafkaEventMap

import os
import warnings
import asyncio
import pytest


@pytest.mark.asyncio
async def test_kafka_mb():
    if os.getenv('TEST_INTEGRATION', '0') != '1':
        warnings.warn(UserWarning('TEST_INTEGRATION != 1, KafkaMessageBus is not tested'))
        return None

    kafka_connection_parameters = create_kafka_connection_parameters(
        bootstrap_servers = os.getenv('TEST_KAFKA_BOOTSTRAP_SERVERS', 'localhost:9092'),
        sasl_mechanism=os.getenv('TEST_KAFKA_SASL_MECHANISM', 'PLAIN'),
        sasl_plain_username=os.getenv('TEST_KAFKA_SASL_PLAIN_USERNAME', ''),
        sasl_plain_password=os.getenv('TEST_KAFKA_SASL_PLAIN_PASSWORD', ''),
        security_protocol=os.getenv('TEST_KAFKA_SECURITY_PROTOCOL', 'PLAINTEXT')
    )
    kafka_event_map = KafkaEventMap({})

    mb = KafkaMessageBus(kafka_connection_parameters, kafka_event_map)
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
