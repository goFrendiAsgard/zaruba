from typing import Optional
from schemas.user import User
from schemas.menu import Menu
from pydantic import BaseModel

class MenuContext(BaseModel):
    current_user: Optional[User]
    current_menu: Optional[Menu]
    accessible_menu: Optional[Menu]