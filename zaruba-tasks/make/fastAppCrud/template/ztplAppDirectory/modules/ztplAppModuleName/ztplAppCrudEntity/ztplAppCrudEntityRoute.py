from typing import Any, List, Mapping, Optional
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData, ZtplAppCrudEntityResult
from schemas.menuContext import MenuContext
from schemas.user import User
from schemas.authType import AuthType

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_ztpl_app_crud_entity_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get('/api/v1/ztpl-app-crud-entities/', response_model=ZtplAppCrudEntityResult)
    async def find_ztpl_app_crud_entities(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:ztpl_app_crud_entity:read'))) -> ZtplAppCrudEntityResult:
        '''
        Serving API to find ztplAppCrudEntities by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_ztpl_app_crud_entity', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ZtplAppCrudEntityResult.parse_obj(result)


    @app.get('/api/v1/ztpl-app-crud-entities/{id}', response_model=ZtplAppCrudEntity)
    async def find_ztpl_app_crud_entity_by_id(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:ztpl_app_crud_entity:read'))) -> ZtplAppCrudEntity:
        '''
        Serving API to find ztplAppCrudEntity by id.
        '''
        ztpl_app_crud_entity = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            ztpl_app_crud_entity = rpc.call('find_ztpl_app_crud_entity_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


    @app.post('/api/v1/ztpl-app-crud-entities/', response_model=ZtplAppCrudEntity)
    async def insert_ztpl_app_crud_entity(ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user: Optional[User] = Depends(auth_service.has_permission('api:ztpl_app_crud_entity:create'))) -> ZtplAppCrudEntity:
        '''
        Serving API to insert new ztplAppCrudEntity.
        '''
        ztpl_app_crud_entity = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            ztpl_app_crud_entity = rpc.call('insert_ztpl_app_crud_entity', ztpl_app_crud_entity_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


    @app.put('/api/v1/ztpl-app-crud-entities/{id}', response_model=ZtplAppCrudEntity)
    async def update_ztpl_app_crud_entity(id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user: Optional[User] = Depends(auth_service.has_permission('api:ztpl_app_crud_entity:update'))) -> ZtplAppCrudEntity:
        '''
        Serving API to update ztplAppCrudEntity by id.
        '''
        ztpl_app_crud_entity = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            ztpl_app_crud_entity = rpc.call('update_ztpl_app_crud_entity', id, ztpl_app_crud_entity_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


    @app.delete('/api/v1/ztpl-app-crud-entities/{id}')
    async def delete_ztpl_app_crud_entity(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:ztpl_app_crud_entity:delete'))) -> ZtplAppCrudEntity:
        '''
        Serving API to delete ztplAppCrudEntity by id.
        '''
        ztpl_app_crud_entity = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            ztpl_app_crud_entity = rpc.call('delete_ztpl_app_crud_entity', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='internal Server Error')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


################################################
# -- 👓 User Interface
################################################
def register_ztpl_app_crud_entity_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # ZtplAppCrudEntity CRUD page
    menu_service.add_menu(name='ztplAppModuleName:ztplAppCrudEntities', title='ZtplAppCrudEntities', url='/ztpl-app-module-name/ztpl-app-crud-entities', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:ztplAppModuleName:ztplAppCrudEntity', parent_name='ztplAppModuleName')
    @app.get('/ztpl-app-module-name/ztpl-app-crud-entities', response_class=HTMLResponse)
    async def manage_ztpl_app_crud_entity(request: Request, context: MenuContext = Depends(menu_service.has_access('ztplAppModuleName:ztplAppCrudEntities'))):
        '''
        Serving user interface for managing ztplAppCrudEntity.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/ztplAppModuleName/crud/ztpl_app_crud_entities.html',
            'request': request, 
            'context': context
        }, status_code=200)
