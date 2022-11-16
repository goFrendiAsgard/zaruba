from module.auth.role import register_role_api_route, register_role_ui_route
from module.auth.user import register_user_api_route, register_user_ui_route
from typing import Mapping, List, Any, Optional
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from schema.menu_context import MenuContext
from schema.auth_type import AuthType

import sys

################################################
# -- ⚙️ API
################################################
def register_auth_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    register_role_api_route(app, mb, rpc, auth_service)
    register_user_api_route(app, mb, rpc, auth_service)

    print('Register auth api route handler', file=sys.stderr)


################################################
# -- 👓 User Interface
################################################
def register_auth_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates, create_access_token_url_path: str):

    # registering menu
    menu_service.add_menu(name='auth', title='Security', url='#', auth_type=AuthType.ANYONE)

    register_role_ui_route(app, mb, rpc, menu_service, page_template)
    register_user_ui_route(app, mb, rpc, menu_service, page_template)

    print('Register auth UI route handler', file=sys.stderr)