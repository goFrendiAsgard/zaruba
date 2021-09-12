from typing import Mapping, List, Any
from fastapi import FastAPI, HTTPException
from helpers.transport import MessageBus, RPC

import traceback

def http_controller(app: FastAPI, mb: MessageBus, rpc: RPC):
    print('Handle http routes for zarubaModuleName')


def event_controller(mb: MessageBus):
    print('Handle events for zarubaModuleName')


def rpc_controller(rpc: RPC):
    print('Handle rpc for zarubaModuleName')
