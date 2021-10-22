# -- common
from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from sqlalchemy import create_engine
from helpers.transport import RMQEventMap, KafkaEventMap
from configs.helper import get_abs_static_dir, get_rmq_connection_parameters, get_kafka_connection_parameters, create_message_bus, create_rpc

import os

db_url = os.getenv('APP_SQLALCHEMY_DATABASE_URL', 'sqlite://')
rmq_connection_parameters = get_rmq_connection_parameters()
rmq_event_map = RMQEventMap({})
kafka_connection_parameters = get_kafka_connection_parameters()
kafka_event_map = KafkaEventMap({})

mb_type = os.getenv('APP_MESSAGE_BUS_TYPE', 'local')
rpc_type = os.getenv('APP_RPC_TYPE', 'local')

enable_route_handler = os.getenv('APP_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_event_handler = os.getenv('APP_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler = os.getenv('APP_ENABLE_RPC_HANDLER', '1') != '0'
static_url = os.getenv('APP_STATIC_URL', '/static')
static_dir = get_abs_static_dir(os.getenv('APP_STATIC_DIR', ''))

engine = create_engine(db_url, echo=True)
app = FastAPI(title='ztplAppName')
mb = create_message_bus(mb_type, rmq_connection_parameters, rmq_event_map, kafka_connection_parameters, kafka_event_map)
rpc = create_rpc(rpc_type, rmq_connection_parameters, rmq_event_map)

@app.on_event('shutdown')
def on_shutdown():
    mb.shutdown()
    rpc.shutdown()
 
if static_dir != '':
    app.mount(static_url, StaticFiles(directory=static_dir), name='static')