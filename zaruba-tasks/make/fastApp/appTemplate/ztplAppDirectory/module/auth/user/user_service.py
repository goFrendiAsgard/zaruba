from typing import Optional
from schema.user import User, UserData, UserResult

import abc

class UserService(abc.ABC):

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int, current_user: Optional[User]) -> UserResult:
        pass

    @abc.abstractmethod
    def find_by_id(self, id: str, current_user: Optional[User]) -> Optional[User]:
        pass

    @abc.abstractmethod
    def find_by_username(self, username: str, current_user: Optional[User]) -> Optional[User]:
        pass

    @abc.abstractmethod
    def find_by_identity_and_password(self, identity: str, password: str, current_user: Optional[User]) -> Optional[User]:
        pass

    @abc.abstractmethod
    def insert(self, user_data: UserData, current_user: User) -> Optional[User]:
        pass

    @abc.abstractmethod
    def update(self, id: str, user_data: UserData, current_user: User) -> Optional[User]:
        pass

    @abc.abstractmethod
    def delete(self, id: str, current_user: User) -> Optional[User]:
        pass

    @abc.abstractclassmethod
    def is_authorized(self, user: User, permission: str) -> bool:
        pass
