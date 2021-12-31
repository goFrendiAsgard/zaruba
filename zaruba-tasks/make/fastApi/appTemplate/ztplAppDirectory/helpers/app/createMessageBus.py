from typing import Mapping, Any
from helpers.transport import MessageBus,KafkaMessageBus, KafkaEventMap, RMQMessageBus, RMQRPC, RMQEventMap, LocalMessageBus
import pika

def create_message_bus(mb_type: str, rmq_connection_parameters: pika.ConnectionParameters, rmq_event_map: RMQEventMap, kafka_connection_parameters: Mapping[str, Any], kafka_event_map: KafkaEventMap) -> MessageBus:
    if mb_type == 'rmq':
        return RMQMessageBus(rmq_connection_parameters, rmq_event_map)
    if mb_type == 'kafka':
        return KafkaMessageBus(kafka_connection_parameters, kafka_event_map)
    return LocalMessageBus()
