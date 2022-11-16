from module.auth.role.repo.role_repo import RoleRepo
from module.auth.role.repo.db_role_repo import DBRoleRepo, DBRoleEntity
from module.auth.role.role_service import RoleService
from module.auth.role.role_route import register_role_api_route, register_role_ui_route
from module.auth.role.role_rpc import register_role_rpc