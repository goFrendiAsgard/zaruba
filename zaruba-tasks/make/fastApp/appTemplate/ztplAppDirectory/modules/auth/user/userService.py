from typing import Optional
from helpers.transport import RPC, MessageBus
from schemas.user import User, UserData, UserResult
from modules.auth.user.repos.userRepo import UserRepo
from modules.auth.role.roleService import RoleService
from fastapi import HTTPException

import abc
import datetime

class UserService(abc.ABC):

    @abc.abstractmethod
    def get_guest(self) -> User:
        '''
        Guest user is only used to fill up `created_by` or `updated_by` anonymously.
        Guest user is not stored in the repository.
        You cannot and should not create user token for guest user.
        You should not use guest user for authentication/authorization.
        '''
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> UserResult:
        pass

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[User]:
        pass

    @abc.abstractmethod
    def find_by_username(self, username: str) -> Optional[User]:
        pass

    @abc.abstractmethod
    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[User]:
        pass

    @abc.abstractmethod
    def insert(self, user_data: UserData) -> Optional[User]:
        pass

    @abc.abstractmethod
    def update(self, id: str, user_data: UserData) -> Optional[User]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[User]:
        pass

    @abc.abstractclassmethod
    def is_authorized(self, user: User, permission: str) -> bool:
        pass

class DefaultUserService(UserService):

    def __init__(self, mb: MessageBus, rpc: RPC, user_repo: UserRepo, role_service: RoleService, guest_username: str, root_permission: str='root'):
        self.mb = mb
        self.rpc = rpc
        self.user_repo = user_repo
        self.role_service = role_service
        self.guest_username = guest_username
        self.earliest_date = datetime.datetime.min
        self.root_permission = root_permission

    def get_guest(self) -> User:
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

    def find_by_id(self, id: str) -> Optional[User]:
        return self.user_repo.find_by_id(id)

    def find_by_username(self, username: str) -> Optional[User]:
        return self.user_repo.find_by_username(username)

    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[User]:
        return self.user_repo.find_by_identity_and_password(identity, password)

    def insert(self, user_data: UserData) -> Optional[User]:
        user_data = self._validate_data(user_data)
        return self.user_repo.insert(user_data)

    def update(self, id: str, user_data: UserData) -> Optional[User]:
        user_data = self._validate_data(user_data, id)
        return self.user_repo.update(id, user_data)

    def delete(self, id: str) -> Optional[User]:
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
            if role is None:
                continue
            if role.has_permission(permission):
                return True 
        return False

    def _validate_data(self, user_data: UserData, id: Optional[str] = None) -> UserData:
        if user_data.username is not None:
            user = self.user_repo.find_by_username(user_data.username)
            if user is not None and (id is None or user.id != id):
                raise HTTPException(
                    status_code=422, 
                    detail='Username already registered: {}'.format(user_data.username)
                )
        if user_data.email is not None:
            user = self.user_repo.find_by_email(user_data.email)
            if user is not None and (id is None or user.id != id):
                raise HTTPException(
                    status_code=422,
                    detail='Email already registered: {}'.format(user_data.email)
                )
        if user_data.phone_number is not None:
            user = self.user_repo.find_by_phone_number(user_data.phone_number)
            if user is not None and (id is None or user.id != id):
                raise HTTPException(
                    status_code=422,
                    detail='Phone number already registered: {}'.format(user_data.phone_number)
                )
        return user_data
