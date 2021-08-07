from typing import Any, List, Mapping
from helpers.transport import MessageBus
from fastapi import FastAPI, HTTPException
from schemas.zarubaEntityName import ZarubaEntityName, ZarubaEntityNameData

import traceback

def handle_route(app: FastAPI, mb: MessageBus):

    @app.get('/zaruba_entity_name/', response_model=List[ZarubaEntityName])
    def find_zaruba_entity_name(keyword: str='', limit: int=100, offset: int=0):
        results = []
        try:
            results = mb.call_rpc('find_zaruba_entity_name', keyword, limit, offset)
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return [ZarubaEntityName.parse_obj(result) for result in results]


    @app.get('/zaruba_entity_name/{id}', response_model=ZarubaEntityName)
    def find_zaruba_entity_name_by_id(id: str):
        result = None
        try:
            result = mb.call_rpc('find_zaruba_entity_name_by_id', id)
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZarubaEntityName.parse_obj(result)


    @app.post('/zaruba_entity_name/', response_model=ZarubaEntityName)
    def insert_zaruba_entity_name(data: ZarubaEntityNameData):
        result = None
        try:
            result = mb.call_rpc('insert_zaruba_entity_name', data.dict())
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZarubaEntityName.parse_obj(result)


    @app.put('/zaruba_entity_name/{id}', response_model=ZarubaEntityName)
    def update_zaruba_entity_name(id: str, data: ZarubaEntityNameData):
        result = None
        try:
            result = mb.call_rpc('update_zaruba_entity_name', id, data.dict())
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZarubaEntityName.parse_obj(result)


    @app.delete('/zaruba_entity_name/{id}')
    def delete_zaruba_entity_name(id: str):
        result = None
        try:
            result = mb.call_rpc('delete_zaruba_entity_name', id)
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return ZarubaEntityName.parse_obj(result)


    print('Handle route for zarubaEntityName')
