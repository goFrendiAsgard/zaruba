from typing import Any
from helpers.transport.kafka_mb import KafkaMessageBus, create_kafka_connection_parameters
from helpers.transport.kafka_config import KafkaEventMap

import os

def test_mb():
    bootstrap_servers = os.getenv('TEST_KAFKA_BOOTSTRAP_SERVERS', 'localhost:9093')
    kafka_connection_parameters = create_kafka_connection_parameters(bootstrap_servers)
    kafka_event_map = KafkaEventMap({})

    mb = KafkaMessageBus(kafka_connection_parameters, kafka_event_map)

    @mb.handle('test_event')
    def handle(message: Any) -> Any:
        assert message == 'test_message'
        mb.shutdown()
    
    mb.publish('test_event', 'test_message')

