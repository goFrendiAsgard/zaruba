from typing import Any, Callable, List, Mapping
from sqlalchemy.orm import sessionmaker, Session
from demo_module import crud, schema, model

import database, transport


def init(mb: transport.MessageBus, DBSession: sessionmaker):

    @transport.handle(mb, 'hit')
    def receive_hit(msg: Any):
        print(msg)


    @transport.handle_rpc(mb, 'hello.rpc')
    def hello(name: str):
        return 'hello ' + name
    

    @transport.handle_rpc(mb, 'list_user')
    @database.handle(DBSession)
    def crud_rpc_list_user(db: Session, skip: int = 0, limit: int = 100) -> List[Mapping[str, Any]]:
        db_user_list = crud.list_user(db, skip = skip, limit = limit)
        return [schema.User.from_orm(db_user).dict() for db_user in db_user_list]

    @transport.handle_rpc(mb, 'get_user')
    @database.handle(DBSession)
    def crud_rpc_get_user(db: Session, user_id: int) -> Mapping[str, Any]:
        db_user = crud.get_user(db, user_id = user_id)
        if db_user is None:
            return None
        return schema.User.from_orm(db_user).dict()

    @transport.handle_rpc(mb, 'create_user')
    @database.handle(DBSession)
    def crud_rpc_create_user(db: Session, user_dict: Mapping[str, Any]) -> Mapping[str, Any]:
        db_user = crud.create_user(db, user_data = schema.UserCreate.parse_obj(user_dict))
        if db_user is None:
            return None
        return schema.User.from_orm(db_user).dict()

    @transport.handle_rpc(mb, 'update_user')
    @database.handle(DBSession)
    def crud_rpc_update_user(db: Session, user_id: int, user_dict: Mapping[str, Any]) -> Mapping[str, Any]:
        db_user = crud.update_user(db, user_id = user_id, user_data = schema.UserUpdate.parse_obj(user_dict))
        if db_user is None:
            return None
        return schema.User.from_orm(db_user).dict()

    @transport.handle_rpc(mb, 'delete_user')
    @database.handle(DBSession)
    def crud_rpc_delete_user(db: Session, user_id: int) -> Mapping[str, Any]:
        db_user = crud.delete_user(db, user_id = user_id)
        if db_user is None:
            return None
        return schema.User.from_orm(db_user).dict()
