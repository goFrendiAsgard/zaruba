from typing import Optional, List, TypeVar
from pydantic import BaseModel
import datetime, re

class UserData(BaseModel):
    username: str = ''
    email: str = ''
    phone_number: str = ''
    permissions: List[str] = []
    role_ids: List[str] = []
    active: bool = False
    password: Optional[str] = ''
    full_name: str = ''

    def has_permission(self, permission: str) -> bool:
        for existing_permission in self.permissions:
            existing_permission_pattern = re.sub('\*', '[0-9a-zA-Z]+', existing_permission)
            if re.search('^{}$'.format(existing_permission_pattern), permission):
                return True
        return False

    def add_permission(self, permission: str):
        for existing_permission in self.permissions:
            if permission == existing_permission:
                return
        self.permissions.append(permission)

    def remove_permission(self, permission: str):
        new_permissions = [existing_permission for existing_permission in self.permissions if existing_permission != permission]
        self.permissions = ' '.join(new_permissions)

    def add_role_id(self, role_id: str):
        for existing_role_id in self.role_ids:
            if role_id == existing_role_id:
                return
        self.role_ids.append(role_id)

    def remove_role_id(self, role_id: str):
        new_role_ids = [existing_role_id for existing_role_id in self.role_ids if existing_role_id != role_id]
        self.role_ids = ' '.join(new_role_ids)


class User(UserData):
    id: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    class Config:
        orm_mode = True