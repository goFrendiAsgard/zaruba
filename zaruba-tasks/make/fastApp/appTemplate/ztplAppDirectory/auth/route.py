from auth.roleRoute import register_role_route
from auth.userRoute import register_user_route
from auth.accountRoute import register_account_route
from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from auth.authService import AuthService
from ui.menuService import MenuService
from helpers.transport import MessageBus, RPC
from schemas.menuContext import MenuContext

import traceback
import sys



def register_auth_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, menu_service: MenuService, templates: Jinja2Templates, enable_ui: bool, enable_api: bool, create_oauth_access_token_url, create_access_token_url: str, refresh_access_token_url: str):

    register_role_route(app, mb, rpc, auth_service, menu_service, templates, enable_ui, enable_api)
    register_user_route(app, mb, rpc, auth_service, menu_service, templates, enable_ui, enable_api)
    register_account_route(app, mb, rpc, menu_service, templates, enable_ui, enable_api, create_oauth_access_token_url, create_access_token_url, refresh_access_token_url)

    print('Register auth route handler')