from typing import Any, Tuple
from modules.auth.auth.noAuthService import NoAuthService
from modules.auth.user.test_util import create_user
from helpers.transport.localRpc import LocalRPC
from schemas.user import User
from starlette.requests import Request

################################################
# -- ⚙️ Helpers
################################################

GUEST_USER = create_user()
GUEST_USER.id = 'mock_guest_user_id'
GUEST_USER.username = 'guest_username'
GUEST_USER.created_by = 'mock_user_id'

class NoAuthMockRPC(LocalRPC):

    def __init__(self):
        super().__init__()

    def call(self, rpc_name: str, *args: Any) -> Any:
        if rpc_name == 'get_guest_user':
            return GUEST_USER
        return super().call(rpc_name, *args)


def init_test_no_auth_service_components() -> Tuple[NoAuthService, LocalRPC]:
    rpc = NoAuthMockRPC()
    auth_service = NoAuthService(rpc)
    return auth_service, rpc


################################################
# -- 🧪 Test
################################################

def test_no_auth_service_authorize_everyone_with_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.everyone(throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_everyone_without_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.everyone(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_unauthenticated_with_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.is_unauthenticated(throw_error = True)
    is_error = False
    try:
        authorize(Request({'type': 'http'}))
    except:
        is_error = True
    assert is_error


def test_no_auth_service_authorize_unauthenticated_without_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.is_unauthenticated(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user is None


def test_no_auth_service_authorize_authenticated_with_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.is_authenticated(throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_authenticated_without_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.is_authenticated(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_authorized_with_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.is_authorized('random_permission', throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER


def test_no_auth_service_authorize_authorized_without_throw_error():
    auth_service, _ = init_test_no_auth_service_components()
    authorize = auth_service.is_authorized('random_permission', throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == GUEST_USER
  