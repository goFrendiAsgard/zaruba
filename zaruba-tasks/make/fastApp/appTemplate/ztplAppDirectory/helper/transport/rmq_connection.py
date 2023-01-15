from typing import List
from pika.adapters.blocking_connection import BlockingConnection

import pika
import logging


def create_rmq_connection_parameters(
    host: str,
    user: str,
    password: str,
    virtual_host: str = '/',
    heartbeat: int = 60,
    blocked_connection_timeout: int = 30
) -> pika.ConnectionParameters:
    return pika.ConnectionParameters(
        host=host,
        credentials=pika.PlainCredentials(user, password),
        virtual_host=virtual_host,
        heartbeat=heartbeat,
        blocked_connection_timeout=blocked_connection_timeout
    )


class RMQConnection():

    def __init__(self, connection_parameters: pika.ConnectionParameters):
        self._connection_parameters = connection_parameters
        self._is_shutdown = False
        self._is_failing = False
        self._connections: List[BlockingConnection] = []

    def is_failing(self) -> bool:
        return self._is_failing

    def create_connection(self) -> BlockingConnection:
        connection: BlockingConnection = pika.BlockingConnection(
            self._connection_parameters)
        self._connections.append(connection)
        return connection

    def remove_connection(self, connection: BlockingConnection):
        self._close_connection(connection)
        self._connections.remove(connection)

    def shutdown(self):
        if self._is_shutdown:
            return
        self._is_shutdown = True
        logging.info('Closing RMQ connections')
        self._stop_connections()

    def _stop_connections(self):
        for connection in self._connections:
            self._close_connection(connection)

    def _close_connection(self, connection: BlockingConnection):
        try:
            if connection is None or connection.is_closed:
                return
            connection.close()
        except Exception:
            logging.error('Cannot closing RMQ connections', exc_info=True)
