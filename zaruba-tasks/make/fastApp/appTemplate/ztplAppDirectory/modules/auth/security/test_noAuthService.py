from starlette.requests import Request
from modules.auth.security.test_util import GUEST_USER, init_test_no_auth_service_components


def test_no_auth_service_authorize_everyone_with_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.anyone(throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_everyone_without_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.anyone(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_unauthenticated_with_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.is_not_user(throw_error = True)
    is_error = False
    try:
        authorize(Request({'type': 'http'}))
    except:
        is_error = True
    assert is_error


def test_no_auth_service_authorize_unauthenticated_without_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.is_not_user(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user is None


def test_no_auth_service_authorize_authenticated_with_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.is_user(throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_authenticated_without_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.is_user(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_authorized_with_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.has_permission('random_permission', throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_authorized_without_throw_error():
    auth_service = init_test_no_auth_service_components()
    authorize = auth_service.has_permission('random_permission', throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER
  