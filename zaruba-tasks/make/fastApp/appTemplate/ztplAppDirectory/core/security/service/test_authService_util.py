from typing import Tuple
from modules.auth.user.test_defaultUserService_util import UNAUTHORIZED_ACTIVE_USER, UNAUTHORIZED_INACTIVE_USER, AUTHORIZED_ACTIVE_USER, AUTHORIZED_INACTIVE_USER
from helpers.transport.localRpc import LocalRPC, RPC
from schemas.user import User, UserData
from fastapi.security import OAuth2, OAuth2PasswordBearer
from core.security.service.authService import AuthService
from core.security.rule.authRule import AuthRule
from core.security.rule.defaultAuthRule import DefaultAuthRule
from core.security.middleware.userFetcher import UserFetcher
from core.security.middleware.defaultUserFetcher import DefaultUserFetcher
from transport import AppRPC


def create_rpc() -> AppRPC:
    rpc = AppRPC(LocalRPC())
    # handle is_user_authorized
    @rpc.handle('is_user_authorized')
    def is_user_authorized(user_data: UserData, permission: str) -> bool:
        user = User.parse_obj(user_data)
        return user.id in [AUTHORIZED_ACTIVE_USER.id, AUTHORIZED_INACTIVE_USER]
    # return rpc
    return rpc


def init_test_auth_service_components() -> Tuple[AuthService, AuthRule, UserFetcher, OAuth2]:
    rpc = create_rpc()
    auth_rule = DefaultAuthRule(rpc)
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
    auth_service = AuthService(auth_rule, user_fetcher, 'root')
    return auth_service, auth_rule, user_fetcher, oauth2_scheme
   