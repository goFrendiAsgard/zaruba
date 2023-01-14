from typing import Optional
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from fastapi.security import OAuth2PasswordRequestForm
from schema.menu_context import MenuContext
from schema.cred_token import (
    CreateCredTokenRequest, CreateCredTokenResponse,
    RenewCredTokenRequest, RenewCredTokenResponse
)
from schema.user import User
from schema.auth_type import AuthType
from core.security.service.auth_service import AuthService
from core.menu.menu_service import MenuService

import logging


################################################
# -- ⚙️ API
################################################
def register_session_api_route(
    app: FastAPI,
    mb: AppMessageBus,
    rpc: AppRPC, auth_service: AuthService,
    create_oauth_cred_token_url: str,
    create_cred_token_url: str,
    renew_cred_token_url: str
):

    @app.post(
        create_oauth_cred_token_url,
        response_model=CreateCredTokenResponse
    )
    async def create_oauth_cred_token(
        form_data: OAuth2PasswordRequestForm = Depends(),
        current_user: Optional[User] = Depends(auth_service.anyone())
    ):
        '''
        Serving API to create new access token.
        Preferable if client use OAuth2 mechanism.
        '''
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            username = form_data.username
            password = form_data.password
            cred_token = rpc.call('create_cred_token', username, password)
            return CreateCredTokenResponse(
                cred_token=cred_token,
                token_type='bearer'
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            logging.error('Non HTTPException error', exc_info=True)
            raise HTTPException(
                status_code=400,
                detail='incorrect identity or password'
            )

    @app.post(
        create_cred_token_url,
        response_model=CreateCredTokenResponse
    )
    async def create_cred_token(
        data: CreateCredTokenRequest,
        current_user: Optional[User] = Depends(auth_service.anyone())
    ):
        '''
        Serving API to create new access token.
        Preferable if client if client user AJAX or similar mechanism.
        '''
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            username = data.username
            password = data.password
            cred_token = rpc.call('create_cred_token', username, password)
            return CreateCredTokenResponse(
                cred_token=cred_token,
                token_type='bearer'
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            logging.error('Non HTTPException error', exc_info=True)
            raise HTTPException(
                status_code=400,
                detail='incorrect identity or password'
            )

    @app.post(
        renew_cred_token_url,
        response_model=RenewCredTokenResponse
    )
    async def renew_cred_token(
        data: RenewCredTokenRequest,
        current_user: Optional[User] = Depends(auth_service.is_user())
    ):
        '''
        Serving API to renew access token.
        '''
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            old_cred_token = data.cred_token
            new_cred_token = rpc.call(
                'renew_cred_token',
                old_cred_token,
                current_user.dict()
            )
            return RenewCredTokenResponse(
                cred_token=new_cred_token,
                token_type='bearer'
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            logging.error('Non HTTPException error', exc_info=True)
            raise HTTPException(status_code=400, detail='invalid token')


################################################
# -- 👓 User Interface
################################################
def register_session_ui_route(
    app: FastAPI,
    mb: AppMessageBus,
    rpc: AppRPC,
    menu_service: MenuService,
    page_template: Jinja2Templates,
    create_cred_token_url: str
):

    # Session menu
    menu_service.add_menu(
        name='account',
        title='Account',
        url='#',
        auth_type=AuthType.ANYONE
    )

    # Login page
    menu_service.add_menu(
        name='account:login',
        title='Log in',
        url='/account/login',
        auth_type=AuthType.VISITOR,
        parent_name='account'
    )

    @app.get('/account/login', response_class=HTMLResponse)
    async def login(
        request: Request,
        context: MenuContext = Depends(
            menu_service.has_access('account:login'))
    ):
        '''
        Serving user interface for login.
        '''
        return page_template.TemplateResponse(
            'default_login.html',
            context={
                'request': request,
                'context': context,
                'create_cred_token_url': create_cred_token_url
            },
            status_code=200
        )

    # Logout page
    menu_service.add_menu(
        name='account:logout',
        title='Log out',
        url='/account/logout',
        auth_type=AuthType.USER,
        parent_name='account'
    )

    @app.get('/account/logout', response_class=HTMLResponse)
    async def logout(
        request: Request,
        context: MenuContext = Depends(
            menu_service.has_access('account:logout')
        )
    ):
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
