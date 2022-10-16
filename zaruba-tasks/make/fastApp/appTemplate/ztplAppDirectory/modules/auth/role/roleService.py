from typing import Optional
from helpers.transport import RPC, MessageBus
from schemas.role import Role, RoleData, RoleResult
from modules.auth.role.repos.roleRepo import RoleRepo
from fastapi import HTTPException

class RoleService():

    def __init__(self, mb: MessageBus, rpc: RPC, role_repo: RoleRepo):
        self.mb = mb
        self.rpc = rpc
        self.role_repo = role_repo

    def find(self, keyword: str, limit: int, offset: int) -> RoleResult:
        count = self.role_repo.count(keyword)
        rows = self.role_repo.find(keyword, limit, offset)
        return RoleResult(count=count, rows=rows)

    def find_by_id(self, id: str) -> Optional[Role]:
        return self.role_repo.find_by_id(id)

    def find_by_name(self, name: str) -> Optional[Role]:
        return self.role_repo.find_by_name(name)

    def insert(self, role_data: RoleData) -> Optional[Role]:
        role_data = self._validate_data(role_data)
        return self.role_repo.insert(role_data)

    def update(self, id: str, role_data: RoleData) -> Optional[Role]:
        role_data = self._validate_data(role_data, id)
        return self.role_repo.update(id, role_data)

    def delete(self, id: str) -> Optional[Role]:
        return self.role_repo.delete(id)

    def _validate_data(self, role_data: RoleData, id: Optional[str] = None) -> RoleData:
        if role_data.name is not None:
            role = self.role_repo.find_by_name(role_data.name)
            if role is not None and (id is None or role.id != id):
                raise HTTPException(
                    status_code=422, 
                    detail='Role already exist: {}'.format(role_data.name)
                )
        return role_data
