from pika.adapters.blocking_connection import BlockingConnection

import pika

def create_rmq_connection_parameters(host: str, user: str, password: str, virtual_host: str = '/', heartbeat: int = 30) -> pika.ConnectionParameters:
    return pika.ConnectionParameters(
        host=host,
        credentials=pika.PlainCredentials(user, password),
        virtual_host=virtual_host,
        heartbeat=heartbeat
    )

class RMQConnection():

    def __init__(self, connection_parameters: pika.ConnectionParameters):
        self.connection_parameters = connection_parameters
        self.is_shutdown = False
        self._connect()

    def _connect(self):
        self.connection: BlockingConnection = pika.BlockingConnection(self.connection_parameters)
        
    def shutdown(self):
        if self.is_shutdown:
            return
        print('closing RMQ connection')
        self.connection.close()
        self.is_shutdown = True
