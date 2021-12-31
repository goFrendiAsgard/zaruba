from typing import Optional, List
from pydantic import BaseModel
import datetime

class UserData(BaseModel):
    username: str
    email: str
    roles: str
    active: bool
    password: Optional[str]
    full_name: str

    def _get_roles(self) -> List[str]:
        return self.roles.split(' ')

    def has_role(self, role: str) -> bool:
        for existing_role in self._get_roles():
            if role == existing_role:
                return True
        return False

    def add_role(self, role: str):
        roles = self._get_roles()
        for existing_role in roles:
            if role == existing_role:
                return
        roles.append(role)
        self.roles = ' '.join(roles)

    def remove_role(self, role: str):
        roles = self._get_roles()
        new_roles = [existing_role for existing_role in roles if existing_role != role]
        self.roles = ' '.join(new_roles)


class User(UserData):
    id: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    class Config:
        orm_mode = True
