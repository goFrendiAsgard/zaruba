from typing import Any, Callable
from fastapi import FastAPI
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

import os
import database, transport

# Import module
import module_demo

# Messagebus selector
def init_mb(message_bus_type: str) -> transport.MessageBus:
    if message_bus_type == 'rmq':
        return transport.RMQMessageBus(
            rmq_host = os.getenv('DEMO_RABBITMQ_HOST', 'localhost'),
            rmq_user = os.getenv('DEMO_RABBITMQ_USER', 'root'),
            rmq_pass = os.getenv('DEMO_RABBITMQ_PASS', 'toor'),
            rmq_vhost = os.getenv('DEMO_RABBITMQ_VHOST', '/')
        )
    return transport.LocalMessageBus()


# Init application component
app = FastAPI()
mb: transport.MessageBus = init_mb(os.getenv('DEMO_MESSAGE_BUS_TYPE', 'local'))
engine = create_engine(
    os.getenv('DEMO_SQLALCHEMY_DATABASE_URL', 'sqlite:///./database.db'),
    connect_args={'check_same_thread': False}
)
DBSession = sessionmaker(autocommit=False, autoflush=False, bind=engine)

# Init module
module_demo.message_handler.init(mb, engine, DBSession)
module_demo.router.init(app, mb)