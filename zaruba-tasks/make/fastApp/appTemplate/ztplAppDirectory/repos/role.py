from typing import List, Mapping, Optional
from schemas.role import Role, RoleData

import abc
import uuid
import datetime

class RoleRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[Role]:
        pass

    @abc.abstractmethod
    def find_by_name(self, name: str) -> Optional[Role]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
        pass

    @abc.abstractmethod
    def insert(self, role_data: RoleData) -> Optional[Role]:
        pass

    @abc.abstractmethod
    def update(self, id: str, role_data: RoleData) -> Optional[Role]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[Role]:
        pass


class MemRoleRepo(RoleRepo):

    def __init__(self):
        self._role_map: Mapping[str, Role] = {}

    def set_storage(self, role_map: Mapping[str, Role]):
        self._role_map = role_map

    def find_by_id(self, id: str) -> Optional[Role]:
        if id not in self._role_map:
            return None
        return self._role_map[id]

    def find_by_name(self, name: str) -> Optional[Role]:
        for _, role in self._role_map.items():
            if role.name == name:
                return role
        return None

    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        mem_roles = list(self._role_map.values())
        roles: List[Role] = []
        for index in range(offset, limit+offset):
            if index >= len(mem_roles):
                break
            mem_role = mem_roles[index]
            roles.append(mem_role)
        return roles

    def count(self, keyword: str) -> List[Role]:
        mem_roles = list(self._role_map.values())
        return len(mem_roles)

    def insert(self, role_data: RoleData) -> Optional[Role]:
        new_role_id=str(uuid.uuid4())
        new_role = Role(
            id=new_role_id,
            name=role_data.name,
            permissions=role_data.permissions,
            created_at=datetime.datetime.utcnow(),
            created_by=role_data.created_by
        )
        self._role_map[new_role_id] = new_role
        return new_role

    def update(self, id: str, role_data: RoleData) -> Optional[Role]:
        if id not in self._role_map:
            return None
        mem_role = self._role_map[id]
        mem_role.name = role_data.name
        mem_role.permissions = role_data.permissions
        mem_role.updated_at = datetime.datetime.utcnow()
        mem_role.updated_by=role_data.updated_by
        self._role_map[id] = mem_role
        return mem_role

    def delete(self, id: str) -> Optional[Role]:
        if id not in self._role_map:
            return None
        mem_role = self._role_map.pop(id)
        return mem_role