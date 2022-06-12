from auth.roleRoute import register_role_route
from auth.userRoute import register_user_route
from typing import Mapping, List, Any
from pydantic import BaseModel
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2PasswordRequestForm
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from auth.authService import AuthService
from ui.menuService import MenuService
from helpers.transport import MessageBus, RPC

import traceback

class TokenResponse(BaseModel):
    access_token: str
    token_type: str


def register_auth_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, access_token_url: str, auth_service: AuthService, menu_service: MenuService, templates: Jinja2Templates, enable_ui: bool):

    register_role_route(app, mb, rpc, auth_service, menu_service, templates, enable_ui)
    register_user_route(app, mb, rpc, auth_service, menu_service, templates, enable_ui)

    @app.post(access_token_url, response_model=TokenResponse)
    async def login(form_data: OAuth2PasswordRequestForm = Depends()):
        try:
            access_token = rpc.call('get_user_token', form_data.username, form_data.password)
            return TokenResponse(access_token = access_token, token_type = 'bearer')
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')

    if enable_ui:
        @app.get('/auth', response_class=HTMLResponse)
        async def user_interface(request: Request, current_user = Depends(auth_service.everyone())):
            accessible_menu = menu_service.get_accessible_menu('auth', current_user)
            return templates.TemplateResponse(
                'default_page.html', 
                context={
                    'request': request, 
                    'menu': accessible_menu
                }, 
                status_code=200
            )

    print('Register auth route handler')