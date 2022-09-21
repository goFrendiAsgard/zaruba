from typing import Any, List, Mapping
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from modules.auth import AuthService
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from modules.ui import MenuService
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData, ZtplAppCrudEntityResult
from schemas.menuContext import MenuContext
from schemas.user import User

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_ztpl_app_crud_entity_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService):

    @app.get('/api/v1/ztpl_app_crud_entities/', response_model=ZtplAppCrudEntityResult)
    def find_ztpl_app_crud_entities(keyword: str='', limit: int=100, offset: int=0, current_user:  User = Depends(auth_service.is_authorized('api:ztpl_app_crud_entity:read'))) -> ZtplAppCrudEntityResult:
        result = {}
        try:
            result = rpc.call('find_ztpl_app_crud_entity', keyword, limit, offset)
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ZtplAppCrudEntityResult.parse_obj(result)


    @app.get('/api/v1/ztpl_app_crud_entities/{id}', response_model=ZtplAppCrudEntity)
    def find_ztpl_app_crud_entity_by_id(id: str, current_user:  User = Depends(auth_service.is_authorized('api:ztpl_app_crud_entity:read'))) -> ZtplAppCrudEntity:
        ztpl_app_crud_entity = None
        try:
            ztpl_app_crud_entity = rpc.call('find_ztpl_app_crud_entity_by_id', id)
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if ztpl_app_crud_entity is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


    @app.post('/api/v1/ztpl_app_crud_entities/', response_model=ZtplAppCrudEntity)
    def insert_ztpl_app_crud_entity(ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user:  User = Depends(auth_service.is_authorized('api:ztpl_app_crud_entity:create'))) -> ZtplAppCrudEntity:
        ztpl_app_crud_entity = None
        try:
            ztpl_app_crud_entity = rpc.call('insert_ztpl_app_crud_entity', ztpl_app_crud_entity_data.dict(), current_user.dict())
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if ztpl_app_crud_entity is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


    @app.put('/api/v1/ztpl_app_crud_entities/{id}', response_model=ZtplAppCrudEntity)
    def update_ztpl_app_crud_entity(id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user:  User = Depends(auth_service.is_authorized('api:ztpl_app_crud_entity:update'))) -> ZtplAppCrudEntity:
        ztpl_app_crud_entity = None
        try:
            ztpl_app_crud_entity = rpc.call('update_ztpl_app_crud_entity', id, ztpl_app_crud_entity_data.dict(), current_user.dict())
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if ztpl_app_crud_entity is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


    @app.delete('/api/v1/ztpl_app_crud_entities/{id}')
    def delete_ztpl_app_crud_entity(id: str, current_user:  User = Depends(auth_service.is_authorized('api:ztpl_app_crud_entity:delete'))) -> ZtplAppCrudEntity:
        ztpl_app_crud_entity = None
        try:
            ztpl_app_crud_entity = rpc.call('delete_ztpl_app_crud_entity', id, current_user.dict())
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if ztpl_app_crud_entity is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


################################################
# -- 👓 User Interface
################################################
def register_ztpl_app_crud_entity_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates):

    @app.get('/ztpl-app-module-name/ztpl-app-crud-entities', response_class=HTMLResponse)
    async def user_interface(request: Request, context: MenuContext = Depends(menu_service.authenticate('ztplAppModuleName:ztplAppCrudEntities'))):
        return page_template.TemplateResponse('default_crud.html', context={
            'api_path': '/api/vi/ztp_app_crud_entities',
            'content_path': 'ztplAppModuleName/crud/ztpl_app_crud_entities.html',
            'request': request, 
            'context': context
        }, status_code=200)
