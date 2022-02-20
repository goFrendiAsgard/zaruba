from typing import List, Mapping
from schemas.user import User, UserData

import abc
import json
import uuid
import datetime

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


class MemUserRepo(UserRepo):

    def __init__(self):
        self._user_map: Mapping[str, User] = {}
        self._password_map: Mapping[str, str] = {}

    def set_storage(self, user_map: Mapping[str, User], password_map: Mapping[str, str]):
        self._user_map = user_map
        self.password_map = password_map

    def find_by_username(self, username: str) -> User:
        mem_users = self._user_map.values()
        for mem_user in mem_users:
            if mem_user.username == username:
                return mem_user

    def find_by_id(self, id: str) -> User:
        return self._user_map[id]

    def find_by_password(self, identity: str, password: str) -> User:
        for id, mem_user in self._user_map.items():
            if mem_user.username != identity and mem_user.email != identity and mem_user.phone_number != identity:
                continue
            if self._password_map[id] != password:
                return None
            return mem_user

    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        mem_users = self._user_map.values()
        users: List[User] = []
        for index in range(limit, limit+offset):
            mem_user = mem_users[index]
            users.append(mem_user)
        return users

    def insert(self, user_data: UserData) -> User:
        id = str(uuid.uuid4())
        new_user_id=str(uuid.uuid4()),
        new_user = User(
            id=new_user_id,
            username=user_data.username,
            email=user_data.email,
            phone_number=user_data.phone_number,
            json_permissions=json.dumps(user_data.permissions),
            active=user_data.active,
            full_name=user_data.full_name,
            created_at=datetime.datetime.utcnow()
        )
        self._password_map[id] = user_data.password
        self._user_map[id] = new_user

    def update(self, id: str, user_data: UserData) -> User:
        mem_user = self._user_map[id]
        mem_user.username = user_data.username
        mem_user.email = user_data.email
        mem_user.phone_number = user_data.phone_number
        mem_user.json_permissions = json.dumps(user_data.permissions)
        mem_user.active = user_data.active
        mem_user.full_name = user_data.full_name
        mem_user.updated_at = datetime.datetime.utcnow()
        if user_data.password:
            self._password_map[id] = user_data.password
        self._user_map[id] = mem_user

    def delete(self, id: str) -> User:
        self._password_map.pop(id)
        mem_user = self._user_map.pop(id)
        return mem_user