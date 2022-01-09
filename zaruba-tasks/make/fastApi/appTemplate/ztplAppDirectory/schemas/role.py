from typing import List
from pydantic import BaseModel
import datetime

class RoleData(BaseModel):
    name: str
    permissions: List[str]

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


class Role(RoleData):
    id: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    class Config:
        orm_mode = True
