from helpers.transport import MessageBus, KafkaMessageBus, KafkaAvroMessageBus, RMQMessageBus, LocalMessageBus
from configs.rmq import rmq_connection_parameters, rmq_event_map
from configs.kafka import kafka_avro_connection_parameters, kafka_avro_event_map, kafka_connection_parameters, kafka_event_map


def create_message_bus(mb_type: str) -> MessageBus:
    if mb_type == 'rmq':
        return RMQMessageBus(rmq_connection_parameters, rmq_event_map)
    if mb_type == 'kafka':
        return KafkaMessageBus(kafka_connection_parameters, kafka_event_map)
    if mb_type == 'kafkaAvro':
        return KafkaAvroMessageBus(kafka_avro_connection_parameters, kafka_avro_event_map)
    return LocalMessageBus()
