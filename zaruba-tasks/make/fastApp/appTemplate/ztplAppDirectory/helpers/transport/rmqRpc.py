from typing import Any, Callable, List, Optional
from pika.adapters.blocking_connection import BlockingChannel, BlockingConnection
from helpers.transport.rpc import RPC
from helpers.transport.rmqConnection import RMQConnection
from helpers.transport.rmqConfig import RMQEventMap
from pydantic import BaseModel

import time
import uuid
import pika
import threading
import traceback
import sys

class RMQRPCReply(BaseModel):
    result: Any
    error_message: Optional[str]

class RMQRPC(RMQConnection, RPC):

    def __init__(self, rmq_connection_parameters: pika.ConnectionParameters, rmq_event_map: RMQEventMap):
        RMQConnection.__init__(self, rmq_connection_parameters)
        self._event_map = rmq_event_map
        self._error_count = 0

    def get_error_count(self) -> int:
        return self._error_count

    def handle(self, rpc_name: str) -> Callable[..., Any]:
        def register_rpc_handler(rpc_handler: Callable[..., Any]):
            exchange = self._event_map.get_exchange_name(rpc_name)
            queue = self._event_map.get_queue_name(rpc_name)
            dead_letter_exchange = self._event_map.get_dead_letter_exchange(rpc_name)
            dead_letter_queue = self._event_map.get_dead_letter_queue(rpc_name)
            auto_ack = self._event_map.get_auto_ack(rpc_name)
            arguments = self._event_map.get_queue_arguments(rpc_name)
            prefetch_count = self._event_map.get_prefetch_count(rpc_name)
            def consume():
                while not self._is_shutdown:
                    connection: BlockingConnection = None
                    try:
                        connection = self.create_connection()
                        ch = connection.channel()
                        if self._event_map.get_ttl(rpc_name) > 0:
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
                    except:
                        self.remove_connection(connection)
                        self._error_count += 1
            thread = threading.Thread(target=consume)
            thread.start()
        return register_rpc_handler

    def _create_rpc_request_handler(self, rpc_name: str, exchange: str, queue: str, auto_ack: bool, rpc_handler: Callable[..., Any]):
        def on_rpc_request(ch, method, props, body):
            try:
                args: List[Any] = self._event_map.get_decoder(rpc_name)(body)
                print({'action': 'handle_rmq_rpc', 'rpc_name': rpc_name, 'args': args, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
                reply = RMQRPCReply()
                try:
                    reply.result = rpc_handler(*args)
                except Exception as e:
                    reply.error_message = getattr(e, 'message', repr(e))
                    self._error_count += 1
                    print(traceback.format_exc(), file=sys.stderr) 
                body: Any = self._event_map.get_encoder(rpc_name)(reply.dict())
                # send reply
                ch.basic_publish(
                    exchange='',
                    routing_key=props.reply_to,
                    properties=pika.BasicProperties(correlation_id=props.correlation_id),
                    body=body
                )
                print({'action': 'send_rmq_rpc_reply', 'rpc_name': rpc_name, 'args': args, 'result': reply.result, 'error': reply.error_message, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
            except Exception as e:
                self._error_count += 1
                print(traceback.format_exc(), file=sys.stderr) 
            finally:
                if not auto_ack:
                    ch.basic_ack(delivery_tag=method.delivery_tag)
        return on_rpc_request

    def call(self, rpc_name: str, *args: Any) -> Any:
        try:
            caller = RMQRPCCaller(self)
            return caller.call(rpc_name, *args)
        except Exception as e:
            self._error_count += 1
            raise e


class RMQRPCCaller():

    def __init__(self, rpc: RMQRPC):
        self.is_timeout: bool = False
        self.reply: Optional[RMQRPCReply] = None
        self.rpc = rpc
        self.event_map = rpc._event_map
        self.connection = rpc.create_connection()
        self.ch = self.connection.channel()
        self.corr_id = str(uuid.uuid4())
        self.replied = False
        self.reply_queue = ''

    def call(self, rpc_name: str, *args: Any) -> Any:
        # consume from reply queue
        self.reply_queue = 'reply.' + rpc_name + self.corr_id
        self._consume_from_reply_queue(rpc_name, self.reply_queue)
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
                reply_to=self.reply_queue,
                correlation_id=self.corr_id,
            ),
            body=body
        )
        # handle timeout
        self._handle_timeout(rpc_name)
        # clean up
        self._clean_up()
        if self.is_timeout:
            raise Exception('Timeout while calling {}'.format(rpc_name))
        if self.reply is None:
            raise Exception('No reply')
        if self.reply.error_message:
            raise Exception(self.reply.error_message)
        return self.reply.result

    def _clean_up(self):
        self.ch.stop_consuming()
        self.ch.queue_delete(self.reply_queue)
        self.ch.close()
        self.rpc.remove_connection(self.connection)

    def _consume_from_reply_queue(self, rpc_name: str, reply_queue: str):
        self.ch.queue_declare(queue=reply_queue, exclusive=True)
        on_rpc_response = self._create_rpc_responder(rpc_name, reply_queue)
        self.ch.basic_consume(queue=reply_queue, on_message_callback=on_rpc_response)

    def _create_rpc_responder(self, rpc_name: str, reply_queue: str):
        def on_rpc_response(ch: BlockingChannel, method, props, body):
            if props.correlation_id == self.corr_id:
                try:
                    body = self.event_map.get_decoder(rpc_name)(body)
                    self.reply = RMQRPCReply.parse_obj(body)
                    print({'action': 'get_rmq_rpc_reply', 'queue': reply_queue, 'correlation_id': self.corr_id, 'result': self.reply.result, 'error': self.reply.error_message})
                except Exception as e:
                    print({'action': 'get_rmq_rpc_reply', 'queue': reply_queue, 'correlation_id': self.corr_id, 'body': body, 'error': getattr(e, 'message', repr(e))})
                    print(traceback.format_exc(), file=sys.stderr) 
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
                break