from typing import Mapping, List, Any, Optional
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from schema.menu_context import MenuContext
from schema.user import User
from schema.auth_type import AuthType

import traceback
import sys


################################################
# -- ⚙️ API
################################################
# Note: 🤖 Don't delete the following line; Zaruba uses it for pattern matching
def register_ztpl_app_module_name_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    print('Register ztplAppModuleName api route handler', file=sys.stderr)


################################################
# -- 👓 User Interface
################################################
# Note: 🤖 Don't delete the following line; Zaruba uses it for pattern matching
def register_ztpl_app_module_name_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # Note: 🤖 Don't delete the following line; Zaruba uses it for pattern matching
    menu_service.add_menu(name='ztplAppModuleName', title='ZtplAppModuleName', url='#', auth_type=AuthType.ANYONE)

    print('Register ztplAppModuleName UI route handler', file=sys.stderr)