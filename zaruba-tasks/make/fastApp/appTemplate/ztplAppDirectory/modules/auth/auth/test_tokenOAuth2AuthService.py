from typing import Any, Optional, Mapping
from modules.auth.auth.tokenOAuth2AuthService import TokenOAuth2AuthService
from helpers.transport.localRpc import LocalRPC
from schemas.user import User
from starlette.requests import Request
from fastapi.security import OAuth2PasswordBearer

import pytest

################################################
# -- ⚙️ Helpers
################################################

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

unauthorized_active_user = create_user()
unauthorized_active_user.id = 'mock_unauthorized_active_user_id'
unauthorized_active_user.username = 'unauthorized_active_username'
unauthorized_active_user.created_by = 'mock_user_id'
unauthorized_active_user.active = True

authorized_active_user = create_user()
authorized_active_user.id = 'mock_authorized_active_user_id'
authorized_active_user.username = 'authorized_active_username'
authorized_active_user.created_by = 'mock_user_id'
authorized_active_user.active = True

unauthorized_inactive_user = create_user()
unauthorized_inactive_user.id = 'mock_unauthorized_inactive_user_id'
unauthorized_inactive_user.username = 'unauthorized_inactive_username'
unauthorized_inactive_user.created_by = 'mock_user_id'
unauthorized_inactive_user.active = False

authorized_inactive_user = create_user()
authorized_inactive_user.id = 'mock_authorized_inactive_user_id'
authorized_inactive_user.username = 'authorized_inactive_username'
authorized_inactive_user.created_by = 'mock_user_id'
authorized_inactive_user.active = False


def get_user_by_token(token: str) -> Optional[User]:
    token_map: Mapping[str, Optional[User]] = {
        'unauthorized_active': unauthorized_active_user,
        'authorized_active': authorized_active_user,
        'unauthorized_inactive': unauthorized_inactive_user,
        'authorized_inactive': authorized_inactive_user,
    }
    if token in token_map:
        return token_map[token]
    if token == 'error':
        raise Exception('Emulating rpc error')
    return None


def is_user_authorized(user_data: Any) -> bool:
    user = User.parse_obj(user_data)
    if user.id in [authorized_active_user.id, authorized_inactive_user]:
        return True
    return False


class MockRpc(LocalRPC):

    def __init__(self):
        super().__init__()

    def call(self, rpc_name: str, *args: Any) -> Any:
        if rpc_name == 'get_user_by_token':
            return get_user_by_token(args[0])
        if rpc_name == 'is_user_authorized':
            return is_user_authorized(args[0])
        return super().call(rpc_name, *args)


################################################
# -- 🧪 Test
################################################

@pytest.mark.asyncio
async def test_no_auth_service_authorize_everyone_with_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authorize = auth_service.everyone(throw_error = True)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == unauthorized_active_user
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == authorized_active_user
    # test access with unauthorized inactive user token 
    user = await authorize(bearer_token='unauthorized_inactive')
    assert user is None
    # test access with authorized inactive user token 
    user = await authorize(bearer_token='authorized_inactive')
    assert user is None
    # test access with trigger-rpc-error token 
    user = await authorize(bearer_token='error')
    assert user is None


@pytest.mark.asyncio
async def test_no_auth_service_authorize_everyone_without_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authorize = auth_service.everyone(throw_error = False)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == unauthorized_active_user
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == authorized_active_user
    # test access with unauthorized inactive user token 
    user = await authorize(bearer_token='unauthorized_inactive')
    assert user is None
    # test access with authorized inactive user token 
    user = await authorize(bearer_token='authorized_inactive')
    assert user is None
    # test access with trigger-rpc-error token 
    user = await authorize(bearer_token='error')
    assert user is None


@pytest.mark.asyncio
async def test_no_auth_service_authorize_unauthenticated_with_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authorize = auth_service.is_unauthenticated(throw_error = True)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    is_error = False
    try:
        user = await authorize(bearer_token='unauthorized_active')
    except:
        is_error = True
    assert is_error
    # test access with authorized active user token 
    is_error = False
    try:
        user = await authorize(bearer_token='authorized_active')
    except:
        is_error = True
    assert is_error
    # test access with unauthorized inactive user token 
    user = await authorize(bearer_token='unauthorized_inactive')
    assert user is None
    # test access with authorized inactive user token 
    user = await authorize(bearer_token='authorized_inactive')
    assert user is None
    # test access with trigger-rpc-error token 
    user = await authorize(bearer_token='error')
    assert user is None


@pytest.mark.asyncio
async def test_no_auth_service_authorize_unauthenticated_without_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authorize = auth_service.is_unauthenticated(throw_error = False)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user is None
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user is None
    # test access with unauthorized inactive user token 
    user = await authorize(bearer_token='unauthorized_inactive')
    assert user is None
    # test access with authorized inactive user token 
    user = await authorize(bearer_token='authorized_inactive')
    assert user is None
    # test access with trigger-rpc-error token 
    user = await authorize(bearer_token='error')
    assert user is None


