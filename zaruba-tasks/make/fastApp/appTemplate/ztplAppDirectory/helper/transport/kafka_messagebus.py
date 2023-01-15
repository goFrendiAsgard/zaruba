from typing import Any, Callable, Mapping
from confluent_kafka import Producer, Consumer
from helper.transport.messagebus import MessageBus
from helper.transport.kafka_helper import create_kafka_topic
from helper.transport.kafka_config import KafkaEventMap

import threading
import logging


def create_kafka_connection_parameters(
    bootstrap_servers: str,
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
        'sasl.mechanism': sasl_mechanism,
        'sasl.username': sasl_plain_username,
        'sasl.password': sasl_plain_password,
        'security.protocol': security_protocol,
        # 'topic.metadata.propagation.max.ms': '100',
        **kwargs
    }


class KafkaMessageBus(MessageBus):

    def __init__(
        self, kafka_connection_parameters: Mapping[str, Any],
        kafka_event_map: KafkaEventMap
    ):
        self._kafka_connection_parameters = kafka_connection_parameters
        self._kafka_event_map = kafka_event_map
        self._consumers: Mapping[str, Consumer] = {}
        self._event_map = kafka_event_map
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
                consumer_args = {**self._kafka_connection_parameters}
                consumer_args['group.id'] = group_id
                # start consume
                thread = threading.Thread(
                    target=self._handle,
                    args=[
                        consumer_args,
                        event_name,
                        topic,
                        group_id,
                        event_handler
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
                    logging.error('Consumer error', exc_info=True)
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
                consumer = Consumer(consumer_args)
                self._consumers[event_name] = consumer
                logging.info('Re subscribe to topic {}'.format(topic))
                consumer.subscribe([topic])

    def _create_consumer(
        self,
        topic: str,
        consumer_args: Any
    ) -> Consumer:
        attempt, limit_attempt = 1, 3
        while attempt < limit_attempt:
            consumer = Consumer(consumer_args)
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
            producer = Producer(self._kafka_connection_parameters)
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
            ' Topic: {}'.format(topic),
            ' Key: {}'.format(key),
            ' Message: {}'.format(message),
            ' Serialized message: {}'.format(serialized_message)
        ]))

    def _create_kafka_topic(self, topic: str):
        connection_param = self._kafka_connection_parameters
        create_kafka_topic(topic, {
            'bootstrap.servers': connection_param['bootstrap.servers'],
            'sasl.mechanism': connection_param['sasl.mechanism'],
            'sasl.username': connection_param['sasl.username'],
            'sasl.password': connection_param['sasl.password']
        })


def _produce_callback(err, msg):
    if err is not None:
        logging.error(
            'Failed to deliver message {}'.format(str(msg)), exc_info=True
        )
        return
    logging.info('Message produced {}'.format(str(msg)))
