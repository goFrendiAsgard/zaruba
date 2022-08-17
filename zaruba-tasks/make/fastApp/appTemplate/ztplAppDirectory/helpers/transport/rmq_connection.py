from typing import List
from pika.adapters.blocking_connection import BlockingConnection, BlockingChannel

import pika
import threading
import traceback
import time

def create_rmq_connection_parameters(host: str, user: str, password: str, virtual_host: str = '/', heartbeat: int = 60, blocked_connection_timeout: int = 30) -> pika.ConnectionParameters:
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
        self._should_check_connection = True
        self._is_shutdown = False
        self._connections: List[BlockingConnection] = []
        self._threads: List[threading.Thread] = []

    
    def create_connection(self) -> BlockingConnection:
        connection: BlockingConnection = pika.BlockingConnection(self._connection_parameters)

        def process_data_events():
            while self._should_check_connection:
                time.sleep(1)
                connection.process_data_events()

        def callback():
            thread = threading.Thread(target=process_data_events)
            self._threads.append(thread)
            thread.start()

        connection.add_callback_threadsafe(callback)
        connection.process_data_events()
        self._connections.append(connection)
        return connection


    def _stop_connections(self):
        for connection in self._connections:
            try:
                connection.process_data_events()
                connection.close()
            except:
                print('connection is not fully closed')
                # print(traceback.format_exc()) 
    

    def _stop_threads(self):
        for thread in self._threads:
            thread.join()        


    def shutdown(self):
        if self._is_shutdown:
            return
        self._should_check_connection = False
        print('closing RMQ connections')
        self._stop_connections()
        self._stop_threads()
        self._is_shutdown = True
