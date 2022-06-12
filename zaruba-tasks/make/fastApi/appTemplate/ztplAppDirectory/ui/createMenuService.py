from ui.menuService import MenuService, DefaultMenuService
from auth.roleService import RoleService

def create_menu_service(role_service: RoleService) -> MenuService:
    menu_service = DefaultMenuService(role_service)
    menu_service.add_menu(name='auth', title='Auth', url='/auth')
    menu_service.add_menu(name='auth/roles', title='Roles', url='/auth/roles', parent_name='auth')
    menu_service.add_menu(name='auth/users', title='Users', url='/auth/users', parent_name='auth')
    return menu_service
