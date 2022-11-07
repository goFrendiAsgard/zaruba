from modules.auth.role.repos.roleRepo import RoleRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo, DBRoleEntity
from modules.auth.role.roleService import RoleService
from modules.auth.role.roleRoute import register_role_api_route, register_role_ui_route
from modules.auth.role.roleRpc import register_role_rpc