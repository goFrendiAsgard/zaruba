from fastapi import FastAPI

import os
import transport

# init app
app = FastAPI()

# init messagebus
mb: transport.MessageBus = transport.RMQMessageBus(
    rmq_host = os.getenv("SERVICE_RABBITMQ_HOST", "localhost"),
    rmq_user = os.getenv("SERVICE_RABBITMQ_USER", "root"),
    rmq_pass = os.getenv("SERVICE_RABBITMQ_PASS", "toor"),
    rmq_vhost = os.getenv("SERVICE_RABBITMQ_VHOST", "/")
)