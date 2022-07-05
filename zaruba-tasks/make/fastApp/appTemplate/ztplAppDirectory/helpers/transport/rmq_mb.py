from typing import Any, Callable
from helpers.transport.interface import MessageBus
from helpers.transport.rmq_connection import RMQConnection
from helpers.transport.rmq_config import RMQEventMap

import pika
import threading
import traceback

class RMQMessageBus(RMQConnection, MessageBus):

    def __init__(self, rmq_connection_parameters: pika.ConnectionParameters, rmq_event_map: RMQEventMap):
        RMQConnection.__init__(self, rmq_connection_parameters)
        self.event_map=rmq_event_map
        self.error_count = 0

    def get_error_count(self) -> int:
        return self.error_count

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            exchange = self.event_map.get_exchange_name(event_name)
            queue = self.event_map.get_queue_name(event_name)
            dead_letter_exchange = self.event_map.get_dead_letter_exchange(event_name)
            dead_letter_queue = self.event_map.get_dead_letter_queue(event_name)
            auto_ack = self.event_map.get_auto_ack(event_name)
            arguments = self.event_map.get_queue_arguments(event_name)
            prefetch_count = self.event_map.get_prefetch_count(event_name)
            def consume():
                ch = self.connection.channel()
                if self.event_map.get_ttl(event_name) > 0:
                    ch.exchange_declare(exchange=dead_letter_exchange, exchange_type='fanout', durable=True)
                    ch.queue_declare(queue=dead_letter_queue, durable=True, exclusive=False)
                    ch.queue_bind(exchange=dead_letter_exchange, queue=dead_letter_queue)
                ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
                ch.queue_declare(queue=queue, exclusive=False, durable=True, arguments=arguments)
                ch.queue_bind(exchange=exchange, queue=queue)
                ch.basic_qos(prefetch_count=prefetch_count)
                # create handler and start consuming
                on_event = self._create_event_handler(event_name, exchange, queue, auto_ack, event_handler)
                ch.basic_consume(queue=queue, on_message_callback=on_event, auto_ack=auto_ack)
            thread = threading.Thread(target=consume, args=[], daemon = True)
            thread.start()
        return register_event_handler

    def _create_event_handler(self, event_name: str, exchange: str, queue: str, auto_ack: bool, event_handler: Callable[[Any], Any]):
        def on_event(ch, method, props, body):
            try:
                message = self.event_map.get_decoder(event_name)(body)
                print({'action': 'handle_rmq_event', 'event_name': event_name, 'message': message, 'exchange': exchange, 'routing_key': queue})
                event_handler(message)
            except Exception as e:
                self.error_count += 1
                print(traceback.format_exc()) 
            finally:
                if not auto_ack:
                    ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_event

    def publish(self, event_name: str, message: Any) -> Any:
        try:
            exchange = self.event_map.get_exchange_name(event_name)
            routing_key = self.event_map.get_queue_name(event_name)
            body = self.event_map.get_encoder(event_name)(message)
            ch = self.connection.channel()
            ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
            print({'action': 'publish_rmq_event', 'event_name': event_name, 'message': message, 'exchange': exchange, 'routing_key': routing_key, 'body': body})
            ch.basic_publish(
                exchange=exchange,
                routing_key=routing_key,
                body=body
            )
        except Exception as e:
            self.error_count += 1
            raise e
