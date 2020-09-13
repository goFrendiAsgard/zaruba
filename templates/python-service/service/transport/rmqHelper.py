from typing import Any, Callable, Dict
from pika.adapters.blocking_connection import BlockingChannel
from pika.spec import Basic, BasicProperties

OnMessageCallback = Callable[
    [BlockingChannel, Basic.Deliver, BasicProperties, str], Any]

def rmq_declare_and_bind_queue(ch: BlockingChannel, exchange_name: str, queue_name: str, durable: bool, auto_delete:bool, arguments: Dict):
    rmq_declare_fanout_exchange(ch, exchange_name)
    rmq_declare_queue(ch, queue_name, durable, auto_delete, arguments)
    ch.queue_bind(exchange=exchange_name, queue=queue_name)


def rmq_declare_queue(ch: BlockingChannel, queue_name: str, durable: bool, auto_delete:bool, arguments: Dict):
    ch.queue_declare(queue_name, durable=durable, auto_delete=auto_delete, arguments=arguments)


def rmq_declare_fanout_exchange(ch: BlockingChannel, exchange_name: str):
    ch.exchange_declare(exchange_name, "fanout", durable=True)


def rmq_consume(ch: BlockingChannel, queue_name: str, durable: bool, auto_delete: bool, auto_ack: bool, arguments: Dict, handler: OnMessageCallback):
    rmq_declare_queue(ch, queue_name, durable, auto_delete, arguments)
    ch.basic_consume(queue=queue_name, on_message_callback=handler, auto_ack=auto_ack)


def rmq_publish(ch: BlockingChannel, exchange_name: str, routing_key: str, data: str, properties: BasicProperties):
    if exchange_name != "":
        rmq_declare_fanout_exchange(ch, exchange_name)
    ch.basic_publish(exchange=exchange_name, routing_key=routing_key, body=data, properties=properties)
