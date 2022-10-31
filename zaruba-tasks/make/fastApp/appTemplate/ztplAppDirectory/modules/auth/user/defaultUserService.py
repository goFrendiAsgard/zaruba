from typing import Optional
from helpers.transport import RPC, MessageBus
from schemas.user import User, UserData, UserResult
from schemas.activity import ActivityData
from modules.auth.user.repos.userRepo import UserRepo
from modules.auth.role.roleService import RoleService
from modules.auth.user.userService import UserService
from fastapi import HTTPException


class DefaultUserService(UserService):

    def __init__(self, mb: MessageBus, rpc: RPC, user_repo: UserRepo, role_service: RoleService, root_permission: str='root'):
        self.mb = mb
        self.rpc = rpc
        self.user_repo = user_repo
        self.role_service = role_service
        self.root_permission = root_permission


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
        new_user = self.user_repo.insert(user_data)
        self.mb.publish('new_activity', ActivityData(
            user_id = current_user.id,
            activity = 'insert',
            object = 'user',
            row = new_user.dict(),
            row_id = new_user.id
        ).dict())
        return new_user


    def update(self, id: str, user_data: UserData, current_user: User) -> Optional[User]:
        self._find_by_id_or_error(id)
        user_data = self._validate_data(user_data, id)
        user_data.updated_by = current_user.id
        updated_user = self.user_repo.update(id, user_data)
        self.mb.publish('new_activity', ActivityData(
            user_id = current_user.id,
            activity = 'update',
            object = 'user',
            row = updated_user.dict(),
            row_id = updated_user.id
        ).dict())
        return updated_user


    def delete(self, id: str, current_user: User) -> Optional[User]:
        self._find_by_id_or_error(id)
        deleted_user = self.user_repo.delete(id)
        self.mb.publish('new_activity', ActivityData(
            user_id = current_user.id,
            activity = 'delete',
            object = 'user',
            row = deleted_user.dict(),
            row_id = deleted_user.id
        ).dict())
        return deleted_user


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
