from modules.auth.role import register_role_api_route, register_role_ui_route
from modules.auth.user import register_user_api_route, register_user_ui_route
from typing import Mapping, List, Any, Optional
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from schemas.menuContext import MenuContext

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

    register_role_ui_route(app, mb, rpc, menu_service, page_template)
    register_user_ui_route(app, mb, rpc, menu_service, page_template)

    print('Register auth UI route handler', file=sys.stderr)