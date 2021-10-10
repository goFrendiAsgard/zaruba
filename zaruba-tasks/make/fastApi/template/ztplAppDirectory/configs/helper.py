from typing import Mapping, Any
from helpers.transport import MessageBus, RPC, KafkaMessageBus, KafkaEventMap, RMQMessageBus, RMQRPC, RMQEventMap, LocalMessageBus, LocalRPC, create_kafka_connection_parameters, create_rmq_connection_parameters
import os, pika

def get_abs_static_dir(raw_static_dir: str) -> str:
    return os.path.abspath(raw_static_dir) if raw_static_dir != '' else ''

def get_rmq_connection_parameters() -> pika.ConnectionParameters:
    return create_rmq_connection_parameters(
        host = os.getenv('APP_RABBITMQ_HOST', 'localhost'),
        user = os.getenv('APP_RABBITMQ_USER', 'root'),
        password = os.getenv('APP_RABBITMQ_PASS', 'toor'),
        virtual_host = os.getenv('APP_RABBITMQ_VHOST', '/'),
        heartbeat=30
    )

def get_kafka_connection_parameters() -> Mapping[str, Any]:
    return create_kafka_connection_parameters(
        bootstrap_servers = os.getenv('APP_KAFKA_BOOTSTRAP_SERVERS', 'localhost:9093'),
        sasl_mechanism=os.getenv('APP_KAFKA_SASL_MECHANISM', 'PLAIN'),
        sasl_plain_username=os.getenv('APP_KAFKA_SASL_PLAIN_USERNAME', ''),
        sasl_plain_password=os.getenv('APP_KAFKA_SASL_PLAIN_PASSWORD', '')
    )

def create_message_bus(mb_type: str, rmq_connection_parameters: pika.ConnectionParameters, rmq_event_map: RMQEventMap, kafka_connection_parameters: Mapping[str, Any], kafka_event_map: KafkaEventMap) -> MessageBus:
    if mb_type == 'rmq':
        return RMQMessageBus(rmq_connection_parameters, rmq_event_map)
    if mb_type == 'kafka':
        return KafkaMessageBus(kafka_connection_parameters, kafka_event_map)
    return LocalMessageBus()

def create_rpc(mb_type: str, rmq_connection_parameters: pika.ConnectionParameters, rmq_event_map: RMQEventMap) -> RPC:
    if mb_type == 'rmq':
        return RMQRPC(rmq_connection_parameters, rmq_event_map)
    return LocalRPC()

