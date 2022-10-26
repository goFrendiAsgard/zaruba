from typing import Any, Optional, Mapping
from helpers.transport import RPC, MessageBus
from schemas.user import UserData
from schemas.user import User
from modules.auth.user.userService import UserService

import sys

def register_user_rpc(mb: MessageBus, rpc: RPC, user_service: UserService):

    @rpc.handle('get_guest_user')
    def get_guest_user(current_user_data: Optional[Mapping[str, Any]] = None) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        return user_service.get_guest_user(current_user).dict()


    @rpc.handle('get_system_user')
    def get_system_user(current_user_data: Optional[Mapping[str, Any]] = None) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        return user_service.get_system_user(current_user).dict()


    @rpc.handle('find_users')
    def find_users(keyword: str, limit: int, offset: int, current_user_data: Optional[Mapping[str, Any]]) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        user_result = user_service.find(keyword, limit, offset, current_user)
        return user_result.dict()


    @rpc.handle('find_user_by_id')
    def find_user_by_id(id: str, current_user_data: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        user = user_service.find_by_id(id, current_user)
        return None if user is None else user.dict()


    @rpc.handle('find_user_by_username')
    def find_user_by_username(id: str, username: str, current_user_data: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        user = user_service.find_by_username(username, current_user)
        return None if user is None else user.dict()


    @rpc.handle('find_user_by_identity_and_password')
    def find_user_by_identity_and_password(id: str, identity: str, password: str, current_user_data: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        user = user_service.find_by_identity_and_password(identity, password, current_user)
        return None if user is None else user.dict()


    @rpc.handle('insert_user')
    def insert_user(data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        user = UserData.parse_obj(data)
        new_user = user_service.insert(user, current_user)
        return None if new_user is None else new_user.dict()


    @rpc.handle('update_user')
    def update_user(id: str, data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        user = UserData.parse_obj(data)
        updated_user = user_service.update(id, user, current_user)
        return None if updated_user is None else updated_user.dict()


    @rpc.handle('delete_user')
    def delete_user(id: str, current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        user = user_service.delete(id, current_user)
        return None if user is None else user.dict()


    @rpc.handle('is_user_authorized')
    def is_authorized(user_data: Mapping[str, Any], permission: str) -> bool:
        user = User.parse_obj(user_data)
        return user_service.is_authorized(user, permission)

    print('Handle RPC for auth.User', file=sys.stderr)