from typing import Tuple
from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from modules.auth.token.tokenService import JWTTokenService
from modules.auth.user.test_util import create_user_data
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine


ROOT_USER_DATA = create_user_data()
ROOT_USER_DATA.username = 'root'
ROOT_USER_DATA.email = 'root@innistrad.com'
ROOT_USER_DATA.phone_number = '+6213456781'
ROOT_USER_DATA.password = 'root'
ROOT_USER_DATA.permissions = ['root']


def init_test_jwt_token_service_components() -> Tuple[JWTTokenService, RoleService, DefaultUserService, DBRoleRepo, DBUserRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    return token_service, role_service, user_service, role_repo, user_repo, mb, rpc

