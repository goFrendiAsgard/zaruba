from .interfaces import Message, RPCClient
from .helpers import rmq_rpc_generate_reply_queue_name, rmq_consume, rmq_create_connection_and_channel, rmq_declare_queue, rmq_close_connection_and_channel, rmq_rpc_call, OnMessageCallback
from .envelopedMessage import EnvelopedMessage
from pika.adapters.blocking_connection import BlockingConnection, BlockingChannel
from pika.spec import Basic, BasicProperties
from logging import Logger, getLogger
from typing import Any, Dict, List, cast


class RmqRPCClient(RPCClient):

    def __init__(self, connection_string: str):
        self.connection_string: str = connection_string
        self.logger: Logger = getLogger()

    def set_logger(self, logger: Logger) -> RPCClient:
        self.logger = logger
        return self

    def call(self, function_name: str, *inputs: Any) -> Any:
        # reply state
        reply_state: Dict[str, Any] = {
            "accepted": False, "output": None, "error_message": ""}

        # reply handler
        def create_on_reply(reply_state: Dict[str, Any]) -> OnMessageCallback:
            def on_reply(ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, body: str):
                if reply_state["accepted"]:
                    return
                reply_state["accepted"] = True
                try:
                    json_enveloped_output = body
                    enveloped_output = EnvelopedMessage(json_enveloped_output)
                    if enveloped_output.error_message:
                        self.logger.info("[ERROR RmqRPCClient] Get Error Reply {} {}: {}".format(
                            function_name, inputs, enveloped_output.error_message))
                        reply_state["error_message"] = enveloped_output.error_message
                        return
                    self.logger.info("[INFO RmqRPCClient] Get Reply {} {}: {}".format(
                        function_name, inputs, enveloped_output.message))
                    message = enveloped_output.message
                    reply_state["output"] = message["output"]
                except Exception as e:
                    self.logger.info("[ERROR RmqRPCClient] Error While Processing Reply {} {}: {}".format(
                        function_name, inputs, e))
                    reply_state["error_message"] = str(e)
            return on_reply

        # algorithm
        reply_to = rmq_rpc_generate_reply_queue_name(function_name)
        conn, ch = rmq_create_connection_and_channel(self.connection_string)
        rmq_declare_queue(ch, reply_to)
        rmq_consume(ch, function_name, create_on_reply(reply_state))
        self.logger.info(
            "[INFO RmqRPCClient] Call {} {}".format(function_name, inputs))
        rmq_rpc_call(ch, function_name, reply_to, cast(List[Any], inputs))
        # waiting
        while not reply_state["accepted"]:
            conn.process_data_events()
        self.delete_queue_and_close_connection(conn, ch, reply_to)
        # return or throw error
        if reply_state["error_message"] != "":
            raise Exception(reply_state["error_message"])
        return reply_state["output"]

    def delete_queue_and_close_connection(self, conn: BlockingConnection, ch: BlockingChannel, queue_name: str):
        try:
            ch.queue_delete(queue_name)
            rmq_close_connection_and_channel(conn, ch)
        except:
            return
