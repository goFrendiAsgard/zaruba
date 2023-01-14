from typing import Any, Callable, Mapping
from helper.transport.messagebus import MessageBus
from helper.transport.kafka_helper import create_kafka_topic
from helper.transport.kafka_avro_config import KafkaAvroEventMap
from confluent_kafka.avro import AvroProducer, AvroConsumer

import logging
import threading


def create_kafka_avro_connection_parameters(
    bootstrap_servers: str,
    schema_registry: str = '',
    sasl_mechanism: str = '',
    sasl_plain_username: str = '',
    sasl_plain_password: str = '',
    security_protocol='', **kwargs
) -> Mapping[str, Any]:
    if sasl_mechanism == '':
        sasl_mechanism = 'PLAIN'
    if security_protocol == '':
        security_protocol = 'PLAINTEXT'
    # https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
    return {
        'bootstrap.servers': bootstrap_servers,
        'schema.registry.url': schema_registry,
        'sasl.mechanism': sasl_mechanism,
        'sasl.username': sasl_plain_username,
        'sasl.password': sasl_plain_password,
        'security.protocol': security_protocol,
        # 'topic.metadata.propagation.max.ms': '100',
        **kwargs
    }


class KafkaAvroMessageBus(MessageBus):

    def __init__(
        self, connection_parameters: Mapping[str, Any],
        event_map: KafkaAvroEventMap
    ):
        self._connection_parameters = connection_parameters
        self._event_map = event_map
        self._consumers: Mapping[str, AvroConsumer] = {}
        self._event_map = event_map
        self._is_shutdown = False
        self._error_count = 0
        self._is_failing = False

    def get_error_count(self) -> int:
        return self._error_count

    def is_failing(self) -> bool:
        return self._is_failing

    def shutdown(self):
        if self._is_shutdown:
            return
        self._is_shutdown = True
        for event_name, consumer in self._consumers.items():
            logging.info('Stop listening to {}'.format(event_name))
            consumer.close()

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            try:
                topic = self._event_map.get_topic(event_name)
                # create topic if not exists
                self._create_kafka_topic(topic)
                # create consumer
                group_id = self._event_map.get_group_id(event_name)
                consumer_args = {**self._connection_parameters}
                consumer_args['group.id'] = group_id
                # start consume
                thread = threading.Thread(
                    target=self._handle,
                    args=[
                        consumer_args, event_name,
                        topic, group_id, event_handler
                    ],
                    daemon=True
                )
                thread.start()
            except Exception:
                logging.error(
                    'Cannot register to event {}'.format(event_name),
                    exc_info=True
                )
                self._is_failing = True
                self.shutdown()
        return register_event_handler

    def _handle(
        self,
        consumer_args,
        event_name: str,
        topic: str,
        group_id: str,
        event_handler: Callable[[Any], Any]
    ):
        consumer = self._create_consumer(topic, consumer_args)
        self._consumers[event_name] = consumer
        while not self._is_shutdown:
            try:
                serialized_message = consumer.poll(1)
                if serialized_message is None:
                    continue
                if serialized_message.error():
                    logging.error('AvroConsumer error', exc_info=True)
                    continue
                if serialized_message.value():
                    self._log_event_handling(
                        event_name, group_id, serialized_message
                    )
                    message = self._event_map.get_decoder(event_name)(
                        serialized_message.value()
                    )
                    event_handler(message)
            except Exception:
                logging.error(
                    'Error while handle event {}'.format(event_name),
                    exc_info=True
                )
                self._error_count += 1
                consumer = AvroConsumer(consumer_args)
                self._consumers[event_name] = consumer
                logging.info('Re subscribe to topic {}'.format(topic))
                consumer.subscribe([topic])

    def _create_consumer(
        self,
        topic: str,
        consumer_args: Any
    ) -> AvroConsumer:
        attempt, limit_attempt = 1, 3
        while attempt < limit_attempt:
            consumer = AvroConsumer(consumer_args)
            try:
                logging.info('Subscribe to topic {} (attempt {} of {})'.format(
                    attempt, limit_attempt, topic))
                consumer.subscribe([topic])
                return consumer
            except Exception:
                logging.error(
                    'Cannot subscribe to topic {} (attempt {} of {})'.format(
                        attempt, limit_attempt, topic
                    ),
                    exc_info=True
                )
            attempt += 1
        raise Exception('Cannot create consumer for topic {}'.format(topic))

    def _log_event_handling(
        self,  event_name: str, group_id: str, serialized_message: Any
    ):
        logging.info(' '.join([
            'Handle event {}'.format(event_name),
            'Topic: {}'.format(serialized_message.topic()),
            'Key: {}'.format(serialized_message.key()),
            'Value: {}'.format(serialized_message.value()),
            'Offset: {}'.format(serialized_message.offset()),
            'Group Id: {}'.format(group_id)
        ]))

    def publish(self, event_name: str, message: Any) -> Any:
        serialized_message = self._event_map.get_encoder(event_name)(message)
        try:
            topic = self._event_map.get_topic(event_name)
            key_schema = self._event_map.get_key_schema(event_name)
            value_schema = self._event_map.get_value_schema(event_name)
            producer = AvroProducer(
                self._connection_parameters,
                default_key_schema=key_schema,
                default_value_schema=value_schema
            )
            key_maker = self._event_map.get_key_maker(event_name)
            key = key_maker(message)
            self._log_event_publish(
                event_name, topic, key, message, serialized_message
            )
            producer.produce(
                topic=topic,
                key=key,
                value=serialized_message,
                callback=_produce_callback
            )
            producer.flush()
        except Exception as exception:
            logging.error(
                'Error publishing event {} with message: {}'.format(
                    event_name, message
                )
            )
            self._error_count += 1
            raise exception

    def _log_event_publish(
        self, event_name: str,  topic: str, key: Any, message: Any,
        serialized_message: Any
    ):
        logging.info(' '.join([
            'Publish event {}'.format(event_name),
            'Topic: {}'.format(topic),
            'Key: {}'.format(key),
            'Message: {}'.format(message),
            'Serialized message: {}'.format(serialized_message)
        ]))

    def _create_kafka_topic(self, topic: str):
        create_kafka_topic(topic, {
            'bootstrap.servers':
                self._connection_parameters['bootstrap.servers'],
            'sasl.mechanism': self._connection_parameters['sasl.mechanism'],
            'sasl.username': self._connection_parameters['sasl.username'],
            'sasl.password': self._connection_parameters['sasl.password']
        })


def _produce_callback(err, msg):
    if err is not None:
        logging.error(
            'Failed to deliver message {}'.format(str(msg)), exc_info=True
        )
        return
    logging.info('Message produced {}'.format(str(msg)))
