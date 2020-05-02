import os


class Config():

    def __init__(self):
        self.http_port: int = int(os.getenv("SERVICENAME_HTTP_PORT", "3000"))
        self.service_name = "servicename"
        self.global_rmq_connection_string = os.getenv(
            "GLOBAL_RMQ_CONNECTION_STRING", "amqp://localhost:5672/")
        self.local_rmq_connection_string = os.getenv(
            "LOCAL_RMQ_CONNECTION_STRING", "amqp://localhost:5672/")

    def __str__(self):
        return "http_port: {}, service_name: {}, global_rmq_connection_string: {}, local_rmq_connection_string: {}".format(
            self.http_port,
            self.service_name,
            self.global_rmq_connection_string,
            self.local_rmq_connection_string
        )
