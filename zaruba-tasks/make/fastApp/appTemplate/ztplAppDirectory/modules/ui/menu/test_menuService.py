from typing import Any, Callable, Optional, Tuple, List, Mapping
from modules.ui.menu.menuService import MenuService
from modules.auth.auth.authService import AuthService
from helpers.transport.localRpc import LocalRPC
from schemas.user import User
from schemas.menu import Menu
from starlette.requests import Request
from fastapi import HTTPException
from schemas.authType import AuthType

import pytest

################################################
# -- ⚙️ Helpers
################################################
def create_user():
    dummy_user = User(
        username='',
        email='',
        password='',
        phone_number='',
        permissions=[],
        role_ids=[],
        active=True,
        full_name='',
        created_by='',
        id=''
    )
    return dummy_user

guest_user = create_user()
guest_user.id = 'mock_guest_user_id'
guest_user.username = 'guest_username'
guest_user.created_by = 'mock_user_id'

unauthorized_user = create_user()
unauthorized_user.id = 'mock_unauthorized_user_id'
unauthorized_user.username = 'unauthorized_username'
unauthorized_user.created_by = 'mock_user_id'

authorized_user = create_user()
authorized_user.id = 'mock_authorized_user_id'
authorized_user.username = 'authorized_username'
authorized_user.created_by = 'mock_user_id'
authorized_user.permissions = ['root']


