from typing import Any, Optional
from modules.auth.auth.noAuthService import NoAuthService
from helpers.transport.localRpc import LocalRPC
from schemas.user import User, UserData, UserWithoutPassword, UserResult
from starlette.requests import Request

################################################
# -- ⚙️ Mock data and objects
################################################

mock_guest_user = User(
    id="mock_guest_user_id",
    username="guest_username",
    email='',
    phone_number='',
    permissions=[],
    role_ids=[],
    active=True,
    full_name='',
    created_by='mock_user_id'
)

class MockRpc(LocalRPC):

    def __init__(self):
        super().__init__()
        self.rpc_name: Optional[str] = None

    def call(self, rpc_name: str, *args: Any) -> Any:
        self.rpc_name = rpc_name
        if rpc_name == 'get_guest_user':
            return mock_guest_user
        return super().call(rpc_name, *args)


################################################
# -- 🧪 Test
################################################

def test_no_auth_service_everyone_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.everyone(throw_error = True)
    user = authenticate(Request({'type': 'http'}))
    # make sure all parameters are passed to user and token service
    assert mock_rpc.rpc_name == 'get_guest_user'
    # make sure token service return correct value
    assert user == mock_guest_user


def test_no_auth_service_everyone_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.everyone(throw_error = False)
    user = authenticate(Request({'type': 'http'}))
    # make sure all parameters are passed to user and token service
    assert mock_rpc.rpc_name == 'get_guest_user'
    # make sure token service return correct value
    assert user == mock_guest_user


def test_no_auth_service_is_unauthenticated_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.is_unauthenticated(throw_error = True)
    is_error = False
    try:
        authenticate(Request({'type': 'http'}))
    except:
        is_error = True
    assert is_error


def test_no_auth_service_is_unauthenticated_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.is_unauthenticated(throw_error = False)
    user = authenticate(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user is None


def test_no_auth_service_is_authenticated_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.is_authenticated(throw_error = True)
    user = authenticate(Request({'type': 'http'}))
    # make sure all parameters are passed to user and token service
    assert mock_rpc.rpc_name == 'get_guest_user'
    # make sure token service return correct value
    assert user == mock_guest_user


def test_no_auth_service_is_authenticated_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.is_authenticated(throw_error = False)
    user = authenticate(Request({'type': 'http'}))
    # make sure all parameters are passed to user and token service
    assert mock_rpc.rpc_name == 'get_guest_user'
    # make sure token service return correct value
    assert user == mock_guest_user


def test_no_auth_service_is_authorized_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.is_authorized('random_permission', throw_error = True)
    user = authenticate(Request({'type': 'http'}))
    # make sure all parameters are passed to user and token service
    assert mock_rpc.rpc_name == 'get_guest_user'
    # make sure token service return correct value
    assert user == mock_guest_user


def test_no_auth_service_is_authorized_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authenticate = auth_service.is_authorized('random_permission', throw_error = False)
    user = authenticate(Request({'type': 'http'}))
    # make sure all parameters are passed to user and token service
    assert mock_rpc.rpc_name == 'get_guest_user'
    # make sure token service return correct value
    assert user == mock_guest_user
  