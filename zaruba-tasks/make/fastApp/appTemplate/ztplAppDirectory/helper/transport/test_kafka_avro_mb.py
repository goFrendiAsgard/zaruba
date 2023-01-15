from typing import Any
from helper.transport.kafka_avro_messagebus import (
    KafkaAvroMessageBus, create_kafka_avro_connection_parameters
)
from helper.transport.kafka_avro_config import KafkaAvroEventMap
from helper.config.boolean import get_boolean_env

import os
import logging
import asyncio
import pytest


@pytest.mark.asyncio
async def test_kafka_avro_mb():
    if not get_boolean_env('TEST_INTEGRATION', False):
        logging.warn('KafkaAvroMessageBus is not tested')
        return None

    kafka_connection_parameters = create_kafka_avro_connection_parameters(
        bootstrap_servers=os.getenv(
            'TEST_KAFKA_BOOTSTRAP_SERVERS', 'localhost:9092'
        ),
        schema_registry=os.getenv(
            'TEST_KAFKA_SCHEMA_REGISTRY', 'http://localhost:8035'
        ),
        sasl_mechanism=os.getenv('TEST_KAFKA_SASL_MECHANISM', 'PLAIN'),
        sasl_plain_username=os.getenv('TEST_KAFKA_SASL_PLAIN_USERNAME', ''),
        sasl_plain_password=os.getenv('TEST_KAFKA_SASL_PLAIN_PASSWORD', ''),
        security_protocol=os.getenv(
            'TEST_KAFKA_SECURITY_PROTOCOL', 'PLAINTEXT'
        )
    )
    kafka_avro_event_map = KafkaAvroEventMap({})

    mb = KafkaAvroMessageBus(kafka_connection_parameters, kafka_avro_event_map)
    # await asyncio.sleep(3)

    result = {}
    try:
        @mb.handle('test_avro_event')
        def handle(message: Any) -> Any:
            result['message'] = message

        await asyncio.sleep(3)
        mb.publish('test_avro_event', 'test_avro_message')

        trial: int = 10
        while trial > 0 and 'message' not in result:
            await asyncio.sleep(1)
            trial -= 1
    finally:
        mb.shutdown()
    assert 'message' in result
    assert result['message'] == 'test_avro_message'
