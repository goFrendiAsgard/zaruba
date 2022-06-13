from typing import Callable, List, Optional, Mapping
from schemas.menu import Menu
from schemas.menuContext import MenuContext
from schemas.user import User
from starlette.requests import Request
from auth.roleService import RoleService
from auth.authService import AuthService
from ui.templateException import TemplateException

import abc
import copy


class MenuService(abc.ABC):

    @abc.abstractmethod
    def add_menu(self, menu: Menu):
        pass

    @abc.abstractmethod
    def get_accessible_menu(self, menu_name: str, user: Optional[User]) -> Optional[Menu]:
        pass

    @abc.abstractmethod
    def is_menu_accessible(self, menu_name: str, user: Optional[User]) -> bool:
        pass

    @abc.abstractmethod
    def validate(self, menu_name, user_fetcher: Callable[[Request], User]) -> Callable[[Request], MenuContext]:
        pass


class DefaultMenuService(MenuService):
    root_menu: Menu
    auth_service: AuthService
    role_service: RoleService
    menu_map = Mapping[str, Menu]
    parent_map = Mapping[str, List[Menu]]

    def __init__(self, auth_service: AuthService, role_service: RoleService, root_menu_name: str = 'root', root_menu_title: str = '', root_menu_url: str = '/', root_menu_role_names: List[str] = []):
        self.auth_service = auth_service
        self.role_service = role_service
        root_role_ids = self.role_service.get_ids_by_names(root_menu_role_names)
        self.root_menu = Menu(name=root_menu_name, title=root_menu_title, url=root_menu_url, role_ids=root_role_ids)
        self.menu_map = {root_menu_name: self.root_menu}
        self.parent_map = {root_menu_name: []}

    def add_menu(self, name: str, title: str, url: str, role_names: List[str]=[], parent_name:Optional[str] = None):
        role_ids = self.role_service.get_ids_by_names(role_names)
        menu = Menu(name=name, title=title, url=url, role_ids=role_ids)
        parent_menu = self.root_menu if parent_name is None else self.menu_map[parent_name]
        if parent_menu is None:
            raise Exception('Menu {} not found'.format(parent_name))
        parent_menu.add_submenu(menu)
        self.menu_map[name] = menu
        self.parent_map[name] = [*self.parent_map[parent_menu.name], parent_menu.name]

    def get_accessible_menu(self, menu_name: str, user: Optional[User] = None) -> Optional[Menu]:
        if not self._is_menu_accessible(self.root_menu, user):
            return None
        parent_names = self.parent_map[menu_name]
        accessible_menu = self._get_accessible_menu(self.root_menu, user)
        accessible_menu = self._get_accessible_menu(self.root_menu, user)
        if menu_name in self.parent_map:
            parent_names = self.parent_map[menu_name]
            highlighted_menu = self._highlight_menu_by_names(accessible_menu, [*parent_names, menu_name])
            return highlighted_menu
        return accessible_menu

    def is_menu_accessible(self, menu_name: str, user: Optional[User]) -> bool:
        if menu_name not in self.menu_map:
            return False
        menu = self.menu_map[menu_name]
        return self._is_menu_accessible(menu, user)

    def validate(self, current_menu_name: str, user_fetcher: Callable[[Request], User]) -> Callable[[Request], MenuContext]:
        async def verify_menu_accessibility(request: Request) -> MenuContext:
            current_user = user_fetcher(request)
            current_menu = copy.deep_copy(self.menu_map[current_menu_name]) if current_menu_name in self.menu_map else None
            accessible_menu = self.get_accessible_menu(current_menu_name, current_user)
            menu_context = MenuContext(current_user=current_user, current_menu=current_menu, accessible_menu=accessible_menu)
            if not self._is_menu_accessible(current_menu_name, current_user):
                raise TemplateException(statuss_code=403, detail='Forbidden', accessible_menu = accessible_menu)
            return MenuContext(current_user=current_user, accessible_menu=accessible_menu)
        return verify_menu_accessibility

    def _highlight_menu_by_names(self, menu: Menu, names: List[str]) -> Menu:
        if menu.name in names:
            menu.is_highlighted = True
        for submenu in menu.submenus:
            self._highlight_menu_by_names(submenu, names)
        return menu

    def _get_accessible_menu(self, original_menu: Menu, user: Optional[User] = None) -> Menu:
        menu: Menu = copy.deepcopy(original_menu)
        new_submenus: List[Menu] = []
        for submenu in menu.submenus:
            if not self._is_menu_accessible(submenu, user):
                continue
            new_submenu = self._get_accessible_menu(submenu, user)
            new_submenus.append(new_submenu)
        menu.submenus = new_submenus
        return menu

    def _is_menu_accessible(self, menu: Menu, user: Optional[User]) -> bool:
        if len(menu.role_ids) == 0:
            return True
        if user is None:
            return False
        for menu_role_id in menu.role_ids:
            if menu_role_id in user.role_ids:
                return True
        return False