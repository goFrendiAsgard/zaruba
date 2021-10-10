from typing import Any, Callable, List, Mapping, TypedDict
from kafka import KafkaProducer, KafkaConsumer
from helpers.transport.interface import MessageBus
from helpers.transport.kafka_config import KafkaEventMap
import threading

def create_kafka_connection_parameters(bootstrap_servers: str, sasl_mechanism: str, sasl_plain_username: str = '', sasl_plain_password: str = '') -> Mapping[str, Any]:
    if sasl_mechanism == '':
        sasl_mechanism = 'PLAIN'
    return {
        'bootstrap_servers': bootstrap_servers,
        'sasl_mechanism': sasl_mechanism,
        'sasl_plain_username': sasl_plain_username,
        'sasl_plain_password': sasl_plain_password
    }

class KafkaMessageBus(MessageBus):

    def __init__(self, kafka_connection_parameters: Mapping[str, Any], kafka_event_map: KafkaEventMap, **kafka_config: Any):
        self.kafka_connection_parameters = kafka_connection_parameters
        self.kafka_event_map = kafka_event_map
        self.consumers: Mapping[str, KafkaConsumer] = {}
        self.event_map = kafka_event_map
    
    def shutdown(self):
        for event_name, consumer in self.consumers.items():
            print('stop listening to {event_name}'.format(event_name=event_name))
            consumer.close()

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            topic = self.event_map.get_topic(event_name)
            group_id = self.event_map.get_group_id(event_name)
            consumer = KafkaConsumer(topic, group_id=group_id, **self.kafka_connection_parameters)
            self.consumers[event_name] = consumer
            thread = threading.Thread(target=self._handle, args=[consumer, event_name, topic, group_id, event_handler])
            thread.start()
        return register_event_handler
    
    def _handle(self, consumer: KafkaConsumer, event_name: str, topic: str, group_id: str, event_handler: Callable[[Any], Any]):
        for serialized_message in consumer:
            message = self.event_map.get_decoder(event_name)(serialized_message.value)
            print({'action': 'handle_kafka_event', 'event_name': event_name, 'message': message, 'topic': topic, 'group_id': group_id, 'serialized': serialized_message})
            event_handler(message)

    def publish(self, event_name: str, message: Any) -> Any:
        producer = KafkaProducer(**self.kafka_connection_parameters)
        topic = self.event_map.get_topic(event_name)
        serialized_message = self.event_map.get_encoder(event_name)(message)
        print({'action': 'publish_kafka_event', 'event_name': event_name, 'message': message, 'topic': topic, 'serialized': serialized_message})
        producer.send(topic, serialized_message)
        producer.close()
