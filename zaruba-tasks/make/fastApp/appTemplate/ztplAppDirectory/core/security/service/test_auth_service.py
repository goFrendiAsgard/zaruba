from core.security.service.test_auth_service_util import UNAUTHORIZED_ACTIVE_USER, AUTHORIZED_ACTIVE_USER, UNAUTHORIZED_INACTIVE_USER, AUTHORIZED_INACTIVE_USER, init_test_auth_service_components

import pytest


@pytest.mark.asyncio
async def test_auth_service_authorize_everyone_with_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.anyone(throw_error = True)
    # test access with no user
    user = await authorize(current_user=None)
    assert user is None
    # test access with unauthorized active user 
    user = await authorize(current_user=UNAUTHORIZED_ACTIVE_USER)
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user 
    user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    assert user == AUTHORIZED_ACTIVE_USER
    # test access with unauthorized inactive user 
    user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    assert user is None
    # test access with authorized inactive user 
    user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    assert user is None


@pytest.mark.asyncio
async def test_auth_service_authorize_everyone_without_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.anyone(throw_error = False)
    # test access with no user
    user = await authorize(current_user=None)
    assert user is None
    # test access with unauthorized active user 
    user = await authorize(current_user=UNAUTHORIZED_ACTIVE_USER)
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user 
    user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    assert user == AUTHORIZED_ACTIVE_USER
    # test access with unauthorized inactive user 
    user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    assert user is None
    # test access with authorized inactive user 
    user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    assert user is None


@pytest.mark.asyncio
async def test_auth_service_authorize_unauthenticated_with_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.is_visitor(throw_error = True)
    # test access without token
    user = await authorize(current_user=None)
    assert user is None
    # test access with authorized active user 
    is_error = False
    try:
        user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    # test access with unauthorized inactive user 
    user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    assert user is None
    # test access with authorized inactive user 
    user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    assert user is None


@pytest.mark.asyncio
async def test_auth_service_authorize_unauthenticated_without_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.is_visitor(throw_error = False)
    # test access without token
    user = await authorize(current_user=None)
    assert user is None
    # test access with unauthorized active user 
    user = await authorize(current_user=UNAUTHORIZED_ACTIVE_USER)
    assert user is None
    # test access with authorized active user 
    user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    assert user is None
    # test access with unauthorized inactive user 
    user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    assert user is None
    # test access with authorized inactive user 
    user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    assert user is None


@pytest.mark.asyncio
async def test_auth_service_authorize_authenticated_with_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.is_user(throw_error = True)
    # test access without token
    is_error = False
    try:
        user = await authorize(current_user=None)
    except:
        is_error = True
    assert is_error
    # test access with unauthorized active user 
    user = await authorize(current_user=UNAUTHORIZED_ACTIVE_USER)
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user 
    user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    assert user == AUTHORIZED_ACTIVE_USER
    # test access with unauthorized inactive user 
    is_error = False
    try:
        user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    except:
        is_error = True
    assert is_error
    # test access with authorized inactive user 
    is_error = False
    try:
        user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    except:
        is_error = True
    assert is_error


@pytest.mark.asyncio
async def test_auth_service_authorize_authenticated_without_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.is_user(throw_error = False)
    # test access without token
    user = await authorize(current_user=None)
    assert user is None
    # test access with unauthorized active user 
    user = await authorize(current_user=UNAUTHORIZED_ACTIVE_USER)
    assert user == UNAUTHORIZED_ACTIVE_USER
    # test access with authorized active user 
    user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    assert user == AUTHORIZED_ACTIVE_USER
    # test access with unauthorized inactive user 
    user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    assert user is None
    # test access with authorized inactive user 
    user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    assert user is None



@pytest.mark.asyncio
async def test_auth_service_authorize_authorized_with_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.has_permission('permission', throw_error = True)
    # test access without token
    is_error = False
    try:
        user = await authorize(current_user=None)
    except:
        is_error = True
    assert is_error
    # test access with unauthorized active user 
    is_error = False
    try:
        user = await authorize(current_user=UNAUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    # test access with authorized active user 
    user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    assert user == AUTHORIZED_ACTIVE_USER
    # test access with unauthorized inactive user 
    is_error = False
    try:
        user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    except:
        is_error = True
    assert is_error
    # test access with authorized inactive user 
    is_error = False
    try:
        user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    except:
        is_error = True
    assert is_error


@pytest.mark.asyncio
async def test_auth_service_authorize_authorized_without_throw_error():
    auth_service, _, _, _ = init_test_auth_service_components()
    authorize = auth_service.has_permission('permission', throw_error = False)
    # test access without token
    user = await authorize(current_user=None)
    assert user is None
    # test access with unauthorized active user 
    user = await authorize(current_user=UNAUTHORIZED_ACTIVE_USER)
    assert user == None
    # test access with authorized active user 
    user = await authorize(current_user=AUTHORIZED_ACTIVE_USER)
    assert user == AUTHORIZED_ACTIVE_USER
    # test access with unauthorized inactive user 
    user = await authorize(current_user=UNAUTHORIZED_INACTIVE_USER)
    assert user is None
    # test access with authorized inactive user 
    user = await authorize(current_user=AUTHORIZED_INACTIVE_USER)
    assert user is None
    