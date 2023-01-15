from module.auth.role import register_role_api_route, register_role_ui_route
from module.auth.user import register_user_api_route, register_user_ui_route
from fastapi import FastAPI
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from schema.auth_type import AuthType


################################################
# -- ⚙️ API
################################################
def register_auth_api_route(
    app: FastAPI,
    mb: AppMessageBus,
    rpc: AppRPC,
    auth_service: AuthService
):

    register_role_api_route(app, mb, rpc, auth_service)
    register_user_api_route(app, mb, rpc, auth_service)


################################################
# -- 👓 User Interface
################################################
def register_auth_ui_route(
    app: FastAPI,
    mb: AppMessageBus,
    rpc: AppRPC,
    menu_service: MenuService,
    page_template: Jinja2Templates
):

    # registering menu
    menu_service.add_menu(
        name='auth', title='Security', url='#', auth_type=AuthType.ANYONE
    )

    register_role_ui_route(app, mb, rpc, menu_service, page_template)
    register_user_ui_route(app, mb, rpc, menu_service, page_template)
