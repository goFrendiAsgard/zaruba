from typing import List
from schemas.user import User, UserData

import abc

class UserRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> User:
        pass

    @abc.abstractmethod
    def find_by_username(self, username: str) -> User:
        pass

    @abc.abstractmethod
    def find_by_password(self, identity: str, password: str) -> User:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        pass

    @abc.abstractmethod
    def insert(self, user_data: UserData) -> User:
        pass

    @abc.abstractmethod
    def update(self, id: str, user_data: UserData) -> User:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> User:
        pass