from fastapi import FastAPI
from fastapi import Depends, FastAPI, HTTPException
from fastapi.responses import HTMLResponse
from helpers.transport import MessageBus, RPC

def register_readiness_handler(app: FastAPI, mb: MessageBus, rpc: RPC, error_threshold: int):

    @app.get('/readiness')
    def handle_readiness():
        if mb.is_failing():
            raise HTTPException(status_code=500, detail='Messagebus is failing')
        if rpc.is_failing():
            raise HTTPException(status_code=500, detail='RPC is failing')
        if mb.get_error_count() > error_threshold:
            raise HTTPException(status_code=500, detail='Messagebus error exceeding threshold')
        if rpc.get_error_count() > error_threshold:
            raise HTTPException(status_code=500, detail='RPC error exceeding threshold')
        return HTMLResponse(content='ready', status_code=200)