from typing import List
from schemas.role import Role, RoleData

import abc

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