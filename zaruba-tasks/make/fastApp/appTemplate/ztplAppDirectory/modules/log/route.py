from typing import Mapping, List, Any, Optional
from core import AuthService, MenuService
from modules.log.activity import register_activity_api_route, register_activity_ui_route
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schemas.menuContext import MenuContext
from schemas.user import User
from schemas.authType import AuthType
from transport import AppMessageBus, AppRPC

import traceback
import sys


################################################
# -- ⚙️ API
################################################
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    register_activity_api_route(app, mb, rpc, auth_service)

    print('Register log api route handler', file=sys.stderr)


################################################
# -- 👓 User Interface
################################################
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # register menu
    menu_service.add_menu(name='log', title='Log', url='#', auth_type=AuthType.ANYONE)

    register_activity_ui_route(app, mb, rpc, menu_service, page_template)

    print('Register log UI route handler', file=sys.stderr)