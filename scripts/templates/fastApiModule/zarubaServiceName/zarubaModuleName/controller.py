from typing import Mapping, List, Any
from fastapi import FastAPI, HTTPException
from helpers.transport import MessageBus

import traceback

def route_controller(app: FastAPI, mb: MessageBus):
    print('Handle routes for zarubaModuleName')


def event_controller(mb: MessageBus):
    print('Handle events for zarubaModuleName')