from typing import Mapping, List, Any
from fastapi import FastAPI, HTTPException
from helpers.transport import MessageBus

import traceback

class Controller():

    def __init__(self, app: FastAPI, mb: MessageBus, enable_route: bool, enable_event: bool):
        self.app = app
        self.mb = mb
        self.enable_route = enable_route
        self.enable_event = enable_event


    def start(self):
        if self.enable_event:
            self.handle_event()
        if self.enable_route:
            self.handle_route()
    

    def handle_event(self):
        print('Handle events for zarubaModuleName')
    

    def handle_route(self):
        print('Handle routes for zarubaModuleName')


