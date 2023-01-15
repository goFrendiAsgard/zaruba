from typing import List
from helper.transport import (
    MessageBus, KafkaMessageBus, KafkaAvroMessageBus,
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
    '''
    Return a new AppMessageBus based on specified mb_type.

    Keyword arguments:
    - mb_type -- Messagebus type (e.g., rmq, kafka, kafkaavro, local).
    - activity_events -- List of event name to be triggered
        whenever AppMessageBus publish any message (default: []).
    '''
    mb = create_message_bus(mb_type)
    return AppMessageBus(mb, activity_events)


def create_message_bus(mb_type: str) -> MessageBus:
    '''
    Return a new MessageBus based on specified mb_type.

    Arguments:
    mb_type -- Messagebus type (e.g., rmq, kafka, kafkaavro).
    '''
    if mb_type.lower() == 'rmq':
        return RMQMessageBus(rmq_connection_parameters, rmq_event_map)
    if mb_type.lower() == 'kafka':
        return KafkaMessageBus(
            kafka_connection_parameters, kafka_event_map
        )
    if mb_type.lower() == 'kafkaavro':
        return KafkaAvroMessageBus(
            kafka_avro_connection_parameters, kafka_avro_event_map
        )
    return LocalMessageBus()
