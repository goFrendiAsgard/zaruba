from typing import Any, Callable, Mapping
from pika.adapters.blocking_connection import BlockingChannel, BlockingConnection
from transport.interface import MessageBus
from transport.rmqconfig import RmqEventMap
from transport.logger import log_info

import jsons, pika, time, threading, uuid

default_event_map = RmqEventMap({})

class RMQMessageBus(MessageBus):

    def __init__(self, rmq_host: str, rmq_user: str, rmq_pass: str, rmq_vhost: str, rmq_event_map: RmqEventMap = default_event_map):
        self.rmq_param = pika.ConnectionParameters(
            host=rmq_host,
            credentials=pika.PlainCredentials(rmq_user, rmq_pass),
            virtual_host=rmq_vhost,
            heartbeat=5
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
            time.sleep(3)
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
        ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        ch.queue_declare(queue=queue, exclusive=False, durable=True, arguments=arguments)
        ch.queue_bind(exchange=exchange, queue=queue)
        ch.basic_qos(prefetch_count=1)
        # create handler and start consuming
        def on_rpc_request(ch, method, props, body):
            decoded_body = body.decode()
            log_info('GET RPC REQUEST', queue=queue, body=decoded_body)
            args = jsons.loads(decoded_body)
            result = handler(*args)
            body = jsons.dumps(result)
            # send reply
            log_info('SEND RPC REPLY', exchange=exchange, routing_key=props.reply_to, correlation_id=props.correlation_id, body=body)
            ch.basic_publish(
                exchange='',
                routing_key=props.reply_to,
                properties=pika.BasicProperties(correlation_id=props.correlation_id),
                body=body.encode()
            )
            if not auto_ack:
                ch.basic_ack(delivery_tag=method.delivery_tag)
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
        ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        ch.queue_declare(queue=queue, exclusive=False, durable=True, arguments=arguments)
        ch.queue_bind(exchange=exchange, queue=queue)
        ch.basic_qos(prefetch_count=1)
        # create handler and start consuming
        def on_request(ch, method, props, body):
            decoded_body = body.decode()
            log_info('HANDLE EVENT', queue=queue, body=decoded_body)
            msg = jsons.loads(decoded_body)
            result = handler(msg)
            if not auto_ack:
                ch.basic_ack(delivery_tag=method.delivery_tag)
            return result
        ch.basic_consume(queue=queue, on_message_callback=on_request, auto_ack=auto_ack)


    def publish(self, event_name: str, msg: Any) -> Any:
        exchange = self.event_map.get_exchange_name(event_name)
        queue = self.event_map.get_queue_name(event_name)
        body = jsons.dumps(msg)
        ch = self.connection.channel()
        ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        log_info('PUBLISH EVENT', exchange=exchange, routing_key=queue, body=body)
        ch.basic_publish(
            exchange=exchange,
            routing_key=queue,
            body=body.encode()
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
        # consume from reply queue
        reply_queue = 'reply.' + event_name + self.corr_id
        self._consume_from_reply_queue(reply_queue)
        # publish message
        exchange = self.event_map.get_exchange_name(event_name)
        queue = self.event_map.get_queue_name(event_name)
        body = jsons.dumps(args)
        self.ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        log_info('CALL RPC', exchange=exchange, routing_key=queue, reply_to=reply_queue, correlation_id=self.corr_id, body=body)
        self.ch.basic_publish(
            exchange=exchange,
            routing_key=queue,
            properties=pika.BasicProperties(
                reply_to=reply_queue,
                correlation_id=self.corr_id,
            ),
            body=body.encode()
        )
        # handle timeout
        self._handle_timeout(event_name)
        # clean up
        self.ch.stop_consuming()
        self.ch.queue_delete(reply_queue)
        return self.result


    def _consume_from_reply_queue(self, reply_queue: str):
        self.ch.queue_declare(queue=reply_queue, exclusive=True)
        def on_rpc_response(ch: BlockingChannel, method, props, body):
            if props.correlation_id == self.corr_id:
                decoded_body = body.decode()
                log_info('GET RPC REPLY', queue=reply_queue, correlation_id=self.corr_id, body=decoded_body)
                self.result = jsons.loads(decoded_body)
            ch.basic_ack(delivery_tag=method.delivery_tag)
        self.ch.basic_consume(queue=reply_queue, on_message_callback=on_rpc_response)

    
    def _handle_timeout(self, event_name: str):
        rpc_timeout = self.event_map.get_rpc_timeout(event_name)
        start = time.time() * 1000
        while self.result is None:
            self.connection.process_data_events()
            if start + rpc_timeout < time.time() * 1000:
                self.is_timeout = True
        if self.is_timeout:
            raise Exception('Timeout while calling {}'.format(event_name))
