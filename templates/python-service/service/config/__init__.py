import os
import json
from transport import RmqEventMap


class Config():

    def __init__(self):
        self.http_port: int = int(os.getenv("SERVICENAME_HTTP_PORT", "3000"))
        self.service_name = "servicename"
        self.default_rmq_connection_string = os.getenv(
            "DEFAULT_RMQ_CONNECTION_STRING", "amqp://localhost:5672/")
        self.rmq_event_map = RmqEventMap({
            "helloRPC": {
                "exchange_name": "servicename.exchange.helloRPC",
                "queue_name": "servicename.queue.helloRPC",
            },
            "hello": {
                "exchange_name": "servicename.exchange.helloEvent",
                "queue_name": "servicename.queue.helloEvent",
            }
        })

    def __str__(self):
        return json.dumps({
            "http_port": self.http_port,
            "service_name": self.service_name,
            "default_rmq_connection_string": self.default_rmq_connection_string,
            "rmq_event_map": self.rmq_event_map.mapping
        })