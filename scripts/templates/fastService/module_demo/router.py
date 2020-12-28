from typing import List
from fastapi import FastAPI, HTTPException
from module_demo import schemas

import transport
import time


def init(app: FastAPI, mb: transport.MessageBus):

    @app.get("/")
    def read_root():
        mb.publish('hit', {"event": "hit", "time": time.gmtime()})
        return {"Hello": "World"}


    @app.get("/hello/{name}")
    def read_item(name: str):
        result = mb.call_rpc("hello.rpc", name)
        return {"name": name, "result": result}


    @app.post("/users/", response_model=schemas.User)
    def create_user(user: schemas.UserCreate):
        return mb.call_rpc('create_user', user)


    @app.get("/users/", response_model=List[schemas.User])
    def read_users(skip: int = 0, limit: int = 100):
        return mb.call_rpc('get_users', skip, limit)


    @app.get("/users/{user_id}", response_model=schemas.User)
    def read_user(user_id: int):
        db_user = mb.call_rpc('get_user', user_id)
        if db_user is None:
            raise HTTPException(status_code=404, detail="User not found")
        return db_user