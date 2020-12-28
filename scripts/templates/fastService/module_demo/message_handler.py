from typing import Any, Callable
from fastapi import FastAPI
from sqlalchemy.orm import sessionmaker, Session
from sqlalchemy.engine import Engine
from module_demo import crud, schemas, models

import database, transport


def init(mb: transport.MessageBus, engine: Engine, DBSession: sessionmaker):

    models.Base.metadata.create_all(bind=engine)

    @transport.handle(mb, 'hit')
    def receive_hit(msg: Any):
        print(msg)


    @transport.handle_rpc(mb, 'hello.rpc')
    def hello(name: str):
        return "hello " + name
    

    @transport.handle_rpc(mb, 'create_user')
    @database.handle(DBSession)
    def create_user(db: Session, user: schemas.UserCreate):
        return crud.create_user(db, user = user)


    @transport.handle_rpc(mb, 'get_user')
    @database.handle(DBSession)
    def get_user(db: Session, user_id: int):
        return crud.get_user(db, user_id = user_id)


    @transport.handle_rpc(mb, 'get_users')
    @database.handle(DBSession)
    def get_users(db: Session, skip: int = 0, limit: int = 100):
        return crud.get_users(db, skip = skip, limit = limit)
