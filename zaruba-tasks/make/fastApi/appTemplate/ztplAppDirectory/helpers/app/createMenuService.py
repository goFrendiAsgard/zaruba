from ui.menuService import MenuService, DefaultMenuService
from auth.authService import AuthService
from auth.userService import UserService

def create_menu_service(auth_service: AuthService, user_service: UserService) -> MenuService:
    menu_service = DefaultMenuService(auth_service, user_service)
    menu_service.add_menu(name='login', title='Log in', url='/login')
    menu_service.add_menu(name='auth', title='Auth', url='/auth', permission_name='ui:auth')
    menu_service.add_menu(name='auth/roles', title='Roles', url='/auth/roles', permission_name='ui:auth:role', parent_name='auth')
    menu_service.add_menu(name='auth/users', title='Users', url='/auth/users', permission_name='ui:auth:user', parent_name='auth')
    return menu_service
