from typing import Any, Callable, List, Mapping, TypedDict
from helpers.transport.interface import MessageBus
from helpers.transport.kafka_avro_config import KafkaAvroEventMap
from confluent_kafka.avro import AvroProducer, AvroConsumer
from confluent_kafka.admin import AdminClient, NewTopic

import threading
import traceback

def create_kafka_avro_connection_parameters(bootstrap_servers: str, schema_registry: str = '', sasl_mechanism: str = '', sasl_plain_username: str = '', sasl_plain_password: str = '', security_protocol='', **kwargs) -> Mapping[str, Any]:
    if sasl_mechanism == '':
        sasl_mechanism = 'PLAIN'
    if security_protocol == '':
        security_protocol = 'SASL_PLAINTEXT'
    return {
        'bootstrap.servers': bootstrap_servers,
        'schema.registry.url': schema_registry,
        'sasl.mechanism': sasl_mechanism,
        'sasl.username': sasl_plain_username,
        'sasl.password': sasl_plain_password,
        # 'security.protocol': security_protocol,
        **kwargs
    }

class KafkaAvroMessageBus(MessageBus):

    def __init__(self, kafka_avro_connection_parameters: Mapping[str, Any], kafka_avro_event_map: KafkaAvroEventMap):
        self.kafka_avro_connection_parameters = kafka_avro_connection_parameters
        self.kafka_avro_event_map = kafka_avro_event_map
        self.consumers: Mapping[str, AvroConsumer] = {}
        self.event_map = kafka_avro_event_map
        self.is_shutdown = False

    def shutdown(self):
        for event_name, consumer in self.consumers.items():
            print('stop listening to {event_name}'.format(event_name=event_name))
            consumer.close()
        self.is_shutdown = True

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            topic = self.event_map.get_topic(event_name)
            # create topic if not exists
            try:
                kafka_admin = AdminClient({
                    'bootstrap.servers': self.kafka_avro_connection_parameters['bootstrap.servers'],
                    'sasl.mechanism': self.kafka_avro_connection_parameters['sasl.mechanism'],
                    'sasl.username': self.kafka_avro_connection_parameters['sasl.username'],
                    'sasl.password': self.kafka_avro_connection_parameters['sasl.password']
                })
                topic_metadata = kafka_admin.list_topics()
                if topic_metadata.topics.get(topic) is None:
                    kafka_admin.create_topics([NewTopic(topic, 1, 1)])
            except:
                print(traceback.format_exc())
            # create consumer
            group_id = self.event_map.get_group_id(event_name)
            consumer_args = {**self.kafka_avro_connection_parameters}
            consumer_args['group.id'] = group_id
            consumer = AvroConsumer(consumer_args)
            # register consumer
            self.consumers[event_name] = consumer
            # start consume
            thread = threading.Thread(target=self._handle, args=[consumer, event_name, topic, group_id, event_handler], daemon = True)
            thread.start()
        return register_event_handler

    def _handle(self, consumer: AvroConsumer, event_name: str, topic: str, group_id: str, event_handler: Callable[[Any], Any]):
        consumer.subscribe([topic])
        while True:
            try:
                message = consumer.poll(1)
                if message is None:
                    continue
                if message.error():
                    print("AvroConsumer error: {}".format(message.error()))
                    continue
                if message.value():
                    print({'action': 'handle_kafka_avro_event', 'event_name': event_name, 'value': message.value(), 'key': message.key(), 'topic': message.topic(), 'partition': message.partition(), 'offset': message.offset(), 'group_id': group_id})
                    event_handler(message.value())
            except:
                print(traceback.format_exc())

    def publish(self, event_name: str, message: Any) -> Any:
        key_schema = self.event_map.get_key_schema(event_name)
        value_schema = self.event_map.get_value_schema(event_name)
        producer = AvroProducer(self.kafka_avro_connection_parameters, default_key_schema=key_schema, default_value_schema=value_schema)
        topic = self.event_map.get_topic(event_name)
        key_maker = self.event_map.get_key_maker(event_name)
        key = key_maker(message)
        print({'action': 'publish_kafka_avro_event', 'event_name': event_name, 'key': key, 'message': message, 'topic': topic})
        producer.produce(topic=topic, key=key, value=message, callback=_produce_callback)
        producer.flush()


def _produce_callback(err, msg):
    if err is not None:
        print("Failed to deliver message: %s: %s" % (str(msg), str(err)))
    else:
        print("Message produced: %s" % (str(msg)))