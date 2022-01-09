from typing import Any, List, Mapping
from helpers.transport import RPC
from schemas.user import UserData
from auth.userModel import UserModel
from auth.tokenModel import TokenModel

def register_user_rpc(rpc: RPC, user_model: UserModel, token_model: TokenModel):

    @rpc.handle('get_user_token')
    def get_user_token(identity: str, password: str) -> str:
        authenticated_user = user_model.find_by_password(identity, password)
        if not authenticated_user:
            raise Exception('Incorrect identity or password')
        return token_model.create_user_token(authenticated_user)

    @rpc.handle('find_user')
    def find_user(keyword: str, limit: int, offset: int, current_user_data: Mapping[str, Any]) -> List[Mapping[str, Any]]:
        results = user_model.find(keyword, limit, offset)
        return [result.dict() for result in results]

    @rpc.handle('find_user_by_id')
    def find_user_by_id(id: str, current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = user_model.find_by_id(id)
        return None if result is None else result.dict()

    @rpc.handle('insert_user')
    def insert_user(data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = user_model.insert(UserData.parse_obj(data))
        return None if result is None else result.dict()

    @rpc.handle('update_user')
    def update_user(id: str, data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = user_model.update(id, UserData.parse_obj(data))
        return None if result is None else result.dict()

    @rpc.handle('delete_user')
    def delete_user(id: str, current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = user_model.delete(id)
        return None if result is None else result.dict()

    print('Handle RPC for auth.User')