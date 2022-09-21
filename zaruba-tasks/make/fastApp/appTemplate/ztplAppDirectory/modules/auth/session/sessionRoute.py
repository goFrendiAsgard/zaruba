from typing import Any, List, Mapping
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from modules.auth.auth.authService import AuthService
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from fastapi.security import OAuth2PasswordRequestForm
from pydantic import BaseModel
from schemas.menuContext import MenuContext
from modules.ui import MenuService

import traceback
import sys

class CreateAccessTokenRequest(BaseModel):
    username: str
    password: str

class CreateAccessTokenResponse(BaseModel):
    access_token: str
    token_type: str

class RefreshAccessTokenRequest(BaseModel):
    access_token: str

class RefreshAccessTokenResponse(BaseModel):
    access_token: str
    token_type: str


################################################
# -- ⚙️ API
################################################
def register_session_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, create_oauth_access_token_url_path: str, create_access_token_url_path: str, renew_access_token_url_path: str):

    @app.post(create_oauth_access_token_url_path, response_model=CreateAccessTokenResponse)
    async def create_oauth_access_token(form_data: OAuth2PasswordRequestForm = Depends()):
        try:
            username = form_data.username
            password = form_data.password
            access_token = rpc.call('create_access_token', username, password)
            return CreateAccessTokenResponse(access_token = access_token, token_type = 'bearer')
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')

    @app.post(create_access_token_url_path, response_model=CreateAccessTokenResponse)
    async def create_access_token(data: CreateAccessTokenRequest):
        try:
            username = data.username
            password = data.password
            access_token = rpc.call('create_access_token', username, password)
            return CreateAccessTokenResponse(access_token = access_token, token_type = 'bearer')
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')

    @app.post(renew_access_token_url_path, response_model=RefreshAccessTokenResponse)
    async def refresh_access_token(data: RefreshAccessTokenRequest):
        try:
            old_access_token = data.access_token
            new_access_token = rpc.call('refresh_access_token', old_access_token)
            return RefreshAccessTokenResponse(access_token = new_access_token, token_type = 'bearer')
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')


################################################
# -- 👓 User Interface
################################################
def register_session_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates, create_access_token_url_path: str):

    @app.get('/account/login', response_class=HTMLResponse)
    async def user_interface(request: Request, context: MenuContext = Depends(menu_service.authenticate('account:login'))):
        return page_template.TemplateResponse(
            'default_login.html', 
            context={
                'request': request, 
                'context': context,
                'create_acess_token_path': create_access_token_url_path
            }, 
            status_code=200
        )

    @app.get('/account/logout', response_class=HTMLResponse)
    async def user_interface(request: Request, context: MenuContext = Depends(menu_service.authenticate('account:logout'))):
        return page_template.TemplateResponse(
            'default_logout.html', 
            context={
                'request': request, 
                'context': context,
            }, 
            status_code=200
        )


