from typing import List, Optional
from schemas.user import User, UserWithoutPassword, UserData

import abc

class UserRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[User]:
        pass

    @abc.abstractmethod
    def find_by_username(self, username: str) -> Optional[User]:
        pass

    @abc.abstractmethod
    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
        pass

    @abc.abstractmethod
    def insert(self, user_data: UserData) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def update(self, id: str, user_data: UserData) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[UserWithoutPassword]:
        pass