from typing import List, Optional
from schema.user import UserWithoutPassword, UserData

import abc


class UserRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def find_by_username(self, username: str) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def find_by_email(self, username: str) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def find_by_phone_number(
        self, username: str
    ) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def find_by_identity_and_password(
        self, identity: str, password: str
    ) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def find(
        self, keyword: str, limit: int, offset: int
    ) -> List[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
        pass

    @abc.abstractmethod
    def insert(self, user_data: UserData) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def update(
        self, id: str, user_data: UserData
    ) -> Optional[UserWithoutPassword]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[UserWithoutPassword]:
        pass
