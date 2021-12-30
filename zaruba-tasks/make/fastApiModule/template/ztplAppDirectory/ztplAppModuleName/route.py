from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, HTTPException
from fastapi.security import OAuth2
from helpers.transport import MessageBus, RPC

import traceback

def register_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, oauth2_scheme: OAuth2):
    # NOTE: follow [this](https://fastapi.tiangolo.com/tutorial/security/first-steps/#how-it-looks) guide for authorization

    print('register ztplAppModuleName route handler')

