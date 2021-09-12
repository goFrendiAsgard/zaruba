from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from sqlalchemy import create_engine
from helpers.transport import RMQEventMap
from configs.helper import get_abs_static_dir, create_message_bus, create_rpc

import os

db_url = os.getenv('ZARUBA_SERVICE_NAME_SQLALCHEMY_DATABASE_URL', 'sqlite://')
rmq_host = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_HOST', 'localhost')
rmq_user = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_USER', 'root')
rmq_pass = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_PASS', 'toor')
rmq_vhost = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_VHOST', '/')
rmq_event_map = RMQEventMap({})

mb_type = os.getenv('ZARUBA_SERVICE_NAME_MESSAGE_BUS_TYPE', 'local')
rpc_type = os.getenv('ZARUBA_SERVICE_NAME_RPC_TYPE', 'local')

enable_http_handler = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_HTTP_HANDLER', '1') != '0'
enable_event_handler = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_RPC_HANDLER', '1') != '0'
static_url = os.getenv('ZARUBA_SERVICE_NAME_STATIC_URL', '/static')
static_dir = get_abs_static_dir(os.getenv('ZARUBA_SERVICE_NAME_STATIC_DIR', ''))

engine = create_engine(db_url, echo=True)
app = FastAPI(title='zarubaServiceName')
mb = create_message_bus(mb_type, rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)
rpc = create_rpc(rpc_type, rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)

@app.on_event('shutdown')
def on_shutdown():
    mb.shutdown()
    rpc.shutdown()
 
if static_dir != '':
    app.mount(static_url, StaticFiles(directory=static_dir), name='static')