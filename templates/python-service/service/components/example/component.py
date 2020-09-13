from transport import Publisher, Subscriber, RPCServer, RPCClient
from core import Comp, App
from config import Config
from .controller import Controller
from flask import Flask


class Component(Comp):

    def __init__(self, config: Config, app: App, router: Flask, publisher: Publisher, subscriber: Subscriber, rpc_server: RPCServer, rpc_client: RPCClient):
        self.config: Config = config
        self.app: App = app
        self.router: Flask = router
        self.publisher: Publisher = publisher
        self.subscriber: Subscriber = subscriber
        self.rpc_server: RPCServer = rpc_server
        self.rpc_client: RPCClient = rpc_client
        self.controller: Controller = Controller(publisher, rpc_client)

    def setup(self):
        self.route()
        self.register_rpc_handler()
        self.register_message_handler()

    def route(self):
        controller = self.controller
        # Use the same HTTP Handler for multiple URLS
        self.router.add_url_rule("/hello", "hello", controller.handle_http_hello, methods=["POST", "GET"])
        self.router.add_url_rule("/hello/<name>", "hello-name", controller.handle_http_hello, methods=["GET"])
        # Use HTTP Handler that take state from component
        self.router.add_url_rule("/hello-all", "hello-all", controller.handle_http_hello_all, methods=["GET"])
        # Trigger RPC Call
        self.router.add_url_rule("/hello-rpc", "hello-rpc", controller.handle_http_hello_rpc, methods=["POST", "GET"])
        self.router.add_url_rule("/hello-rpc/<name>", "hello-rpc-name", controller.handle_http_hello_rpc)
        # Trigger Publisher
        self.router.add_url_rule("/hello-pub", "hello-pub", controller.handle_http_hello_pub, methods=["POST", "GET"])
        self.router.add_url_rule("/hello-pub/<name>", "hello-pub-name", controller.handle_http_hello_pub)

    def register_rpc_handler(self):
        controller = self.controller
        self.rpc_server.register_handler("helloRPC", controller.handle_rpc_hello)

    def register_message_handler(self):
        controller = self.controller
        self.subscriber.register_handler("hello", controller.handle_event_hello)