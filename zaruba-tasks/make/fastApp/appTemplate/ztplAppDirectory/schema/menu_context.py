from typing import Optional
from schema.user import User
from schema.menu import Menu
from pydantic import BaseModel

class MenuContext(BaseModel):
    current_user: Optional[User]
    current_menu: Optional[Menu]
    accessible_menu: Optional[Menu]
    guest_user: User