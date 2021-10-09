from typing import Mapping, List, Any
from fastapi import FastAPI, HTTPException
from helpers.transport import MessageBus, RPC

import traceback

def register_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC):
    print('register ztplAppModuleName route handler')

