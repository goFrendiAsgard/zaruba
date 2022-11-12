from typing import Any, List, Mapping, Optional
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schemas.content import Content, ContentData, ContentResult
from schemas.menuContext import MenuContext
from schemas.user import User
from schemas.authType import AuthType

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_content_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get('/api/v1/contents/', response_model=ContentResult)
    async def find_contents(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:content:read'))) -> ContentResult:
        '''
        Serving API to find contents by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_content', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ContentResult.parse_obj(result)


    @app.get('/api/v1/contents/{id}', response_model=Content)
    async def find_content_by_id(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:content:read'))) -> Content:
        '''
        Serving API to find content by id.
        '''
        content = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content = rpc.call('find_content_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return Content.parse_obj(content)


    @app.post('/api/v1/contents/', response_model=Content)
    async def insert_content(content_data: ContentData, current_user: Optional[User] = Depends(auth_service.has_permission('api:content:create'))) -> Content:
        '''
        Serving API to insert new content.
        '''
        content = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content = rpc.call('insert_content', content_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return Content.parse_obj(content)


    @app.put('/api/v1/contents/{id}', response_model=Content)
    async def update_content(id: str, content_data: ContentData, current_user: Optional[User] = Depends(auth_service.has_permission('api:content:update'))) -> Content:
        '''
        Serving API to update content by id.
        '''
        content = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content = rpc.call('update_content', id, content_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return Content.parse_obj(content)


    @app.delete('/api/v1/contents/{id}')
    async def delete_content(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:content:delete'))) -> Content:
        '''
        Serving API to delete content by id.
        '''
        content = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            content = rpc.call('delete_content', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return Content.parse_obj(content)


################################################
# -- 👓 User Interface
################################################
def register_content_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # register menu
    menu_service.add_menu(name='cms:contents', title='Contents', url='/cms/contents', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:cms:content', parent_name='cms')


    @app.get('/cms/contents', response_class=HTMLResponse)
    async def manage_content(request: Request, context: MenuContext = Depends(menu_service.has_access('cms:contents'))):
        '''
        Serving user interface for managing content.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/cms/crud/contents.html',
            'request': request, 
            'context': context
        }, status_code=200)
