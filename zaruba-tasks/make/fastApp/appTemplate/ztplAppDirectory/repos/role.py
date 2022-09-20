from typing import List, Optional
from schemas.role import Role, RoleData

import abc

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
