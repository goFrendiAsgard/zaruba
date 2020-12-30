from fastapi import FastAPI
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

import os
import database, transport


# Handle app shutdown event
def handle_shutdown(app: FastAPI, mb: transport.MessageBus):
    @app.on_event('shutdown')
    def on_shutdown():
        mb.shutdown()

# Messagebus selector
def init_mb(message_bus_type: str) -> transport.MessageBus:
    if message_bus_type == 'rmq':
        return transport.RMQMessageBus(
            rmq_host = os.getenv('ZARUBA_ENV_PREFIX_RABBITMQ_HOST', 'localhost'),
            rmq_user = os.getenv('ZARUBA_ENV_PREFIX_RABBITMQ_USER', 'root'),
            rmq_pass = os.getenv('ZARUBA_ENV_PREFIX_RABBITMQ_PASS', 'toor'),
            rmq_vhost = os.getenv('ZARUBA_ENV_PREFIX_RABBITMQ_VHOST', '/')
        )
    return transport.LocalMessageBus()


# init application component
app = FastAPI()
mb: transport.MessageBus = init_mb(os.getenv('ZARUBA_ENV_PREFIX_MESSAGE_BUS_TYPE', 'local'))
engine = create_engine(
    os.getenv('DEMO_SQLALCHEMY_DATABASE_URL', 'sqlite:///./database.db'),
    connect_args={'check_same_thread': False}
)
DBSession = sessionmaker(autocommit=False, autoflush=False, bind=engine)
handle_shutdown(app, mb)

