from typing import Any, List, Mapping
from helpers.transport import RPC
from schemas.role import RoleData
from auth.roleService import RoleService

def register_role_rpc(rpc: RPC, role_service: RoleService):

    @rpc.handle('find_role')
    def find_role(keyword: str, limit: int, offset: int, current_user_data: Mapping[str, Any]) -> List[Mapping[str, Any]]:
        roles = role_service.find(keyword, limit, offset)
        return [role.dict() for role in roles]

    @rpc.handle('find_role_by_id')
    def find_role_by_id(id: str, current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        role = role_service.find_by_id(id)
        return None if role is None else role.dict()

    @rpc.handle('insert_role')
    def insert_role(role_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        role = role_service.insert(RoleData.parse_obj(role_data))
        return None if role is None else role.dict()

    @rpc.handle('update_role')
    def update_role(id: str, role_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        role = role_service.update(id, RoleData.parse_obj(role_data))
        return None if role is None else role.dict()

    @rpc.handle('delete_role')
    def delete_role(id: str, current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        role = role_service.delete(id)
        return None if role is None else role.dict()

    print('Handle RPC for auth.Role')