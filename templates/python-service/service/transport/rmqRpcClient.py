import pika
from .interfaces import Message, RPCClient
from .helpers import rmq_rpc_generate_reply_queue_name, rmq_consume, rmq_declare_queue, rmq_rpc_call, OnMessageCallback
from .envelopedMessage import EnvelopedMessage
from pika.adapters.blocking_connection import BlockingChannel
from pika.spec import Basic, BasicProperties
from logging import Logger
from typing import Any, Dict, List, cast


class RmqRPCClient(RPCClient):

    def __init__(self, logger: Logger, connection: pika.BlockingConnection):
        self.connection: pika.BlockingConnection = connection
        self.logger: Logger = logger

    def create_default_reply_state(self) -> Dict[str, Any]:
        return {"accepted": False, "output": None, "error_message": ""}

    def call(self, function_name: str, *inputs: Any) -> Any:
        # reply state
        reply_state: Dict[str, Any] = self.create_default_reply_state()

        # reply handler
        def create_on_reply(reply_state: Dict[str, Any]) -> OnMessageCallback:
            def on_reply(ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, json_enveloped_output: str):
                if reply_state["accepted"]:
                    return
                reply_state["accepted"] = True
                try:
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

        # send message
        reply_to = rmq_rpc_generate_reply_queue_name(function_name)
        ch = self.connection.channel()
        rmq_declare_queue(ch, reply_to)
        rmq_consume(ch, reply_to, create_on_reply(reply_state))
        self.logger.info(
            "[INFO RmqRPCClient] Call {} {}".format(function_name, inputs))
        rmq_rpc_call(ch, function_name, reply_to, cast(List[Any], inputs))

        # waiting
        while not reply_state["accepted"]:
            self.connection.process_data_events()
        self.delete_queue_and_close_channel(ch, reply_to)
        # return or throw error
        if reply_state["error_message"] != "":
            raise Exception(reply_state["error_message"])
        return reply_state["output"]

    def delete_queue_and_close_channel(self, ch: BlockingChannel, queue_name: str):
        try:
            ch.queue_delete(queue_name)
            ch.close()
        except:
            return
