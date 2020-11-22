from typing import Optional, Any
from fastapi import FastAPI

import transport
import time

def init(app: FastAPI, mb: transport.MessageBus):

    # handle message
    def receive(msg: Any):
        print(msg)
    mb.handle('hit', receive)


    # handle rpc
    def hello(name: str):
        return "hello " + name
    mb.handle_rpc("hello.rpc", hello)


    @app.get("/")
    def read_root():
        mb.publish('hit', {"event": "hit", "time": time.gmtime()})
        return {"Hello": "World"}


    @app.get("/hello/{name}")
    def read_item(name: str, q: Optional[str] = None):
        mb.call_rpc("hello.rpc", name)
        return {"name": name, "q": q}
