from typing import Any, List, Mapping
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from auth.authService import AuthService
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from fastapi.security import OAuth2PasswordRequestForm
from pydantic import BaseModel
from schemas.menuContext import MenuContext
from ui.menuService import MenuService

import traceback

class TokenResponse(BaseModel):
    access_token: str
    token_type: str


def register_login_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, menu_service: MenuService, templates: Jinja2Templates, enable_ui: bool, access_token_url: str):

    @app.post(access_token_url, response_model=TokenResponse)
    async def login(form_data: OAuth2PasswordRequestForm = Depends()):
        try:
            access_token = rpc.call('get_user_token', form_data.username, form_data.password)
            return TokenResponse(access_token = access_token, token_type = 'bearer')
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')
    

    ################################################
    # -- 👓 User Interface
    ################################################
    if enable_ui:
        @app.get('/login', response_class=HTMLResponse)
        async def user_interface(request: Request, context: MenuContext = Depends(menu_service.everyone('login'))):
            return templates.TemplateResponse(
                'default_login.html', 
                context={
                    'request': request, 
                    'context': context,
                    'access_token_url': access_token_url
                }, 
                status_code=200
            )
