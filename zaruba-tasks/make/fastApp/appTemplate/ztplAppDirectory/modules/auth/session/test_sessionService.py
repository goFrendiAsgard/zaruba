from modules.auth.session.sessionService import SessionService
from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from modules.auth.token.tokenService import JWTTokenService
from modules.auth.user.test_defaultUserService import create_user_data
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine


def test_session_service_with_authenticated_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    session_service = SessionService(user_service, token_service)
    # Init existing user
    root_user_data = create_user_data()
    root_user_data.username = 'root'
    root_user_data.email = 'root@innistrad.com'
    root_user_data.phone_number = '+6213456781'
    root_user_data.password = 'root_password'
    root_user_data.permissions = ['root']
    root_user = user_repo.insert(root_user_data)
    # test create token
    token = session_service.create_access_token('root', 'root_password')
    # make sure token service return correct value
    assert token_service.get_user_by_token(token).id == root_user.id
    # test refresh token
    new_token = session_service.refresh_access_token(token)
    # make sure token service return correct value
    assert token_service.get_user_by_token(new_token).id == root_user.id


def test_session_service_create_access_token_for_unauthenticated_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    session_service = SessionService(user_service, token_service)
    is_error_happened = False
    try:
        session_service.create_access_token('invalid_identity', 'invalid_password')
    except:
        is_error_happened = True
    # make sure token service throw error
    assert is_error_happened