@pytest.mark.asyncio
async def test_no_auth_service_authorize_authenticated_with_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authorize = auth_service.is_authenticated(throw_error = True)
    # test access without token
    is_error = False
    try:
        user = await authorize(bearer_token=None, app_access_token=None)
    except:
        is_error = True
    assert is_error
    # test access with invalid token 
    is_error = False
    try:
        user = await authorize(bearer_token='invalid')
    except:
        is_error = True
    assert is_error
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == unauthorized_active_user
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == authorized_active_user
    # test access with unauthorized inactive user token 
    is_error = False
    try:
        user = await authorize(bearer_token='unauthorized_inactive')
    except:
        is_error = True
    assert is_error
    # test access with authorized inactive user token 
    is_error = False
    try:
        user = await authorize(bearer_token='authorized_inactive')
    except:
        is_error = True
    assert is_error
    # test access with trigger-rpc-error token 
    is_error = False
    try:
        user = await authorize(bearer_token='error')
    except:
        is_error = True
    assert is_error


@pytest.mark.asyncio
async def test_no_auth_service_authorize_authenticated_without_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authorize = auth_service.is_authenticated(throw_error = False)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == unauthorized_active_user
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == authorized_active_user
    # test access with unauthorized inactive user token 
    user = await authorize(bearer_token='unauthorized_inactive')
    assert user is None
    # test access with authorized inactive user token 
    user = await authorize(bearer_token='authorized_inactive')
    assert user is None
    # test access with trigger-rpc-error token 
    user = await authorize(bearer_token='error')
    assert user is None



@pytest.mark.asyncio
async def test_no_auth_service_authorize_authorized_with_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authenticate = auth_service.is_authenticated(throw_error = True)
    authorize = auth_service.is_authorized('random_permission', throw_error = True)
    # test access without token
    is_error = False
    try:
        authenticated_user = await authenticate(bearer_token=None, app_access_token=None)
    except:
        is_error = True
    assert is_error
    # test access with invalid token 
    is_error = False
    try:
        authenticated_user = await authenticate(bearer_token='invalid')
        user = await authorize(current_user = authenticated_user)
    except:
        is_error = True
    assert is_error
    # test access with unauthorized active user token 
    is_error = False
    try:
        authenticated_user = await authenticate(bearer_token='unauthorized_active')
        user = await authorize(current_user = authenticated_user)
    except:
        is_error = True
    assert is_error
    # test access with authorized active user token 
    authenticated_user = await authenticate(bearer_token='authorized_active')
    user = await authorize(current_user = authenticated_user)
    assert user == authorized_active_user
    # test access with unauthorized inactive user token 
    is_error = False
    try:
        authenticated_user = await authenticate(bearer_token='unauthorized_inactive')
        user = await authorize(current_user = authenticated_user)
    except:
        is_error = True
    assert is_error
    # test access with authorized inactive user token 
    is_error = False
    try:
        authenticated_user = await authenticate(bearer_token='authorized_inactive')
        user = await authorize(current_user = authenticated_user)
    except:
        is_error = True
    assert is_error
    # test access with trigger-rpc-error token 
    is_error = False
    try:
        authenticated_user = await authenticate(bearer_token='error')
        user = await authorize(current_user = authenticated_user)
    except:
        is_error = True
    assert is_error


@pytest.mark.asyncio
async def test_no_auth_service_authorize_authorized_without_throw_error():
    mock_rpc = MockRpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(mock_rpc, oauth2_scheme)
    authenticate = auth_service.is_authenticated(throw_error = False)
    authorize = auth_service.is_authorized('random_permission', throw_error = False)
    # test access without token
    authenticated_user = await authenticate(bearer_token=None, app_access_token=None)
    user = await authorize(current_user = authenticated_user)
    assert user is None
    # test access with invalid token 
    authenticated_user = await authenticate(bearer_token='invalid')
    user = await authorize(current_user = authenticated_user)
    assert user is None
    # test access with unauthorized active user token 
    authenticated_user = await authenticate(bearer_token='unauthorized_active')
    user = await authorize(current_user = authenticated_user)
    assert user == None
    # test access with authorized active user token 
    authenticated_user = await authenticate(bearer_token='authorized_active')
    user = await authorize(current_user = authenticated_user)
    assert user == authorized_active_user
    # test access with unauthorized inactive user token 
    authenticated_user = await authenticate(bearer_token='unauthorized_inactive')
    user = await authorize(current_user = authenticated_user)
    assert user is None
    # test access with authorized inactive user token 
    authenticated_user = await authenticate(bearer_token='authorized_inactive')
    user = await authorize(current_user = authenticated_user)
    assert user is None
    # test access with trigger-rpc-error token 
    authenticated_user = await authenticate(bearer_token='error')
    user = await authorize(current_user = authenticated_user)
    assert user is None
