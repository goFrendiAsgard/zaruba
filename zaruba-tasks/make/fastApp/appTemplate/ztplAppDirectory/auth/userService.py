from typing import List
from schemas.user import User, UserData, UserResult
from repos.user import UserRepo
from auth.roleService import RoleService

import abc
import datetime

class UserService(abc.ABC):

    @abc.abstractmethod
    def get_guest_user(self) -> User:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> UserResult:
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

    @abc.abstractclassmethod
    def is_authorized(self, user: User, permission: str) -> bool:
        pass

class DefaultUserService(UserService):

    def __init__(self, user_repo: UserRepo, role_service: RoleService, guest_username: str, root_permission: str='root'):
        self.user_repo = user_repo
        self.role_service = role_service
        self.guest_username = guest_username
        self.earliest_date = datetime.datetime.min
        self.root_permission = root_permission

    def get_guest_user(self) -> User:
        return User(
            id = 'guest',
            username = self.guest_username, 
            active = True,
            updated_at = self.earliest_date,
            created_at = self.earliest_date,
        )

    def find(self, keyword: str, limit: int, offset: int) -> UserResult:
        count = self.user_repo.count(keyword)
        rows = self.user_repo.find(keyword, limit, offset)
        return UserResult(count=count, rows=rows)

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
    
    def is_authorized(self, user: User, permission: str) -> bool:
        # user has root permission
        if user.has_permission(self.root_permission):
            return True
        # user has any required permission
        if user.has_permission(permission):
            return True
        # user has any role that has any required permission
        role_ids = user.role_ids
        for role_id in role_ids:
            role = self.role_service.find_by_id(role_id)
            if role.has_permission(permission):
                return True 
        return False
