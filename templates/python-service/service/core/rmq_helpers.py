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


def link_rmq_status_to_app(app: App, rmq_connection_list: List[pika.BlockingConnection]) -> None:

    def check_closed():
        is_ok = True
        while is_ok:
            for rmq_connection in rmq_connection_list:
                if rmq_connection.is_closed:
                    app.set_liveness(False)
                    app.set_readiness(False)
                    is_ok = False
                    break

    thread = threading.Thread(target=check_closed)
    thread.start()
