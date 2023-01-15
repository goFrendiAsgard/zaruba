from typing import Any, Optional, Mapping
from transport import AppMessageBus, AppRPC
from core.security.service.auth_service import AuthService
from schema.role import RoleData, Role
from schema.user import User
from module.auth.role.role_service import RoleService

import logging


def register_role_rpc(
    mb: AppMessageBus,
    rpc: AppRPC,
    auth_service: AuthService,
    role_service: RoleService
):

    @rpc.handle('find_roles')
    def find_roles(
        keyword: str,
        limit: int,
        offset: int,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Mapping[str, Any]:
        current_user = _get_user_from_dict(current_user_data)
        role_result = role_service.find(keyword, limit, offset, current_user)
        return role_result.dict()

    @rpc.handle('find_role_by_id')
    def find_role_by_id(
        id: str,
        current_user_data: Optional[Mapping[str, Any]] = None
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user_from_dict(current_user_data)
        role = role_service.find_by_id(id, current_user)
        return _role_as_dict(role)

    @rpc.handle('find_role_by_name')
    def find_role_by_name(
        name: str,
        current_user_data: Optional[Mapping[str, Any]] = None
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user_from_dict(current_user_data)
        role = role_service.find_by_name(name, current_user)
        return _role_as_dict(role)

    @rpc.handle('insert_role')
    def insert_role(
        role_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        role = RoleData.parse_obj(role_data)
        role.created_by = current_user.id
        role.updated_by = current_user.id
        new_role = role_service.insert(role, current_user)
        return _role_as_dict(new_role)

    @rpc.handle('update_role')
    def update_role(
        id: str,
        role_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        role = RoleData.parse_obj(role_data)
        role.updated_by = current_user.id
        updated_role = role_service.update(id, role, current_user)
        return _role_as_dict(updated_role)

    @rpc.handle('delete_role')
    def delete_role(
        id: str,
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        deleted_role = role_service.delete(id, current_user)
        return _role_as_dict(deleted_role)

    def _get_user_from_dict(
        user_data: Optional[Mapping[str, Any]]
    ) -> Optional[User]:
        if user_data is None:
            return None
        return User.parse_obj(user_data)

    def _role_as_dict(
        role: Optional[Role]
    ) -> Optional[Mapping[str, Any]]:
        if role is None:
            return None
        return role.dict()

    logging.info('Register auth.role RPC handler')
