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
    def get_guest_user(self, current_user: Optional[User]) -> User:
        '''
        Guest user is only used to fill up `created_by` or `updated_by` anonymously.
        Guest user is not stored in the repository.
        You cannot and should not create user token for guest user.
        You should not use guest user for authentication/authorization.
        '''
        pass

    @abc.abstractmethod
    def get_system_user(self, current_user: Optional[User]) -> User:
        '''
        System user is only used to fill up `created_by` or `updated_by` anonymously.
        System user is not stored in the repository.
        You cannot and should not create user token for system user.
        You should not use system user for authentication/authorization.
        '''
        pass

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

class DefaultUserService(UserService):

    def __init__(self, mb: MessageBus, rpc: RPC, user_repo: UserRepo, role_service: RoleService, root_permission: str='root'):
        self.mb = mb
        self.rpc = rpc
        self.user_repo = user_repo
        self.role_service = role_service
        self.earliest_date = datetime.datetime(1970, 1, 1, 0, 0, 0, 0, datetime.timezone.utc)
        self.root_permission = root_permission


    def get_guest_user(self, current_user: Optional[User] = None) -> User:
        return User(
            id = 'guest',
            username = 'guest', 
            active = True,
            updated_at = self.earliest_date,
            created_at = self.earliest_date,
        )


    def get_system_user(self, current_user: Optional[User] = None) -> User:
        return User(
            id = 'system',
            username = 'system', 
            active = True,
            permissions=[self.root_permission],
            updated_at = self.earliest_date,
            created_at = self.earliest_date,
        )


    def find(self, keyword: str, limit: int, offset: int, current_user: Optional[User] = None) -> UserResult:
        count = self.user_repo.count(keyword)
        rows = self.user_repo.find(keyword, limit, offset)
        return UserResult(count=count, rows=rows)


    def find_by_id(self, id: str, current_user: Optional[User] = None) -> Optional[User]:
        user = self._find_by_id_or_error(id)
        return user


    def find_by_username(self, username: str, current_user: Optional[User] = None) -> Optional[User]:
        user = self.user_repo.find_by_username(username)
        if user is None:
            raise HTTPException(
                status_code=404, 
                detail='Username not found: {}'.format(username)
            )
        return user


    def find_by_identity_and_password(self, identity: str, password: str, current_user: Optional[User] = None) -> Optional[User]:
        user = self.user_repo.find_by_identity_and_password(identity, password)
        if user is None:
            raise HTTPException(
                status_code=404, 
                detail='Identity or password does not match: {}'.format(identity)
            )
        return user


    def insert(self, user_data: UserData, current_user: User) -> Optional[User]:
        user_data.created_by = current_user.id
        user_data.updated_by = current_user.id
        user_data = self._validate_data(user_data)
        return self.user_repo.insert(user_data)


    def update(self, id: str, user_data: UserData, current_user: User) -> Optional[User]:
        self._find_by_id_or_error(id)
        user_data = self._validate_data(user_data, id)
        user_data.updated_by = current_user.id
        return self.user_repo.update(id, user_data)


    def delete(self, id: str, current_user: User) -> Optional[User]:
        self._find_by_id_or_error(id)
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
            try:
                role = self.role_service.find_by_id(role_id)
                if role.has_permission(permission):
                    return True 
            except:
                continue
        return False


    def _find_by_id_or_error(self, id: Optional[str] = None) -> Optional[User]:
        user = self.user_repo.find_by_id(id)
        if user is None:
            raise HTTPException(
                status_code=404, 
                detail='User id not found: {}'.format(id)
            )
        return user


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
