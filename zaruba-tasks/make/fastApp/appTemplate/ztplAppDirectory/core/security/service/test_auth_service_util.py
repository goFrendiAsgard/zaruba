from typing import Tuple
from module.auth.user.test_default_user_service_util import UNAUTHORIZED_ACTIVE_USER, UNAUTHORIZED_INACTIVE_USER, AUTHORIZED_ACTIVE_USER, AUTHORIZED_INACTIVE_USER
from helper.transport.local_rpc import LocalRPC, RPC
from schema.user import User, UserData
from fastapi.security import OAuth2, OAuth2PasswordBearer
from core.security.service.auth_service import AuthService
from core.security.rule.auth_rule import AuthRule
from core.security.rule.default_auth_rule import DefaultAuthRule
from core.security.middleware.user_fetcher import UserFetcher
from core.security.middleware.default_user_fetcher import DefaultUserFetcher
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
   