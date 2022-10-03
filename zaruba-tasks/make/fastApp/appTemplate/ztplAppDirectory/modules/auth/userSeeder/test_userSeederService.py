from typing import Optional
from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.userSeeder.userSeederService import UserSeederService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from schemas.user import User, UserData
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine


################################################
# -- ⚙️ Helpers
################################################

def create_user_data():
    dummy_user_data = UserData(
        username='',
        email='',
        password='',
        phone_number='',
        permissions=[],
        role_ids=[],
        active=True,
        full_name='',
        created_by=''
    )
    return dummy_user_data


################################################
# -- 🧪 Test
################################################

def test_user_seeder_service_with_existing_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    user_seeder_service = UserSeederService(user_service)
    # Init existing user
    root_user_data = create_user_data()
    root_user_data.username = 'root'
    root_user_data.email = 'root@innistrad.com'
    root_user_data.phone_number = '+6213456781'
    root_user_data.password = 'root'
    root_user_data.permissions = ['root']
    user_repo.insert(root_user_data)
    # This should yield no error
    user_seeder_service.seed(root_user_data)
    assert user_service.find(keyword='', limit=10, offset=0).count == 1


def test_user_seeder_service_with_non_existing_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    user_seeder_service = UserSeederService(user_service)
    # Init existing user
    root_user_data = create_user_data()
    root_user_data.username = 'root'
    root_user_data.email = 'root@innistrad.com'
    root_user_data.phone_number = '+6213456781'
    root_user_data.password = 'root'
    root_user_data.permissions = ['root']
    # This should yield no error
    user_seeder_service.seed(root_user_data)
    assert user_service.find(keyword='', limit=10, offset=0).count == 1