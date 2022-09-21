from helpers.transport import RMQEventMap, create_rmq_connection_parameters

import os

################################################
# -- 🐇 Rabbitmq setting
################################################

rmq_connection_parameters = create_rmq_connection_parameters(
    host = os.getenv('APP_RABBITMQ_HOST', 'localhost'),
    user = os.getenv('APP_RABBITMQ_USER', ''),
    password = os.getenv('APP_RABBITMQ_PASS', ''),
    virtual_host = os.getenv('APP_RABBITMQ_VHOST', '/'),
    heartbeat=60,
    blocked_connection_timeout=30
)

rmq_event_map = RMQEventMap({})

