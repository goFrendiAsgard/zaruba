from ui.menuService import MenuService, DefaultMenuService
from auth.authService import AuthService

def create_menu_service(auth_service: AuthService) -> MenuService:
    menu_service = DefaultMenuService(auth_service)
    menu_service.add_menu(name='auth', title='Auth', url='/auth')
    menu_service.add_menu(name='auth/roles', title='Roles', url='/auth/roles', parent_name='auth')
    menu_service.add_menu(name='auth/users', title='Users', url='/auth/users', parent_name='auth')
    return menu_service
