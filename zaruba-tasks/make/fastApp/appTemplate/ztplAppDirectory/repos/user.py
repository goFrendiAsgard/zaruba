from typing import List, Mapping, Optional
from schemas.user import User, UserData

import abc
import json
import uuid
import datetime

class UserRepo(abc.ABC):

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
    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
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


class MemUserRepo(UserRepo):

    def __init__(self):
        self._user_map: Mapping[str, User] = {}
        self._password_map: Mapping[str, str] = {}

    def set_storage(self, user_map: Mapping[str, User]):
        self._user_map = user_map
        self._password_map = {}
        for id, password in self._user_map.items():
            self._password_map[id] = password

    def find_by_username(self, username: str) -> Optional[User]:
        mem_users = list(self._user_map.values())
        for mem_user in mem_users:
            if mem_user.username == username:
                return mem_user

    def find_by_id(self, id: str) -> Optional[User]:
        if id not in self._user_map:
            return None
        return self._user_map[id]

    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[User]:
        for id, mem_user in self._user_map.items():
            if mem_user.username != identity and mem_user.email != identity and mem_user.phone_number != identity:
                continue
            if self._password_map[id] != password:
                return None
            return mem_user

    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        mem_users = list(self._user_map.values())
        users: List[User] = []
        for index in range(offset, limit+offset):
            if index >= len(mem_users):
                break
            mem_user = mem_users[index]
            users.append(mem_user)
        return users

    def count(self, keyword: str) -> List[User]:
        mem_users = list(self._user_map.values())
        return len(mem_users)

    def insert(self, user_data: UserData) -> Optional[User]:
        new_user_id=str(uuid.uuid4())
        new_user = User(
            id=new_user_id,
            username=user_data.username,
            email=user_data.email,
            phone_number=user_data.phone_number,
            permissions=user_data.permissions,
            active=user_data.active,
            password=user_data.password,
            full_name=user_data.full_name,
            created_at=datetime.datetime.now(),
            created_by=user_data.created_by
        )
        self._password_map[new_user_id] = user_data.password
        self._user_map[new_user_id] = new_user
        return new_user

    def update(self, id: str, user_data: UserData) -> Optional[User]:
        if id not in self._user_map:
            return None
        mem_user = self._user_map[id]
        mem_user.username = user_data.username
        mem_user.email = user_data.email
        mem_user.phone_number = user_data.phone_number
        mem_user.permissions = user_data.permissions
        mem_user.active = user_data.active
        mem_user.full_name = user_data.full_name
        mem_user.updated_at = datetime.datetime.now()
        mem_user.updated_by = user_data.updated_by
        if user_data.password:
            mem_user.password = user_data.password
            self._password_map[id] = user_data.password
        self._user_map[id] = mem_user
        return mem_user

    def delete(self, id: str) -> Optional[User]:
        if id not in self._user_map:
            return None
        self._password_map.pop(id)
        mem_user = self._user_map.pop(id)
        return mem_user