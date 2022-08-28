from typing import List, Optional
from pydantic import BaseModel
import datetime, re

class UserDataWithoutPassword(BaseModel):
    username: str = ''
    email: str = ''
    phone_number: str = ''
    permissions: List[str] = []
    role_ids: List[str] = []
    active: bool = False
    full_name: str = ''
    created_at: Optional[datetime.datetime]
    created_by: Optional[str]
    updated_at: Optional[datetime.datetime]
    updated_by: Optional[str]

    def has_permission(self, permission: str) -> bool:
        for existing_permission in self.permissions:
            permission_pattern = re.sub(r'\*', '[0-9a-zA-Z\*]+', permission)
            if re.search('^{}$'.format(permission_pattern), existing_permission):
                return True
        return False

    def add_permission(self, permission: str):
        for existing_permission in self.permissions:
            if permission == existing_permission:
                return
        self.permissions.append(permission)

    def remove_permission(self, permission: str):
        new_permissions = [existing_permission for existing_permission in self.permissions if existing_permission != permission]
        self.permissions = new_permissions

    def add_role_id(self, role_id: str):
        for existing_role_id in self.role_ids:
            if role_id == existing_role_id:
                return
        self.role_ids.append(role_id)

    def remove_role_id(self, role_id: str):
        new_role_ids = [existing_role_id for existing_role_id in self.role_ids if existing_role_id != role_id]
        self.role_ids = new_role_ids


class UserWithoutPassword(UserDataWithoutPassword):
    id: str
    class Config:
        orm_mode = True


class UserData(UserDataWithoutPassword):
    password: Optional[str] = ''


class User(UserData):
    id: str
    class Config:
        orm_mode = True


class UserResult(BaseModel):
    count: int
    rows: List[UserWithoutPassword]