from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, HTTPException
from fastapi.responses import HTMLResponse
from schemas.user import User
from auth.authService import AuthService
from helpers.transport import MessageBus, RPC

import traceback

def register_ztpl_app_module_name_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService):
    # NOTE: follow [this](https://fastapi.tiangolo.com/tutorial/security/first-steps/#how-it-looks) guide for authorization

    print('Register ztplAppModuleName route handler')

