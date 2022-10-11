from typing import Any, Optional, Mapping, Tuple
from modules.auth.auth.tokenOAuth2AuthService import TokenOAuth2AuthService
from modules.auth.user.test_util import create_user
from helpers.transport.localRpc import LocalRPC
from schemas.user import User
from fastapi.security import OAuth2, OAuth2PasswordBearer

import pytest

################################################
# -- ⚙️ Helpers
################################################

UNAUTHORIZED_ACTIVE_USER = create_user()
UNAUTHORIZED_ACTIVE_USER.id = 'mock_unauthorized_active_user_id'
UNAUTHORIZED_ACTIVE_USER.username = 'unauthorized_active_username'
UNAUTHORIZED_ACTIVE_USER.created_by = 'mock_user_id'
UNAUTHORIZED_ACTIVE_USER.active = True

AUTHORIZED_ACTIVE_USER = create_user()
AUTHORIZED_ACTIVE_USER.id = 'mock_authorized_active_user_id'
AUTHORIZED_ACTIVE_USER.username = 'authorized_active_username'
AUTHORIZED_ACTIVE_USER.created_by = 'mock_user_id'
AUTHORIZED_ACTIVE_USER.active = True

UNAUTHORIZED_INACTIVE_USER = create_user()
UNAUTHORIZED_INACTIVE_USER.id = 'mock_unauthorized_inactive_user_id'
UNAUTHORIZED_INACTIVE_USER.username = 'unauthorized_inactive_username'
UNAUTHORIZED_INACTIVE_USER.created_by = 'mock_user_id'
UNAUTHORIZED_INACTIVE_USER.active = False

AUTHORIZED_INACTIVE_USER = create_user()
AUTHORIZED_INACTIVE_USER.id = 'mock_authorized_inactive_user_id'
AUTHORIZED_INACTIVE_USER.username = 'authorized_inactive_username'
AUTHORIZED_INACTIVE_USER.created_by = 'mock_user_id'
AUTHORIZED_INACTIVE_USER.active = False


def get_user_by_token(token: str) -> Optional[User]:
    token_map: Mapping[str, Optional[User]] = {
        'unauthorized_active': UNAUTHORIZED_ACTIVE_USER,
        'authorized_active': AUTHORIZED_ACTIVE_USER,
        'unauthorized_inactive': UNAUTHORIZED_INACTIVE_USER,
        'authorized_inactive': AUTHORIZED_INACTIVE_USER,
    }
    if token in token_map:
        return token_map[token]
    if token == 'error':
        raise Exception('Emulating rpc error')
    return None


def is_user_authorized(user_data: Any) -> bool:
    user = User.parse_obj(user_data)
    if user.id in [AUTHORIZED_ACTIVE_USER.id, AUTHORIZED_INACTIVE_USER]:
        return True
    return False


class TokenOAuth2AuthMockRPC(LocalRPC):

    def __init__(self):
        super().__init__()

    def call(self, rpc_name: str, *args: Any) -> Any:
        if rpc_name == 'get_user_by_token':
            return get_user_by_token(args[0])
        if rpc_name == 'is_user_authorized':
            return is_user_authorized(args[0])
        return super().call(rpc_name, *args)


def init_test_oauth2_auth_service_components() -> Tuple[TokenOAuth2AuthService, TokenOAuth2AuthMockRPC, OAuth2]:
    rpc = TokenOAuth2AuthMockRPC()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenOAuth2AuthService(rpc, oauth2_scheme)
    return auth_service, rpc, oauth2_scheme

################################################
# -- 🧪 Test
################################################

@pytest.mark.asyncio
async def test_token_oauth2_auth_service_authorize_everyone_with_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.everyone(throw_error = True)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == AUTHORIZED_ACTIVE_USER
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
async def test_token_oauth2_auth_service_authorize_everyone_without_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.everyone(throw_error = False)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == AUTHORIZED_ACTIVE_USER
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
async def test_token_oauth2_auth_service_authorize_unauthenticated_with_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
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
async def test_token_oauth2_auth_service_authorize_unauthenticated_without_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
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
async def test_token_oauth2_auth_service_authorize_authenticated_with_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
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
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == AUTHORIZED_ACTIVE_USER
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
async def test_token_oauth2_auth_service_authorize_authenticated_without_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.is_authenticated(throw_error = False)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == AUTHORIZED_ACTIVE_USER
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
async def test_token_oauth2_auth_service_authorize_authorized_with_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.is_authorized('random_permission', throw_error = True)
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
    is_error = False
    try:
        user = await authorize(bearer_token='unauthorized_active')
    except:
        is_error = True
    assert is_error
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == AUTHORIZED_ACTIVE_USER
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
async def test_token_oauth2_auth_service_authorize_authorized_without_throw_error():
    auth_service, _, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.is_authorized('random_permission', throw_error = False)
    # test access without token
    user = await authorize(bearer_token=None, app_access_token=None)
    assert user is None
    # test access with invalid token 
    user = await authorize(bearer_token='invalid')
    assert user is None
    # test access with unauthorized active user token 
    user = await authorize(bearer_token='unauthorized_active')
    assert user == None
    # test access with authorized active user token 
    user = await authorize(bearer_token='authorized_active')
    assert user == AUTHORIZED_ACTIVE_USER
    # test access with unauthorized inactive user token 
    user = await authorize(bearer_token='unauthorized_inactive')
    assert user is None
    # test access with authorized inactive user token 
    user = await authorize(bearer_token='authorized_inactive')
    assert user is None
    # test access with trigger-rpc-error token 
    user = await authorize(bearer_token='error')
    assert user is None
