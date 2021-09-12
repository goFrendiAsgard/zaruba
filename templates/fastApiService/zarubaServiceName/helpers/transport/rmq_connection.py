from pika.adapters.blocking_connection import BlockingConnection

import pika, time, threading

class RMQConnection():

    def __init__(self, rmq_host: str, rmq_user: str, rmq_pass: str, rmq_vhost: str):
        self.rmq_param = pika.ConnectionParameters(
            host=rmq_host,
            credentials=pika.PlainCredentials(rmq_user, rmq_pass),
            virtual_host=rmq_vhost,
            heartbeat=30
        )
        self.should_check_connection = True
        self._connect()

    def _connect(self):
        self.connection: BlockingConnection = pika.BlockingConnection(self.rmq_param)
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
        self.thread = threading.Thread(target=self._process_data_events)
        self.thread.start()

    def shutdown(self):
        self.should_check_connection = False
        self.connection.close()
        self.thread.join()
