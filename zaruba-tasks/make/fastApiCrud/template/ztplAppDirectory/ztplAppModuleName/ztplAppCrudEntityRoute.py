from typing import Any, List, Mapping
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, HTTPException
from fastapi.security import OAuth2
from auth.authModel import AuthModel
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData

import traceback

def register_ztpl_app_crud_entity_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_model: AuthModel):

    @app.get('/ztpl_app_crud_entities/', response_model=List[ZtplAppCrudEntity])
    def find_ztpl_app_crud_entity(keyword: str='', limit: int=100, offset: int=0, current_user = Depends(auth_model.has_any_permissions( 'ztpl_app_crud_entity:read'))) -> List[ZtplAppCrudEntity]:
        results = []
        try:
            results = rpc.call('find_ztpl_app_crud_entity', keyword, limit, offset, current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return [ZtplAppCrudEntity.parse_obj(result) for result in results]

    @app.get('/ztpl_app_crud_entities/{id}', response_model=ZtplAppCrudEntity)
    def find_ztpl_app_crud_entity_by_id(id: str, current_user = Depends(auth_model.has_any_permissions( 'ztpl_app_crud_entity:read'))) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('find_ztpl_app_crud_entity_by_id', id, current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)

    @app.post('/ztpl_app_crud_entities/', response_model=ZtplAppCrudEntity)
    def insert_ztpl_app_crud_entity(ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user = Depends(auth_model.has_any_permissions( 'ztpl_app_crud_entity:create'))) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('insert_ztpl_app_crud_entity', ztpl_app_crud_entity_data.dict(), current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)

    @app.put('/ztpl_app_crud_entities/{id}', response_model=ZtplAppCrudEntity)
    def update_ztpl_app_crud_entity(id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user = Depends(auth_model.has_any_permissions( 'ztpl_app_crud_entity:update'))) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('update_ztpl_app_crud_entity', id, ztpl_app_crud_entity_data.dict(), current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)

    @app.delete('/ztpl_app_crud_entities/{id}')
    def delete_ztpl_app_crud_entity(id: str, current_user = Depends(auth_model.has_any_permissions( 'ztpl_app_crud_entity:delete'))) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('delete_ztpl_app_crud_entity', id, current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)

    print('Handle HTTP routes for ztplAppModuleName.ZtplAppCrudEntity')