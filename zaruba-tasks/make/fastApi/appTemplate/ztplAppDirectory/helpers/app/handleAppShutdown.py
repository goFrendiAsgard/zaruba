from fastapi import FastAPI
from helpers.transport import MessageBus, RPC

def handle_app_shutdown(app: FastAPI, mb: MessageBus, rpc: RPC):
    @app.on_event('shutdown')
    def on_shutdown():
        mb.shutdown()
        rpc.shutdown()
    print('Register app shutdown handler')
