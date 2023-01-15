from typing import Any, Optional, Mapping
from transport import AppMessageBus, AppRPC
from core.security.service.auth_service import AuthService
from schema.user import UserData
from schema.user import User
from module.auth.user.user_service import UserService

import logging


def register_user_rpc(
    mb: AppMessageBus, 
    rpc: AppRPC, 
    auth_service: AuthService, 
    user_service: UserService
):

    @rpc.handle('find_users')
    def find_users(
        keyword: str,
        limit: int,
        offset: int,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Mapping[str, Any]:
        current_user = _get_user_from_dict(current_user_data)
        user_result = user_service.find(keyword, limit, offset, current_user)
        return user_result.dict()

    @rpc.handle('find_user_by_id')
    def find_user_by_id(
        id: str,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user_from_dict(current_user_data)
        user = user_service.find_by_id(id, current_user)
        return user_as_dict(user)

    @rpc.handle('find_user_by_username')
    def find_user_by_username(
        id: str,
        username: str,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user_from_dict(current_user_data)
        user = user_service.find_by_username(username, current_user)
        return user_as_dict(user)

    @rpc.handle('find_user_by_identity_and_password')
    def find_user_by_identity_and_password(
        id: str,
        identity: str,
        password: str,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user_from_dict(current_user_data)
        user = user_service.find_by_identity_and_password(
            identity, password, current_user
        )
        return user_as_dict(user)

    @rpc.handle('insert_user')
    def insert_user(
        user_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        user = UserData.parse_obj(user_data)
        new_user = user_service.insert(user, current_user)
        return user_as_dict(new_user)

    @rpc.handle('update_user')
    def update_user(
        id: str,
        user_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        user_data = UserData.parse_obj(user_data)
        updated_user = user_service.update(id, user_data, current_user)
        return user_as_dict(updated_user)

    @rpc.handle('delete_user')
    def delete_user(
        id: str,
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        deleted_user = user_service.delete(id, current_user)
        return user_as_dict(deleted_user)

    @rpc.handle('is_user_authorized')
    def is_authorized(
        user_data: Mapping[str, Any],
        permission: str
    ) -> bool:
        user = User.parse_obj(user_data)
        return user_service.is_authorized(user, permission)

    def _get_user_from_dict(
        user_data: Optional[Mapping[str, Any]]
    ) -> Optional[User]:
        if user_data is None:
            return None
        return User.parse_obj(user_data)

    def user_as_dict(
        user: Optional[User]
    ) -> Optional[Mapping[str, Any]]:
        if user is None:
            return None
        return user.dict()

    logging.info('Register auth.user RPC handler')
