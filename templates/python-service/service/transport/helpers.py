from typing import List, Any, Tuple, Callable
from .envelopedMessage import EnvelopedMessage
import uuid
import pika
from pika.adapters.blocking_connection import BlockingConnection, BlockingChannel
from pika.spec import Basic, BasicProperties

OnMessageCallback = Callable[
    [BlockingChannel, Basic.Deliver, BasicProperties, str], Any]


def rpc_create_enveloped_input(inputs: List[Any]) -> EnvelopedMessage:
    enveloped_message = EnvelopedMessage()
    enveloped_message.message = {"inputs": inputs}
    return enveloped_message


def rpc_inputs_to_json(inputs: List[Any]) -> str:
    return rpc_create_enveloped_input(inputs).to_json()


def rpc_create_enveloped_error(enveloped_input: EnvelopedMessage, err: Exception) -> EnvelopedMessage:
    enveloped_error = EnvelopedMessage().set_correlation_id(
        enveloped_input.correlation_id)
    error_message: str = str(Exception)
    enveloped_error.message = {"output": "", "error": error_message}
    enveloped_error.error_message = error_message
    return enveloped_error


def rpc_create_enveloped_output(enveloped_input: EnvelopedMessage, output: Any) -> EnvelopedMessage:
    enveloped_output = EnvelopedMessage().set_correlation_id(
        enveloped_input.correlation_id)
    enveloped_output.message = {"output": output, "error": ""}
    return enveloped_output


def rmq_rpc_generate_reply_queue_name(function_name: str) -> str:
    random_id = "".join(str(uuid.uuid4()).split("-"))
    return "{}.reply.{}".format(function_name, random_id)


def rmq_rpc_call(ch: BlockingChannel, function_name: str, reply_to: str, inputs: List[Any]):
    enveloped_input = rpc_create_enveloped_input(inputs)
    json_input = enveloped_input.to_json()
    rmq_publish(ch, function_name, "", json_input, BasicProperties(
        content_type="text/json",
        correlation_id=enveloped_input.correlation_id,
        reply_to=reply_to
    ))


def rmq_rpc_reply_output(ch: BlockingChannel, reply_to: str, enveloped_input: EnvelopedMessage, output: Any):
    json_output = rpc_create_enveloped_output(
        enveloped_input, output).to_json()
    rmq_publish(ch, "", reply_to, json_output, BasicProperties(
        content_type="text/json",
        correlation_id=enveloped_input.correlation_id
    ))


def rmq_rpc_reply_error(ch: BlockingChannel, reply_to: str, enveloped_input: EnvelopedMessage, err: Exception):
    json_error = rpc_create_enveloped_error(enveloped_input, err).to_json()
    rmq_publish(ch, "", reply_to, json_error, BasicProperties(
        content_type="text/json",
        correlation_id=enveloped_input.correlation_id
    ))


def rmq_create_connection_and_channel(connection_string: str) -> Tuple[BlockingConnection, BlockingChannel]:
    connection = pika.BlockingConnection(pika.URLParameters(connection_string))
    channel = connection.channel()
    return connection, channel


def rmq_declare_queue_and_bind_to_default_exchange(ch: BlockingChannel, queue_name: str):
    exchange_name = queue_name
    rmq_declare_fanout_exchange(ch, exchange_name)
    rmq_declare_queue(ch, queue_name)
    ch.queue_declare(exchange=exchange_name, queue=queue_name)


def rmq_declare_queue(ch: BlockingChannel, queue_name: str):
    ch.queue_declare(queue_name)


def rmq_declare_fanout_exchange(ch: BlockingChannel, exchange_name: str):
    ch.exchange_declare(exchange_name, "fanout")


def rmq_consume(ch: BlockingChannel, queue_name: str, handler: OnMessageCallback):
    ch.basic_consume(queue=queue_name, on_message_callback=handler)


def rmq_publish(ch: BlockingChannel, exchange_name: str, routing_key: str, data: str, properties: BasicProperties):
    ch.basic_publish(
        exchange=exchange_name, routing_key=routing_key, body=data, properties=properties)


def rmq_close_connection_and_channel(conn: BlockingConnection, ch: BlockingChannel):
    try:
        ch.close()
    except:
        "do nothing"
    try:
        conn.close()
    except:
        "do nothing"
