from typing import Any, List, Mapping
from helpers.transport import RPC
from schemas.user import UserData
from auth.userModel import UserModel

def register_user_rpc(rpc: RPC, user_model: UserModel):

    @rpc.handle('find_user')
    def find_user(keyword: str, limit: int, offset: int) -> List[Mapping[str, Any]]:
        results = user_model.find(keyword, limit, offset)
        return [result.dict() for result in results]

    @rpc.handle('find_user_by_id')
    def find_user_by_id(id: str) -> Mapping[str, Any]:
        result = user_model.find_by_id(id)
        return None if result is None else result.dict()

    @rpc.handle('insert_user')
    def insert_user(data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = user_model.insert(UserData.parse_obj(data))
        return None if result is None else result.dict()

    @rpc.handle('update_user')
    def update_user(id: str, data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = user_model.update(id, UserData.parse_obj(data))
        return None if result is None else result.dict()

    @rpc.handle('delete_user')
    def delete_user(id: str) -> Mapping[str, Any]:
        result = user_model.delete(id)
        return None if result is None else result.dict()

    print('Handle RPC for auth.User')