from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schemas.menuContext import MenuContext
from schemas.user import User
from modules.auth import AuthService
from modules.ui import MenuService
from helpers.transport import MessageBus, RPC

import traceback
import sys

def register_ztpl_app_module_name_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, menu_service: MenuService, page_template: Jinja2Templates, enable_ui: bool, enable_api:bool):
    # NOTE: follow [this](https://fastapi.tiangolo.com/tutorial/security/first-steps/#how-it-looks) guide for authorization

    print('Register ztplAppModuleName route handler')