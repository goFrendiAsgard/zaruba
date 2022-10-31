from typing import Mapping, List, Any, Optional
from core import AuthService, MenuService
from modules.log.activity import register_activity_api_route, register_activity_ui_route
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schemas.menuContext import MenuContext
from schemas.user import User
from helpers.transport import MessageBus, RPC

import traceback
import sys


################################################
# -- ⚙️ API
################################################
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService):

    register_activity_api_route(app, mb, rpc, auth_service)

    print('Register log api route handler', file=sys.stderr)


################################################
# -- 👓 User Interface
################################################
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates):

    register_activity_ui_route(app, mb, rpc, menu_service, page_template)

    print('Register log UI route handler', file=sys.stderr)