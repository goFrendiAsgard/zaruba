from typing import Any, Callable

from pika.adapters.blocking_connection import BlockingChannel, BlockingConnection
from .interface import MessageBus
from .rmqconfig import RmqEventMap
import pika
import uuid
import json
import time
import threading
import signal

default_event_map = RmqEventMap({})

class RMQMessageBus(MessageBus):

    def __init__(self, rmq_host: str, rmq_user: str, rmq_pass: str, rmq_vhost: str, rmq_event_map: RmqEventMap = default_event_map):
        rmq_param = pika.ConnectionParameters(
            host=rmq_host,
            credentials=pika.PlainCredentials(rmq_user, rmq_pass),
            virtual_host=rmq_vhost,
            heartbeat=5
        )
        self.event_map = rmq_event_map
        self.connection: BlockingConnection = pika.BlockingConnection(rmq_param)
        self.connection.add_callback_threadsafe(self._callback)
        self.connection.process_data_events()
        signal.signal(signal.SIGINT, self._handle_signal) 
        signal.signal(signal.SIGTERM, self._handle_signal) 


    def _handle_signal(self, sig, frame):
        self.connection.close()
        self.thread.join()
   

    def _process_data_events(self):
        while True:
            time.sleep(1)
            self.connection.process_data_events()


    def _callback(self):
        self.thread = threading.Thread(target=self._process_data_events)
        self.thread.start()


    def handle_rpc(self, event_name: str, handler: Callable[..., Any]) -> Any:
        exchange = self.event_map.get_exchange_name(event_name)
        queue = self.event_map.get_queue_name(event_name)
        dead_letter_exchange = self.event_map.get_dead_letter_exchange(event_name)
        dead_letter_queue = self.event_map.get_dead_letter_queue(event_name)
        auto_ack = self.event_map.get_auto_ack(event_name)
        arguments = self.event_map.get_queue_arguments(event_name)
        ch = self.connection.channel()
        if self.event_map.get_ttl(event_name) > 0:
            ch.exchange_declare(exchange=dead_letter_exchange, exchange_type='fanout', durable=True)
            ch.queue_declare(queue=dead_letter_queue, durable=True, exclusive=False)
            ch.queue_bind(exchange=dead_letter_exchange, queue=dead_letter_queue)
        def on_rpc_request(ch, method, props, body):
            args = json.loads(body)
            result = handler(*args)
            ch.basic_publish(
                exchange='',
                routing_key=props.reply_to,
                properties=pika.BasicProperties(correlation_id=props.correlation_id),
                body=str(result)
            )
            if not auto_ack:
                ch.basic_ack(delivery_tag=method.delivery_tag)
        ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        ch.queue_declare(queue=queue, exclusive=False, durable=True, arguments=arguments)
        ch.queue_bind(exchange=exchange, queue=queue)
        ch.basic_qos(prefetch_count=1)
        ch.basic_consume(queue=queue, on_message_callback=on_rpc_request, auto_ack=auto_ack)


    def call_rpc(self, event_name: str, *args: Any) -> Any:
        caller = RMQRPCCaller(self)
        return caller.call(event_name, *args)

    
    def handle(self, event_name: str, handler: Callable[[Any], Any]):
        exchange = self.event_map.get_exchange_name(event_name)
        queue = self.event_map.get_queue_name(event_name)
        dead_letter_exchange = self.event_map.get_dead_letter_exchange(event_name)
        dead_letter_queue = self.event_map.get_dead_letter_queue(event_name)
        auto_ack = self.event_map.get_auto_ack(event_name)
        arguments = self.event_map.get_queue_arguments(event_name)
        ch = self.connection.channel()
        if self.event_map.get_ttl(event_name) > 0:
            ch.exchange_declare(exchange=dead_letter_exchange, exchange_type='fanout', durable=True)
            ch.queue_declare(queue=dead_letter_queue, durable=True, exclusive=False)
            ch.queue_bind(exchange=dead_letter_exchange, queue=dead_letter_queue)
        def on_request(ch, method, props, body):
            msg = json.loads(body)
            result = handler(msg)
            if not auto_ack:
                ch.basic_ack(delivery_tag=method.delivery_tag)
            return result
        ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        ch.queue_declare(queue=queue, exclusive=False, durable=True, arguments=arguments)
        ch.queue_bind(exchange=exchange, queue=queue)
        ch.basic_qos(prefetch_count=1)
        ch.basic_consume(queue=queue, on_message_callback=on_request, auto_ack=auto_ack)

    
    def publish(self, event_name: str, msg: Any) -> Any:
        exchange = self.event_map.get_exchange_name(event_name)
        queue = self.event_map.get_queue_name(event_name)
        ch = self.connection.channel()
        ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        ch.basic_publish(
            exchange=exchange,
            routing_key=queue,
            body=json.dumps(msg)
        )


class RMQRPCCaller():

    def __init__(self, messagebus: RMQMessageBus):
        self.is_timeout: bool = False
        self.result = None
        self.event_map = messagebus.event_map
        self.connection = messagebus.connection
        self.ch = self.connection.channel()
        self.corr_id = str(uuid.uuid4())


    def call(self, event_name: str, *args: Any) -> Any:
        exchange = self.event_map.get_exchange_name(event_name)
        queue = self.event_map.get_queue_name(event_name)
        rpc_timeout = self.event_map.get_rpc_timeout(event_name)
        reply_queue = 'reply' + event_name + self.corr_id
        self.ch.queue_declare(queue=reply_queue, exclusive=True)
        self.ch.basic_consume(queue=reply_queue, on_message_callback=self._on_rpc_response)
        body = json.dumps(args)
        self.ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        self.ch.basic_publish(
            exchange=exchange,
            routing_key=queue,
            properties=pika.BasicProperties(
                reply_to=reply_queue,
                correlation_id=self.corr_id,
            ),
            body=body
        )
        start = time.time() * 1000
        while self.result is None:
            self.connection.process_data_events()
            if start + rpc_timeout < time.time() * 1000:
                self.is_timeout = True
        if self.is_timeout:
            raise Exception("Timeout while calling {}".format(event_name))
        self.ch.stop_consuming()
        self.ch.queue_delete(reply_queue)
        return self.result


    def _on_rpc_response(self, ch: BlockingChannel, method, props, body):
        if props.correlation_id == self.corr_id:
            self.result = str(body)
        ch.basic_ack(delivery_tag=method.delivery_tag)
