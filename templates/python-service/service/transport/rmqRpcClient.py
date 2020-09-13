from typing import Any, Dict, List, Tuple, cast
from pika.adapters.blocking_connection import BlockingConnection, BlockingChannel
from pika.spec import Basic, BasicProperties
from logging import Logger
from .interfaces import Message, RPCClient
from .rmqRpcHelper import rmq_rpc_generate_reply_queue_name, rmq_rpc_call
from .rmqHelper import rmq_consume, rmq_declare_queue, OnMessageCallback
from .rmqEventMap import RmqEventMap
from .envelopedMessage import EnvelopedMessage

import time


class RmqRPCClient(RPCClient):

    def __init__(self, logger: Logger, connection: BlockingConnection, event_map: RmqEventMap):
        self.connection: BlockingConnection = connection
        self.logger: Logger = logger
        self.event_map: RmqEventMap = event_map

    def call(self, function_name: str, *inputs: Any) -> Any:
        # reply state
        reply_state: Dict[str, Any] = self._create_default_reply_state()
        exchange_name = self.event_map.get_exchange_name(function_name)
        queue_name = self.event_map.get_queue_name(function_name)
        # send message
        reply_to = rmq_rpc_generate_reply_queue_name(queue_name)
        ch = self.connection.channel()
        rmq_declare_queue(ch, reply_to, True, True, {})
        rmq_consume(ch, reply_to, True, True, True, {}, self._create_reply_handler(function_name, inputs, reply_state))
        self.logger.info(
            "[INFO RmqRPCClient] Call {} {}".format(function_name, inputs))
        rmq_rpc_call(ch, exchange_name, reply_to, cast(List[Any], inputs))
        # waiting
        start = time.time() * 1000
        timeout = self.event_map.get_rpc_timeout(function_name)
        while not reply_state["accepted"]:
            self.connection.process_data_events()
            if start + timeout < time.time() * 1000:
                reply_state["error_message"] = "Timeout {}".format(timeout)
                self.logger.info("[ERROR RmqRPCClient] Get timeout {} {}: {} ms".format(function_name, inputs, timeout))
                break
        # return or throw error
        if reply_state["error_message"] != "":
            raise Exception(reply_state["error_message"])
        return reply_state["output"]

    def _create_default_reply_state(self) -> Dict[str, Any]:
        return {"accepted": False, "output": None, "error_message": ""}

    def _create_reply_handler(self, function_name: str, inputs: Tuple, reply_state: Dict[str, Any]) -> OnMessageCallback:
        def on_reply(ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, json_enveloped_output: str) -> Any:
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