from typing import Optional
from schemas.role import Role, RoleData, RoleResult
from modules.auth.role.repos.roleRepo import RoleRepo

class RoleService():

    def __init__(self, role_repo: RoleRepo):
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
        return self.role_repo.insert(role_data)

    def update(self, id: str, role_data: RoleData) -> Optional[Role]:
        return self.role_repo.update(id, role_data)

    def delete(self, id: str) -> Optional[Role]:
        return self.role_repo.delete(id)