from typing import List
from helper.transport import MessageBus, KafkaMessageBus, KafkaAvroMessageBus, RMQMessageBus, LocalMessageBus
from config.rmq import rmq_connection_parameters, rmq_event_map
from config.kafka import kafka_avro_connection_parameters, kafka_avro_event_map, kafka_connection_parameters, kafka_event_map
from transport import AppMessageBus


def create_message_bus(mb_type: str, activity_events: List[str]) -> AppMessageBus:
    if mb_type == 'rmq':
        return AppMessageBus(RMQMessageBus(rmq_connection_parameters, rmq_event_map), activity_events)
    if mb_type == 'kafka':
        return AppMessageBus(KafkaMessageBus(kafka_connection_parameters, kafka_event_map), activity_events)
    if mb_type == 'kafkaAvro':
        return AppMessageBus(KafkaAvroMessageBus(kafka_avro_connection_parameters, kafka_avro_event_map), activity_events)
    return AppMessageBus(LocalMessageBus(), activity_events)
