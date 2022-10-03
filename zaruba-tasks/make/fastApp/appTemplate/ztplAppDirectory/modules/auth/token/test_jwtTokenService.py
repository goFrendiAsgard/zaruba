from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from modules.auth.token.tokenService import JWTTokenService
from schemas.user import UserData, User
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


def create_user():
    dummy_user = User(
        username='',
        email='',
        password='',
        phone_number='',
        permissions=[],
        role_ids=[],
        active=True,
        full_name='',
        created_by='',
        id=''
    )
    return dummy_user


################################################
# -- 🧪 Test
################################################

def test_jwt_token_service_with_empty_token():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    user = token_service.get_user_by_token('')
    # make sure token service return correct value
    assert user is None
    

def test_jwt_token_service_with_null_token():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    user = token_service.get_user_by_token(None)
    # make sure token service return correct value
    assert user is None


def test_jwt_token_service_with_invalid_token():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    user = token_service.get_user_by_token('invalid token')
    # make sure token service return correct value
    assert user is None


def test_jwt_token_service_with_existing_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    # Init existing user
    root_user_data = create_user_data()
    root_user_data.username = 'root'
    root_user_data.email = 'root@innistrad.com'
    root_user_data.phone_number = '+6213456781'
    root_user_data.password = 'root'
    root_user_data.permissions = ['root']
    root_user = user_repo.insert(root_user_data)
    # Creating token
    token = token_service.create_user_token(root_user)
    # make sure token service return correct value
    assert token is not None
    assert token != ''
    # Get user
    user = token_service.get_user_by_token(token)
    # make sure token service return correct value
    assert user == root_user


def test_jwt_token_service_with_non_existing_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    # Init existing user
    root_user_data = create_user_data()
    root_user_data.username = 'root'
    root_user_data.email = 'root@innistrad.com'
    root_user_data.phone_number = '+6213456781'
    root_user_data.password = 'root'
    root_user_data.permissions = ['root']
    root_user = user_repo.insert(root_user_data)
    # Init non existing user
    inexist_user = create_user()
    inexist_user.id='inexist'
    inexist_user.username = 'inexist'
    inexist_user.email = 'inexist@innistrad.com'
    inexist_user.phone_number = '+6213456784'
    inexist_user.password = 'root'
    # Creating token
    token = token_service.create_user_token(inexist_user)
    # make sure token service return correct value
    assert token is not None
    assert token != ''
    # Get user
    user = token_service.get_user_by_token(token)
    # make sure token service return correct value
    assert user is None

