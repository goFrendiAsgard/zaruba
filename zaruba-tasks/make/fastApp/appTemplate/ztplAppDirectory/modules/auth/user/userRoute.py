from os import access
from typing import Any, List, Mapping
from modules.auth.auth.authService import AuthService
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from modules.ui import MenuService
from schemas.menuContext import MenuContext
from schemas.user import User, UserData, UserResult

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_user_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService):

    @app.get('/api/v1/users/', response_model=UserResult)
    def find_user(keyword: str='', limit: int=100, offset: int=0, current_user: User = Depends(auth_service.is_authorized('api:user:read'))) -> UserResult:
        result = {}
        try:
            result = rpc.call('find_users', keyword, limit, offset)
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return UserResult.parse_obj(result)

    @app.get('/api/v1/users/{id}', response_model=User)
    def find_user_by_id(id: str, current_user: User = Depends(auth_service.is_authorized('api:user:read'))) -> User:
        result = None
        try:
            result = rpc.call('find_user_by_id', id)
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)

    @app.post('/api/v1/users/', response_model=User)
    def insert_user(data: UserData, current_user: User = Depends(auth_service.is_authorized('api:user:create'))) -> User:
        result = None
        try:
            result = rpc.call('insert_user', data.dict(), current_user.dict())
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)

    @app.put('/api/v1/users/{id}', response_model=User)
    def update_user(id: str, data: UserData, current_user: User = Depends(auth_service.is_authorized('api:user:update'))) -> User:
        result = None
        try:
            result = rpc.call('update_user', id, data.dict(), current_user.dict())
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)

    @app.delete('/api/v1/users/{id}')
    def delete_user(id: str, current_user: User = Depends(auth_service.is_authorized('api:user:delete'))) -> User:
        result = None
        try:
            result = rpc.call('delete_user', id, current_user.dict())
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)


################################################
# -- 👓 User Interface
################################################
def register_user_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates):

    @app.get('/auth/users', response_class=HTMLResponse)
    async def user_interface(request: Request, context: MenuContext = Depends(menu_service.authenticate('auth:users'))):
        return page_template.TemplateResponse('default_crud.html', context={
            'api_path': '/api/vi/users',
            'content_path': 'auth/crud/users.html',
            'request': request, 
            'context': context
        }, status_code=200)
