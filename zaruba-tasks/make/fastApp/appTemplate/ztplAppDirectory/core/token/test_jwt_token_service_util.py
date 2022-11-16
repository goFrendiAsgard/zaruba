from typing import Tuple
from module.auth.user.default_user_service import DefaultUserService
from module.auth.role.role_service import RoleService
from module.auth.user.repo.db_user_repo import DBUserRepo
from module.auth.role.repo.db_role_repo import DBRoleRepo
from core.token.jwt_token_service import JWTTokenService
from module.auth.user.test_default_user_service_util import create_user_data
from helper.transport import LocalRPC, LocalMessageBus
from transport import AppMessageBus, AppRPC

from sqlalchemy import create_engine


ROOT_USER_DATA = create_user_data()
ROOT_USER_DATA.username = 'root'
ROOT_USER_DATA.email = 'root@innistrad.com'
ROOT_USER_DATA.phone_number = '+6213456781'
ROOT_USER_DATA.password = 'root'
ROOT_USER_DATA.permissions = ['root']


def init_test_jwt_token_service_components() -> Tuple[JWTTokenService, RoleService, DefaultUserService, DBRoleRepo, DBUserRepo, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = AppMessageBus(LocalMessageBus())
    rpc = AppRPC(LocalRPC())
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    return token_service, role_service, user_service, role_repo, user_repo, mb, rpc

