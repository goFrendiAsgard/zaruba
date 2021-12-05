from typing import Any, List, Mapping
from helpers.transport import MessageBus, RPC
from fastapi import FastAPI, HTTPException
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData

import traceback

def register_ztpl_app_crud_entity_route(app: FastAPI, mb: MessageBus, rpc: RPC):

    @app.get('/ztpl_app_crud_entity/', response_model=List[ZtplAppCrudEntity])
    def find_ztpl_app_crud_entity(keyword: str='', limit: int=100, offset: int=0) -> List[ZtplAppCrudEntity]:
        results = []
        try:
            results = rpc.call('find_ztpl_app_crud_entity', keyword, limit, offset)
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return [ZtplAppCrudEntity.parse_obj(result) for result in results]


    @app.get('/ztpl_app_crud_entity/{id}', response_model=ZtplAppCrudEntity)
    def find_ztpl_app_crud_entity_by_id(id: str) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('find_ztpl_app_crud_entity_by_id', id)
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)


    @app.post('/ztpl_app_crud_entity/', response_model=ZtplAppCrudEntity)
    def insert_ztpl_app_crud_entity(data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('insert_ztpl_app_crud_entity', data.dict())
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)


    @app.put('/ztpl_app_crud_entity/{id}', response_model=ZtplAppCrudEntity)
    def update_ztpl_app_crud_entity(id: str, data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('update_ztpl_app_crud_entity', id, data.dict())
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)


    @app.delete('/ztpl_app_crud_entity/{id}')
    def delete_ztpl_app_crud_entity(id: str) -> ZtplAppCrudEntity:
        result = None
        try:
            result = rpc.call('delete_ztpl_app_crud_entity', id)
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZtplAppCrudEntity.parse_obj(result)


    print('Handle HTTP routes for ztplAppModuleName.ZtplAppCrudEntity')
