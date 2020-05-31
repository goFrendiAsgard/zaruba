import os
import json


class Config():

    def __init__(self):
        self.http_port: int = int(os.getenv("SERVICENAME_HTTP_PORT", "3000"))
        self.service_name = "servicename"
        self.default_rmq_connection_string = os.getenv(
            "DEFAULT_RMQ_CONNECTION_STRING", "amqp://localhost:5672/")

    def __str__(self):
        return "http_port: {}, service_name: {}, default_rmq_connection_string: {}".format(
            self.http_port,
            self.service_name,
            self.default_rmq_connection_string,
        )
