from typing import Any, List, Mapping, Optional
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from fastapi.security import OAuth2PasswordRequestForm
from pydantic import BaseModel
from schemas.menuContext import MenuContext
from schemas.user import User
from core.security.authService import AuthService
from core.menu.menuService import MenuService

import traceback
import sys

class CreateAccessTokenRequest(BaseModel):
    username: str
    password: str

class CreateAccessTokenResponse(BaseModel):
    access_token: str
    token_type: str

class RenewAccessTokenRequest(BaseModel):
    access_token: str

class RenewAccessTokenResponse(BaseModel):
    access_token: str
    token_type: str


################################################
# -- ⚙️ API
################################################
def register_session_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, create_oauth_access_token_url_path: str, create_access_token_url_path: str, renew_access_token_url_path: str):

    @app.post(create_oauth_access_token_url_path, response_model=CreateAccessTokenResponse)
    async def create_oauth_access_token(form_data: OAuth2PasswordRequestForm = Depends(), current_user: Optional[User] = Depends(auth_service.anyone())):
        '''
        Serving API to create new access token.
        Preferable if client use OAuth2 mechanism.
        '''
        try:
            if not current_user:
                current_user = User.parse_obj(rpc.call('get_guest_user'))
            username = form_data.username
            password = form_data.password
            access_token = rpc.call('create_access_token', username, password)
            return CreateAccessTokenResponse(access_token = access_token, token_type = 'bearer')
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')


    @app.post(create_access_token_url_path, response_model=CreateAccessTokenResponse)
    async def create_access_token(data: CreateAccessTokenRequest, current_user: Optional[User] = Depends(auth_service.anyone())):
        '''
        Serving API to create new access token.
        Preferable if client if client user AJAX or similar mechanism. 
        '''
        try:
            if not current_user:
                current_user = User.parse_obj(rpc.call('get_guest_user'))
            username = data.username
            password = data.password
            access_token = rpc.call('create_access_token', username, password)
            return CreateAccessTokenResponse(access_token = access_token, token_type = 'bearer')
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')


    @app.post(renew_access_token_url_path, response_model=RenewAccessTokenResponse)
    async def renew_access_token(data: RenewAccessTokenRequest, current_user: Optional[User] = Depends(auth_service.is_user())):
        '''
        Serving API to renew access token.
        '''
        try:
            if not current_user:
                current_user = User.parse_obj(rpc.call('get_guest_user'))
            old_access_token = data.access_token
            new_access_token = rpc.call('renew_access_token', old_access_token, current_user.dict())
            return RenewAccessTokenResponse(access_token = new_access_token, token_type = 'bearer')
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=400, detail='Invalid token')


################################################
# -- 👓 User Interface
################################################
def register_session_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates, create_access_token_url_path: str):

    @app.get('/account/login', response_class=HTMLResponse)
    async def login(request: Request, context: MenuContext = Depends(menu_service.has_access('account:login'))):
        '''
        Serving user interface for login.
        '''
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
    async def logout(request: Request, context: MenuContext = Depends(menu_service.has_access('account:logout'))):
        '''
        Serving user interface for logout.
        '''
        return page_template.TemplateResponse(
            'default_logout.html', 
            context={
                'request': request, 
                'context': context,
            }, 
            status_code=200
        )


