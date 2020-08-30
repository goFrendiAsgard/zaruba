from typing import Dict, cast
import threading
from logging import Logger
from pika.adapters.blocking_connection import BlockingConnection, BlockingChannel
from pika.spec import Basic, BasicProperties
from .interfaces import RPCHandler, RPCServer
from .envelopedMessage import EnvelopedMessage
from .rmqHelper import rmq_declare_and_bind_queue, rmq_consume, OnMessageCallback
from .rmqEventMap import RmqEventMap
from .rmqRpcHelper import rmq_rpc_reply_error, rmq_rpc_reply_output


class RmqRPCServer(RPCServer):

    def __init__(self, logger: Logger, connection: BlockingConnection, event_map: RmqEventMap):
        self.connection: BlockingConnection = connection
        self.logger: Logger = logger
        self.handlers: Dict[str, RPCHandler] = cast(
            Dict[str, RPCHandler], {})
        self.event_map: RmqEventMap = event_map

    def register_handler(self, event_name: str, handler: RPCHandler) -> RPCServer:
        self.handlers[event_name] = handler
        return self

    def serve(self):
        ch = self.connection.channel()
        for key in self.handlers:
            function_name = key
            exchange_name = self.event_map.get_exchange_name(function_name)
            queue_name = self.event_map.get_queue_name(function_name)
            handler = self.handlers[function_name]
            rmq_declare_and_bind_queue(ch, exchange_name, queue_name)
            self.logger.info(
                "[INFO RmqRPCServer] Serve {}".format(function_name))
            rmq_consume(ch, queue_name, self._create_rmq_handler(
                function_name, handler))
        thread = threading.Thread(target=ch.start_consuming)
        thread.start()

    def _create_rmq_handler(self, function_name: str, handler: RPCHandler) -> OnMessageCallback:
        def on_message(ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, body: str):
            try:
                reply_to = properties.reply_to
                json_enveloped_input = body
                enveloped_input = EnvelopedMessage(json_enveloped_input)
                try:
                    inputs = enveloped_input.message["inputs"]
                    output = handler(*inputs)
                    self.logger.info("[INFO RmqRPCServer] Reply {}, {}: {}".format(
                        function_name, inputs, output))
                    rmq_rpc_reply_output(ch, reply_to, enveloped_input, output)
                except Exception as e:
                    self.logger.error(
                        "[ERROR RmqRPCServer] Reply {}: {}".format(function_name, e))
                    rmq_rpc_reply_error(ch, reply_to, enveloped_input, e)
            except Exception as e:
                self.logger.error(
                    "[ERROR RmqRPCServer] Error Replying {}: {}".format(function_name, e))
        return on_message
