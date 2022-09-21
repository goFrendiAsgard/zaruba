from helpers.transport import KafkaEventMap, KafkaAvroEventMap, create_kafka_connection_parameters, create_kafka_avro_connection_parameters

import os

################################################
# -- 🪠 Kafka setting
################################################

kafka_connection_parameters = create_kafka_connection_parameters(
    bootstrap_servers = os.getenv('APP_KAFKA_BOOTSTRAP_SERVERS', 'localhost:29092'),
    sasl_mechanism = os.getenv('APP_KAFKA_SASL_MECHANISM', 'PLAIN'),
    sasl_plain_username = os.getenv('APP_KAFKA_SASL_PLAIN_USERNAME', ''),
    sasl_plain_password = os.getenv('APP_KAFKA_SASL_PLAIN_PASSWORD', ''),
    security_protocol = os.getenv('APP_KAFKA_SECURITY_PROTOCOL', 'PLAINTEXT')
)

kafka_event_map = KafkaEventMap({})


################################################
# -- 🪠 Kafka avro setting
################################################

kafka_avro_connection_parameters = create_kafka_avro_connection_parameters(
    bootstrap_servers = os.getenv('APP_KAFKA_BOOTSTRAP_SERVERS', 'localhost:29092'),
    schema_registry = os.getenv('APP_KAFKA_SCHEMA_REGISTRY', 'http://localhost:8035'),
    sasl_mechanism = os.getenv('APP_KAFKA_SASL_MECHANISM', 'PLAIN'),
    sasl_plain_username = os.getenv('APP_KAFKA_SASL_PLAIN_USERNAME', ''),
    sasl_plain_password = os.getenv('APP_KAFKA_SASL_PLAIN_PASSWORD', ''),
    security_protocol = os.getenv('APP_KAFKA_SECURITY_PROTOCOL', 'PLAINTEXT')
)

kafka_avro_event_map = KafkaAvroEventMap({})

