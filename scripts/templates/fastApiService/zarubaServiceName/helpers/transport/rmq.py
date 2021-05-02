from typing import Any, Callable, Mapping, TypedDict, Dict
from pika.adapters.blocking_connection import BlockingChannel, BlockingConnection
from helpers.transport.interface import MessageBus

import jsons, pika, time, threading, uuid

class RMQEventConfig(TypedDict):
    queue_name: str
    exchange_name: str
    rpc_timeout: int
    dead_letter_exchange: str
    dead_letter_queue: str
    ttl: int
    auto_ack: bool
    prefetch_count: int


class RMQEventMap:

    def __init__(self, mapping: Mapping[str, RMQEventConfig]):
        self.mapping = mapping

    def get_exchange_name(self, event_name: str) -> str:
        if event_name in self.mapping and 'exchange_name' in self.mapping[event_name] and self.mapping[event_name]['exchange_name'] != '':
            return self.mapping[event_name]['exchange_name']
        return event_name

    def get_queue_name(self, event_name: str) -> str:
        if event_name in self.mapping and 'queue_name' in self.mapping[event_name] and self.mapping[event_name]['queue_name'] != '':
            return self.mapping[event_name]['queue_name']
        return event_name

    def get_dead_letter_exchange(self, event_name: str) -> str:
        if event_name in self.mapping and 'dead_letter_exchange' in self.mapping[event_name] and self.mapping[event_name]['dead_letter_exchange'] != '':
            return self.mapping[event_name]['dead_letter_exchange']
        return '{}.dlx'.format(self.get_exchange_name(event_name))

    def get_dead_letter_queue(self, event_name: str) -> str:
        if event_name in self.mapping and 'dead_letter_queue' in self.mapping[event_name] and self.mapping[event_name]['dead_letter_queue'] != '':
            return self.mapping[event_name]['dead_letter_queue']
        return '{}.dlx'.format(self.get_queue_name(event_name))

    def get_ttl(self, event_name: str) -> int:
        if event_name in self.mapping and 'ttl' in self.mapping[event_name] and self.mapping[event_name]['ttl'] > 0:
            return self.mapping[event_name]['ttl']
        return 0

    def get_queue_arguments(self, event_name: str) -> Dict:
        args = {}
        if self.get_ttl(event_name) <= 0:
            return {}
        args['x-dead-letter-exchange'] = self.get_dead_letter_exchange(event_name)
        args['x-message-ttl'] = self.get_ttl(event_name)
        return args

    def get_rpc_timeout(self, event_name: str) -> int:
        if event_name in self.mapping and 'rpc_timeout' in self.mapping[event_name] and self.mapping[event_name]['rpc_timeout'] > 0:
            return self.mapping[event_name]['rpc_timeout']
        return 5000
    
    def get_prefetch_count(self, event_name: str) -> int:
        if event_name in self.mapping and 'prefetch_count' in self.mappping[event_name] and self.mapping[event_name]['prefetch_count'] > 0:
            return self.mapping[event_name]['prefetch_count']
        return 10

    def get_auto_ack(self, event_name: str) -> bool:
        if event_name in self.mapping and 'auto_ack' in self.mapping[event_name]:
            return self.mapping[event_name]['auto_ack']
        return False


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
            # create handler and start consuming
            def on_rpc_request(ch, method, props, body):
                try:
                    decoded_body = body.decode()
                    args = jsons.loads(decoded_body)
                    print({'action': 'handle_rmq_rpc', 'event_name': event_name, 'args': args, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
                    result = rpc_handler(*args)
                    body = jsons.dumps(result)
                    # send reply
                    ch.basic_publish(
                        exchange='',
                        routing_key=props.reply_to,
                        properties=pika.BasicProperties(correlation_id=props.correlation_id),
                        body=body.encode()
                    )
                    print({'action': 'send_rmq_rpc_reply', 'event_name': event_name, 'args': args, 'result': result, 'exchange': exchange, 'routing_key': queue, 'correlation_id': props.correlation_id})
                finally:
                    if not auto_ack:
                        ch.basic_ack(delivery_tag=method.delivery_tag)
            ch.basic_consume(queue=queue, on_message_callback=on_rpc_request, auto_ack=auto_ack)
        return register_rpc_handler

    def call_rpc(self, event_name: str, *args: Any) -> Any:
        caller = RMQRPCCaller(self)
        return caller.call(event_name, *args)

    def handle(self, event_name: str) -> Callable[..., Any]:
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
            def on_event(ch, method, props, body):
                try:
                    decoded_body = body.decode()
                    message = jsons.loads(decoded_body)
                    print({'action': 'handle_rmq_event', 'event_name': event_name, 'message': message, 'exchange': exchange, 'routing_key': queue})
                    result = event_handler(message)
                finally:
                    if not auto_ack:
                        ch.basic_ack(delivery_tag=method.delivery_tag)
            ch.basic_consume(queue=queue, on_message_callback=on_event, auto_ack=auto_ack)
        return register_event_handler

    def publish(self, event_name: str, message: Any) -> Any:
        exchange = self.event_map.get_exchange_name(event_name)
        routing_key = self.event_map.get_queue_name(event_name)
        body = jsons.dumps(message)
        ch = self.connection.channel()
        ch.exchange_declare(exchange=exchange, exchange_type='fanout', durable=True)
        print({'action': 'publish_rmq_event', 'event_name': event_name, 'message': message, 'exchange': exchange, 'routing_key': routing_key, 'body': body})
        ch.basic_publish(
            exchange=exchange,
            routing_key=routing_key,
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
        self.replied = False

    def call(self, event_name: str, *args: Any) -> Any:
        # consume from reply queue
        reply_queue = 'reply.' + event_name + self.corr_id
        self._consume_from_reply_queue(reply_queue)
        # publish message
        exchange = self.event_map.get_exchange_name(event_name)
        routing_key = self.event_map.get_queue_name(event_name)
        body = jsons.dumps(args).encode()
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

    def _consume_from_reply_queue(self, reply_queue: str):
        self.ch.queue_declare(queue=reply_queue, exclusive=True)
        def on_rpc_response(ch: BlockingChannel, method, props, body):
            if props.correlation_id == self.corr_id:
                decoded_body = body.decode()
                self.result = jsons.loads(decoded_body)
                print({'action': 'get_rmq_rpc_reply', 'queue': reply_queue, 'correlation_id': self.corr_id, 'result': self.result})
                self.replied = True
            ch.basic_ack(delivery_tag=method.delivery_tag)
        self.ch.basic_consume(queue=reply_queue, on_message_callback=on_rpc_response)
    
    def _handle_timeout(self, event_name: str):
        rpc_timeout = self.event_map.get_rpc_timeout(event_name)
        start = time.time() * 1000
        while not self.replied:
            self.connection.process_data_events()
            if start + rpc_timeout < time.time() * 1000:
                self.is_timeout = True
        if self.is_timeout:
            raise Exception('Timeout while calling {}'.format(event_name))