from typing import Tuple
from fastapi.security import OAuth2PasswordBearer
from core.security.service.auth_service import AuthService
from core.security.rule.default_auth_rule import DefaultAuthRule
from core.security.middleware.default_user_fetcher import DefaultUserFetcher
from module.auth.user.default_user_service import DefaultUserService
from module.auth.role.role_service import RoleService
from module.auth.user_seeder.user_seeder_service import UserSeederService
from module.auth.user.repo.db_user_repo import DBUserRepo
from module.auth.role.repo.db_role_repo import DBRoleRepo
from module.auth.user.test_default_user_service import create_user_data
from helper.transport import LocalRPC, LocalMessageBus, MessageBus
from transport import AppMessageBus, AppRPC

from sqlalchemy import create_engine


ROOT_USER_DATA = create_user_data()
ROOT_USER_DATA.username = 'root'
ROOT_USER_DATA.email = 'root@innistrad.com'
ROOT_USER_DATA.phone_number = '+6213456781'
ROOT_USER_DATA.password = 'root'
ROOT_USER_DATA.permissions = ['root']


def create_mb() -> AppMessageBus:
    mb = AppMessageBus(LocalMessageBus())
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_user_seeder_service_components() -> Tuple[UserSeederService, RoleService, DefaultUserService, DBRoleRepo, DBUserRepo, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'root')
    auth_rule = DefaultAuthRule(rpc)
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
    auth_service = AuthService(auth_rule, user_fetcher, 'root')
    user_seeder_service = UserSeederService(auth_service, user_service)
    return user_seeder_service, role_service, user_service, role_repo, user_repo, mb, rpc

