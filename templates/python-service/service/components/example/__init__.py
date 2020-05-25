from typing import List, Any
from transport import Message, Publisher, Subscriber, RPCServer, RPCClient
from config import Config
from .services import greet, greet_everyone
from .helpers import get_name
from flask import Flask, Response


class Component():

    def __init__(self, config: Config, router: Flask, publisher: Publisher, subscriber: Subscriber, rpc_server: RPCServer, rpc_client: RPCClient):
        self.names: List[str] = []
        self.config: Config = config
        self.router: Flask = router
        self.publisher: Publisher = publisher
        self.subscriber: Subscriber = subscriber
        self.rpc_server: RPCServer = rpc_server
        self.rpc_client: RPCClient = rpc_client

    def setup(self):

        # Use the same HTTP Handler for multiple URLS
        self.router.add_url_rule(
            "/hello", "hello", self.handle_http_hello, methods=["POST", "GET"])
        self.router.add_url_rule(
            "/hello/<name>", "hello-name", self.handle_http_hello, methods=["GET"])

        # Use HTTP Handler that take state from component
        self.router.add_url_rule(
            "/hello-all", "hello-all", self.handle_http_hello_all, methods=["GET"])

        # Trigger RPC Call
        self.router.add_url_rule(
            "/hello-rpc", "hello-rpc", self.handle_http_hello_rpc, methods=["POST", "GET"])
        self.router.add_url_rule(
            "/hello-rpc/<name>", "hello-rpc-name", self.handle_http_hello_rpc)

        # Trigger Publisher
        self.router.add_url_rule(
            "/hello-pub", "hello-pub", self.handle_http_hello_pub, methods=["POST", "GET"])
        self.router.add_url_rule(
            "/hello-pub/<name>", "hello-pub-name", self.handle_http_hello_pub)

        # Serve RPC
        self.rpc_server.register_handler(
            "servicename.helloRPC", self.handle_rpc_hello)

        # Event
        self.subscriber.register_handler(
            "servicename.helloEvent", self.handle_event_hello)

    def handle_http_hello(self, name: str = ""):
        name = get_name(name)
        return Response(greet(name))

    def handle_http_hello_all(self):
        return Response(greet_everyone(self.names))

    def handle_http_hello_rpc(self, name: str = ""):
        name = get_name(name)
        try:
            greeting = self.rpc_client.call("servicename.helloRPC", name)
            return Response(greeting)
        except Exception as e:
            return Response(str(e), status=500)

    def handle_http_hello_pub(self, name: str = ""):
        name = get_name(name)
        try:
            self.publisher.publish("servicename.helloEvent", {"name": name})
            return Response("Message sent")
        except Exception as e:
            return Response(str(e), status=500)

    def handle_rpc_hello(self, *inputs: Any):
        if len(inputs) == 0:
            raise Exception("Message accepted but input is invalid")
        name = str(inputs[0])
        return greet(name)

    def handle_event_hello(self, msg: Message):
        name = msg["name"]
        self.names.append(name)
