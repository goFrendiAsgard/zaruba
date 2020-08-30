from typing import List, Any
import uuid
from pika.adapters.blocking_connection import BlockingChannel
from pika.spec import BasicProperties

from .envelopedMessage import EnvelopedMessage
from .rpcHelper import rpc_create_enveloped_error, rpc_create_enveloped_input, rpc_create_enveloped_output
from .rmqHelper import rmq_publish

def rmq_rpc_generate_reply_queue_name(function_name: str) -> str:
    random_id = "".join(str(uuid.uuid4()).split("-"))
    return "{}.reply.{}".format(function_name, random_id)


def rmq_rpc_call(ch: BlockingChannel, exchange_name: str, reply_to: str, inputs: List[Any]):
    enveloped_input = rpc_create_enveloped_input(inputs)
    json_input = enveloped_input.to_json()
    rmq_publish(ch, exchange_name, "", json_input, BasicProperties(
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

