from typing import List
from pika.adapters.blocking_connection import BlockingConnection

import pika
import threading
import traceback
import time

def create_rmq_connection_parameters(host: str, user: str, password: str, virtual_host: str = '/', heartbeat: int = 30) -> pika.ConnectionParameters:
    return pika.ConnectionParameters(
        host=host,
        credentials=pika.PlainCredentials(user, password),
        virtual_host=virtual_host,
        heartbeat=heartbeat
    )

class RMQConnection():

    def __init__(self, connection_parameters: pika.ConnectionParameters):
        self._connection_parameters = connection_parameters
        self._should_check_connection = True
        self._is_shutdown = False
        self._connections: List[BlockingConnection] = []
        self._threads: List[threading.Thread] = []

    
    def create_connection(self) -> pika.BlockingConnection:
        connection: BlockingConnection = pika.BlockingConnection(self._connection_parameters)

        def process_data_events():
            while self._should_check_connection:
                time.sleep(5)
                connection.process_data_events()

        def callback():
            thread = threading.Thread(target=process_data_events)
            self._threads.append(thread)
            thread.start()

        connection.add_callback_threadsafe(callback)
        connection.process_data_events()
        self._connections.append(connection)
        return connection


    def shutdown(self):
        if self._is_shutdown:
            return
        self._should_check_connection = False
        print('closing RMQ connection')
        for connection in self._connections:
            try:
                connection.close()
            except:
                print(traceback.format_exc())
        for thread in self._threads:
            thread.join()
        self._is_shutdown = True
