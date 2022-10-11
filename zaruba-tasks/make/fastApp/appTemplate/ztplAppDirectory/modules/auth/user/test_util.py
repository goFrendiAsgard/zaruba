from typing import Optional, List, Tuple
from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from schemas.user import User, UserData
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine


def create_user_data():
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_user_data = UserData(
        username='',
        email='',
        phone_number='',
        permissions=[],
        role_ids=[],
        active=True,
        full_name='',
        created_by=''
    )
    return dummy_user_data


def create_user():
    user_data_dict = create_user_data().dict()
    dummy_user = User(id='', **user_data_dict)
    return dummy_user


def init_test_default_user_service_components() -> Tuple[DefaultUserService, RoleService, DBUserRepo, DBRoleRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    return user_service, role_service, user_repo, role_repo, mb, rpc


def init_user_data(user_repo: DBUserRepo, index: Optional[int] = None, permissions: List[str] = [], role_ids: List[str] = [], password: str = '', active: bool = True) -> User:
    user_data = create_user_data()
    user_data.username = 'user' if index is None else 'user-{index}'.format(index=index)
    user_data.email = '{username}@innistrad.com'.format(username=user_data.username)
    user_data.password = password
    user_data.phone_number = '+628123456789' if index is None else '+6281234567890{index}'.format(index=index)
    user_data.permissions = permissions
    user_data.role_ids = role_ids
    user_data.active = active
    user_data.full_name = '{username} Nguyen'.format(username=user_data.username)
    user_data.created_by = 'original_user'
    user_data.updated_by = 'original_user'
    return user_repo.insert(user_data)

