from typing import List, Any
from transport import Message, Publisher, RPCClient
from .services import greet, greet_everyone
from flask import Response, request

class Controller:

    def __init__(self, publisher: Publisher, rpc_client: RPCClient):
        self.publisher = publisher
        self.rpc_client = rpc_client
        self.names: List[str] = []
    
    def _get_name(self, name: str) -> str:
        if not name:
            name = request.args.get("name")
        if not name:
            name = request.form.get("name")
        if not name:
            name = ""
        return name

    def handle_http_hello(self, name: str = ""):
        name = self._get_name(name)
        return Response(greet(name))

    def handle_http_hello_all(self):
        return Response(greet_everyone(self.names))

    def handle_http_hello_rpc(self, name: str = ""):
        name = self._get_name(name)
        try:
            greeting = self.rpc_client.call("helloRPC", name)
            return Response(greeting)
        except Exception as e:
            return Response(str(e), status=500)

    def handle_http_hello_pub(self, name: str = ""):
        name = self._get_name(name)
        try:
            self.publisher.publish("hello", {"name": name})
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
