from typing import Any, List, Mapping, Optional
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schemas.contentAttribute import ContentAttribute, ContentAttributeData, ContentAttributeResult
from schemas.menuContext import MenuContext
from schemas.user import User

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_content_attribute_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get('/api/v1/content_attributes/', response_model=ContentAttributeResult)
    async def find_content_attributes(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_attribute:read'))) -> ContentAttributeResult:
        '''
        Serving API to find content_attributes by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_content_attribute', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ContentAttributeResult.parse_obj(result)


    @app.get('/api/v1/content_attributes/{id}', response_model=ContentAttribute)
    async def find_content_attribute_by_id(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_attribute:read'))) -> ContentAttribute:
        '''
        Serving API to find contentAttribute by id.
        '''
        content_attribute = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_attribute = rpc.call('find_content_attribute_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ContentAttribute.parse_obj(content_attribute)


    @app.post('/api/v1/content_attributes/', response_model=ContentAttribute)
    async def insert_content_attribute(content_attribute_data: ContentAttributeData, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_attribute:create'))) -> ContentAttribute:
        '''
        Serving API to insert new contentAttribute.
        '''
        content_attribute = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_attribute = rpc.call('insert_content_attribute', content_attribute_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ContentAttribute.parse_obj(content_attribute)


    @app.put('/api/v1/content_attributes/{id}', response_model=ContentAttribute)
    async def update_content_attribute(id: str, content_attribute_data: ContentAttributeData, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_attribute:update'))) -> ContentAttribute:
        '''
        Serving API to update contentAttribute by id.
        '''
        content_attribute = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_attribute = rpc.call('update_content_attribute', id, content_attribute_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ContentAttribute.parse_obj(content_attribute)


    @app.delete('/api/v1/content_attributes/{id}')
    async def delete_content_attribute(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:content_attribute:delete'))) -> ContentAttribute:
        '''
        Serving API to delete contentAttribute by id.
        '''
        content_attribute = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content_attribute = rpc.call('delete_content_attribute', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ContentAttribute.parse_obj(content_attribute)


################################################
# -- 👓 User Interface
################################################
def register_content_attribute_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    @app.get('/cms/content-attributes', response_class=HTMLResponse)
    async def manage_content_attribute(request: Request, context: MenuContext = Depends(menu_service.has_access('cms:contentAttributes'))):
        '''
        Serving user interface for managing contentAttribute.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'api_path': '/api/vi/ztp_app_crud_entities',
            'content_path': 'modules/cms/crud/content_attributes.html',
            'request': request, 
            'context': context
        }, status_code=200)
