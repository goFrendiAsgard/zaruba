from fastapi import FastAPI

import os
import transport


def init_mb(message_bus_type: str) -> transport.MessageBus:
    if message_bus_type == 'rmq':
        return transport.RMQMessageBus(
            rmq_host = os.getenv("ZARUBA_ENV_PREFIX_RABBITMQ_HOST", "localhost"),
            rmq_user = os.getenv("ZARUBA_ENV_PREFIX_RABBITMQ_USER", "root"),
            rmq_pass = os.getenv("ZARUBA_ENV_PREFIX_RABBITMQ_PASS", "toor"),
            rmq_vhost = os.getenv("ZARUBA_ENV_PREFIX_RABBITMQ_VHOST", "/")
        )
    return transport.LocalMessageBus()


# init application component
app = FastAPI()
mb: transport.MessageBus = init_mb(os.getenv("ZARUBA_ENV_PREFIX_MESSAGE_BUS_TYPE", 'local'))