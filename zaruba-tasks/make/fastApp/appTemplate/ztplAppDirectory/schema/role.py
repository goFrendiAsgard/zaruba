from typing import List, Optional
from pydantic import BaseModel
import datetime, re

class RoleData(BaseModel):
    name: str
    permissions: List[str] = []
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


class Role(RoleData):
    id: str
    class Config:
        orm_mode = True


class RoleResult(BaseModel):
    count: int
    rows: List[Role]