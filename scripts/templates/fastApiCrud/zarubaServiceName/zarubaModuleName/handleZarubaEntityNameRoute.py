from fastapi import FastAPI, HTTPException
from schemas.zarubaEntityName import ZarubaEntityName, ZarubaEntityNameData

def handle_route(app: FastAPI):

    @app.get('/zaruba_entity_name/', response_model=List[ZarubaEntityName])
    def find_zaruba_entity_name(keyword: str='', limit: int=100, offset: int=0):
        try:
            results = mb.call_rpc('find_zaruba_entity_name', keyword, limit, offset)
            return [ZarubaEntityName.parse_obj(result) for result in results]
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')

    @app.get('/zaruba_entity_name/{id}', response_model=ZarubaEntityName)
    def find_zaruba_entity_name_by_id(id: str):
        try:
            result = mb.call_rpc('find_zaruba_entity_name_by_id', id)
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return ZarubaEntityName.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')

    @app.post('/zaruba_entity_name/', response_model=ZarubaEntityName)
    def insert_zaruba_entity_name(data: ZarubaEntityNameData):
        try:
            result = mb.call_rpc('insert_zaruba_entity_name', data.dict())
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return ZarubaEntityName.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')

    @app.put('/zaruba_entity_name/{id}', response_model=ZarubaEntityName)
    def update_zaruba_entity_name(id: str, data: ZarubaEntityNameData):
        try:
            result = mb.call_rpc('update_zaruba_entity_name', id, data.dict())
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return ZarubaEntityName.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')

    @app.delete('/zaruba_entity_name/{id}')
    def delete_zaruba_entity_name(id: str):
        try:
            result = mb.call_rpc('delete_zaruba_entity_name', id)
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return ZarubaEntityName.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
