from auth.roleRoute import register_role_route
from auth.userRoute import register_user_route
from auth.loginRoute import register_login_route
from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from auth.authService import AuthService
from ui.menuService import MenuService
from helpers.transport import MessageBus, RPC
from schemas.menuContext import MenuContext

import traceback



def register_auth_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, menu_service: MenuService, templates: Jinja2Templates, enable_ui: bool, enable_api: bool, access_token_url: str):

    register_role_route(app, mb, rpc, auth_service, menu_service, templates, enable_ui, enable_api)
    register_user_route(app, mb, rpc, auth_service, menu_service, templates, enable_ui, enable_api)
    register_login_route(app, mb, rpc, auth_service, menu_service, templates, enable_ui, enable_api, access_token_url)

    ################################################
    # -- 👓 User Interface
    ################################################
    if enable_ui:

        @app.get('/auth', response_class=HTMLResponse)
        async def user_interface(request: Request, context: MenuContext = Depends(menu_service.is_authorized('auth'))):
            return templates.TemplateResponse(
                'default_page.html', 
                context={
                    'request': request, 
                    'context': context
                }, 
                status_code=200
            )

    print('Register auth route handler')