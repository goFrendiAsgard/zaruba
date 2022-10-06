from typing import Any, Optional
from modules.auth.auth.noAuthService import NoAuthService
from helpers.transport.localRpc import LocalRPC
from schemas.user import User
from starlette.requests import Request

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

guest_user = create_user()
guest_user.id = 'mock_guest_user_id'
guest_user.username = 'guest_username'
guest_user.created_by = 'mock_user_id'

class MockRpc(LocalRPC):

    def __init__(self):
        super().__init__()

    def call(self, rpc_name: str, *args: Any) -> Any:
        if rpc_name == 'get_guest_user':
            return guest_user
        return super().call(rpc_name, *args)


################################################
# -- 🧪 Test
################################################

def test_no_auth_service_authorize_everyone_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.everyone(throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == guest_user


def test_no_auth_service_authorize_everyone_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.everyone(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == guest_user


def test_no_auth_service_authorize_unauthenticated_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.is_unauthenticated(throw_error = True)
    is_error = False
    try:
        authorize(Request({'type': 'http'}))
    except:
        is_error = True
    assert is_error


def test_no_auth_service_authorize_unauthenticated_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.is_unauthenticated(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user is None


def test_no_auth_service_authorize_authenticated_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.is_authenticated(throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == guest_user


def test_no_auth_service_authorize_authenticated_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.is_authenticated(throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == guest_user


def test_no_auth_service_authorize_authorized_with_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.is_authorized('random_permission', throw_error = True)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == guest_user


def test_no_auth_service_authorize_authorized_without_throw_error():
    mock_rpc = MockRpc()
    auth_service = NoAuthService(mock_rpc)
    authorize = auth_service.is_authorized('random_permission', throw_error = False)
    user = authorize(Request({'type': 'http'}))
    # make sure token service return correct value
    assert user == guest_user
  