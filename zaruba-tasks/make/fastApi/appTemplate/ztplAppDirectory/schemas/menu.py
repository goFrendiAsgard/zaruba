from typing import List, Mapping
from pydantic import BaseModel, Field

class MenuNode(BaseModel):
    name: str
    title: str
    url: str
    role_ids: List[str] = []


class Menu(MenuNode):
    submenus: List[MenuNode] = []
    is_highlighted = False

    def add_submenu(self, submenu: MenuNode):
        self.submenus.append(submenu)