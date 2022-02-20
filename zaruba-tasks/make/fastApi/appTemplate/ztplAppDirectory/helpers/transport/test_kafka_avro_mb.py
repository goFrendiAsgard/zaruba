from typing import Any
from helpers.transport.kafka_avro_mb import KafkaAvroMessageBus, create_kafka_avro_connection_parameters
from helpers.transport.kafka_avro_config import KafkaAvroEventMap

import os
import warnings

def test_mb():
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
    kafka_event_map = KafkaAvroEventMap({
        'test_avro_event': {
            'value_schema_str': '{"type": "string"}'
        }
    })

    mb = KafkaAvroMessageBus(kafka_connection_parameters, kafka_event_map)

    @mb.handle('test_avro_event')
    def handle(message: Any) -> Any:
        assert message == 'test_avro_message'
        mb.shutdown()
    
    mb.publish('test_avro_event', 'test_avro_message')

