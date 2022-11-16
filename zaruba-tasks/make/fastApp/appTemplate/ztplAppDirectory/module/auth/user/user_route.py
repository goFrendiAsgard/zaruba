from os import access
from typing import Any, List, Mapping, Optional
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from schema.menu_context import MenuContext
from schema.user import User, UserData, UserResult
from schema.auth_type import AuthType

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_user_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get('/api/v1/users/', response_model=UserResult)
    def find_user(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:user:read'))) -> UserResult:
        '''
        Serving API to find users by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_users', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return UserResult.parse_obj(result)


    @app.get('/api/v1/users/{id}', response_model=User)
    def find_user_by_id(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:user:read'))) -> User:
        '''
        Serving API to find user by id.
        '''
        result = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_user_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return User.parse_obj(result)


    @app.post('/api/v1/users/', response_model=User)
    def insert_user(data: UserData, current_user: Optional[User] = Depends(auth_service.has_permission('api:user:create'))) -> User:
        '''
        Serving API to insert new user.
        '''
        result = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('insert_user', data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return User.parse_obj(result)


    @app.put('/api/v1/users/{id}', response_model=User)
    def update_user(id: str, data: UserData, current_user: Optional[User] = Depends(auth_service.has_permission('api:user:update'))) -> User:
        '''
        Serving API to update user by id.
        '''
        result = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('update_user', id, data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return User.parse_obj(result)


    @app.delete('/api/v1/users/{id}')
    def delete_user(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:user:delete'))) -> User:
        '''
        Serving API to delete user by id.
        '''
        result = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('delete_user', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return User.parse_obj(result)


################################################
# -- 👓 User Interface
################################################
def register_user_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # User CRUD page
    menu_service.add_menu(name='auth:users', title='Users', url='/auth/users', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:auth:user', parent_name='auth')
    @app.get('/auth/users', response_class=HTMLResponse)
    async def manage_user(request: Request, context: MenuContext = Depends(menu_service.has_access('auth:users'))):
        '''
        Serving user interface to manage user.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/auth/crud/users.html',
            'request': request, 
            'context': context
        }, status_code=200)
