from typing import Any, List, Mapping
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from auth.authService import AuthService
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from ui.menuService import MenuService
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from schemas.menuContext import MenuContext
from schemas.user import User

import traceback

def register_ztpl_app_crud_entity_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, menu_service: MenuService, templates: Jinja2Templates, enable_ui: bool, enable_api:bool):

    ################################################
    # -- ⚙️ API
    ################################################
    if enable_api:

        @app.get('/api/v1/ztpl_app_crud_entities/', response_model=List[ZtplAppCrudEntity])
        def find_ztpl_app_crud_entity(keyword: str='', limit: int=100, offset: int=0, current_user:  User = Depends(auth_service.is_authorized('api:ztpl_app_crud_entity:read'))) -> List[ZtplAppCrudEntity]:
            ztpl_app_crud_entities = []
            try:
                ztpl_app_crud_entities = rpc.call('find_ztpl_app_crud_entity', keyword, limit, offset, current_user.dict())
            except:
                print(traceback.format_exc()) 
                raise HTTPException(status_code=500, detail='Internal Server Error')
            return [ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity) for ztpl_app_crud_entity in ztpl_app_crud_entities]


        @app.get('/api/v1/ztpl_app_crud_entities/{id}', response_model=ZtplAppCrudEntity)
        def find_ztpl_app_crud_entity_by_id(id: str, current_user:  User = Depends(auth_service.is_authorized('api:ztpl_app_crud_entity:read'))) -> ZtplAppCrudEntity:
            ztpl_app_crud_entity = None
            try:
                ztpl_app_crud_entity = rpc.call('find_ztpl_app_crud_entity_by_id', id, current_user.dict())
            except:
                print(traceback.format_exc()) 
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
                print(traceback.format_exc()) 
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
                print(traceback.format_exc()) 
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
                print(traceback.format_exc()) 
                raise HTTPException(status_code=500, detail='Internal Server Error')
            if ztpl_app_crud_entity is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)


    ################################################
    # -- 👓 User Interface
    ################################################
    if enable_ui:

        @app.get('/ztpl-app-module-name/ztpl-app-crud-entities', response_class=HTMLResponse)
        async def user_interface(request: Request, context: MenuContext = Depends(menu_service.authenticate('ztplAppModuleName:ztplAppCrudEntities'))):
            return templates.TemplateResponse('default_crud.html', context={
                'api_path': '/api/vi/ztp_app_crud_entities',
                'content_path': 'ztplAppModuleName/crud/ztpl_app_crud_entities.html',
                'request': request, 
                'context': context
            }, status_code=200)

    print('Handle HTTP routes for ztplAppModuleName.ZtplAppCrudEntity')