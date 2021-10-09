from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from sqlalchemy import create_engine
from helpers.transport import RMQEventMap, KafkaEventMap, get_rmq_connection_parameters, get_kafka_connection_parameters
from configs.helper import get_abs_static_dir, create_message_bus, create_rpc

import os

db_url = os.getenv('ZARUBA_SERVICE_NAME_SQLALCHEMY_DATABASE_URL', 'sqlite://')
rmq_connection_parameters = get_rmq_connection_parameters(
    host = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_HOST', 'localhost'),
    user = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_USER', 'root'),
    password = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_PASS', 'toor'),
    virtual_host = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_VHOST', '/'),
    heartbeat=30
)
rmq_event_map = RMQEventMap({})
kafka_connection_parameters = get_kafka_connection_parameters(
    bootstrap_servers = os.getenv('ZARUBA_SERVICE_NAME_KAFKA_BOOTSTRAP_SERVERS', 'localhost:9093'),
    sasl_mechanism=os.getenv('ZARUBA_SERVICE_NAME_KAFKA_SASL_MECHANISM', 'PLAIN'),
    sasl_plain_username=os.getenv('ZARUBA_SERVICE_NAME_KAFKA_SASL_PLAIN_USERNAME', ''),
    sasl_plain_password=os.getenv('ZARUBA_SERVICE_NAME_KAFKA_SASL_PLAIN_PASSWORD', '')
)
kafka_event_map = KafkaEventMap({})

mb_type = os.getenv('ZARUBA_SERVICE_NAME_MESSAGE_BUS_TYPE', 'local')
rpc_type = os.getenv('ZARUBA_SERVICE_NAME_RPC_TYPE', 'local')

enable_http_handler = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_HTTP_HANDLER', '1') != '0'
enable_event_handler = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_RPC_HANDLER', '1') != '0'
static_url = os.getenv('ZARUBA_SERVICE_NAME_STATIC_URL', '/static')
static_dir = get_abs_static_dir(os.getenv('ZARUBA_SERVICE_NAME_STATIC_DIR', ''))

engine = create_engine(db_url, echo=True)
app = FastAPI(title='zarubaServiceName')
mb = create_message_bus(mb_type, rmq_connection_parameters, rmq_event_map, kafka_connection_parameters, kafka_event_map)
rpc = create_rpc(rpc_type, rmq_connection_parameters, rmq_event_map)

@app.on_event('shutdown')
def on_shutdown():
    mb.shutdown()
    rpc.shutdown()
 
if static_dir != '':
    app.mount(static_url, StaticFiles(directory=static_dir), name='static')