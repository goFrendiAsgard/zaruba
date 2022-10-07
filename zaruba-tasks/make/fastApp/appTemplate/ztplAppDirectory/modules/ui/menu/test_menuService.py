from typing import Any, Callable, Optional, Tuple
from modules.ui.menu.menuService import MenuService
from modules.auth.auth.authService import AuthService
from helpers.transport.localRpc import LocalRPC
from schemas.user import User
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



def init_test_menu_service_components(user: Optional[User]) -> Tuple[MenuService, MockAuthService, LocalRPC]:
    rpc = LocalRPC()
    auth_service = MockAuthService(user)
    menu_service = MenuService(rpc, auth_service)
    return menu_service, auth_service, rpc


def init_test_menu_data(menu_service: MenuService):
    menu_service.add_menu('everyone', 'Everyone', '/everyone', AuthType.EVERYONE)
    menu_service.add_menu('submenu-everyone', 'Submenu Everyone', '/everyone/everyone', AuthType.EVERYONE, parent_name = 'everyone')
    menu_service.add_menu('submenu-unauthenticated', 'Submenu Unauthenticated', '/everyone/unauthenticated', AuthType.UNAUTHENTICATED, parent_name = 'everyone')
    menu_service.add_menu('submenu-authenticated', 'Submenu Authenticated', '/everyone/authenticated', AuthType.AUTHENTICATED, parent_name = 'everyone')
    menu_service.add_menu('submenu-authorized', 'Submenu Authorized', '/everyone/authorized', AuthType.AUTHORIZED, permission_name = 'root', parent_name = 'everyone')
    menu_service.add_menu('unauthenticated', 'Unauthenticated', '/unauthenticated', AuthType.UNAUTHENTICATED)
    menu_service.add_menu('authenticated', 'Authenticated', '/authenticated', AuthType.AUTHENTICATED)
    menu_service.add_menu('authorized', 'Authorized', '/authorized', AuthType.AUTHORIZED, permission_name = 'root')


################################################
# -- 🧪 Test
################################################

@pytest.mark.asyncio
async def test_menu_service_no_user_get_accessible_menu():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)

    # test get accessible menu for authorized subbmenu
    accessible_menu = menu_service.get_accessible_menu('submenu-everyone', user)
    assert accessible_menu.name == 'root'
    assert len(accessible_menu.submenus) == 2
    # everyone
    assert accessible_menu.submenus[0].name == 'everyone'
    assert accessible_menu.submenus[0].is_highlighted
    assert len(accessible_menu.submenus[0].submenus) == 2
    # authenticated
    assert accessible_menu.submenus[1].name == 'unauthenticated'
    assert not accessible_menu.submenus[1].is_highlighted
    # submenu-everyone
    assert accessible_menu.submenus[0].submenus[0].name == 'submenu-everyone'
    assert accessible_menu.submenus[0].submenus[0].is_highlighted
    # submenu-authenticated
    assert accessible_menu.submenus[0].submenus[1].name == 'submenu-unauthenticated'
    assert not accessible_menu.submenus[0].submenus[1].is_highlighted


@pytest.mark.asyncio
async def test_menu_service_no_user_authorize():
    user = None
    menu_service, _, _ = init_test_menu_service_components(user)
    init_test_menu_data(menu_service)

    # test menu 'everyone'
    authorize = menu_service.is_authorized('everyone')
    menu_context = await authorize(current_user = user)
    assert menu_context.current_user == user

    # test menu 'unauthenticated'
    authorize = menu_service.is_authorized('unauthenticated')
    menu_context = await authorize(current_user = user)
    assert menu_context.current_user == user
 
    # test menu 'authenticated'
    is_error = False
    try:
        authorize = menu_service.is_authorized('authenticated')
        menu_context = await authorize(current_user = user)
    except:
        is_error = True
    assert is_error
     
    # test menu 'authorized'
    is_error = False
    try:
        authorize = menu_service.is_authorized('authorized')
        menu_context = await authorize(current_user = user)
    except:
        is_error = True
    assert is_error
     
    # test menu 'submenu-everyone'
    authorize = menu_service.is_authorized('submenu-everyone')
    menu_context = await authorize(current_user = user)
    assert menu_context.current_user == user

    # test menu 'submenu-unauthenticated'
    authorize = menu_service.is_authorized('submenu-unauthenticated')
    menu_context = await authorize(current_user = user)
    assert menu_context.current_user == user
     
    # test menu 'submenu-authenticated'
    is_error = False
    try:
        authorize = menu_service.is_authorized('submenu-authenticated')
        menu_context = await authorize(current_user = user)
    except:
        is_error = True
    assert is_error
     
    # test menu 'submenu-authorized'
    is_error = False
    try:
        authorize = menu_service.is_authorized('submenu-authorized')
        menu_context = await authorize(current_user = user)
    except:
        is_error = True
    assert is_error
     
     
      
     