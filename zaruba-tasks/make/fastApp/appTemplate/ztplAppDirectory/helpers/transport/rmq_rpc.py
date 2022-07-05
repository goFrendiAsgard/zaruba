from typing import Any, Callable, List
from pika.adapters.blocking_connection import BlockingChannel
from helpers.transport.interface import RPC
from helpers.transport.rmq_connection import RMQConnection
from helpers.transport.rmq_config import RMQEventMap

import time
import uuid
import pika
import threading
import traceback

class RMQRPC(RMQConnection, RPC):

    def __init__(self, rmq_connection_parameters: pika.ConnectionParameters, rmq_event_map: RMQEventMap):
        RMQConnection.__init__(self, rmq_connection_parameters)
        self.event_map = rmq_event_map
        self.error_count = 0
        self.publish_connection = self.create_connection()

    def get_error_count(self) -> int:
        return self.error_count

    def handle(self, rpc_name: str) -> Callable[..., Any]:
        def register_rpc_handler(rpc_handler: Callable[..., Any]):
            exchange = self.event_map.get_exchange_name(rpc_name)
            queue = self.event_map.get_queue_name(rpc_name)
            dead_letter_exchange = self.event_map.get_dead_letter_exchange(rpc_name)
            dead_letter_queue = self.event_map.get_dead_letter_queue(rpc_name)
            auto_ack = self.event_map.get_auto_ack(rpc_name)
            arguments = self.event_map.get_queue_arguments(rpc_name)
            prefetch_count = self.event_map.get_prefetch_count(rpc_name)
            def consume():
                connection = self.create_connection()
                ch = connection.channel()
                if self.event_map.get_ttl(rpc_name) > 0:
                    ch.exchange_declare(exchange=dead_letter_exchange, exchange_type='fanout', durable=True)
                    ch.queue_declare(queue=dead_letter_queue, durable=True, exclusive=False)
                    ch.queue_bind(exchange=dead_letter_exchange, queue=dead_letter_queue)
                ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
                ch.queue_declare(queue=queue, exclusive=False, durable=True, arguments=arguments)
                ch.queue_bind(exchange=exchange, queue=queue)
                ch.basic_qos(prefetch_count=prefetch_count)
                on_rpc_request = self._create_rpc_request_handler(rpc_name, exchange, queue, auto_ack, rpc_handler)
                ch.basic_consume(queue=queue, on_message_callback=on_rpc_request, auto_ack=auto_ack)
                ch.start_consuming()
            thread = threading.Thread(target=consume)
            thread.start()
        return register_rpc_handler

    def _create_rpc_request_handler(self, rpc_name: str, exchange: str, queue: str, auto_ack: bool, rpc_handler: Callable[..., Any]):
        def on_rpc_request(ch, method, props, body):
            try:
                args: List[Any] = self.event_map.get_decoder(rpc_name)(body)
                print({'action': 'handle_rmq_rpc', 'rpc_name': rpc_name, 'args': args, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
                result = rpc_handler(*args)
                body: Any = self.event_map.get_encoder(rpc_name)(result)
                # send reply
                ch.basic_publish(
                    exchange='',
                    routing_key=props.reply_to,
                    properties=pika.BasicProperties(correlation_id=props.correlation_id),
                    body=body
                )
                print({'action': 'send_rmq_rpc_reply', 'rpc_name': rpc_name, 'args': args, 'result': result, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
            except Exception as e:
                self.error_count += 1
                print(traceback.format_exc()) 
            finally:
                if not auto_ack:
                    ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_rpc_request

    def call(self, rpc_name: str, *args: Any) -> Any:
        try:
            caller = RMQRPCCaller(self)
            return caller.call(rpc_name, *args)
        except Exception as e:
            self.error_count += 1
            raise e


class RMQRPCCaller():

    def __init__(self, rpc: RMQRPC):
        self.is_timeout: bool = False
        self.result = None
        self.event_map = rpc.event_map
        self.connection = rpc.publish_connection
        self.ch = self.connection.channel()
        self.corr_id = str(uuid.uuid4())
        self.replied = False

    def call(self, rpc_name: str, *args: Any) -> Any:
        # consume from reply queue
        reply_queue = 'reply.' + rpc_name + self.corr_id
        self._consume_from_reply_queue(rpc_name, reply_queue)
        # publish message
        exchange = self.event_map.get_exchange_name(rpc_name)
        routing_key = self.event_map.get_queue_name(rpc_name)
        body = self.event_map.get_encoder(rpc_name)(args)
        self.ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        print({'action': 'call_rmq_rpc', 'rpc_name': rpc_name, 'args': args, 'exchange': exchange, 'routing_key': routing_key, 'correlation_id': self.corr_id, 'body': body})
        self.ch.basic_publish(
            exchange=exchange,
            routing_key=routing_key,
            properties=pika.BasicProperties(
                reply_to=reply_queue,
                correlation_id=self.corr_id,
            ),
            body=body
        )
        # handle timeout
        self._handle_timeout(rpc_name)
        # clean up
        self.ch.stop_consuming()
        self.ch.queue_delete(reply_queue)
        return self.result

    def _consume_from_reply_queue(self, rpc_name: str, reply_queue: str):
        self.ch.queue_declare(queue=reply_queue, exclusive=True)
        on_rpc_response = self._create_rpc_responder(rpc_name, reply_queue)
        self.ch.basic_consume(queue=reply_queue, on_message_callback=on_rpc_response)

    def _create_rpc_responder(self, rpc_name: str, reply_queue: str):
        def on_rpc_response(ch: BlockingChannel, method, props, body):
            if props.correlation_id == self.corr_id:
                self.result = self.event_map.get_decoder(rpc_name)(body)
                print({'action': 'get_rmq_rpc_reply', 'queue': reply_queue, 'correlation_id': self.corr_id, 'result': self.result})
                self.replied = True
            ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_rpc_response

    def _handle_timeout(self, rpc_name: str):
        rpc_timeout = self.event_map.get_rpc_timeout(rpc_name)
        start = time.time() * 1000
        while not self.replied:
            self.connection.process_data_events()
            if start + rpc_timeout < time.time() * 1000:
                self.is_timeout = True
        if self.is_timeout:
            raise Exception('Timeout while calling {}'.format(rpc_name))