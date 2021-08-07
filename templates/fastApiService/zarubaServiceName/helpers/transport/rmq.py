from typing import Any, Callable, List, Mapping, TypedDict, Dict
from pika.adapters.blocking_connection import BlockingChannel, BlockingConnection
from helpers.transport.interface import MessageBus
from helpers.transport.rmq_config import RMQEventConfig, RMQEventMap

import pika, time, threading, uuid

class RMQMessageBus(MessageBus):

    def __init__(self, rmq_host: str, rmq_user: str, rmq_pass: str, rmq_vhost: str, rmq_event_map: RMQEventMap):
        self.rmq_param = pika.ConnectionParameters(
            host=rmq_host,
            credentials=pika.PlainCredentials(rmq_user, rmq_pass),
            virtual_host=rmq_vhost,
            heartbeat=30
        )
        self.should_check_connection = True
        self.event_map = rmq_event_map
        self._connect()
    
    def _connect(self):
        self.connection: BlockingConnection = pika.BlockingConnection(self.rmq_param)
        self.connection.add_callback_threadsafe(self._callback)
        self.connection.process_data_events()

    def _process_data_events(self):
        while True:
            time.sleep(5)
            if not self.should_check_connection:
                break
            try:
                self.connection.process_data_events()
            except:
                self._connect()

    def _callback(self):
        self.thread = threading.Thread(target=self._process_data_events)
        self.thread.start()
 
    def shutdown(self):
        self.should_check_connection = False
        self.connection.close()
        self.thread.join()

    def handle_rpc(self, event_name: str) -> Callable[..., Any]:
        def register_rpc_handler(rpc_handler: Callable[..., Any]):
            exchange = self.event_map.get_exchange_name(event_name)
            queue = self.event_map.get_queue_name(event_name)
            dead_letter_exchange = self.event_map.get_dead_letter_exchange(event_name)
            dead_letter_queue = self.event_map.get_dead_letter_queue(event_name)
            auto_ack = self.event_map.get_auto_ack(event_name)
            arguments = self.event_map.get_queue_arguments(event_name)
            prefetch_count = self.event_map.get_prefetch_count(event_name)
            ch = self.connection.channel()
            if self.event_map.get_ttl(event_name) > 0:
                ch.exchange_declare(exchange=dead_letter_exchange, exchange_type='fanout', durable=True)
                ch.queue_declare(queue=dead_letter_queue, durable=True, exclusive=False)
                ch.queue_bind(exchange=dead_letter_exchange, queue=dead_letter_queue)
            ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
            ch.queue_declare(queue=queue, exclusive=False, durable=True, arguments=arguments)
            ch.queue_bind(exchange=exchange, queue=queue)
            ch.basic_qos(prefetch_count=prefetch_count)
            on_rpc_request = self._create_rpc_request_handler(event_name, exchange, queue, auto_ack, rpc_handler)
            ch.basic_consume(queue=queue, on_message_callback=on_rpc_request, auto_ack=auto_ack)
        return register_rpc_handler

    def _create_rpc_request_handler(self, event_name: str, exchange: str, queue: str, auto_ack: bool, rpc_handler: Callable[..., Any]):
        def on_rpc_request(ch, method, props, body):
            try:
                args: List[Any] = self.event_map.get_decoder(event_name)(body)
                print({'action': 'handle_rmq_rpc', 'event_name': event_name, 'args': args, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
                result = rpc_handler(*args)
                body: Any = self.event_map.get_encoder(event_name)(result)
                # send reply
                ch.basic_publish(
                    exchange='',
                    routing_key=props.reply_to,
                    properties=pika.BasicProperties(correlation_id=props.correlation_id),
                    body=body
                )
                print({'action': 'send_rmq_rpc_reply', 'event_name': event_name, 'args': args, 'result': result, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
            finally:
                if not auto_ack:
                    ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_rpc_request

    def call_rpc(self, event_name: str, *args: Any) -> Any:
        caller = RMQRPCCaller(self)
        return caller.call(event_name, *args)

    def handle_event(self, event_name: str) -> Callable[..., Any]:
        def register_event_handler(event_handler: Callable[[Any], Any]):
            exchange = self.event_map.get_exchange_name(event_name)
            queue = self.event_map.get_queue_name(event_name)
            dead_letter_exchange = self.event_map.get_dead_letter_exchange(event_name)
            dead_letter_queue = self.event_map.get_dead_letter_queue(event_name)
            auto_ack = self.event_map.get_auto_ack(event_name)
            arguments = self.event_map.get_queue_arguments(event_name)
            prefetch_count = self.event_map.get_prefetch_count(event_name)
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
            on_event = self._create_rpc_handler(event_name, exchange, queue, auto_ack, event_handler)
            ch.basic_consume(queue=queue, on_message_callback=on_event, auto_ack=auto_ack)
        return register_event_handler
    
    def _create_rpc_handler(self, event_name: str, exchange: str, queue: str, auto_ack: bool, event_handler: Callable[[Any], Any]):
        def on_event(ch, method, props, body):
            try:
                message = self.event_map.get_decoder(event_name)(body)
                print({'action': 'handle_rmq_event', 'event_name': event_name, 'message': message, 'exchange': exchange, 'routing_key': queue})
                result = event_handler(message)
            finally:
                if not auto_ack:
                    ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_event

    def publish(self, event_name: str, message: Any) -> Any:
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

class RMQRPCCaller():

    def __init__(self, messagebus: RMQMessageBus):
        self.is_timeout: bool = False
        self.result = None
        self.event_map = messagebus.event_map
        self.connection = messagebus.connection
        self.ch = self.connection.channel()
        self.corr_id = str(uuid.uuid4())
        self.replied = False

    def call(self, event_name: str, *args: Any) -> Any:
        # consume from reply queue
        reply_queue = 'reply.' + event_name + self.corr_id
        self._consume_from_reply_queue(event_name, reply_queue)
        # publish message
        exchange = self.event_map.get_exchange_name(event_name)
        routing_key = self.event_map.get_queue_name(event_name)
        body = self.event_map.get_encoder(event_name)(args)
        self.ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        print({'action': 'call_rmq_rpc', 'event_name': event_name, 'args': args, 'exchange': exchange, 'routing_key': routing_key, 'correlation_id': self.corr_id, 'body': body})
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
        self._handle_timeout(event_name)
        # clean up
        self.ch.stop_consuming()
        self.ch.queue_delete(reply_queue)
        return self.result

    def _consume_from_reply_queue(self, event_name: str, reply_queue: str):
        self.ch.queue_declare(queue=reply_queue, exclusive=True)
        on_rpc_response = self._create_rpc_responder(event_name, reply_queue)
        self.ch.basic_consume(queue=reply_queue, on_message_callback=on_rpc_response)

    def _create_rpc_responder(self, event_name: str, reply_queue: str):
        def on_rpc_response(ch: BlockingChannel, method, props, body):
            if props.correlation_id == self.corr_id:
                self.result = self.event_map.get_decoder(event_name)(body)
                print({'action': 'get_rmq_rpc_reply', 'queue': reply_queue, 'correlation_id': self.corr_id, 'result': self.result})
                self.replied = True
            ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_rpc_response

    def _handle_timeout(self, event_name: str):
        rpc_timeout = self.event_map.get_rpc_timeout(event_name)
        start = time.time() * 1000
        while not self.replied:
            self.connection.process_data_events()
            if start + rpc_timeout < time.time() * 1000:
                self.is_timeout = True
        if self.is_timeout:
            raise Exception('Timeout while calling {}'.format(event_name))