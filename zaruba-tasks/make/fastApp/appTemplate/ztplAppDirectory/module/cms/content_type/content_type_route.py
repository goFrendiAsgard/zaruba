from typing import Any, List, Mapping, Optional
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schema.content_type import ContentType, ContentTypeData, ContentTypeResult
from schema.menu_context import MenuContext
from schema.user import User
from schema.auth_type import AuthType

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_content_type_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get('/api/v1/content-types/', response_model=ContentTypeResult)
    async def find_content_types(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_type:read'))) -> ContentTypeResult:
        '''
        Serving API to find content_types by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_content_type', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ContentTypeResult.parse_obj(result)


    @app.get('/api/v1/content-types/{id}', response_model=ContentType)
    async def find_content_type_by_id(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_type:read'))) -> ContentType:
        '''
        Serving API to find content_type by id.
        '''
        content_type = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_type = rpc.call('find_content_type_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ContentType.parse_obj(content_type)


    @app.post('/api/v1/content-types/', response_model=ContentType)
    async def insert_content_type(content_type_data: ContentTypeData, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_type:create'))) -> ContentType:
        '''
        Serving API to insert new content_type.
        '''
        content_type = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_type = rpc.call('insert_content_type', content_type_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ContentType.parse_obj(content_type)


    @app.put('/api/v1/content-types/{id}', response_model=ContentType)
    async def update_content_type(id: str, content_type_data: ContentTypeData, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_type:update'))) -> ContentType:
        '''
        Serving API to update content_type by id.
        '''
        content_type = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_type = rpc.call('update_content_type', id, content_type_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ContentType.parse_obj(content_type)


    @app.delete('/api/v1/content-types/{id}')
    async def delete_content_type(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_type:delete'))) -> ContentType:
        '''
        Serving API to delete content_type by id.
        '''
        content_type = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_type = rpc.call('delete_content_type', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ContentType.parse_obj(content_type)


################################################
# -- 👓 User Interface
################################################
def register_content_type_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # ContentType CRUD page
    menu_service.add_menu(name='cms:content_types', title='Content Types', url='/cms/content-types', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:cms:content_type', parent_name='cms')
    @app.get('/cms/content-types', response_class=HTMLResponse)
    async def manage_content_type(request: Request, context: MenuContext = Depends(menu_service.has_access('cms:content_types'))):
        '''
        Serving user interface for managing content_type.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/cms/crud/content_types.html',
            'request': request, 
            'context': context
        }, status_code=200)