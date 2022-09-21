from typing import Any, Callable, List, Optional, Mapping
from schemas.authType import AuthType
from schemas.menu import Menu
from schemas.menuContext import MenuContext
from schemas.user import User
from fastapi import Depends, status
from starlette.requests import Request
from modules.auth.auth.authService import AuthService
from helpers.transport import RPC
from modules.ui.page.pageTemplateException import PageTemplateException

import copy

class MenuService():

    def __init__(self, rpc: RPC, auth_service: AuthService, root_menu_name: str = 'root', root_menu_title: str = '', root_menu_url: str = '/', permission_name: Optional[str] = None):
        self.auth_service: AuthService = auth_service
        self.rpc: RPC = rpc
        self.root_menu: Menu = Menu(name=root_menu_name, title=root_menu_title, url=root_menu_url, auth_type=AuthType.EVERYONE, permission_name=permission_name)
        self.menu_map: Mapping[str, Menu] = {root_menu_name: self.root_menu}
        self.parent_map: Mapping[str, List[Menu]] = {root_menu_name: []}

    def add_menu(self, name: str, title: str, url: str, auth_type: int, permission_name: Optional[str] = None, parent_name:Optional[str] = None):
        menu = Menu(name=name, title=title, url=url, auth_type=auth_type, permission_name=permission_name)
        parent_menu = self.root_menu if parent_name is None else self.menu_map[parent_name]
        if parent_menu is None:
            raise Exception('Menu {} not found'.format(parent_name))
        parent_menu.add_submenu(menu)
        self.menu_map[name] = menu
        self.parent_map[name] = [*self.parent_map[parent_menu.name], parent_menu.name]

    def get_accessible_menu(self, menu_name: str, user: User) -> Optional[Menu]:
        if not self._is_menu_accessible(self.root_menu, user):
            return None
        parent_names = self.parent_map[menu_name]
        accessible_menu = self._get_accessible_menu(self.root_menu, user)
        if menu_name in self.parent_map:
            parent_names = self.parent_map[menu_name]
            highlighted_menu = self._highlight_menu_by_names(accessible_menu, [*parent_names, menu_name])
            return highlighted_menu
        return accessible_menu

    def authenticate(self, menu_name: str) -> Callable[[Callable[[Request], Optional[User]]], MenuContext]:
        if menu_name not in self.menu_map:
            raise Exception('Menu {} is not registered'.format(menu_name))
        menu = self.menu_map[menu_name]
        authenticator = self._get_menu_authenticator(menu)
        async def verify(current_user: Optional[User] = Depends(authenticator)) -> MenuContext:
            return self._get_menu_context(menu_name, current_user)
        return verify

    def _get_menu_authenticator(self, menu: Menu) -> Callable[[Request], Optional[User]]:
        if menu.auth_type == AuthType.EVERYONE:
            return self.auth_service.everyone(throw_error=False)
        if menu.auth_type == AuthType.UNAUTHENTICATED:
            return self.auth_service.is_unauthenticated(throw_error=False)
        if menu.auth_type == AuthType.AUTHENTICATED:
            return self.auth_service.is_authenticated(throw_error=False)
        if menu.auth_type == AuthType.AUTHORIZED:
            return self.auth_service.is_authorized(menu.permission_name, throw_error=False)
        raise Exception('Menu {} has invalid auth_type: {}'.format(menu.name, menu.auth_type))

    def _get_menu_context(self, current_menu_name: str, current_user: Optional[User]) -> MenuContext:
        current_menu = copy.deepcopy(self.menu_map[current_menu_name]) if current_menu_name in self.menu_map else None
        accessible_menu = self.get_accessible_menu(current_menu_name, current_user)
        menu_context = MenuContext(
            current_menu = current_menu,
            current_user = current_user,
            accessible_menu = accessible_menu
        )
        if not self._is_menu_accessible(current_menu, current_user):
            raise PageTemplateException(status_code=status.HTTP_403_FORBIDDEN, detail='Forbidden', menu_context = menu_context)
        return menu_context

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

    def _is_menu_accessible(self, menu: Optional[Menu], user: Optional[User]) -> bool:
        if menu is None:
            return False
        if menu.auth_type == AuthType.EVERYONE:
            return True
        if menu.auth_type == AuthType.UNAUTHENTICATED and user is None:
            return True
        if menu.auth_type == AuthType.AUTHENTICATED and user is not None:
            return True
        if menu.auth_type == AuthType.AUTHORIZED and user is not None:
            return self.rpc.call('is_user_authorized', user.dict(), menu.permission_name)
        return False