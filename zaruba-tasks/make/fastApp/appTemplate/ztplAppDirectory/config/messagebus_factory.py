from typing import List
from helper.transport import (
    KafkaMessageBus, KafkaAvroMessageBus,
    RMQMessageBus, LocalMessageBus
)
from config.rmq import rmq_connection_parameters, rmq_event_map
from config.kafka import (
    kafka_avro_connection_parameters, kafka_avro_event_map,
    kafka_connection_parameters, kafka_event_map
)
from transport import AppMessageBus


def create_app_message_bus(
    mb_type: str, activity_events: List[str]
) -> AppMessageBus:
    if mb_type == 'rmq':
        rmq_message_bus = RMQMessageBus(
            rmq_connection_parameters, rmq_event_map
        )
        return AppMessageBus(rmq_message_bus, activity_events)
    if mb_type == 'kafka':
        kafka_message_bus = KafkaMessageBus(
            kafka_connection_parameters, kafka_event_map
        )
        return AppMessageBus(kafka_message_bus, activity_events)
    if mb_type == 'kafkaAvro':
        kafka_avro_message_bus = KafkaAvroMessageBus(
            kafka_avro_connection_parameters, kafka_avro_event_map
        )
        return AppMessageBus(kafka_avro_message_bus, activity_events)
    return AppMessageBus(LocalMessageBus(), activity_events)
