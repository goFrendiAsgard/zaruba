from typing import Tuple
from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.userSeeder.userSeederService import UserSeederService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from modules.auth.user.test_defaultUserService import create_user_data
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine


################################################
# -- ⚙️ Helpers
################################################

ROOT_USER_DATA = create_user_data()
ROOT_USER_DATA.username = 'root'
ROOT_USER_DATA.email = 'root@innistrad.com'
ROOT_USER_DATA.phone_number = '+6213456781'
ROOT_USER_DATA.password = 'root'
ROOT_USER_DATA.permissions = ['root']


def init_test_user_seeder_service_components() -> Tuple[UserSeederService, RoleService, DefaultUserService, DBRoleRepo, DBUserRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    user_seeder_service = UserSeederService(user_service)
    return user_seeder_service, role_service, user_service, role_repo, user_repo, mb, rpc


################################################
# -- 🧪 Test
################################################

def test_user_seeder_service_with_existing_user():
    user_seeder_service, _, user_service, _, user_repo, _, _ = init_test_user_seeder_service_components()
    # Init existing user
    user_repo.insert(ROOT_USER_DATA)
    # This should yield no error
    user_seeder_service.seed(ROOT_USER_DATA)
    assert user_service.find(keyword='', limit=10, offset=0).count == 1


def test_user_seeder_service_with_non_existing_user():
    user_seeder_service, _, user_service, _, _, _, _ = init_test_user_seeder_service_components()
    # This should yield no error
    user_seeder_service.seed(ROOT_USER_DATA)
    assert user_service.find(keyword='', limit=10, offset=0).count == 1