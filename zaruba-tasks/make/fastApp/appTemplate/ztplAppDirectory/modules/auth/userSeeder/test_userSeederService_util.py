from typing import Tuple
from fastapi.security import OAuth2PasswordBearer
from core.security.service.authService import AuthService
from core.security.rule.defaultAuthRule import DefaultAuthRule
from core.security.middleware.defaultUserFetcher import DefaultUserFetcher
from modules.auth.user.defaultUserService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.userSeeder.userSeederService import UserSeederService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from modules.auth.user.test_defaultUserService import create_user_data
from helpers.transport import LocalRPC, LocalMessageBus, MessageBus

from sqlalchemy import create_engine


ROOT_USER_DATA = create_user_data()
ROOT_USER_DATA.username = 'root'
ROOT_USER_DATA.email = 'root@innistrad.com'
ROOT_USER_DATA.phone_number = '+6213456781'
ROOT_USER_DATA.password = 'root'
ROOT_USER_DATA.permissions = ['root']


def create_mb() -> MessageBus:
    mb = LocalMessageBus()
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_user_seeder_service_components() -> Tuple[UserSeederService, RoleService, DefaultUserService, DBRoleRepo, DBUserRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'root')
    auth_rule = DefaultAuthRule(rpc)
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
    auth_service = AuthService(auth_rule, user_fetcher, 'root')
    user_seeder_service = UserSeederService(auth_service, user_service)
    return user_seeder_service, role_service, user_service, role_repo, user_repo, mb, rpc

