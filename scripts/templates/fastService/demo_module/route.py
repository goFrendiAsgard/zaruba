from typing import List
from fastapi import FastAPI, HTTPException
from demo_module import schema

import transport
import time


def init(app: FastAPI, mb: transport.MessageBus):

    @app.get('/')
    def read_root():
        mb.publish('hit', {'event': 'hit', 'time': time.gmtime()})
        return {'Hello': 'World'}


    @app.get('/hello/{name}')
    def read_item(name: str):
        result = mb.call_rpc('hello.rpc', name)
        return {'name': name, 'result': result}


    @app.get('/users/', response_model=List[schema.User])
    def crud_list_user(skip: int = 0, limit: int = 100):
        db_users = mb.call_rpc('list_user', skip, limit)
        return [schema.User.parse_obj(db_user) for db_user in db_users]

    @app.get('/users/{user_id}', response_model=schema.User)
    def crud_get_user(user_id: int):
        db_user = mb.call_rpc('get_user', user_id)
        if db_user is None:
            raise HTTPException(status_code=404, error='User not found')
        return schema.User.parse_obj(db_user)

    @app.post('/users/', response_model=schema.User)
    def crud_create_user(user_data: schema.UserCreate):
        db_user = mb.call_rpc('create_user', user_data.dict())
        if db_user is None:
            raise HTTPException(status_code=404, error='User not found')
        return schema.User.parse_obj(db_user)

    @app.put('/users/{user_id}', response_model=schema.User)
    def crud_create_user(user_id: int, user_data: schema.UserCreate):
        db_user = mb.call_rpc('update_user', user_id, user_data.dict())
        if db_user is None:
            raise HTTPException(status_code=404, error='User not found')
        return schema.User.parse_obj(db_user)

    @app.delete('/users/{user_id}', response_model=schema.User)
    def crud_get_user(user_id: int):
        db_user = mb.call_rpc('delete_user', user_id)
        if db_user is None:
            raise HTTPException(status_code=404, error='User not found')
        return schema.User.parse_obj(db_user)
