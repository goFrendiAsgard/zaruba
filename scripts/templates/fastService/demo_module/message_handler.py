from typing import Any, Callable, List, Mapping
from sqlalchemy.orm import sessionmaker, Session
from sqlalchemy.engine import Engine
from demo_module import crud, schemas, models

import database, transport


def init(mb: transport.MessageBus, engine: Engine, DBSession: sessionmaker):

    models.Base.metadata.create_all(bind=engine)

    @transport.handle(mb, 'hit')
    def receive_hit(msg: Any):
        print(msg)


    @transport.handle_rpc(mb, 'hello.rpc')
    def hello(name: str):
        return 'hello ' + name
    

    @transport.handle_rpc(mb, 'create_user')
    @database.handle(DBSession)
    def create_user(db: Session, user_dict: Mapping[str, Any]) -> Mapping[str, Any]:
        db_user = crud.create_user(db, user = schemas.UserCreate.parse_obj(user_dict))
        return schemas.User.from_orm(db_user).dict()


    @transport.handle_rpc(mb, 'get_user')
    @database.handle(DBSession)
    def get_user(db: Session, user_id: int) -> Mapping[str, Any]:
        db_user = crud.get_user(db, user_id = user_id)
        return schemas.User.from_orm(db_user).dict()


    @transport.handle_rpc(mb, 'list_user')
    @database.handle(DBSession)
    def list_user(db: Session, skip: int = 0, limit: int = 100) -> List[Mapping[str, Any]]:
        db_user_list = crud.list_user(db, skip = skip, limit = limit)
        return [schemas.User.from_orm(db_user).dict() for db_user in db_user_list]
