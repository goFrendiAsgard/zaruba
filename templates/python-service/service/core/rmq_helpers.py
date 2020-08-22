from typing import List
import threading
import time
import pika
import random

from .interfaces import App


def create_rmq_connection(rmq_connection_string: str, heartbeat: int) -> pika.BlockingConnection:
    rmq_connection_url = pika.URLParameters(rmq_connection_string)
    rmq_connection = pika.BlockingConnection(rmq_connection_url)
    add_rmq_heartbeat(rmq_connection, heartbeat)
    return rmq_connection


def add_rmq_heartbeat(rmq_connection: pika.BlockingConnection, heartbeat: int) -> None:

    def process_data_events():
        while True:
            time.sleep(heartbeat)
            rmq_connection.process_data_events()

    def callback():
        thread = threading.Thread(target=process_data_events)
        thread.start()

    rmq_connection.add_callback_threadsafe(callback)
    rmq_connection.process_data_events()