class MockAuthService(AuthService):

    def __init__(self, user: Optional[User]):
        self.user = user
    
    def _return_none_or_throw_error(throw_error: bool):
        if not throw_error:
            return False
        raise HTTPException(status_code=402, detail='Unauthroized')

    def everyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def verify_everyone(request: Optional[Request]) -> Optional[User]:
            if self.user is None:
                return guest_user
            return self.user
        return verify_everyone

    def is_authenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def verify_authenticated(request: Optional[Request]) -> Optional[User]:
            if self.user == unauthorized_user or self.user == authorized_user:
                return self.user
            return self._return_none_or_throw_error(throw_error)
        return verify_authenticated

    def is_unauthenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def verify_unauthenticated(request: Optional[Request]) -> Optional[User]:
            if self.user is None or self.user == guest_user:
                return self.user
            return self._return_none_or_throw_error(throw_error)
        return verify_unauthenticated

    def is_authorized(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def verify_authorized(request: Optional[Request]) -> Optional[User]:
            if self.user == authorized_user:
                return self.user
            return self._return_none_or_throw_error(throw_error)
        return verify_authorized


class MockRPC(LocalRPC):

    def __init__(self):
        super().__init__()

    def call(self, rpc_name: str, *args: Any) -> Any:
        if rpc_name == 'is_user_authorized':
            user_data = args[0]
            user = User.parse_obj(user_data)
            return user.id == authorized_user.id
        return super().call(rpc_name, *args)


class SingleMenuTestCase():
    def __init__(self, name: str, is_highlighted: bool, submenus_count: int = 0):
        self.name = name
        self.is_highlighted = is_highlighted
        self.submenus_count = submenus_count

    def assert_menu(self, menu = Menu):
        assert menu.name == self.name
        assert menu.is_highlighted == self.is_highlighted
        assert len(menu.submenus) == self.submenus_count 


class MenuTestCase(SingleMenuTestCase):
    def __init__(self, name: str, is_highlighted: bool, submenus_count: int = 0, children: List[SingleMenuTestCase] = []):
        super().__init__(name, is_highlighted, submenus_count)
        self.children = children
    
    def assert_menu(self, menu = Menu):
        super().assert_menu(menu)
        for child_index, child in enumerate(self.children):
            child.assert_menu(menu.submenus[child_index])


def init_test_menu_service_components(user: Optional[User]) -> Tuple[MenuService, MockAuthService, MockRPC]:
    rpc = MockRPC()
    auth_service = MockAuthService(user)
    menu_service = MenuService(rpc, auth_service)
    return menu_service, auth_service, rpc


def init_test_menu_data(menu_service: MenuService):
    '''
    This will generate all possible combination of two level menu:
        - everyone (AuthType.EVERYONE)
            - everyone-everyone (AuthType.EVERYONE)
            - everyone-unauthenticated (AuthType.UNAUTHENTICATED)
            - everyone-authenticated (AuthType.AUTHENTICATED)
            - everyone-authorized (AuthType.AUTHORIZED)
        - unauthenticated (AuthType.UNAUTHENTICATED)
            - unauthenticated-everyone (AuthType.EVERYONE)
            - unauthenticated-unauthenticated (AuthType.UNAUTHENTICATED)
            - unauthenticated-authenticated (AuthType.AUTHENTICATED)
            - unauthenticated-authorized (AuthType.AUTHORIZED)
        ...
        - authorized (AuthType.AUTHORIZED)
            ...
            - authorized-authorized (AuthType.AUTHORIZED)
    '''
    auth_type_map = {
        'everyone': AuthType.EVERYONE,
        'unauthenticated': AuthType.UNAUTHENTICATED,
        'authenticated': AuthType.AUTHENTICATED,
        'authorized': AuthType.AUTHORIZED,
    }
    for parent_key in auth_type_map:
        parent_menu_name = parent_key
        parent_menu_title = parent_menu_name
        parent_menu_url = '/{}'.format(parent_key)
        parent_menu_auth_type = auth_type_map[parent_key]
        menu_service.add_menu(parent_menu_name, title=parent_menu_title, url=parent_menu_url, auth_type=parent_menu_auth_type)
        for child_key in auth_type_map:
            child_menu_name = '{}-{}'.format(parent_key, child_key)
            child_menu_title = child_menu_name
            child_menu_url = '/{}/{}'.format(parent_key, child_key)
            child_menu_auth_type = auth_type_map[child_key]
            menu_service.add_menu(child_menu_name, title=child_menu_title, url=child_menu_url, auth_type=child_menu_auth_type, parent_name=parent_menu_name)


async def check_is_authorized(menu_service: MenuService, user: Optional[User], accessibility_test_cases: Mapping[str, bool] = {}):
    for menu_name, expectation in accessibility_test_cases.items():
        if expectation:
            authorize = menu_service.is_authorized(menu_name)
            menu_context = await authorize(current_user = user)
            assert menu_context.current_user == user
            continue
        is_error = False
        try:
            authorize = menu_service.is_authorized(menu_name)
            menu_context = await authorize(current_user = user)
        except:
            is_error = True
        assert is_error

################################################
# -- 🧪 Test
################################################

def test_menu_service_add_menu_with_invalid_auth_type():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    is_error = False
    try:
        menu_service.add_menu(name='test', title='test', url='/test', auth_type= 500)
    except:
        is_error = True
    assert is_error


def test_menu_service_add_menu_with_non_existing_parent():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    is_error = False
    try:
        menu_service.add_menu(name='test', title='test', url='/test', auth_type= AuthType.EVERYONE, parent_name='invalid')
    except:
        is_error = True
    assert is_error


def test_menu_service_get_accessible_non_existing_menu():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    # test get accessible menu for not existing menu
    is_error = False
    try:
        menu_service.get_accessible_menu('invalid', user)
    except:
        is_error = True
    assert is_error


@pytest.mark.asyncio
async def test_menu_service_no_user_get_accessible_everyone_menu():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('everyone-everyone', user)
    assert len(root.submenus) == 2
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=True, submenus_count=2, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=True),
            MenuTestCase(name='everyone-unauthenticated', is_highlighted=False),
        ]),
        MenuTestCase(name='unauthenticated', is_highlighted=False, submenus_count=2, children = [
            MenuTestCase(name='unauthenticated-everyone', is_highlighted=False),
            MenuTestCase(name='unauthenticated-unauthenticated', is_highlighted=False),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_no_user_get_accessible_unauthenticated_menu():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('unauthenticated-unauthenticated', user)
    assert len(root.submenus) == 2
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=False, submenus_count=2, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=False),
            MenuTestCase(name='everyone-unauthenticated', is_highlighted=False),
        ]),
        MenuTestCase(name='unauthenticated', is_highlighted=True, submenus_count=2, children = [
            MenuTestCase(name='unauthenticated-everyone', is_highlighted=False),
            MenuTestCase(name='unauthenticated-unauthenticated', is_highlighted=True),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_authenticated_user_get_accessible_authenticated_menu():
    user = unauthorized_user
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('authenticated-authenticated', user)
    assert len(root.submenus) == 2
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=False, submenus_count=2, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=False),
            MenuTestCase(name='everyone-authenticated', is_highlighted=False),
        ]),
        MenuTestCase(name='authenticated', is_highlighted=True, submenus_count=2, children = [
            MenuTestCase(name='authenticated-everyone', is_highlighted=False),
            MenuTestCase(name='authenticated-authenticated', is_highlighted=True),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_authorized_user_get_accessible_authorized_menu():
    user = authorized_user
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('authorized-authorized', user)
    assert len(root.submenus) == 3
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=False, submenus_count=3, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=False),
            MenuTestCase(name='everyone-authenticated', is_highlighted=False),
            MenuTestCase(name='everyone-authorized', is_highlighted=False),
        ]),
        MenuTestCase(name='authenticated', is_highlighted=False, submenus_count=3, children = [
            MenuTestCase(name='authenticated-everyone', is_highlighted=False),
            MenuTestCase(name='authenticated-authenticated', is_highlighted=False),
            MenuTestCase(name='authenticated-authorized', is_highlighted=False),
        ]),
        MenuTestCase(name='authorized', is_highlighted=True, submenus_count=3, children = [
            MenuTestCase(name='authorized-everyone', is_highlighted=False),
            MenuTestCase(name='authorized-authenticated', is_highlighted=False),
            MenuTestCase(name='authorized-authorized', is_highlighted=True),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_no_user_authorize():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    await check_is_authorized(menu_service, user, {
        'everyone' : True,
        'unauthenticated' : True,
        'authenticated' : False,
        'authorized': False,
        'invalid': False
    })


@pytest.mark.asyncio
async def test_menu_service_authenticated_user_authorize():
    user = unauthorized_user
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    await check_is_authorized(menu_service, user, {
        'everyone' : True,
        'unauthenticated' : False,
        'authenticated' : True,
        'authorized': False,
        'invalid': False
    })


@pytest.mark.asyncio
async def test_menu_service_authorized_user_authorize():
    user = authorized_user
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)
    await check_is_authorized(menu_service, user, {
        'everyone' : True,
        'unauthenticated' : False,
        'authenticated' : True,
        'authorized': True,
        'invalid': False
    })
