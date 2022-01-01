from typing import Optional, List
from pydantic import BaseModel
import datetime, re

class UserData(BaseModel):
    username: str
    email: str
    permissions: List[str]
    active: bool
    password: Optional[str]
    full_name: str

    def has_permission(self, permission: str) -> bool:
        for existing_permission in self.permissions:
            if re.search('^{}$'.format(existing_permission), permission):
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


class User(UserData):
    id: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    class Config:
        orm_mode = True
