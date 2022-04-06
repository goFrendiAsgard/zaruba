from typing import Any, List, Mapping
from schemas.role import Role, RoleData
from repos.role import RoleRepo

class RoleService():

    def __init__(self, role_repo: RoleRepo):
        self.role_repo = role_repo

    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        return self.role_repo.find(keyword, limit, offset)

    def find_by_id(self, id: str) -> Role:
        return self.role_repo.find_by_id(id)

    def insert(self, role_data: RoleData) -> Role:
        return self.role_repo.insert(role_data)

    def update(self, id: str, role_data: RoleData) -> Role:
        return self.role_repo.update(id, role_data)

    def delete(self, id: str) -> Role:
        return self.role_repo.delete(id)