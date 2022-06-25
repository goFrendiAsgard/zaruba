from pika.adapters.blocking_connection import BlockingConnection

import pika, time, threading

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
        self.should_check_connection = True
        self.is_shutdown = False
        self._connect()

    def _connect(self):
        self.connection: BlockingConnection = pika.BlockingConnection(self.connection_parameters)
        self.connection.add_callback_threadsafe(self._callback)
        self.connection.process_data_events()

    def _process_data_events(self):
        while True:
            time.sleep(5)
            if not self.should_check_connection:
                break
            try:
                self.connection.process_data_events()
            except:
                self._connect()

    def _callback(self):
        self.thread = threading.Thread(target=self._process_data_events, daemon=True)
        self.thread.start()

    def shutdown(self):
        if self.is_shutdown:
            return
        self.should_check_connection = False
        print('closing RMQ connection')
        self.connection.close()
        self.thread.join()
        self.is_shutdown = True
