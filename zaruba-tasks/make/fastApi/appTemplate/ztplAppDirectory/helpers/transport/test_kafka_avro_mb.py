from typing import Any
from helpers.transport.kafka_avro_mb import KafkaAvroMessageBus, create_kafka_avro_connection_parameters
from helpers.transport.kafka_avro_config import KafkaAvroEventMap

import os
import warnings
import asyncio


def test_kafka_avro_mb():
    asyncio.run(_test_kafka_avro_mb())


async def _test_kafka_avro_mb():
    if os.getenv('TEST_INTEGRATION', '0') != '1':
        warnings.warn(UserWarning('TEST_INTEGRATION != 1, KafkaAvroMessageBus is not tested'))
        return None

    kafka_connection_parameters = create_kafka_avro_connection_parameters(
        bootstrap_servers = os.getenv('TEST_KAFKA_BOOTSTRAP_SERVERS', 'localhost:9092'),
        schema_registry = os.getenv('TEST_KAFKA_SCHEMA_REGISTRY', 'http://localhost:8035'),
        sasl_mechanism=os.getenv('TEST_KAFKA_SASL_MECHANISM', 'PLAIN'),
        sasl_plain_username=os.getenv('TEST_KAFKA_SASL_PLAIN_USERNAME', ''),
        sasl_plain_password=os.getenv('TEST_KAFKA_SASL_PLAIN_PASSWORD', '')
    )
    kafka_avro_event_map = KafkaAvroEventMap({})
    
    mb = KafkaAvroMessageBus(kafka_connection_parameters, kafka_avro_event_map)

    result = {}
    @mb.handle('test_avro_event')
    def handle(message: Any) -> Any:
        result['message'] = message
        mb.shutdown()

    mb.publish('test_avro_event', 'test_avro_message')
    await asyncio.sleep(5)
    assert 'message' in result
    assert result['message'] == 'test_avro_message'