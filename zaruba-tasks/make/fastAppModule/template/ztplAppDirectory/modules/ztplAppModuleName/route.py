from typing import Mapping, List, Any, Optional
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


################################################
# -- ⚙️ API
################################################
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_ztpl_app_module_name_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService):

    print('Register ztplAppModuleName api route handler')


################################################
# -- 👓 User Interface
################################################
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_ztpl_app_module_name_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates):

    print('Register ztplAppModuleName api route handler')