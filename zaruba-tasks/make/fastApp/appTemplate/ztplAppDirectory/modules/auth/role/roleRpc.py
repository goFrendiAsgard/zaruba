from typing import Any, Optional, Mapping
from helpers.transport import RPC, MessageBus
from core.security.service.authService import AuthService
from schemas.role import RoleData
from schemas.user import User
from modules.auth.role.roleService import RoleService

import sys

def register_role_rpc(mb: MessageBus, rpc: RPC, auth_service: AuthService, role_service: RoleService):

    @rpc.handle('find_roles')
    def find_roles(keyword: str, limit: int, offset: int, current_user_data: Optional[Mapping[str, Any]] = None) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        role_result = role_service.find(keyword, limit, offset, current_user)
        return role_result.dict()


    @rpc.handle('find_role_by_id')
    def find_role_by_id(id: str, current_user_data: Optional[Mapping[str, Any]] = None) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        role = role_service.find_by_id(id, current_user)
        return None if role is None else role.dict()


    @rpc.handle('find_role_by_name')
    def find_role_by_name(name: str, current_user_data: Optional[Mapping[str, Any]] = None) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        role = role_service.find_by_name(name, current_user)
        return None if role is None else role.dict()


    @rpc.handle('insert_role')
    def insert_role(role_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        role = RoleData.parse_obj(role_data)
        role.created_by = current_user.id
        role.updated_by = current_user.id
        new_role = role_service.insert(role, current_user)
        return None if new_role is None else new_role.dict()


    @rpc.handle('update_role')
    def update_role(id: str, role_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        role = RoleData.parse_obj(role_data)
        role.updated_by = current_user.id
        updated_role = role_service.update(id, role, current_user)
        return None if updated_role is None else updated_role.dict()


    @rpc.handle('delete_role')
    def delete_role(id: str, current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        role = role_service.delete(id, current_user)
        return None if role is None else role.dict()

    print('Handle RPC for auth.Role', file=sys.stderr)