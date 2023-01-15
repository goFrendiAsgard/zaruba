from typing import Callable, List, Optional, Mapping
from schema.auth_type import AuthType
from schema.menu import Menu
from schema.menu_context import MenuContext
from schema.user import User
from core.security.service.auth_service import AuthService
from core.page.page_template_exception import PageTemplateException
from fastapi import Depends, status
from starlette.requests import Request

import copy


class MenuService():
    '''
    Service to handle menu.

    You can use MenuService to:
    - Add menu.
    - Validate whether a user is allowed to access a page or not.
    '''

    def __init__(
        self,
        auth_service: AuthService,
        root_menu_name: str = 'root',
        root_menu_title: str = '',
        root_menu_url: str = '/'
    ):
        '''
        Initiate a new MenuService.

        Keyword arguments:
        - auth_service -- Instance of core.security.service.auth_service.
        - root_menu_name -- Root menu name (default: 'root')
        - root_menu_title -- Root menu title (default: '')
        - root_menu_url -- Root menu URL (default: '/')

        You can use MenuService to:
        - Add menu.
        - Validate whether a user is allowed to access a page or not.
        '''
        self.auth_service: AuthService = auth_service
        self.root_menu: Menu = Menu(
            name=root_menu_name,
            title=root_menu_title,
            url=root_menu_url,
            auth_type=AuthType.ANYONE
        )
        self.menu_map: Mapping[str, Menu] = {root_menu_name: self.root_menu}
        self.parent_map: Mapping[str, List[Menu]] = {root_menu_name: []}

    def add_menu(
        self,
        name: str,
        title: str,
        url: str,
        auth_type: int,
        permission_name: Optional[str] = None,
        parent_name: Optional[str] = None
    ):
        '''
        Add menu to MenuService.

        Keyword arguments:
        - name: Menu name.
        - title: Menu title.
        - url: Menu URL.
        - auth_type: Authentication type. Valid values are:
            - schema.auth_type.AuthType.ANYONE
            - schema.auth_type.AuthType.VISITOR
            - schema.auth_type.AuthType.USER
            - schema.auth_type.AuthType.HAS_PERMISSION
        - permission_name: Permission name to access the inew menu.
        - parent_name: Name of menu's parent.
            If not specified, the new menu will be added to root_menu.
        '''
        if auth_type not in (
            AuthType.ANYONE,
            AuthType.VISITOR,
            AuthType.USER,
            AuthType.HAS_PERMISSION
        ):
            raise Exception(' '.join([
                'Cannot adding menu {}'.format(name),
                'because it has invalid auth_type {}'.format(auth_type)
            ]))
        menu = Menu(
            name=name,
            title=title,
            url=url,
            auth_type=auth_type,
            permission_name=permission_name
        )
        parent_menu = self._get_parent_menu_by_parent_name(parent_name)
        if parent_menu is None:
            raise Exception(' '.join([
                'Cannot adding menu {}'.format(name),
                'because the parent menu {} is not found'.format(parent_name)
            ]))
        parent_menu.add_submenu(menu)
        self.menu_map[name] = menu
        self.parent_map[name] = [
            *self.parent_map[parent_menu.name], parent_menu.name
        ]

    def _get_parent_menu_by_parent_name(
        self,
        parent_name: str
    ) -> Optional[Menu]:
        if parent_name is None:
            return self.root_menu
        return self.menu_map.get(parent_name, None)

    def get_accessible_menu(
        self,
        menu_name: str,
        user: Optional[User]
    ) -> Optional[Menu]:
        '''
        Get accessible menu for a user (can be None).

        If menu_name is accessible,
        then menu_name and it's parents will have is_hightlighted == True

        Keyword Arguments:
        - menu_name -- Menu name to be acessed
        - user -- User accessing the menu (optional, default: None)
        '''
        parent_names = self.parent_map[menu_name]
        accessible_menu = self._get_accessible_menu(self.root_menu, user)
        if menu_name in self.parent_map:
            parent_names = self.parent_map[menu_name]
            highlighted_menu = self._highlight_menu_by_names(
                accessible_menu, [*parent_names, menu_name]
            )
            return highlighted_menu
        return accessible_menu

    def has_access(
        self,
        menu_name: str
    ) -> Callable[[Callable[[Request], Optional[User]]], MenuContext]:
        '''
        Return MenuContext if user is allowed to access menu_name.
        Otherwise, it will raise a PageException error.

        Keyword arguments:
        - menu_name -- Menu name to be accessed
        '''
        if menu_name not in self.menu_map:
            raise Exception('Menu {} is not registered'.format(menu_name))
        menu = self.menu_map[menu_name]
        user_fetcher = self.auth_service.get_user_fetcher(throw_error=False)

        async def check(
            current_user: Optional[User] = Depends(user_fetcher)
        ) -> MenuContext:
            menu_context = self._get_menu_context(menu_name, current_user)
            is_authorized = self.auth_service.check_user_access(
                current_user, menu.auth_type, menu.permission_name
            )
            if not is_authorized:
                raise PageTemplateException(
                    status_code=status.HTTP_403_FORBIDDEN,
                    detail='forbidden',
                    menu_context=menu_context
                )
            return menu_context
        return check

    def _get_menu_context(
        self, current_menu_name: str,
        current_user: Optional[User]
    ) -> MenuContext:
        current_menu: Optional[Mapping[str, Menu]] = None
        if current_menu_name in self.menu_map:
            current_menu = copy.deepcopy(
                self.menu_map[current_menu_name]
            )
        accessible_menu = self.get_accessible_menu(
            current_menu_name, current_user)
        return MenuContext(
            current_menu=current_menu,
            current_user=current_user,
            accessible_menu=accessible_menu,
            guest_user=self.auth_service.get_guest_user()
        )

    def _highlight_menu_by_names(self, menu: Menu, names: List[str]) -> Menu:
        if menu.name in names:
            menu.is_highlighted = True
        for submenu in menu.submenus:
            self._highlight_menu_by_names(submenu, names)
        return menu

    def _get_accessible_menu(
        self, original_menu: Menu,
        user: Optional[User] = None
    ) -> Menu:
        menu: Menu = copy.deepcopy(original_menu)
        new_submenus: List[Menu] = []
        for submenu in menu.submenus:
            is_authorized = self.auth_service.check_user_access(
                user, submenu.auth_type, submenu.permission_name
            )
            if not is_authorized:
                continue
            new_submenu = self._get_accessible_menu(submenu, user)
            new_submenus.append(new_submenu)
        menu.submenus = new_submenus
        return menu
