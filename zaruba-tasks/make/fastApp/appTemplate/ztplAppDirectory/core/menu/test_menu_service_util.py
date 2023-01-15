from typing import Optional, Tuple, List, Mapping, Any
from core.menu.menu_service import MenuService
from core.security.service.auth_service import AuthService
from module.auth.user.test_default_user_service_util import (
    UNAUTHORIZED_INACTIVE_USER, UNAUTHORIZED_ACTIVE_USER,
    AUTHORIZED_ACTIVE_USER, AUTHORIZED_INACTIVE_USER
)
from helper.transport.local_rpc import LocalRPC
from schema.user import User
from schema.menu import Menu
from fastapi.security import OAuth2PasswordBearer
from schema.auth_type import AuthType
from core.security.middleware.default_user_fetcher import DefaultUserFetcher
from core.security.rule.default_auth_rule import DefaultAuthRule
from transport import AppRPC


def create_rpc() -> AppRPC:
    rpc = AppRPC(LocalRPC())

    # handle get_user_by_cred_token
    @rpc.handle('get_user_by_cred_token')
    def get_user_by_token(token: str) -> Optional[User]:
        token_map: Mapping[str, Optional[User]] = {
            'unauthorized_active': UNAUTHORIZED_ACTIVE_USER,
            'unauthorized_inactive': UNAUTHORIZED_INACTIVE_USER,
            'authorized_active': AUTHORIZED_ACTIVE_USER,
            'authorized_inactive': AUTHORIZED_INACTIVE_USER,
        }
        if token in token_map:
            return token_map[token]
        if token == 'error':
            raise Exception('Emulating rpc error')
        return None

    # handle is_user_authorized
    @rpc.handle('is_user_authorized')
    def is_user_authorized(user_data: Any, permission: str) -> bool:
        user = User.parse_obj(user_data)
        return user.id in [AUTHORIZED_ACTIVE_USER.id, AUTHORIZED_INACTIVE_USER]

    # return rpc
    return rpc


def init_test_menu_service_components() -> Tuple[MenuService, AuthService]:
    rpc = create_rpc()
    auth_rule = DefaultAuthRule(rpc)
    oauth2_scheme = OAuth2PasswordBearer(
        tokenUrl='/', auto_error=False
    )
    user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
    auth_service = AuthService(auth_rule, user_fetcher, 'root')
    menu_service = MenuService(auth_service)
    return menu_service, auth_service


class SingleMenuTestCase():
    def __init__(
        self, name: str,
        is_highlighted: bool,
        submenus_count: int = 0
    ):
        self.name = name
        self.is_highlighted = is_highlighted
        self.submenus_count = submenus_count

    def assert_menu(self, menu=Menu):
        assert menu.name == self.name
        assert menu.is_highlighted == self.is_highlighted
        assert len(menu.submenus) == self.submenus_count


class MenuTestCase(SingleMenuTestCase):
    def __init__(
        self, name: str,
        is_highlighted: bool,
        submenus_count: int = 0,
        children: List[SingleMenuTestCase] = []
    ):
        super().__init__(name, is_highlighted, submenus_count)
        self.children = children

    def assert_menu(self, menu=Menu):
        super().assert_menu(menu)
        for child_index, child in enumerate(self.children):
            child.assert_menu(menu.submenus[child_index])


async def check_has_access(
    menu_service: MenuService,
    user: Optional[User],
    accessibility_test_cases: Mapping[str, bool] = {}
):
    for menu_name, expectation in accessibility_test_cases.items():
        if expectation:
            authorize = menu_service.has_access(menu_name)
            menu_context = await authorize(current_user=user)
            assert menu_context.current_user == user
            continue
        is_error = False
        try:
            authorize = menu_service.has_access(menu_name)
            menu_context = await authorize(current_user=user)
        except Exception:
            is_error = True
        assert is_error


def init_test_menu_data(menu_service: MenuService):
    '''
    This will generate all possible combination of two level menu:
        - everyone (AuthType.ANYONE)
            - everyone-everyone (AuthType.ANYONE)
            - everyone-unauthenticated (AuthType.VISITOR)
            - everyone-authenticated (AuthType.USER)
            - everyone-authorized (AuthType.HAS_PERMISSION)
        - unauthenticated (AuthType.VISITOR)
            - unauthenticated-everyone (AuthType.ANYONE)
            - unauthenticated-unauthenticated (AuthType.VISITOR)
            - unauthenticated-authenticated (AuthType.USER)
            - unauthenticated-authorized (AuthType.HAS_PERMISSION)
        ...
        - authorized (AuthType.HAS_PERMISSION)
            ...
            - authorized-authorized (AuthType.HAS_PERMISSION)
    '''
    auth_type_map = {
        'everyone': AuthType.ANYONE,
        'unauthenticated': AuthType.VISITOR,
        'authenticated': AuthType.USER,
        'authorized': AuthType.HAS_PERMISSION,
    }
    for parent_key in auth_type_map:
        parent_menu_name = parent_key
        parent_menu_title = parent_menu_name
        parent_menu_url = '/{}'.format(parent_key)
        parent_menu_auth_type = auth_type_map[parent_key]
        menu_service.add_menu(
            parent_menu_name,
            title=parent_menu_title,
            url=parent_menu_url,
            auth_type=parent_menu_auth_type
        )
        for child_key in auth_type_map:
            child_menu_name = '{}-{}'.format(parent_key, child_key)
            child_menu_title = child_menu_name
            child_menu_url = '/{}/{}'.format(parent_key, child_key)
            child_menu_auth_type = auth_type_map[child_key]
            menu_service.add_menu(
                child_menu_name,
                title=child_menu_title,
                url=child_menu_url,
                auth_type=child_menu_auth_type,
                parent_name=parent_menu_name
            )
