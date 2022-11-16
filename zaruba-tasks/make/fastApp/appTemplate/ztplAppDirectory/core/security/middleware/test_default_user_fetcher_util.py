from typing import Tuple, Mapping, Optional
from helper.transport.local_rpc import RPC, LocalRPC
from schema.user import User
from module.auth.user.test_default_user_service_util import AUTHORIZED_ACTIVE_USER
from core.security.middleware.default_user_fetcher import DefaultUserFetcher
from fastapi.security import OAuth2, OAuth2PasswordBearer


def create_rpc() -> RPC:
    rpc = LocalRPC()
    # handle get_user_by_access_token
    @rpc.handle('get_user_by_access_token')
    def get_user_by_token(token: str) -> Optional[User]:
        token_map: Mapping[str, Optional[User]] = {
            'authorized_active': AUTHORIZED_ACTIVE_USER,
        }
        if token in token_map:
            return token_map[token]
        if token == 'error':
            raise Exception('Emulating rpc error')
        return None
    # return rpc
    return rpc


def init_test_default_user_fetcher_components() -> Tuple[DefaultUserFetcher, RPC, OAuth2]:
    rpc = create_rpc()
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
    return user_fetcher, rpc, oauth2_scheme
 