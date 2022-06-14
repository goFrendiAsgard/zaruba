from ui.menuService import MenuService, DefaultMenuService
from auth.authService import AuthService

def create_menu_service(auth_service: AuthService) -> MenuService:
    menu_service = DefaultMenuService(auth_service)
    menu_service.add_menu(name='auth', title='Auth', url='/auth', permission_name='ui:auth')
    menu_service.add_menu(name='auth/roles', title='Roles', url='/auth/roles', permission_name='ui:auth:role', parent_name='auth')
    menu_service.add_menu(name='auth/users', title='Users', url='/auth/users', permission_name='ui:auth:user', parent_name='auth')
    return menu_service
