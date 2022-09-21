from modules.auth.role.roleRoute import register_role_api_route, register_role_ui_route
from modules.auth.user.userRoute import register_user_api_route, register_user_ui_route
from modules.auth.session.sessionRoute import register_session_api_route, register_session_ui_route
from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from modules.auth.auth.authService import AuthService
from modules.ui import MenuService
from helpers.transport import MessageBus, RPC
from schemas.menuContext import MenuContext

################################################
# -- ⚙️ API
################################################
def register_auth_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, create_oauth_access_token_url_path, create_access_token_url_path: str, renew_access_token_url_path: str):

    register_role_api_route(app, mb, rpc, auth_service)
    register_user_api_route(app, mb, rpc, auth_service)
    register_session_api_route(app, mb, rpc, create_oauth_access_token_url_path, create_access_token_url_path, renew_access_token_url_path)

    print('Register auth api route handler')


################################################
# -- 👓 User Interface
################################################
def register_auth_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates, create_access_token_url_path: str):

    register_role_ui_route(app, mb, rpc, menu_service, page_template)
    register_user_ui_route(app, mb, rpc, menu_service, page_template)
    register_session_ui_route(app, mb, rpc, menu_service, page_template, create_access_token_url_path)

    print('Register auth api route handler')