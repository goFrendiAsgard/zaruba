from core.session.session_service import SessionService
from core.token.jwt_token_service import JWTTokenService
from module.auth.user.default_user_service import DefaultUserService
from module.auth.role.role_service import RoleService
from module.auth.user.repo.db_user_repo import DBUserRepo
from module.auth.role.repo.db_role_repo import DBRoleRepo
from module.auth.user.test_default_user_service import create_user_data
from helper.transport import LocalRPC, LocalMessageBus
from transport import AppMessageBus, AppRPC

from sqlalchemy import create_engine


def test_session_service_with_authenticated_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = AppMessageBus(LocalMessageBus())
    rpc = AppRPC(LocalRPC())
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'root')
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
    new_token = session_service.renew_access_token(token)
    # make sure token service return correct value
    assert token_service.get_user_by_token(new_token).id == root_user.id


def test_session_service_create_access_token_for_unauthenticated_user():
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = AppMessageBus(LocalMessageBus())
    rpc = AppRPC(LocalRPC())
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'root')
    token_service = JWTTokenService(user_service, 'secret', 'HS256', 1800)
    session_service = SessionService(user_service, token_service)
    is_error_happened = False
    try:
        session_service.create_access_token('invalid_identity', 'invalid_password')
    except:
        is_error_happened = True
    # make sure token service throw error
    assert is_error_happened
