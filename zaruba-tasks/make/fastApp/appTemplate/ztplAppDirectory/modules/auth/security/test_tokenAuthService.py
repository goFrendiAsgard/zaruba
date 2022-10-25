from modules.auth.security.test_util import UNAUTHORIZED_ACTIVE_USER, AUTHORIZED_ACTIVE_USER, init_test_oauth2_auth_service_components

import pytest


@pytest.mark.asyncio
async def test_token_oauth2_auth_service_authorize_everyone_with_throw_error():
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.anyone(throw_error = True)
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
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.anyone(throw_error = False)
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
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.is_not_user(throw_error = True)
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
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.is_not_user(throw_error = False)
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
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.is_user(throw_error = True)
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
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.is_user(throw_error = False)
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
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.has_permission('random_permission', throw_error = True)
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
    auth_service, _ = init_test_oauth2_auth_service_components()
    authorize = auth_service.has_permission('random_permission', throw_error = False)
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
