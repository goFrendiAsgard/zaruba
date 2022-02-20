from typing import List, Mapping
from schemas.role import Role, RoleData

import abc
import json
import uuid
import datetime

class RoleRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Role:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        pass

    @abc.abstractmethod
    def insert(self, role_data: RoleData) -> Role:
        pass

    @abc.abstractmethod
    def update(self, id: str, role_data: RoleData) -> Role:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Role:
        pass


class MemRoleRepo(RoleRepo):

    def __init__(self):
        self._role_map: Mapping[str, Role] = {}

    def set_storage(self, role_map: Mapping[str, Role]):
        self._role_map = role_map

    def find_by_id(self, id: str) -> Role:
        return self._role_map[id]

    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        storage_roles = self._role_map.values()
        roles: List[Role] = []
        for index in range(limit, limit+offset):
            mem_role = storage_roles[index]
            roles.append(mem_role)
        return roles

    def insert(self, role_data: RoleData) -> Role:
        new_role_id=str(uuid.uuid4())
        new_role = Role(
            id=new_role_id,
            name=role_data.name,
            json_permissions=json.dumps(role_data.permissions),
            created_at=datetime.datetime.utcnow()
        )
        self._role_map[id] = new_role

    def update(self, id: str, role_data: RoleData) -> Role:
        mem_role = self._role_map[id]
        mem_role.name = role_data.name
        mem_role.json_permissions = json.dumps(role_data.permissions)
        mem_role.updated_at = datetime.datetime.utcnow()
        self._role_map[id] = mem_role

    def delete(self, id: str) -> Role:
        mem_role = self._role_map.pop(id)
        return mem_role