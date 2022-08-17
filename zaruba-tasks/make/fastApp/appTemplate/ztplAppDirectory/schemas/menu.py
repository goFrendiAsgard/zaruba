from typing import List, Optional
from pydantic import BaseModel

class MenuNode(BaseModel):
    name: str
    title: str
    url: str
    auth_type: int
    permission_name: Optional[str]


class Menu(MenuNode):
    submenus: List[MenuNode] = []
    is_highlighted = False

    def add_submenu(self, submenu: MenuNode):
        self.submenus.append(submenu)