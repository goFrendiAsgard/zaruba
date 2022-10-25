from modules.auth import AuthService
from helpers.transport import RPC
from schemas.authType import AuthType
from modules.ui import MenuService

def create_menu_service(rpc: RPC, auth_service: AuthService) -> MenuService:
    menu_service = MenuService(rpc, auth_service)
    menu_service.add_menu(name='account', title='Account', url='#', auth_type=AuthType.ANYONE)
    menu_service.add_menu(name='account:login', title='Log in', url='/account/login', auth_type=AuthType.NON_USER, parent_name='account')
    menu_service.add_menu(name='account:logout', title='Log out', url='/account/logout', auth_type=AuthType.USER, parent_name='account')
    menu_service.add_menu(name='auth', title='Security', url='#', auth_type=AuthType.ANYONE)
    menu_service.add_menu(name='auth:roles', title='Roles', url='/auth/roles', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:auth:role', parent_name='auth')
    menu_service.add_menu(name='auth:users', title='Users', url='/auth/users', auth_type=AuthType.HAS_PERMISSION, permission_name='ui:auth:user', parent_name='auth')
    return menu_service # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
