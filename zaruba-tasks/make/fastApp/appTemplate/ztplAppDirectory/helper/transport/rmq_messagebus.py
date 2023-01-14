from typing import Any, Callable
from pika.adapters.blocking_connection import (
    BlockingConnection, BlockingChannel
)
from helper.transport.messagebus import MessageBus
from helper.transport.rmq_connection import RMQConnection
from helper.transport.rmq_config import RMQEventMap

import pika
import threading
import logging


class RMQMessageBus(RMQConnection, MessageBus):

    def __init__(
        self,
        rmq_connection_parameters: pika.ConnectionParameters,
        rmq_event_map: RMQEventMap
    ):
        RMQConnection.__init__(self, rmq_connection_parameters)
        self._event_map = rmq_event_map
        self._error_count = 0

    def get_error_count(self) -> int:
        return self._error_count

    def handle(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            consume = self._create_consumer(event_name, event_handler)
            thread = threading.Thread(target=consume)
            thread.start()
        return register_event_handler

    def _create_consumer(
        self,
        event_name: str,
        event_handler: Callable[[Any], Any]
    ) -> Callable[[Any], Any]:
        def consume():
            exchange = self._event_map.get_exchange_name(event_name)
            queue = self._event_map.get_queue_name(event_name)
            auto_ack = self._event_map.get_auto_ack(event_name)
            while not self._is_shutdown:
                connection: BlockingConnection = None
                try:
                    connection = self.create_connection()
                    ch = self._create_consumer_channel(connection, event_name)
                    # create handler and start consuming
                    ch.basic_consume(
                        queue=queue,
                        on_message_callback=self._create_event_handler(
                            event_name, exchange, queue, auto_ack,
                            event_handler
                        ),
                        auto_ack=auto_ack
                    )
                    ch.start_consuming()
                except Exception:
                    logging.error(
                        'Cannot consume from queue {}'.format(queue),
                        exc_info=True
                    )
                    self.remove_connection(connection)
                    self._error_count += 1
        return consume

    def _create_consumer_channel(
        self, connection: BlockingConnection, event_name: str
    ) -> BlockingChannel:
        exchange = self._event_map.get_exchange_name(event_name)
        queue = self._event_map.get_queue_name(event_name)
        dlx = self._event_map.get_dead_letter_exchange(
            event_name
        )
        dlq = self._event_map.get_dead_letter_queue(
            event_name
        )
        arguments = self._event_map.get_queue_arguments(event_name)
        prefetch_count = self._event_map.get_prefetch_count(event_name)
        ttl = self._event_map.get_ttl(event_name)
        ch = connection.channel()
        if ttl > 0:
            ch.exchange_declare(
                exchange=dlx, exchange_type='fanout', durable=True
            )
            ch.queue_declare(queue=dlq, durable=True, exclusive=False)
            ch.queue_bind(exchange=dlx, queue=dlq)
        ch.exchange_declare(
            exchange=exchange, exchange_type='fanout', durable=True
        )
        ch.queue_declare(
            queue=queue, exclusive=False, durable=True, arguments=arguments
        )
        ch.queue_bind(exchange=exchange, queue=queue)
        ch.basic_qos(prefetch_count=prefetch_count)
        return ch

    def _create_event_handler(
        self, event_name: str, exchange: str, queue: str, auto_ack: bool,
        event_handler: Callable[[Any], Any]
    ):
        def on_event(ch, method, props, body):
            try:
                message = self._event_map.get_decoder(event_name)(body)
                self._log_event_handling(event_name, message, exchange, queue)
                event_handler(message)
            except Exception:
                logging.error(
                    'Cannot handle event {}'.format(event_name), exc_info=True
                )
                self._error_count += 1
            finally:
                if not auto_ack:
                    ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_event

    def _log_event_handling(
        self, event_name: str, message: Any, exchange: str, routing_key: str
    ):
        logging.info(' '.join([
            'Handle event {}'.format(event_name),
            'Message: {}'.format(message),
            'Exchange: {}'.format(exchange),
            'Routing key: {}'.format(routing_key)
        ]))

    def publish(self, event_name: str, message: Any) -> Any:
        try:
            connection = self.create_connection()
            exchange = self._event_map.get_exchange_name(event_name)
            routing_key = self._event_map.get_queue_name(event_name)
            body = self._event_map.get_encoder(event_name)(message)
            ch = connection.channel()
            ch.exchange_declare(
                exchange=exchange, exchange_type='fanout', durable=True
            )
            self._log_event_publish(
                event_name, message, exchange, routing_key, body
            )
            ch.basic_publish(
                exchange=exchange,
                routing_key=routing_key,
                body=body
            )
            self.remove_connection(connection)
        except Exception as exception:
            logging.error(
                'Error publishing event {} with message: {}'.format(
                    event_name, message
                )
            )
            self._error_count += 1
            raise exception

    def _log_event_publish(
        self, event_name: str, message: Any, exchange: str,
        routing_key: str, body: Any
    ):
        logging.info(' '.join([
            'Publish event {}'.format(event_name),
            'Message: {}'.format(message),
            'Exchange: {}'.format(exchange),
            'Routing key: {}'.format(routing_key),
            'Body: {}'.format(body)
        ]))
