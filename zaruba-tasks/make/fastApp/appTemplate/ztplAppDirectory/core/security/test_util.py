from typing import Any, Tuple, Optional, Mapping
from core.security.noAuthService import NoAuthService
from core.security.tokenAuthService import TokenAuthService
from modules.auth.user.test_util import GUEST_USER, UNAUTHORIZED_ACTIVE_USER, AUTHORIZED_ACTIVE_USER, UNAUTHORIZED_INACTIVE_USER, AUTHORIZED_INACTIVE_USER
from helpers.transport.localRpc import LocalRPC
from schemas.user import User
from fastapi.security import OAuth2, OAuth2PasswordBearer

no_auth_mock_rpc = LocalRPC()
token_oauth2_auth_mock_rpc = LocalRPC()


@no_auth_mock_rpc.handle('get_guest_user')
def get_guest_user() -> User:
    return GUEST_USER


@token_oauth2_auth_mock_rpc.handle('get_user_by_access_token')
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


@token_oauth2_auth_mock_rpc.handle('is_user_authorized')
def is_user_authorized(user_data: Any, permission: str) -> bool:
    user = User.parse_obj(user_data)
    if user.id in [AUTHORIZED_ACTIVE_USER.id, AUTHORIZED_INACTIVE_USER]:
        return True
    return False


def init_test_no_auth_service_components() -> NoAuthService:
    auth_service = NoAuthService(no_auth_mock_rpc)
    return auth_service


def init_test_oauth2_auth_service_components() -> Tuple[TokenAuthService, OAuth2]:
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    auth_service = TokenAuthService(token_oauth2_auth_mock_rpc, oauth2_scheme)
    return auth_service, oauth2_scheme
