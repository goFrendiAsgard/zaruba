from typing import List
from schemas.user import User, UserData
from repos.user import UserRepo

import abc
import datetime

class UserModel(abc.ABC):

    @abc.abstractmethod
    def get_guest_user(self) -> User:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        pass

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
    def insert(self, user_data: UserData) -> User:
        pass

    @abc.abstractmethod
    def update(self, id: str, user_data: UserData) -> User:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> User:
        pass

class DefaultUserModel(UserModel):

    def __init__(self, user_repo: UserRepo, guest_username: str):
        self.user_repo = user_repo
        self.guest_username = guest_username
        self.earliest_date = datetime.datetime.min

    def get_guest_user(self) -> User:
        return User(
            id = 'guest',
            username = self.guest_username, 
            active = True,
            updated_at = self.earliest_date,
            created_at = self.earliest_date,
        )

    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        return self.user_repo.find(keyword, limit, offset)

    def find_by_id(self, id: str) -> User:
        return self.user_repo.find_by_id(id)

    def find_by_username(self, username: str) -> User:
        return self.user_repo.find_by_username(username)

    def find_by_password(self, identity: str, password: str) -> User:
        return self.user_repo.find_by_password(identity, password)

    def insert(self, user_data: UserData) -> User:
        return self.user_repo.insert(user_data)

    def update(self, id: str, user_data: UserData) -> User:
        return self.user_repo.update(id, user_data)

    def delete(self, id: str) -> User:
        return self.user_repo.delete(id)